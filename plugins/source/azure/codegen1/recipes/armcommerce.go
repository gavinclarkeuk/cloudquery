// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commerce/armcommerce"

func Armcommerce() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcommerce.NewUsageAggregatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commerce/armcommerce",
		},
		{
			NewFunc: armcommerce.NewRateCardClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commerce/armcommerce",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcommerce())
}