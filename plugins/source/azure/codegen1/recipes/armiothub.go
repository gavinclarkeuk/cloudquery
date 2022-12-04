// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"

func Armiothub() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armiothub.NewCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
		{
			NewFunc: armiothub.NewResourceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
		{
			NewFunc: armiothub.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
		{
			NewFunc: armiothub.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
		{
			NewFunc: armiothub.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
		{
			NewFunc: armiothub.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
		{
			NewFunc: armiothub.NewResourceProviderCommonClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armiothub())
}