// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Dyleme/Notifier/internal/service/service (interfaces: TgImagesRepository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/tg_images_mocks.go -package=mocks . TgImagesRepository
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domains "github.com/Dyleme/Notifier/internal/domains"
	gomock "go.uber.org/mock/gomock"
)

// MockTgImagesRepository is a mock of TgImagesRepository interface.
type MockTgImagesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTgImagesRepositoryMockRecorder
}

// MockTgImagesRepositoryMockRecorder is the mock recorder for MockTgImagesRepository.
type MockTgImagesRepositoryMockRecorder struct {
	mock *MockTgImagesRepository
}

// NewMockTgImagesRepository creates a new mock instance.
func NewMockTgImagesRepository(ctrl *gomock.Controller) *MockTgImagesRepository {
	mock := &MockTgImagesRepository{ctrl: ctrl}
	mock.recorder = &MockTgImagesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTgImagesRepository) EXPECT() *MockTgImagesRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTgImagesRepository) Add(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockTgImagesRepositoryMockRecorder) Add(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTgImagesRepository)(nil).Add), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockTgImagesRepository) Get(arg0 context.Context, arg1 string) (domains.TgImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(domains.TgImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTgImagesRepositoryMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTgImagesRepository)(nil).Get), arg0, arg1)
}
