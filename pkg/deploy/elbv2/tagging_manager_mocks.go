// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pjtatlow/aws-load-balancer-controller/pkg/deploy/elbv2 (interfaces: TaggingManager)

// Package elbv2 is a generated GoMock package.
package elbv2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tracking "github.com/pjtatlow/aws-load-balancer-controller/pkg/deploy/tracking"
)

// MockTaggingManager is a mock of TaggingManager interface.
type MockTaggingManager struct {
	ctrl     *gomock.Controller
	recorder *MockTaggingManagerMockRecorder
}

// MockTaggingManagerMockRecorder is the mock recorder for MockTaggingManager.
type MockTaggingManagerMockRecorder struct {
	mock *MockTaggingManager
}

// NewMockTaggingManager creates a new mock instance.
func NewMockTaggingManager(ctrl *gomock.Controller) *MockTaggingManager {
	mock := &MockTaggingManager{ctrl: ctrl}
	mock.recorder = &MockTaggingManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaggingManager) EXPECT() *MockTaggingManagerMockRecorder {
	return m.recorder
}

// ListListenerRules mocks base method.
func (m *MockTaggingManager) ListListenerRules(arg0 context.Context, arg1 string) ([]ListenerRuleWithTags, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListListenerRules", arg0, arg1)
	ret0, _ := ret[0].([]ListenerRuleWithTags)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListListenerRules indicates an expected call of ListListenerRules.
func (mr *MockTaggingManagerMockRecorder) ListListenerRules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListListenerRules", reflect.TypeOf((*MockTaggingManager)(nil).ListListenerRules), arg0, arg1)
}

// ListListeners mocks base method.
func (m *MockTaggingManager) ListListeners(arg0 context.Context, arg1 string) ([]ListenerWithTags, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListListeners", arg0, arg1)
	ret0, _ := ret[0].([]ListenerWithTags)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListListeners indicates an expected call of ListListeners.
func (mr *MockTaggingManagerMockRecorder) ListListeners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListListeners", reflect.TypeOf((*MockTaggingManager)(nil).ListListeners), arg0, arg1)
}

// ListLoadBalancers mocks base method.
func (m *MockTaggingManager) ListLoadBalancers(arg0 context.Context, arg1 ...tracking.TagFilter) ([]LoadBalancerWithTags, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLoadBalancers", varargs...)
	ret0, _ := ret[0].([]LoadBalancerWithTags)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoadBalancers indicates an expected call of ListLoadBalancers.
func (mr *MockTaggingManagerMockRecorder) ListLoadBalancers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoadBalancers", reflect.TypeOf((*MockTaggingManager)(nil).ListLoadBalancers), varargs...)
}

// ListTargetGroups mocks base method.
func (m *MockTaggingManager) ListTargetGroups(arg0 context.Context, arg1 ...tracking.TagFilter) ([]TargetGroupWithTags, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTargetGroups", varargs...)
	ret0, _ := ret[0].([]TargetGroupWithTags)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTargetGroups indicates an expected call of ListTargetGroups.
func (mr *MockTaggingManagerMockRecorder) ListTargetGroups(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTargetGroups", reflect.TypeOf((*MockTaggingManager)(nil).ListTargetGroups), varargs...)
}

// ReconcileTags mocks base method.
func (m *MockTaggingManager) ReconcileTags(arg0 context.Context, arg1 string, arg2 map[string]string, arg3 ...ReconcileTagsOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReconcileTags", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTags indicates an expected call of ReconcileTags.
func (mr *MockTaggingManagerMockRecorder) ReconcileTags(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTags", reflect.TypeOf((*MockTaggingManager)(nil).ReconcileTags), varargs...)
}
