/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockvmclient

import (
	"context"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/golang/mock/gomock"
	"sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, VMName string, expand compute.InstanceViewTypes) (compute.VirtualMachine, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, VMName, expand)
	ret0, _ := ret[0].(compute.VirtualMachine)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, VMName, expand interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, VMName, expand)
}

// List mocks base method
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]compute.VirtualMachine, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]compute.VirtualMachine)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
}

// CreateOrUpdate mocks base method
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, VMName string, parameters compute.VirtualMachine, source string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, VMName, parameters, source)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, VMName, parameters, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, VMName, parameters, source)
}

// Update mocks base method
func (m *MockInterface) Update(ctx context.Context, resourceGroupName, VMName string, parameters compute.VirtualMachineUpdate, source string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, resourceGroupName, VMName, parameters, source)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockInterfaceMockRecorder) Update(ctx, resourceGroupName, VMName, parameters, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInterface)(nil).Update), ctx, resourceGroupName, VMName, parameters, source)
}

// UpdateAsync mocks base method
func (m *MockInterface) UpdateAsync(ctx context.Context, resourceGroupName, VMName string, parameters compute.VirtualMachineUpdate, source string) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAsync", ctx, resourceGroupName, VMName, parameters, source)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// UpdateAsync indicates an expected call of UpdateAsync
func (mr *MockInterfaceMockRecorder) UpdateAsync(ctx, resourceGroupName, VMName, parameters, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAsync", reflect.TypeOf((*MockInterface)(nil).UpdateAsync), ctx, resourceGroupName, VMName, parameters, source)
}

// WaitForUpdateResult waits for the response of the update request
func (m *MockInterface) WaitForUpdateResult(ctx context.Context, future *azure.Future, resourceGroupName, source string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForUpdateResult", ctx, future, resourceGroupName, source)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// WaitForUpdateResult waits for the response of the update request
func (mr *MockInterfaceMockRecorder) WaitForUpdateResult(ctx, future, resourceGroupName, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForUpdateResult", reflect.TypeOf((*MockInterface)(nil).WaitForUpdateResult), ctx, future, resourceGroupName, source)
}

// Delete mocks base method
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, VMName string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, VMName)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, VMName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, VMName)
}
