// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges"

func Armchanges() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armchanges.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armchanges())
}