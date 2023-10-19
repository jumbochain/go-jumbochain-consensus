package node

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/blockchain/testing"
	"github.com/jumbochain/go-jumbochain-consensus/beacon-chain/rpc/testutil"
	syncmock "github.com/jumbochain/go-jumbochain-consensus/beacon-chain/sync/initial-sync/testing"
	"github.com/jumbochain/go-jumbochain-consensus/consensus-types/primitives"
	"github.com/jumbochain/go-jumbochain-consensus/testing/assert"
	"github.com/jumbochain/go-jumbochain-consensus/testing/require"
	"github.com/jumbochain/go-jumbochain-consensus/testing/util"
)

func TestSyncStatus(t *testing.T) {
	currentSlot := new(primitives.Slot)
	*currentSlot = 110
	state, err := util.NewBeaconState()
	require.NoError(t, err)
	err = state.SetSlot(100)
	require.NoError(t, err)
	chainService := &mock.ChainService{Slot: currentSlot, State: state, Optimistic: true}
	syncChecker := &syncmock.Sync{}
	syncChecker.IsSyncing = true

	s := &Server{
		HeadFetcher:               chainService,
		GenesisTimeFetcher:        chainService,
		OptimisticModeFetcher:     chainService,
		SyncChecker:               syncChecker,
		ExecutionChainInfoFetcher: &testutil.MockExecutionChainInfoFetcher{},
	}

	request := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	writer := httptest.NewRecorder()
	writer.Body = &bytes.Buffer{}

	s.GetSyncStatus(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	resp := &SyncStatusResponse{}
	require.NoError(t, json.Unmarshal(writer.Body.Bytes(), resp))
	require.NotNil(t, resp)
	assert.Equal(t, "100", resp.Data.HeadSlot)
	assert.Equal(t, "10", resp.Data.SyncDistance)
	assert.Equal(t, true, resp.Data.IsSyncing)
	assert.Equal(t, true, resp.Data.IsOptimistic)
	assert.Equal(t, false, resp.Data.ElOffline)
}
