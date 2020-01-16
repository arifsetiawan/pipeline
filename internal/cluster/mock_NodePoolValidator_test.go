// Code generated by mockery v1.0.0. DO NOT EDIT.

package cluster

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockNodePoolValidator is an autogenerated mock type for the NodePoolValidator type
type MockNodePoolValidator struct {
	mock.Mock
}

// ValidateNew provides a mock function with given fields: ctx, cluster, rawNodePool
func (_m *MockNodePoolValidator) ValidateNew(ctx context.Context, cluster Cluster, rawNodePool NewRawNodePool) error {
	ret := _m.Called(ctx, cluster, rawNodePool)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Cluster, NewRawNodePool) error); ok {
		r0 = rf(ctx, cluster, rawNodePool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}