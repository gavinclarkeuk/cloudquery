// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"

func Armresourcemover() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armresourcemover.NewMoveCollectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover",
			URL: "",
		},
		{
			NewFunc: armresourcemover.NewMoveResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources",
		},
		{
			NewFunc: armresourcemover.NewOperationsDiscoveryClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover",
			URL: "",
		},
		{
			NewFunc: armresourcemover.NewUnresolvedDependenciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armresourcemover())
}