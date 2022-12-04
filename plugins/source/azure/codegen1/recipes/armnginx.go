// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx"

func Armnginx() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armnginx.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx",
		},
		{
			NewFunc: armnginx.NewDeploymentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx",
		},
		{
			NewFunc: armnginx.NewCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx",
		},
		{
			NewFunc: armnginx.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armnginx())
}