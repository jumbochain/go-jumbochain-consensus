// Package main defines a validator client, a critical actor in Ethereum which manages
// a keystore of private keys, connects to a beacon node to receive assignments,
// and submits blocks/attestations as needed.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	runtimeDebug "runtime/debug"

	joonix "github.com/joonix/log"
	"github.com/jumbochain/go-jumbochain-consensus/cmd"
	accountcommands "github.com/jumbochain/go-jumbochain-consensus/cmd/validator/accounts"
	dbcommands "github.com/jumbochain/go-jumbochain-consensus/cmd/validator/db"
	"github.com/jumbochain/go-jumbochain-consensus/cmd/validator/flags"
	slashingprotectioncommands "github.com/jumbochain/go-jumbochain-consensus/cmd/validator/slashing-protection"
	walletcommands "github.com/jumbochain/go-jumbochain-consensus/cmd/validator/wallet"
	"github.com/jumbochain/go-jumbochain-consensus/cmd/validator/web"
	"github.com/jumbochain/go-jumbochain-consensus/config/features"
	"github.com/jumbochain/go-jumbochain-consensus/io/file"
	"github.com/jumbochain/go-jumbochain-consensus/io/logs"
	"github.com/jumbochain/go-jumbochain-consensus/monitoring/journald"
	"github.com/jumbochain/go-jumbochain-consensus/runtime/debug"
	prefixed "github.com/jumbochain/go-jumbochain-consensus/runtime/logging/logrus-prefixed-formatter"
	_ "github.com/jumbochain/go-jumbochain-consensus/runtime/maxprocs"
	"github.com/jumbochain/go-jumbochain-consensus/runtime/tos"
	"github.com/jumbochain/go-jumbochain-consensus/runtime/version"
	"github.com/jumbochain/go-jumbochain-consensus/validator/node"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func startNode(ctx *cli.Context) error {
	// verify if ToS accepted
	if err := tos.VerifyTosAcceptedOrPrompt(ctx); err != nil {
		return err
	}

	validatorClient, err := node.NewValidatorClient(ctx)
	if err != nil {
		return err
	}
	validatorClient.Start()
	return nil
}

var appFlags = []cli.Flag{
	flags.BeaconRPCProviderFlag,
	flags.BeaconRPCGatewayProviderFlag,
	flags.BeaconRESTApiProviderFlag,
	flags.CertFlag,
	flags.GraffitiFlag,
	flags.DisablePenaltyRewardLogFlag,
	flags.InteropStartIndex,
	flags.InteropNumValidators,
	flags.EnableRPCFlag,
	flags.RPCHost,
	flags.RPCPort,
	flags.GRPCGatewayPort,
	flags.GRPCGatewayHost,
	flags.GrpcRetriesFlag,
	flags.GrpcRetryDelayFlag,
	flags.GrpcHeadersFlag,
	flags.GPRCGatewayCorsDomain,
	flags.DisableAccountMetricsFlag,
	flags.MonitoringPortFlag,
	flags.SlasherRPCProviderFlag,
	flags.SlasherCertFlag,
	flags.WalletPasswordFileFlag,
	flags.WalletDirFlag,
	flags.EnableWebFlag,
	flags.GraffitiFileFlag,
	// Consensys' Web3Signer flags
	flags.Web3SignerURLFlag,
	flags.Web3SignerPublicValidatorKeysFlag,
	flags.SuggestedFeeRecipientFlag,
	flags.ProposerSettingsURLFlag,
	flags.ProposerSettingsFlag,
	flags.EnableBuilderFlag,
	flags.BuilderGasLimitFlag,
	////////////////////
	cmd.DisableMonitoringFlag,
	cmd.MonitoringHostFlag,
	cmd.BackupWebhookOutputDir,
	cmd.EnableBackupWebhookFlag,
	cmd.MinimalConfigFlag,
	cmd.E2EConfigFlag,
	cmd.VerbosityFlag,
	cmd.DataDirFlag,
	cmd.ClearDB,
	cmd.ForceClearDB,
	cmd.EnableTracingFlag,
	cmd.TracingProcessNameFlag,
	cmd.TracingEndpointFlag,
	cmd.TraceSampleFractionFlag,
	cmd.LogFormat,
	cmd.LogFileName,
	cmd.ConfigFileFlag,
	cmd.ChainConfigFileFlag,
	cmd.GrpcMaxCallRecvMsgSizeFlag,
	cmd.ApiTimeoutFlag,
	debug.PProfFlag,
	debug.PProfAddrFlag,
	debug.PProfPortFlag,
	debug.MemProfileRateFlag,
	debug.CPUProfileFlag,
	debug.TraceFlag,
	debug.BlockProfileRateFlag,
	debug.MutexProfileFractionFlag,
	cmd.AcceptTosFlag,
}

func init() {
	appFlags = cmd.WrapFlags(append(appFlags, features.ValidatorFlags...))
}

func main() {
	app := cli.App{}
	app.Name = "validator"
	app.Usage = `launches an Ethereum validator client that interacts with a beacon chain, starts proposer and attester services, p2p connections, and more`
	app.Version = version.Version()
	app.Action = func(ctx *cli.Context) error {
		if err := startNode(ctx); err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return nil
	}
	app.Commands = []*cli.Command{
		walletcommands.Commands,
		accountcommands.Commands,
		slashingprotectioncommands.Commands,
		dbcommands.Commands,
		web.Commands,
	}

	app.Flags = appFlags

	app.Before = func(ctx *cli.Context) error {
		// Load flags from config file, if specified.
		if err := cmd.LoadFlagsFromConfig(ctx, app.Flags); err != nil {
			return err
		}

		format := ctx.String(cmd.LogFormat.Name)
		switch format {
		case "text":
			formatter := new(prefixed.TextFormatter)
			formatter.TimestampFormat = "2006-01-02 15:04:05"
			formatter.FullTimestamp = true
			// If persistent log files are written - we disable the log messages coloring because
			// the colors are ANSI codes and seen as Gibberish in the log files.
			formatter.DisableColors = ctx.String(cmd.LogFileName.Name) != ""
			logrus.SetFormatter(formatter)
		case "fluentd":
			f := joonix.NewFormatter()
			if err := joonix.DisableTimestampFormat(f); err != nil {
				panic(err)
			}
			logrus.SetFormatter(f)
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case "journald":
			if err := journald.Enable(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown log format %s", format)
		}

		logFileName := ctx.String(cmd.LogFileName.Name)
		if logFileName != "" {
			if err := logs.ConfigurePersistentLogging(logFileName); err != nil {
				log.WithError(err).Error("Failed to configuring logging to disk.")
			}
		}

		// Fix data dir for Windows users.
		outdatedDataDir := filepath.Join(file.HomeDir(), "AppData", "Roaming", "Eth2Validators")
		currentDataDir := flags.DefaultValidatorDir()
		if err := cmd.FixDefaultDataDir(outdatedDataDir, currentDataDir); err != nil {
			log.WithError(err).Error("Cannot update data directory")
		}

		if err := debug.Setup(ctx); err != nil {
			return err
		}
		return cmd.ValidateNoArgs(ctx)
	}

	app.After = func(ctx *cli.Context) error {
		debug.Exit(ctx)
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			log.Errorf("Runtime panic: %v\n%v", x, string(runtimeDebug.Stack()))
			panic(x)
		}
	}()

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
	}
}
