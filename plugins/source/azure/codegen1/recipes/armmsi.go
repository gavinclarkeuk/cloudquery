// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"

func Armmsi() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmsi.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
		},
		{
			NewFunc: armmsi.NewFederatedIdentityCredentialsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
		},
		{
			NewFunc: armmsi.NewSystemAssignedIdentitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
		},
		{
			NewFunc: armmsi.NewUserAssignedIdentitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmsi())
}