// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	types "simple_text_editor/core/v3/types"

	mock "github.com/stretchr/testify/mock"
)

// IDialogHelper is an autogenerated mock type for the IDialogHelper type
type IDialogHelper struct {
	mock.Mock
}

// OkCancelMessageDialog provides a mock function with given fields: title, message
func (_m *IDialogHelper) OkCancelMessageDialog(title string, message string) (types.Button, error) {
	ret := _m.Called(title, message)

	var r0 types.Button
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (types.Button, error)); ok {
		return rf(title, message)
	}
	if rf, ok := ret.Get(0).(func(string, string) types.Button); ok {
		r0 = rf(title, message)
	} else {
		r0 = ret.Get(0).(types.Button)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(title, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenFileDialog provides a mock function with given fields:
func (_m *IDialogHelper) OpenFileDialog() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveFileDialog provides a mock function with given fields: defaultFileNameWithExt
func (_m *IDialogHelper) SaveFileDialog(defaultFileNameWithExt string) (string, error) {
	ret := _m.Called(defaultFileNameWithExt)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(defaultFileNameWithExt)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(defaultFileNameWithExt)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(defaultFileNameWithExt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIDialogHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDialogHelper creates a new instance of IDialogHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDialogHelper(t mockConstructorTestingTNewIDialogHelper) *IDialogHelper {
	mock := &IDialogHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
