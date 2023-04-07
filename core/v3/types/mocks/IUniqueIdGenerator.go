// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IUniqueIdGenerator is an autogenerated mock type for the IUniqueIdGenerator type
type IUniqueIdGenerator struct {
	mock.Mock
}

// GenerateId provides a mock function with given fields:
func (_m *IUniqueIdGenerator) GenerateId() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type mockConstructorTestingTNewIUniqueIdGenerator interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUniqueIdGenerator creates a new instance of IUniqueIdGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUniqueIdGenerator(t mockConstructorTestingTNewIUniqueIdGenerator) *IUniqueIdGenerator {
	mock := &IUniqueIdGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
