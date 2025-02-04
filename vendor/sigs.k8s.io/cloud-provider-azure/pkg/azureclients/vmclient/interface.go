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

package vmclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"

	"sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

const (
	// APIVersion is the API version for VirtualMachine.
	APIVersion = "2020-12-01"
	// AzureStackCloudAPIVersion is the API version for Azure Stack
	AzureStackCloudAPIVersion = "2017-12-01"
	// AzureStackCloudName is the cloud name of Azure Stack
	AzureStackCloudName = "AZURESTACKCLOUD"
)

// Interface is the client interface for VirtualMachines.
// Don't forget to run the following command to generate the mock client:
// mockgen -source=$GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/vmclient/interface.go -package=mockvmclient Interface > $GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/vmclient/mockvmclient/interface.go
type Interface interface {
	// Get gets a VirtualMachine.
	Get(ctx context.Context, resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (compute.VirtualMachine, *retry.Error)

	// List gets a list of VirtualMachines in the resourceGroupName.
	List(ctx context.Context, resourceGroupName string) ([]compute.VirtualMachine, *retry.Error)

	// CreateOrUpdate creates or updates a VirtualMachine.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachine, source string) *retry.Error

	// Update updates a VirtualMachine.
	Update(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineUpdate, source string) *retry.Error

	// UpdateAsync updates a VirtualMachine asynchronously
	UpdateAsync(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineUpdate, source string) (*azure.Future, *retry.Error)

	// WaitForUpdateResult waits for the response of the update request
	WaitForUpdateResult(ctx context.Context, future *azure.Future, resourceGroupName, source string) *retry.Error

	// Delete deletes a VirtualMachine.
	Delete(ctx context.Context, resourceGroupName string, VMName string) *retry.Error
}
