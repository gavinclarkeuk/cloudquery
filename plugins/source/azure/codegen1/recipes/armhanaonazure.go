// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"

func Armhanaonazure() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhanaonazure.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure",
		},
		{
			NewFunc: armhanaonazure.NewProviderInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure",
		},
		{
			NewFunc: armhanaonazure.NewSapMonitorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhanaonazure())
}