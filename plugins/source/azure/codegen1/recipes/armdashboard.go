// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"

func Armdashboard() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdashboard.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard",
		},
		{
			NewFunc: armdashboard.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard",
		},
		{
			NewFunc: armdashboard.NewGrafanaClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard",
		},
		{
			NewFunc: armdashboard.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdashboard())
}