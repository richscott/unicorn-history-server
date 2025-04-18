// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/G-Research/unicorn-history-server/internal/yunikorn (interfaces: Client)
//
// Generated by this command:
//
//	mockgen -destination=mock_client.go -package=yunikorn github.com/G-Research/unicorn-history-server/internal/yunikorn Client
//

// Package yunikorn is a generated GoMock package.
package yunikorn

import (
	context "context"
	http "net/http"
	reflect "reflect"

	webservice "github.com/G-Research/yunikorn-core/pkg/webservice"
	dao "github.com/G-Research/yunikorn-core/pkg/webservice/dao"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetApplication mocks base method.
func (m *MockClient) GetApplication(arg0 context.Context, arg1, arg2, arg3 string) (*dao.ApplicationDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*dao.ApplicationDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockClientMockRecorder) GetApplication(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockClient)(nil).GetApplication), arg0, arg1, arg2, arg3)
}

// GetApplications mocks base method.
func (m *MockClient) GetApplications(arg0 context.Context, arg1, arg2 string) ([]*dao.ApplicationDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplications", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*dao.ApplicationDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplications indicates an expected call of GetApplications.
func (mr *MockClientMockRecorder) GetApplications(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplications", reflect.TypeOf((*MockClient)(nil).GetApplications), arg0, arg1, arg2)
}

// GetAppsHistory mocks base method.
func (m *MockClient) GetAppsHistory(arg0 context.Context) ([]*dao.ApplicationHistoryDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppsHistory", arg0)
	ret0, _ := ret[0].([]*dao.ApplicationHistoryDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppsHistory indicates an expected call of GetAppsHistory.
func (mr *MockClientMockRecorder) GetAppsHistory(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppsHistory", reflect.TypeOf((*MockClient)(nil).GetAppsHistory), arg0)
}

// GetClusters mocks base method.
func (m *MockClient) GetClusters(arg0 context.Context) ([]*dao.ClusterDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusters", arg0)
	ret0, _ := ret[0].([]*dao.ClusterDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters.
func (mr *MockClientMockRecorder) GetClusters(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockClient)(nil).GetClusters), arg0)
}

// GetContainersHistory mocks base method.
func (m *MockClient) GetContainersHistory(arg0 context.Context) ([]*dao.ContainerHistoryDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainersHistory", arg0)
	ret0, _ := ret[0].([]*dao.ContainerHistoryDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainersHistory indicates an expected call of GetContainersHistory.
func (mr *MockClientMockRecorder) GetContainersHistory(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainersHistory", reflect.TypeOf((*MockClient)(nil).GetContainersHistory), arg0)
}

// GetEventStream mocks base method.
func (m *MockClient) GetEventStream(arg0 context.Context) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventStream", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventStream indicates an expected call of GetEventStream.
func (mr *MockClientMockRecorder) GetEventStream(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventStream", reflect.TypeOf((*MockClient)(nil).GetEventStream), arg0)
}

// GetFullStateDump mocks base method.
func (m *MockClient) GetFullStateDump(arg0 context.Context) (*webservice.AggregatedStateInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullStateDump", arg0)
	ret0, _ := ret[0].(*webservice.AggregatedStateInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullStateDump indicates an expected call of GetFullStateDump.
func (mr *MockClientMockRecorder) GetFullStateDump(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullStateDump", reflect.TypeOf((*MockClient)(nil).GetFullStateDump), arg0)
}

// GetPartitionNodes mocks base method.
func (m *MockClient) GetPartitionNodes(arg0 context.Context, arg1 string) ([]*dao.NodeDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionNodes", arg0, arg1)
	ret0, _ := ret[0].([]*dao.NodeDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionNodes indicates an expected call of GetPartitionNodes.
func (mr *MockClientMockRecorder) GetPartitionNodes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionNodes", reflect.TypeOf((*MockClient)(nil).GetPartitionNodes), arg0, arg1)
}

// GetPartitionQueue mocks base method.
func (m *MockClient) GetPartitionQueue(arg0 context.Context, arg1, arg2 string) (*dao.PartitionQueueDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionQueue", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dao.PartitionQueueDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionQueue indicates an expected call of GetPartitionQueue.
func (mr *MockClientMockRecorder) GetPartitionQueue(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionQueue", reflect.TypeOf((*MockClient)(nil).GetPartitionQueue), arg0, arg1, arg2)
}

// GetPartitionQueues mocks base method.
func (m *MockClient) GetPartitionQueues(arg0 context.Context, arg1 string) (*dao.PartitionQueueDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitionQueues", arg0, arg1)
	ret0, _ := ret[0].(*dao.PartitionQueueDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionQueues indicates an expected call of GetPartitionQueues.
func (mr *MockClientMockRecorder) GetPartitionQueues(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionQueues", reflect.TypeOf((*MockClient)(nil).GetPartitionQueues), arg0, arg1)
}

// GetPartitions mocks base method.
func (m *MockClient) GetPartitions(arg0 context.Context) ([]*dao.PartitionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitions", arg0)
	ret0, _ := ret[0].([]*dao.PartitionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitions indicates an expected call of GetPartitions.
func (mr *MockClientMockRecorder) GetPartitions(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitions", reflect.TypeOf((*MockClient)(nil).GetPartitions), arg0)
}

// Healthcheck mocks base method.
func (m *MockClient) Healthcheck(arg0 context.Context) (*dao.SchedulerHealthDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck", arg0)
	ret0, _ := ret[0].(*dao.SchedulerHealthDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck.
func (mr *MockClientMockRecorder) Healthcheck(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockClient)(nil).Healthcheck), arg0)
}

// NodeUtilizations mocks base method.
func (m *MockClient) NodeUtilizations(arg0 context.Context) ([]*dao.PartitionNodesUtilDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeUtilizations", arg0)
	ret0, _ := ret[0].([]*dao.PartitionNodesUtilDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeUtilizations indicates an expected call of NodeUtilizations.
func (mr *MockClientMockRecorder) NodeUtilizations(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeUtilizations", reflect.TypeOf((*MockClient)(nil).NodeUtilizations), arg0)
}
