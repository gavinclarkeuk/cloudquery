// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"

func Armaad() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armaad.NewPrivateLinkForAzureAdClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad",
		},
		{
			NewFunc: armaad.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad",
		},
		{
			NewFunc: armaad.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armaad())
}