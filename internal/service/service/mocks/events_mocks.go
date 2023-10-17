// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Dyleme/Notifier/internal/service/service (interfaces: EventRepository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/events_mocks.go -package=mocks . EventRepository
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	domains "github.com/Dyleme/Notifier/internal/domains"
	service "github.com/Dyleme/Notifier/internal/service/service"
	gomock "go.uber.org/mock/gomock"
)

// MockEventRepository is a mock of EventRepository interface.
type MockEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryMockRecorder
}

// MockEventRepositoryMockRecorder is the mock recorder for MockEventRepository.
type MockEventRepositoryMockRecorder struct {
	mock *MockEventRepository
}

// NewMockEventRepository creates a new mock instance.
func NewMockEventRepository(ctrl *gomock.Controller) *MockEventRepository {
	mock := &MockEventRepository{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepository) EXPECT() *MockEventRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockEventRepository) Add(arg0 context.Context, arg1 domains.Event) (domains.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(domains.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockEventRepositoryMockRecorder) Add(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEventRepository)(nil).Add), arg0, arg1)
}

// Delay mocks base method.
func (m *MockEventRepository) Delay(arg0 context.Context, arg1, arg2 int, arg3 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delay indicates an expected call of Delay.
func (mr *MockEventRepositoryMockRecorder) Delay(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockEventRepository)(nil).Delay), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockEventRepository) Delete(arg0 context.Context, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEventRepositoryMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventRepository)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockEventRepository) Get(arg0 context.Context, arg1, arg2 int) (domains.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(domains.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEventRepositoryMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEventRepository)(nil).Get), arg0, arg1, arg2)
}

// GetNearestEventSendTime mocks base method.
func (m *MockEventRepository) GetNearestEventSendTime(arg0 context.Context) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearestEventSendTime", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearestEventSendTime indicates an expected call of GetNearestEventSendTime.
func (mr *MockEventRepositoryMockRecorder) GetNearestEventSendTime(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearestEventSendTime", reflect.TypeOf((*MockEventRepository)(nil).GetNearestEventSendTime), arg0)
}

// List mocks base method.
func (m *MockEventRepository) List(arg0 context.Context, arg1 int, arg2 service.ListParams) ([]domains.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domains.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockEventRepositoryMockRecorder) List(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEventRepository)(nil).List), arg0, arg1, arg2)
}

// ListEventsBefore mocks base method.
func (m *MockEventRepository) ListEventsBefore(arg0 context.Context, arg1 time.Time) ([]domains.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventsBefore", arg0, arg1)
	ret0, _ := ret[0].([]domains.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventsBefore indicates an expected call of ListEventsBefore.
func (mr *MockEventRepositoryMockRecorder) ListEventsBefore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventsBefore", reflect.TypeOf((*MockEventRepository)(nil).ListEventsBefore), arg0, arg1)
}

// ListInPeriod mocks base method.
func (m *MockEventRepository) ListInPeriod(arg0 context.Context, arg1 int, arg2, arg3 time.Time, arg4 service.ListParams) ([]domains.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInPeriod", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]domains.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInPeriod indicates an expected call of ListInPeriod.
func (mr *MockEventRepositoryMockRecorder) ListInPeriod(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInPeriod", reflect.TypeOf((*MockEventRepository)(nil).ListInPeriod), arg0, arg1, arg2, arg3, arg4)
}

// MarkNotified mocks base method.
func (m *MockEventRepository) MarkNotified(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkNotified", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkNotified indicates an expected call of MarkNotified.
func (mr *MockEventRepositoryMockRecorder) MarkNotified(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkNotified", reflect.TypeOf((*MockEventRepository)(nil).MarkNotified), arg0, arg1)
}

// Update mocks base method.
func (m *MockEventRepository) Update(arg0 context.Context, arg1 domains.Event) (domains.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(domains.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEventRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventRepository)(nil).Update), arg0, arg1)
}

// UpdateNotificationParams mocks base method.
func (m *MockEventRepository) UpdateNotificationParams(arg0 context.Context, arg1, arg2 int, arg3 domains.NotificationParams) (domains.NotificationParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotificationParams", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(domains.NotificationParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNotificationParams indicates an expected call of UpdateNotificationParams.
func (mr *MockEventRepositoryMockRecorder) UpdateNotificationParams(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotificationParams", reflect.TypeOf((*MockEventRepository)(nil).UpdateNotificationParams), arg0, arg1, arg2, arg3)
}
