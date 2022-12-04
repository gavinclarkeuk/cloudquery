// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"

func Armwindowsesu() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armwindowsesu.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu",
		},
		{
			NewFunc: armwindowsesu.NewMultipleActivationKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armwindowsesu())
}