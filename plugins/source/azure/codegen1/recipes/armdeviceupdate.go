// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"

func Armdeviceupdate() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdeviceupdate.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
		{
			NewFunc: armdeviceupdate.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
		{
			NewFunc: armdeviceupdate.NewInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
		{
			NewFunc: armdeviceupdate.NewPrivateEndpointConnectionProxiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
		{
			NewFunc: armdeviceupdate.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
		{
			NewFunc: armdeviceupdate.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
		{
			NewFunc: armdeviceupdate.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdeviceupdate())
}