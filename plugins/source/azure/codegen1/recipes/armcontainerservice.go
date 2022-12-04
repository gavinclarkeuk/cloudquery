// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"

func Armcontainerservice() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcontainerservice.NewManagedClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewResolvePrivateLinkServiceIDClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewAgentPoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewMaintenanceConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
		{
			NewFunc: armcontainerservice.NewSnapshotsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcontainerservice())
}