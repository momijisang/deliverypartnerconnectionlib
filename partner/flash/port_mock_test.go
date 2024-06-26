// Code generated by MockGen. DO NOT EDIT.
// Source: port.go
//
// Generated by this command:
//
//	mockgen -source=port.go -package=flash -destination=port_mock_test.go
//

// Package flash is a generated GoMock package.
package flash

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFlashCreateOrderAPI is a mock of FlashCreateOrderAPI interface.
type MockFlashCreateOrderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFlashCreateOrderAPIMockRecorder
}

// MockFlashCreateOrderAPIMockRecorder is the mock recorder for MockFlashCreateOrderAPI.
type MockFlashCreateOrderAPIMockRecorder struct {
	mock *MockFlashCreateOrderAPI
}

// NewMockFlashCreateOrderAPI creates a new mock instance.
func NewMockFlashCreateOrderAPI(ctrl *gomock.Controller) *MockFlashCreateOrderAPI {
	mock := &MockFlashCreateOrderAPI{ctrl: ctrl}
	mock.recorder = &MockFlashCreateOrderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlashCreateOrderAPI) EXPECT() *MockFlashCreateOrderAPIMockRecorder {
	return m.recorder
}

// PostForm mocks base method.
func (m *MockFlashCreateOrderAPI) PostForm(endpoint string, form map[string]string) (FlashCreateOrderAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", endpoint, form)
	ret0, _ := ret[0].(FlashCreateOrderAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostForm indicates an expected call of PostForm.
func (mr *MockFlashCreateOrderAPIMockRecorder) PostForm(endpoint, form any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockFlashCreateOrderAPI)(nil).PostForm), endpoint, form)
}

// MockFlashUpdateShipmentInfo is a mock of FlashUpdateShipmentInfo interface.
type MockFlashUpdateShipmentInfo struct {
	ctrl     *gomock.Controller
	recorder *MockFlashUpdateShipmentInfoMockRecorder
}

// MockFlashUpdateShipmentInfoMockRecorder is the mock recorder for MockFlashUpdateShipmentInfo.
type MockFlashUpdateShipmentInfoMockRecorder struct {
	mock *MockFlashUpdateShipmentInfo
}

// NewMockFlashUpdateShipmentInfo creates a new mock instance.
func NewMockFlashUpdateShipmentInfo(ctrl *gomock.Controller) *MockFlashUpdateShipmentInfo {
	mock := &MockFlashUpdateShipmentInfo{ctrl: ctrl}
	mock.recorder = &MockFlashUpdateShipmentInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlashUpdateShipmentInfo) EXPECT() *MockFlashUpdateShipmentInfoMockRecorder {
	return m.recorder
}

// PostForm mocks base method.
func (m *MockFlashUpdateShipmentInfo) PostForm(endpoint string, form map[string]string) (FlashOrderUpdateAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", endpoint, form)
	ret0, _ := ret[0].(FlashOrderUpdateAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostForm indicates an expected call of PostForm.
func (mr *MockFlashUpdateShipmentInfoMockRecorder) PostForm(endpoint, form any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockFlashUpdateShipmentInfo)(nil).PostForm), endpoint, form)
}

// MockFlashDeleteOrderAPI is a mock of FlashDeleteOrderAPI interface.
type MockFlashDeleteOrderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFlashDeleteOrderAPIMockRecorder
}

// MockFlashDeleteOrderAPIMockRecorder is the mock recorder for MockFlashDeleteOrderAPI.
type MockFlashDeleteOrderAPIMockRecorder struct {
	mock *MockFlashDeleteOrderAPI
}

// NewMockFlashDeleteOrderAPI creates a new mock instance.
func NewMockFlashDeleteOrderAPI(ctrl *gomock.Controller) *MockFlashDeleteOrderAPI {
	mock := &MockFlashDeleteOrderAPI{ctrl: ctrl}
	mock.recorder = &MockFlashDeleteOrderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlashDeleteOrderAPI) EXPECT() *MockFlashDeleteOrderAPIMockRecorder {
	return m.recorder
}

// PostForm mocks base method.
func (m *MockFlashDeleteOrderAPI) PostForm(endpoint string, form map[string]string) (FlashOrderDeleteAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", endpoint, form)
	ret0, _ := ret[0].(FlashOrderDeleteAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostForm indicates an expected call of PostForm.
func (mr *MockFlashDeleteOrderAPIMockRecorder) PostForm(endpoint, form any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockFlashDeleteOrderAPI)(nil).PostForm), endpoint, form)
}
