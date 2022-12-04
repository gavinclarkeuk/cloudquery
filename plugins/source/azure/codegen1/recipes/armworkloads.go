// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func Armworkloads() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armworkloads.NewSAPApplicationServerInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewWordpressInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewPhpWorkloadsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewProviderInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewSAPCentralInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewMonitorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewSAPDatabaseInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
		{
			NewFunc: armworkloads.NewSAPVirtualInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armworkloads())
}