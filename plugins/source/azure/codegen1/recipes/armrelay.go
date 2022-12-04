// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"

func Armrelay() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armrelay.NewHybridConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
		},
		{
			NewFunc: armrelay.NewNamespacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
		},
		{
			NewFunc: armrelay.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
		},
		{
			NewFunc: armrelay.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
		},
		{
			NewFunc: armrelay.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
		},
		{
			NewFunc: armrelay.NewWCFRelaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armrelay())
}