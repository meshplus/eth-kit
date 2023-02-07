// Code generated by MockGen. DO NOT EDIT.
// Source: chain_ledger.go

// Package mock_ledger is a generated GoMock package.
package mock_ledger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/meshplus/bitxhub-kit/types"
	pb "github.com/meshplus/bitxhub-model/pb"
)

// MockChainLedger is a mock of ChainLedger interface.
type MockChainLedger struct {
	ctrl     *gomock.Controller
	recorder *MockChainLedgerMockRecorder
}

// MockChainLedgerMockRecorder is the mock recorder for MockChainLedger.
type MockChainLedgerMockRecorder struct {
	mock *MockChainLedger
}

// NewMockChainLedger creates a new mock instance.
func NewMockChainLedger(ctrl *gomock.Controller) *MockChainLedger {
	mock := &MockChainLedger{ctrl: ctrl}
	mock.recorder = &MockChainLedgerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainLedger) EXPECT() *MockChainLedgerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockChainLedger) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockChainLedgerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChainLedger)(nil).Close))
}

// GetBlock mocks base method.
func (m *MockChainLedger) GetBlock(height uint64, fullTx bool) (*pb.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", height, fullTx)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockChainLedgerMockRecorder) GetBlock(height, fullTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockChainLedger)(nil).GetBlock), height, fullTx)
}

// GetBlockByHash mocks base method.
func (m *MockChainLedger) GetBlockByHash(hash *types.Hash, fullTx bool) (*pb.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHash", hash, fullTx)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash.
func (mr *MockChainLedgerMockRecorder) GetBlockByHash(hash, fullTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockChainLedger)(nil).GetBlockByHash), hash, fullTx)
}

// GetBlockHash mocks base method.
func (m *MockChainLedger) GetBlockHash(height uint64) *types.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHash", height)
	ret0, _ := ret[0].(*types.Hash)
	return ret0
}

// GetBlockHash indicates an expected call of GetBlockHash.
func (mr *MockChainLedgerMockRecorder) GetBlockHash(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHash", reflect.TypeOf((*MockChainLedger)(nil).GetBlockHash), height)
}

// GetBlockSign mocks base method.
func (m *MockChainLedger) GetBlockSign(height uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockSign", height)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockSign indicates an expected call of GetBlockSign.
func (mr *MockChainLedgerMockRecorder) GetBlockSign(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockSign", reflect.TypeOf((*MockChainLedger)(nil).GetBlockSign), height)
}

// GetChainMeta mocks base method.
func (m *MockChainLedger) GetChainMeta() *pb.ChainMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainMeta")
	ret0, _ := ret[0].(*pb.ChainMeta)
	return ret0
}

// GetChainMeta indicates an expected call of GetChainMeta.
func (mr *MockChainLedgerMockRecorder) GetChainMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainMeta", reflect.TypeOf((*MockChainLedger)(nil).GetChainMeta))
}

// GetInterchainMeta mocks base method.
func (m *MockChainLedger) GetInterchainMeta(height uint64) (*pb.InterchainMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterchainMeta", height)
	ret0, _ := ret[0].(*pb.InterchainMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInterchainMeta indicates an expected call of GetInterchainMeta.
func (mr *MockChainLedgerMockRecorder) GetInterchainMeta(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterchainMeta", reflect.TypeOf((*MockChainLedger)(nil).GetInterchainMeta), height)
}

// GetReceipt mocks base method.
func (m *MockChainLedger) GetReceipt(hash *types.Hash) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", hash)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt.
func (mr *MockChainLedgerMockRecorder) GetReceipt(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockChainLedger)(nil).GetReceipt), hash)
}

// GetTransaction mocks base method.
func (m *MockChainLedger) GetTransaction(hash *types.Hash) (pb.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", hash)
	ret0, _ := ret[0].(pb.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockChainLedgerMockRecorder) GetTransaction(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockChainLedger)(nil).GetTransaction), hash)
}

// GetTransactionCount mocks base method.
func (m *MockChainLedger) GetTransactionCount(height uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionCount", height)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCount indicates an expected call of GetTransactionCount.
func (mr *MockChainLedgerMockRecorder) GetTransactionCount(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCount", reflect.TypeOf((*MockChainLedger)(nil).GetTransactionCount), height)
}

// GetTransactionMeta mocks base method.
func (m *MockChainLedger) GetTransactionMeta(hash *types.Hash) (*pb.TransactionMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionMeta", hash)
	ret0, _ := ret[0].(*pb.TransactionMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionMeta indicates an expected call of GetTransactionMeta.
func (mr *MockChainLedgerMockRecorder) GetTransactionMeta(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionMeta", reflect.TypeOf((*MockChainLedger)(nil).GetTransactionMeta), hash)
}

// LoadChainMeta mocks base method.
func (m *MockChainLedger) LoadChainMeta() *pb.ChainMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadChainMeta")
	ret0, _ := ret[0].(*pb.ChainMeta)
	return ret0
}

// LoadChainMeta indicates an expected call of LoadChainMeta.
func (mr *MockChainLedgerMockRecorder) LoadChainMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadChainMeta", reflect.TypeOf((*MockChainLedger)(nil).LoadChainMeta))
}

// PersistExecutionResult mocks base method.
func (m *MockChainLedger) PersistExecutionResult(block *pb.Block, receipts []*pb.Receipt, meta *pb.InterchainMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistExecutionResult", block, receipts, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// PersistExecutionResult indicates an expected call of PersistExecutionResult.
func (mr *MockChainLedgerMockRecorder) PersistExecutionResult(block, receipts, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistExecutionResult", reflect.TypeOf((*MockChainLedger)(nil).PersistExecutionResult), block, receipts, meta)
}

// PutBlock mocks base method.
func (m *MockChainLedger) PutBlock(height uint64, block *pb.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBlock", height, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutBlock indicates an expected call of PutBlock.
func (mr *MockChainLedgerMockRecorder) PutBlock(height, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBlock", reflect.TypeOf((*MockChainLedger)(nil).PutBlock), height, block)
}

// RollbackBlockChain mocks base method.
func (m *MockChainLedger) RollbackBlockChain(height uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackBlockChain", height)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackBlockChain indicates an expected call of RollbackBlockChain.
func (mr *MockChainLedgerMockRecorder) RollbackBlockChain(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackBlockChain", reflect.TypeOf((*MockChainLedger)(nil).RollbackBlockChain), height)
}

// UpdateChainMeta mocks base method.
func (m *MockChainLedger) UpdateChainMeta(arg0 *pb.ChainMeta) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateChainMeta", arg0)
}

// UpdateChainMeta indicates an expected call of UpdateChainMeta.
func (mr *MockChainLedgerMockRecorder) UpdateChainMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChainMeta", reflect.TypeOf((*MockChainLedger)(nil).UpdateChainMeta), arg0)
}