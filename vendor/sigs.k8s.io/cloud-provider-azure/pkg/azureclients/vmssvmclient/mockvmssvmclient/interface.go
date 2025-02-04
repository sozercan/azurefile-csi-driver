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

package mockvmssvmclient

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
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, VMScaleSetName, instanceID string, expand compute.InstanceViewTypes) (compute.VirtualMachineScaleSetVM, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, VMScaleSetName, instanceID, expand)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVM)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, VMScaleSetName, instanceID, expand interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, VMScaleSetName, instanceID, expand)
}

// List mocks base method
func (m *MockInterface) List(ctx context.Context, resourceGroupName, virtualMachineScaleSetName, expand string) ([]compute.VirtualMachineScaleSetVM, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName, virtualMachineScaleSetName, expand)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSetVM)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName, virtualMachineScaleSetName, expand interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName, virtualMachineScaleSetName, expand)
}

// Update mocks base method
func (m *MockInterface) Update(ctx context.Context, resourceGroupName, VMScaleSetName, instanceID string, parameters compute.VirtualMachineScaleSetVM, source string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, resourceGroupName, VMScaleSetName, instanceID, parameters, source)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockInterfaceMockRecorder) Update(ctx, resourceGroupName, VMScaleSetName, instanceID, parameters, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInterface)(nil).Update), ctx, resourceGroupName, VMScaleSetName, instanceID, parameters, source)
}

// UpdateSync mocks base method
func (m *MockInterface) UpdateAsync(ctx context.Context, resourceGroupName, VMScaleSetName, instanceID string, parameters compute.VirtualMachineScaleSetVM, source string) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAsync", ctx, resourceGroupName, VMScaleSetName, instanceID, parameters, source)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// UpdateSync indicates an expected call of UpdateSync
func (mr *MockInterfaceMockRecorder) UpdateAsync(ctx, resourceGroupName, VMScaleSetName, instanceID, parameters, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAsync", reflect.TypeOf((*MockInterface)(nil).UpdateAsync), ctx, resourceGroupName, VMScaleSetName, instanceID, parameters, source)
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

// UpdateVMs mocks base method
func (m *MockInterface) UpdateVMs(ctx context.Context, resourceGroupName, VMScaleSetName string, instances map[string]compute.VirtualMachineScaleSetVM, source string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVMs", ctx, resourceGroupName, VMScaleSetName, instances, source)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// UpdateVMs indicates an expected call of UpdateVMs
func (mr *MockInterfaceMockRecorder) UpdateVMs(ctx, resourceGroupName, VMScaleSetName, instances, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVMs", reflect.TypeOf((*MockInterface)(nil).UpdateVMs), ctx, resourceGroupName, VMScaleSetName, instances, source)
}
