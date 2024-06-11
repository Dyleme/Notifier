// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Dyleme/Notifier/internal/service/service (interfaces: PeriodicTasksRepository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/periodic_tasks_mocks.go -package=mocks . PeriodicTasksRepository
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domains "github.com/Dyleme/Notifier/internal/domains"
	service "github.com/Dyleme/Notifier/internal/service/service"
	gomock "go.uber.org/mock/gomock"
)

// MockPeriodicTasksRepository is a mock of PeriodicTasksRepository interface.
type MockPeriodicTasksRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPeriodicTasksRepositoryMockRecorder
}

// MockPeriodicTasksRepositoryMockRecorder is the mock recorder for MockPeriodicTasksRepository.
type MockPeriodicTasksRepositoryMockRecorder struct {
	mock *MockPeriodicTasksRepository
}

// NewMockPeriodicTasksRepository creates a new mock instance.
func NewMockPeriodicTasksRepository(ctrl *gomock.Controller) *MockPeriodicTasksRepository {
	mock := &MockPeriodicTasksRepository{ctrl: ctrl}
	mock.recorder = &MockPeriodicTasksRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeriodicTasksRepository) EXPECT() *MockPeriodicTasksRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPeriodicTasksRepository) Add(arg0 context.Context, arg1 domains.PeriodicTask) (domains.PeriodicTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(domains.PeriodicTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockPeriodicTasksRepositoryMockRecorder) Add(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPeriodicTasksRepository)(nil).Add), arg0, arg1)
}

// Delete mocks base method.
func (m *MockPeriodicTasksRepository) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPeriodicTasksRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPeriodicTasksRepository)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockPeriodicTasksRepository) Get(arg0 context.Context, arg1 int) (domains.PeriodicTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(domains.PeriodicTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPeriodicTasksRepositoryMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPeriodicTasksRepository)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockPeriodicTasksRepository) List(arg0 context.Context, arg1 int, arg2 service.ListParams) ([]domains.PeriodicTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domains.PeriodicTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPeriodicTasksRepositoryMockRecorder) List(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPeriodicTasksRepository)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockPeriodicTasksRepository) Update(arg0 context.Context, arg1 domains.PeriodicTask) (domains.PeriodicTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(domains.PeriodicTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPeriodicTasksRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPeriodicTasksRepository)(nil).Update), arg0, arg1)
}