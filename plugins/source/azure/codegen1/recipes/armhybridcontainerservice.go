// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"

func Armhybridcontainerservice() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhybridcontainerservice.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
		{
			NewFunc: armhybridcontainerservice.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
		{
			NewFunc: armhybridcontainerservice.NewProvisionedClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
		{
			NewFunc: armhybridcontainerservice.NewStorageSpacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
		{
			NewFunc: armhybridcontainerservice.NewAgentPoolClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
		{
			NewFunc: armhybridcontainerservice.NewHybridIdentityMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
		{
			NewFunc: armhybridcontainerservice.NewVirtualNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhybridcontainerservice())
}