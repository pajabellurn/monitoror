// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SimpleValidator is an autogenerated mock type for the SimpleValidator type
type SimpleValidator struct {
	mock.Mock
}

// IsValid provides a mock function with given fields:
func (_m *SimpleValidator) IsValid() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}