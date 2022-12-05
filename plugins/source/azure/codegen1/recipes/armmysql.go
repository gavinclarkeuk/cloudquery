// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"

func Armmysql() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmysql.NewAdvisorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewCheckNameAvailabilityClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewLocationBasedPerformanceTierClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/performanceTiers",
		},
		{
			NewFunc: armmysql.NewLocationBasedRecommendedActionSessionsOperationStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewLocationBasedRecommendedActionSessionsResultClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/recommendedActionSessionsOperationResults/{operationId}",
		},
		{
			NewFunc: armmysql.NewLogFilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewQueryTextsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewRecommendedActionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewRecoverableServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewReplicasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewServerAdministratorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators",
		},
		{
			NewFunc: armmysql.NewServerBasedPerformanceTierClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/performanceTiers",
		},
		{
			NewFunc: armmysql.NewServerKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/keys",
		},
		{
			NewFunc: armmysql.NewServerParametersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewServerSecurityAlertPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/servers",
		},
		{
			NewFunc: armmysql.NewTopQueryStatisticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewVirtualNetworkRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
		{
			NewFunc: armmysql.NewWaitStatisticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmysql())
}