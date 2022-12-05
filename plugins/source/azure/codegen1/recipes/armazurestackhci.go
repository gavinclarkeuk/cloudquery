// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"

func Armazurestackhci() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armazurestackhci.NewArcSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci",
			URL: "",
		},
		{
			NewFunc: armazurestackhci.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci",
			URL: "",
		},
		{
			NewFunc: armazurestackhci.NewExtensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armazurestackhci())
}