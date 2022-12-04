// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"

func Armnotificationhubs() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armnotificationhubs.NewNamespacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs",
		},
		{
			NewFunc: armnotificationhubs.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs",
		},
		{
			NewFunc: armnotificationhubs.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armnotificationhubs())
}