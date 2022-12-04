// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"

func Armdatabox() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatabox.NewManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
		},
		{
			NewFunc: armdatabox.NewJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
		},
		{
			NewFunc: armdatabox.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
		},
		{
			NewFunc: armdatabox.NewServiceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatabox())
}