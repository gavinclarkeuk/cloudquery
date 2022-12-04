// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

func Armmonitor() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmonitor.NewDataCollectionRuleAssociationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewDiagnosticSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewLogProfilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewMetricsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewActivityLogAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewVMInsightsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewAutoscaleSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewDataCollectionEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewDataCollectionRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewBaselinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewEventCategoriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewPrivateLinkScopesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewMetricDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewScheduledQueryRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewActivityLogsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewAlertRuleIncidentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewAlertRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewMetricAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewMetricNamespacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewTenantActivityLogsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewActionGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewDiagnosticSettingsCategoryClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewPredictiveMetricClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewPrivateLinkScopedResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewMetricAlertsStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
		{
			NewFunc: armmonitor.NewPrivateLinkScopeOperationStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmonitor())
}