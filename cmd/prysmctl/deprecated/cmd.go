package deprecated

import (
	"github.com/jumbochain/go-jumbochain-consensus/cmd/prysmctl/deprecated/checkpoint"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{}

func init() {
	Commands = append(Commands, checkpoint.Commands...)
}
