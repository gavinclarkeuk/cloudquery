// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"

func Armazuredata() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armazuredata.NewSQLServerRegistrationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlServerRegistrations",
		},
		{
			NewFunc: armazuredata.NewSQLServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armazuredata())
}