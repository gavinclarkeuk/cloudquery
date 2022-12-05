// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"

func Armconnectedvmware() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armconnectedvmware.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/clusters",
		},
		{
			NewFunc: armconnectedvmware.NewDatastoresClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/datastores",
		},
		{
			NewFunc: armconnectedvmware.NewGuestAgentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "",
		},
		{
			NewFunc: armconnectedvmware.NewHostsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts",
		},
		{
			NewFunc: armconnectedvmware.NewHybridIdentityMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "",
		},
		{
			NewFunc: armconnectedvmware.NewInventoryItemsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "",
		},
		{
			NewFunc: armconnectedvmware.NewMachineExtensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{name}/extensions",
		},
		{
			NewFunc: armconnectedvmware.NewResourcePoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/resourcePools",
		},
		{
			NewFunc: armconnectedvmware.NewVCentersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/vcenters",
		},
		{
			NewFunc: armconnectedvmware.NewVirtualMachineTemplatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates",
		},
		{
			NewFunc: armconnectedvmware.NewVirtualMachinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
		},
		{
			NewFunc: armconnectedvmware.NewVirtualNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armconnectedvmware())
}