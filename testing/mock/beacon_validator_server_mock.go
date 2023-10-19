// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1 (interfaces: BeaconNodeValidatorServer,BeaconNodeValidator_WaitForActivationServer,BeaconNodeValidator_WaitForChainStartServer,BeaconNodeValidator_StreamDutiesServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	eth "github.com/jumbochain/go-jumbochain-consensus/proto/prysm/v1alpha1"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockBeaconNodeValidatorServer is a mock of BeaconNodeValidatorServer interface.
type MockBeaconNodeValidatorServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidatorServerMockRecorder
}

// MockBeaconNodeValidatorServerMockRecorder is the mock recorder for MockBeaconNodeValidatorServer.
type MockBeaconNodeValidatorServerMockRecorder struct {
	mock *MockBeaconNodeValidatorServer
}

// NewMockBeaconNodeValidatorServer creates a new mock instance.
func NewMockBeaconNodeValidatorServer(ctrl *gomock.Controller) *MockBeaconNodeValidatorServer {
	mock := &MockBeaconNodeValidatorServer{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidatorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconNodeValidatorServer) EXPECT() *MockBeaconNodeValidatorServerMockRecorder {
	return m.recorder
}

// AggregatedSigAndAggregationBits mocks base method.
func (m *MockBeaconNodeValidatorServer) AggregatedSigAndAggregationBits(arg0 context.Context, arg1 *eth.AggregatedSigAndAggregationBitsRequest) (*eth.AggregatedSigAndAggregationBitsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatedSigAndAggregationBits", arg0, arg1)
	ret0, _ := ret[0].(*eth.AggregatedSigAndAggregationBitsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AggregatedSigAndAggregationBits indicates an expected call of AggregatedSigAndAggregationBits.
func (mr *MockBeaconNodeValidatorServerMockRecorder) AggregatedSigAndAggregationBits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatedSigAndAggregationBits", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).AggregatedSigAndAggregationBits), arg0, arg1)
}

// AssignValidatorToSubnet mocks base method.
func (m *MockBeaconNodeValidatorServer) AssignValidatorToSubnet(arg0 context.Context, arg1 *eth.AssignValidatorToSubnetRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignValidatorToSubnet", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignValidatorToSubnet indicates an expected call of AssignValidatorToSubnet.
func (mr *MockBeaconNodeValidatorServerMockRecorder) AssignValidatorToSubnet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignValidatorToSubnet", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).AssignValidatorToSubnet), arg0, arg1)
}

// CheckDoppelGanger mocks base method.
func (m *MockBeaconNodeValidatorServer) CheckDoppelGanger(arg0 context.Context, arg1 *eth.DoppelGangerRequest) (*eth.DoppelGangerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDoppelGanger", arg0, arg1)
	ret0, _ := ret[0].(*eth.DoppelGangerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDoppelGanger indicates an expected call of CheckDoppelGanger.
func (mr *MockBeaconNodeValidatorServerMockRecorder) CheckDoppelGanger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDoppelGanger", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).CheckDoppelGanger), arg0, arg1)
}

// DomainData mocks base method.
func (m *MockBeaconNodeValidatorServer) DomainData(arg0 context.Context, arg1 *eth.DomainRequest) (*eth.DomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainData", arg0, arg1)
	ret0, _ := ret[0].(*eth.DomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainData indicates an expected call of DomainData.
func (mr *MockBeaconNodeValidatorServerMockRecorder) DomainData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainData", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).DomainData), arg0, arg1)
}

// GetAttestationData mocks base method.
func (m *MockBeaconNodeValidatorServer) GetAttestationData(arg0 context.Context, arg1 *eth.AttestationDataRequest) (*eth.AttestationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttestationData", arg0, arg1)
	ret0, _ := ret[0].(*eth.AttestationData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttestationData indicates an expected call of GetAttestationData.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetAttestationData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttestationData", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetAttestationData), arg0, arg1)
}

// GetBeaconBlock mocks base method.
func (m *MockBeaconNodeValidatorServer) GetBeaconBlock(arg0 context.Context, arg1 *eth.BlockRequest) (*eth.GenericBeaconBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBeaconBlock", arg0, arg1)
	ret0, _ := ret[0].(*eth.GenericBeaconBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBeaconBlock indicates an expected call of GetBeaconBlock.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetBeaconBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBeaconBlock", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetBeaconBlock), arg0, arg1)
}

// GetDuties mocks base method.
func (m *MockBeaconNodeValidatorServer) GetDuties(arg0 context.Context, arg1 *eth.DutiesRequest) (*eth.DutiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDuties", arg0, arg1)
	ret0, _ := ret[0].(*eth.DutiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDuties indicates an expected call of GetDuties.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetDuties(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDuties", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetDuties), arg0, arg1)
}

// GetFeeRecipientByPubKey mocks base method.
func (m *MockBeaconNodeValidatorServer) GetFeeRecipientByPubKey(arg0 context.Context, arg1 *eth.FeeRecipientByPubKeyRequest) (*eth.FeeRecipientByPubKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeeRecipientByPubKey", arg0, arg1)
	ret0, _ := ret[0].(*eth.FeeRecipientByPubKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeeRecipientByPubKey indicates an expected call of GetFeeRecipientByPubKey.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetFeeRecipientByPubKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeeRecipientByPubKey", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetFeeRecipientByPubKey), arg0, arg1)
}

// GetSyncCommitteeContribution mocks base method.
func (m *MockBeaconNodeValidatorServer) GetSyncCommitteeContribution(arg0 context.Context, arg1 *eth.SyncCommitteeContributionRequest) (*eth.SyncCommitteeContribution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncCommitteeContribution", arg0, arg1)
	ret0, _ := ret[0].(*eth.SyncCommitteeContribution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncCommitteeContribution indicates an expected call of GetSyncCommitteeContribution.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetSyncCommitteeContribution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncCommitteeContribution", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetSyncCommitteeContribution), arg0, arg1)
}

// GetSyncMessageBlockRoot mocks base method.
func (m *MockBeaconNodeValidatorServer) GetSyncMessageBlockRoot(arg0 context.Context, arg1 *emptypb.Empty) (*eth.SyncMessageBlockRootResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncMessageBlockRoot", arg0, arg1)
	ret0, _ := ret[0].(*eth.SyncMessageBlockRootResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncMessageBlockRoot indicates an expected call of GetSyncMessageBlockRoot.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetSyncMessageBlockRoot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncMessageBlockRoot", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetSyncMessageBlockRoot), arg0, arg1)
}

// GetSyncSubcommitteeIndex mocks base method.
func (m *MockBeaconNodeValidatorServer) GetSyncSubcommitteeIndex(arg0 context.Context, arg1 *eth.SyncSubcommitteeIndexRequest) (*eth.SyncSubcommitteeIndexResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncSubcommitteeIndex", arg0, arg1)
	ret0, _ := ret[0].(*eth.SyncSubcommitteeIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncSubcommitteeIndex indicates an expected call of GetSyncSubcommitteeIndex.
func (mr *MockBeaconNodeValidatorServerMockRecorder) GetSyncSubcommitteeIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncSubcommitteeIndex", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).GetSyncSubcommitteeIndex), arg0, arg1)
}

// MultipleValidatorStatus mocks base method.
func (m *MockBeaconNodeValidatorServer) MultipleValidatorStatus(arg0 context.Context, arg1 *eth.MultipleValidatorStatusRequest) (*eth.MultipleValidatorStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipleValidatorStatus", arg0, arg1)
	ret0, _ := ret[0].(*eth.MultipleValidatorStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultipleValidatorStatus indicates an expected call of MultipleValidatorStatus.
func (mr *MockBeaconNodeValidatorServerMockRecorder) MultipleValidatorStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipleValidatorStatus", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).MultipleValidatorStatus), arg0, arg1)
}

// PrepareBeaconProposer mocks base method.
func (m *MockBeaconNodeValidatorServer) PrepareBeaconProposer(arg0 context.Context, arg1 *eth.PrepareBeaconProposerRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareBeaconProposer", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareBeaconProposer indicates an expected call of PrepareBeaconProposer.
func (mr *MockBeaconNodeValidatorServerMockRecorder) PrepareBeaconProposer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareBeaconProposer", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).PrepareBeaconProposer), arg0, arg1)
}

// ProposeAttestation mocks base method.
func (m *MockBeaconNodeValidatorServer) ProposeAttestation(arg0 context.Context, arg1 *eth.Attestation) (*eth.AttestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposeAttestation", arg0, arg1)
	ret0, _ := ret[0].(*eth.AttestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeAttestation indicates an expected call of ProposeAttestation.
func (mr *MockBeaconNodeValidatorServerMockRecorder) ProposeAttestation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeAttestation", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).ProposeAttestation), arg0, arg1)
}

// ProposeBeaconBlock mocks base method.
func (m *MockBeaconNodeValidatorServer) ProposeBeaconBlock(arg0 context.Context, arg1 *eth.GenericSignedBeaconBlock) (*eth.ProposeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposeBeaconBlock", arg0, arg1)
	ret0, _ := ret[0].(*eth.ProposeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeBeaconBlock indicates an expected call of ProposeBeaconBlock.
func (mr *MockBeaconNodeValidatorServerMockRecorder) ProposeBeaconBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeBeaconBlock", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).ProposeBeaconBlock), arg0, arg1)
}

// ProposeExit mocks base method.
func (m *MockBeaconNodeValidatorServer) ProposeExit(arg0 context.Context, arg1 *eth.SignedVoluntaryExit) (*eth.ProposeExitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposeExit", arg0, arg1)
	ret0, _ := ret[0].(*eth.ProposeExitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeExit indicates an expected call of ProposeExit.
func (mr *MockBeaconNodeValidatorServerMockRecorder) ProposeExit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeExit", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).ProposeExit), arg0, arg1)
}

// StreamBlocksAltair mocks base method.
func (m *MockBeaconNodeValidatorServer) StreamBlocksAltair(arg0 *eth.StreamBlocksRequest, arg1 eth.BeaconNodeValidator_StreamBlocksAltairServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamBlocksAltair", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StreamBlocksAltair indicates an expected call of StreamBlocksAltair.
func (mr *MockBeaconNodeValidatorServerMockRecorder) StreamBlocksAltair(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamBlocksAltair", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).StreamBlocksAltair), arg0, arg1)
}

// StreamDuties mocks base method.
func (m *MockBeaconNodeValidatorServer) StreamDuties(arg0 *eth.DutiesRequest, arg1 eth.BeaconNodeValidator_StreamDutiesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamDuties", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StreamDuties indicates an expected call of StreamDuties.
func (mr *MockBeaconNodeValidatorServerMockRecorder) StreamDuties(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamDuties", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).StreamDuties), arg0, arg1)
}

// SubmitAggregateSelectionProof mocks base method.
func (m *MockBeaconNodeValidatorServer) SubmitAggregateSelectionProof(arg0 context.Context, arg1 *eth.AggregateSelectionRequest) (*eth.AggregateSelectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitAggregateSelectionProof", arg0, arg1)
	ret0, _ := ret[0].(*eth.AggregateSelectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitAggregateSelectionProof indicates an expected call of SubmitAggregateSelectionProof.
func (mr *MockBeaconNodeValidatorServerMockRecorder) SubmitAggregateSelectionProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitAggregateSelectionProof", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).SubmitAggregateSelectionProof), arg0, arg1)
}

// SubmitSignedAggregateSelectionProof mocks base method.
func (m *MockBeaconNodeValidatorServer) SubmitSignedAggregateSelectionProof(arg0 context.Context, arg1 *eth.SignedAggregateSubmitRequest) (*eth.SignedAggregateSubmitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitSignedAggregateSelectionProof", arg0, arg1)
	ret0, _ := ret[0].(*eth.SignedAggregateSubmitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitSignedAggregateSelectionProof indicates an expected call of SubmitSignedAggregateSelectionProof.
func (mr *MockBeaconNodeValidatorServerMockRecorder) SubmitSignedAggregateSelectionProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitSignedAggregateSelectionProof", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).SubmitSignedAggregateSelectionProof), arg0, arg1)
}

// SubmitSignedContributionAndProof mocks base method.
func (m *MockBeaconNodeValidatorServer) SubmitSignedContributionAndProof(arg0 context.Context, arg1 *eth.SignedContributionAndProof) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitSignedContributionAndProof", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitSignedContributionAndProof indicates an expected call of SubmitSignedContributionAndProof.
func (mr *MockBeaconNodeValidatorServerMockRecorder) SubmitSignedContributionAndProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitSignedContributionAndProof", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).SubmitSignedContributionAndProof), arg0, arg1)
}

// SubmitSyncMessage mocks base method.
func (m *MockBeaconNodeValidatorServer) SubmitSyncMessage(arg0 context.Context, arg1 *eth.SyncCommitteeMessage) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitSyncMessage", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitSyncMessage indicates an expected call of SubmitSyncMessage.
func (mr *MockBeaconNodeValidatorServerMockRecorder) SubmitSyncMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitSyncMessage", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).SubmitSyncMessage), arg0, arg1)
}

// SubmitValidatorRegistrations mocks base method.
func (m *MockBeaconNodeValidatorServer) SubmitValidatorRegistrations(arg0 context.Context, arg1 *eth.SignedValidatorRegistrationsV1) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitValidatorRegistrations", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitValidatorRegistrations indicates an expected call of SubmitValidatorRegistrations.
func (mr *MockBeaconNodeValidatorServerMockRecorder) SubmitValidatorRegistrations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitValidatorRegistrations", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).SubmitValidatorRegistrations), arg0, arg1)
}

// SubscribeCommitteeSubnets mocks base method.
func (m *MockBeaconNodeValidatorServer) SubscribeCommitteeSubnets(arg0 context.Context, arg1 *eth.CommitteeSubnetsSubscribeRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeCommitteeSubnets", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeCommitteeSubnets indicates an expected call of SubscribeCommitteeSubnets.
func (mr *MockBeaconNodeValidatorServerMockRecorder) SubscribeCommitteeSubnets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeCommitteeSubnets", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).SubscribeCommitteeSubnets), arg0, arg1)
}

// ValidatorIndex mocks base method.
func (m *MockBeaconNodeValidatorServer) ValidatorIndex(arg0 context.Context, arg1 *eth.ValidatorIndexRequest) (*eth.ValidatorIndexResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorIndex", arg0, arg1)
	ret0, _ := ret[0].(*eth.ValidatorIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorIndex indicates an expected call of ValidatorIndex.
func (mr *MockBeaconNodeValidatorServerMockRecorder) ValidatorIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorIndex", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).ValidatorIndex), arg0, arg1)
}

// ValidatorStatus mocks base method.
func (m *MockBeaconNodeValidatorServer) ValidatorStatus(arg0 context.Context, arg1 *eth.ValidatorStatusRequest) (*eth.ValidatorStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorStatus", arg0, arg1)
	ret0, _ := ret[0].(*eth.ValidatorStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorStatus indicates an expected call of ValidatorStatus.
func (mr *MockBeaconNodeValidatorServerMockRecorder) ValidatorStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorStatus", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).ValidatorStatus), arg0, arg1)
}

// WaitForActivation mocks base method.
func (m *MockBeaconNodeValidatorServer) WaitForActivation(arg0 *eth.ValidatorActivationRequest, arg1 eth.BeaconNodeValidator_WaitForActivationServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForActivation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForActivation indicates an expected call of WaitForActivation.
func (mr *MockBeaconNodeValidatorServerMockRecorder) WaitForActivation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForActivation", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).WaitForActivation), arg0, arg1)
}

// WaitForChainStart mocks base method.
func (m *MockBeaconNodeValidatorServer) WaitForChainStart(arg0 *emptypb.Empty, arg1 eth.BeaconNodeValidator_WaitForChainStartServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForChainStart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForChainStart indicates an expected call of WaitForChainStart.
func (mr *MockBeaconNodeValidatorServerMockRecorder) WaitForChainStart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForChainStart", reflect.TypeOf((*MockBeaconNodeValidatorServer)(nil).WaitForChainStart), arg0, arg1)
}

// MockBeaconNodeValidator_WaitForActivationServer is a mock of BeaconNodeValidator_WaitForActivationServer interface.
type MockBeaconNodeValidator_WaitForActivationServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidator_WaitForActivationServerMockRecorder
}

// MockBeaconNodeValidator_WaitForActivationServerMockRecorder is the mock recorder for MockBeaconNodeValidator_WaitForActivationServer.
type MockBeaconNodeValidator_WaitForActivationServerMockRecorder struct {
	mock *MockBeaconNodeValidator_WaitForActivationServer
}

// NewMockBeaconNodeValidator_WaitForActivationServer creates a new mock instance.
func NewMockBeaconNodeValidator_WaitForActivationServer(ctrl *gomock.Controller) *MockBeaconNodeValidator_WaitForActivationServer {
	mock := &MockBeaconNodeValidator_WaitForActivationServer{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidator_WaitForActivationServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconNodeValidator_WaitForActivationServer) EXPECT() *MockBeaconNodeValidator_WaitForActivationServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) Send(arg0 *eth.ValidatorActivationResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockBeaconNodeValidator_WaitForActivationServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockBeaconNodeValidator_WaitForActivationServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBeaconNodeValidator_WaitForActivationServer)(nil).SetTrailer), arg0)
}

// MockBeaconNodeValidator_WaitForChainStartServer is a mock of BeaconNodeValidator_WaitForChainStartServer interface.
type MockBeaconNodeValidator_WaitForChainStartServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder
}

// MockBeaconNodeValidator_WaitForChainStartServerMockRecorder is the mock recorder for MockBeaconNodeValidator_WaitForChainStartServer.
type MockBeaconNodeValidator_WaitForChainStartServerMockRecorder struct {
	mock *MockBeaconNodeValidator_WaitForChainStartServer
}

// NewMockBeaconNodeValidator_WaitForChainStartServer creates a new mock instance.
func NewMockBeaconNodeValidator_WaitForChainStartServer(ctrl *gomock.Controller) *MockBeaconNodeValidator_WaitForChainStartServer {
	mock := &MockBeaconNodeValidator_WaitForChainStartServer{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidator_WaitForChainStartServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) EXPECT() *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) Send(arg0 *eth.ChainStartResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockBeaconNodeValidator_WaitForChainStartServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockBeaconNodeValidator_WaitForChainStartServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBeaconNodeValidator_WaitForChainStartServer)(nil).SetTrailer), arg0)
}

// MockBeaconNodeValidator_StreamDutiesServer is a mock of BeaconNodeValidator_StreamDutiesServer interface.
type MockBeaconNodeValidator_StreamDutiesServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconNodeValidator_StreamDutiesServerMockRecorder
}

// MockBeaconNodeValidator_StreamDutiesServerMockRecorder is the mock recorder for MockBeaconNodeValidator_StreamDutiesServer.
type MockBeaconNodeValidator_StreamDutiesServerMockRecorder struct {
	mock *MockBeaconNodeValidator_StreamDutiesServer
}

// NewMockBeaconNodeValidator_StreamDutiesServer creates a new mock instance.
func NewMockBeaconNodeValidator_StreamDutiesServer(ctrl *gomock.Controller) *MockBeaconNodeValidator_StreamDutiesServer {
	mock := &MockBeaconNodeValidator_StreamDutiesServer{ctrl: ctrl}
	mock.recorder = &MockBeaconNodeValidator_StreamDutiesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconNodeValidator_StreamDutiesServer) EXPECT() *MockBeaconNodeValidator_StreamDutiesServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) Send(arg0 *eth.DutiesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockBeaconNodeValidator_StreamDutiesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockBeaconNodeValidator_StreamDutiesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBeaconNodeValidator_StreamDutiesServer)(nil).SetTrailer), arg0)
}
