// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"

func Armmaintenance() []Table {
	tables := []Table{
		{
      Name: "apply_update",
      Struct: &armmaintenance.ApplyUpdate{},
      ResponseStruct: &armmaintenance.ApplyUpdateForResourceGroupClientListResponse{},
      Client: &armmaintenance.ApplyUpdateForResourceGroupClient{},
      ListFunc: (&armmaintenance.ApplyUpdateForResourceGroupClient{}).NewListPager,
			NewFunc: armmaintenance.NewApplyUpdateForResourceGroupClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maintenance/applyUpdates",
		},
		{
      Name: "apply_update",
      Struct: &armmaintenance.ApplyUpdate{},
      ResponseStruct: &armmaintenance.ApplyUpdatesClientListResponse{},
      Client: &armmaintenance.ApplyUpdatesClient{},
      ListFunc: (&armmaintenance.ApplyUpdatesClient{}).NewListPager,
			NewFunc: armmaintenance.NewApplyUpdatesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/applyUpdates",
		},
		{
      Name: "configuration_assignment",
      Struct: &armmaintenance.ConfigurationAssignment{},
      ResponseStruct: &armmaintenance.ConfigurationAssignmentsClientListResponse{},
      Client: &armmaintenance.ConfigurationAssignmentsClient{},
      ListFunc: (&armmaintenance.ConfigurationAssignmentsClient{}).NewListPager,
			NewFunc: armmaintenance.NewConfigurationAssignmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments",
		},
		{
      Name: "configuration",
      Struct: &armmaintenance.Configuration{},
      ResponseStruct: &armmaintenance.ConfigurationsClientListResponse{},
      Client: &armmaintenance.ConfigurationsClient{},
      ListFunc: (&armmaintenance.ConfigurationsClient{}).NewListPager,
			NewFunc: armmaintenance.NewConfigurationsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/maintenanceConfigurations",
		},
		{
      Name: "configuration",
      Struct: &armmaintenance.Configuration{},
      ResponseStruct: &armmaintenance.ConfigurationsForResourceGroupClientListResponse{},
      Client: &armmaintenance.ConfigurationsForResourceGroupClient{},
      ListFunc: (&armmaintenance.ConfigurationsForResourceGroupClient{}).NewListPager,
			NewFunc: armmaintenance.NewConfigurationsForResourceGroupClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maintenance/maintenanceConfigurations",
		},
		{
      Name: "configuration",
      Struct: &armmaintenance.Configuration{},
      ResponseStruct: &armmaintenance.PublicMaintenanceConfigurationsClientListResponse{},
      Client: &armmaintenance.PublicMaintenanceConfigurationsClient{},
      ListFunc: (&armmaintenance.PublicMaintenanceConfigurationsClient{}).NewListPager,
			NewFunc: armmaintenance.NewPublicMaintenanceConfigurationsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
		},
		{
      Name: "update",
      Struct: &armmaintenance.Update{},
      ResponseStruct: &armmaintenance.UpdatesClientListResponse{},
      Client: &armmaintenance.UpdatesClient{},
      ListFunc: (&armmaintenance.UpdatesClient{}).NewListPager,
			NewFunc: armmaintenance.NewUpdatesClient,
			URL: "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates",
		},
	}

	for i := range tables {
		tables[i].Service = "armmaintenance"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmaintenance()...)
}