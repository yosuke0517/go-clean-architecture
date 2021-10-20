// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/todo.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	model "todo/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockTodoRepository is a mock of TodoRepository interface
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockTodoRepository) FindAll() ([]*model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockTodoRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTodoRepository)(nil).FindAll))
}

// Find mocks base method
func (m *MockTodoRepository) Find(word string) ([]*model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", word)
	ret0, _ := ret[0].([]*model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockTodoRepositoryMockRecorder) Find(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTodoRepository)(nil).Find), word)
}

// Create mocks base method
func (m *MockTodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", todo)
	ret0, _ := ret[0].(*model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTodoRepositoryMockRecorder) Create(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoRepository)(nil).Create), todo)
}

// Update mocks base method
func (m *MockTodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", todo)
	ret0, _ := ret[0].(*model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockTodoRepositoryMockRecorder) Update(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), todo)
}