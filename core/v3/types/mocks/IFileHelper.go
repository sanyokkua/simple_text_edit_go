// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	types "simple_text_editor/core/v3/types"

	mock "github.com/stretchr/testify/mock"
)

// IFileHelper is an autogenerated mock type for the IFileHelper type
type IFileHelper struct {
	mock.Mock
}

// CreateNewFileEmpty provides a mock function with given fields:
func (_m *IFileHelper) CreateNewFileEmpty() (*types.FileStruct, error) {
	ret := _m.Called()

	var r0 *types.FileStruct
	var r1 error
	if rf, ok := ret.Get(0).(func() (*types.FileStruct, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *types.FileStruct); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.FileStruct)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNewFileWithData provides a mock function with given fields: path, originalContent
func (_m *IFileHelper) CreateNewFileWithData(path string, originalContent string) (*types.FileStruct, error) {
	ret := _m.Called(path, originalContent)

	var r0 *types.FileStruct
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*types.FileStruct, error)); ok {
		return rf(path, originalContent)
	}
	if rf, ok := ret.Get(0).(func(string, string) *types.FileStruct); ok {
		r0 = rf(path, originalContent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.FileStruct)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, originalContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFileExtensionFromPath provides a mock function with given fields: filePath
func (_m *IFileHelper) GetFileExtensionFromPath(filePath string) (types.FileTypeExtension, error) {
	ret := _m.Called(filePath)

	var r0 types.FileTypeExtension
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.FileTypeExtension, error)); ok {
		return rf(filePath)
	}
	if rf, ok := ret.Get(0).(func(string) types.FileTypeExtension); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(types.FileTypeExtension)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFileNameFromPath provides a mock function with given fields: filePath
func (_m *IFileHelper) GetFileNameFromPath(filePath string) (string, error) {
	ret := _m.Called(filePath)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(filePath)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFileTypeFromExtension provides a mock function with given fields: fileExtension
func (_m *IFileHelper) GetFileTypeFromExtension(fileExtension types.FileTypeExtension) (types.FileTypeKey, error) {
	ret := _m.Called(fileExtension)

	var r0 types.FileTypeKey
	var r1 error
	if rf, ok := ret.Get(0).(func(types.FileTypeExtension) (types.FileTypeKey, error)); ok {
		return rf(fileExtension)
	}
	if rf, ok := ret.Get(0).(func(types.FileTypeExtension) types.FileTypeKey); ok {
		r0 = rf(fileExtension)
	} else {
		r0 = ret.Get(0).(types.FileTypeKey)
	}

	if rf, ok := ret.Get(1).(func(types.FileTypeExtension) error); ok {
		r1 = rf(fileExtension)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIFileHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewIFileHelper creates a new instance of IFileHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIFileHelper(t mockConstructorTestingTNewIFileHelper) *IFileHelper {
	mock := &IFileHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}