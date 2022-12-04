// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"

func Armmysqlflexibleservers() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmysqlflexibleservers.NewCheckVirtualNetworkSubnetUsageClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewReplicasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewCheckNameAvailabilityClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewBackupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewGetPrivateDNSZoneSuffixClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
		{
			NewFunc: armmysqlflexibleservers.NewLocationBasedCapabilitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmysqlflexibleservers())
}