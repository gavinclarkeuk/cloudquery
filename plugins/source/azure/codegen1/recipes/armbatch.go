// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"

func Armbatch() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armbatch.NewAccountClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts",
		},
		{
			NewFunc: armbatch.NewApplicationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications",
		},
		{
			NewFunc: armbatch.NewApplicationPackageClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions",
		},
		{
			NewFunc: armbatch.NewCertificateClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "",
		},
		{
			NewFunc: armbatch.NewLocationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "",
		},
		{
			NewFunc: armbatch.NewPoolClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "",
		},
		{
			NewFunc: armbatch.NewPrivateEndpointConnectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "",
		},
		{
			NewFunc: armbatch.NewPrivateLinkResourceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armbatch())
}