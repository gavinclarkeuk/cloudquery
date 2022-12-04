// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"

func Armdeviceprovisioningservices() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdeviceprovisioningservices.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices",
		},
		{
			NewFunc: armdeviceprovisioningservices.NewDpsCertificateClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices",
		},
		{
			NewFunc: armdeviceprovisioningservices.NewIotDpsResourceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdeviceprovisioningservices())
}