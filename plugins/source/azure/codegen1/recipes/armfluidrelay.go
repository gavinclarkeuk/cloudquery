// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay"

func Armfluidrelay() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armfluidrelay.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay",
		},
		{
			NewFunc: armfluidrelay.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay",
		},
		{
			NewFunc: armfluidrelay.NewContainersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armfluidrelay())
}