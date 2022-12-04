// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"

func Armelastic() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armelastic.NewUpgradableVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewCreateAndAssociateIPFilterClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewDeploymentInfoClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewDetachTrafficFilterClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewExternalUserClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewVMCollectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewVMHostClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewAllTrafficFiltersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewAssociateTrafficFilterClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewCreateAndAssociatePLFilterClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewListAssociatedTrafficFiltersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewTagRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewTrafficFiltersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewVMIngestionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewDetachAndDeleteTrafficFilterClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewMonitorClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewMonitoredResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewMonitorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
		{
			NewFunc: armelastic.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armelastic())
}