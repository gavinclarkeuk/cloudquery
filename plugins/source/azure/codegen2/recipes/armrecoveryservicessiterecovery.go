// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"

func Armrecoveryservicessiterecovery() []Table {
	tables := []Table{
		{
      Name: "alert",
      Struct: &armrecoveryservicessiterecovery.Alert{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationAlertSettingsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationAlertSettingsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationAlertSettingsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationAlertSettingsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings",
		},
		{
      Name: "replication_appliance",
      Struct: &armrecoveryservicessiterecovery.ReplicationAppliance{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationAppliancesClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationAppliancesClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationAppliancesClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationAppliancesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAppliances",
		},
		{
      Name: "event",
      Struct: &armrecoveryservicessiterecovery.Event{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationEventsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationEventsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationEventsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationEventsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents",
		},
		{
      Name: "fabric",
      Struct: &armrecoveryservicessiterecovery.Fabric{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationFabricsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationFabricsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationFabricsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationFabricsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics",
		},
		{
      Name: "job",
      Struct: &armrecoveryservicessiterecovery.Job{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationJobsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationJobsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationJobsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationJobsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs",
		},
		{
      Name: "migration_item",
      Struct: &armrecoveryservicessiterecovery.MigrationItem{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationMigrationItemsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationMigrationItemsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationMigrationItemsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationMigrationItems",
		},
		{
      Name: "network_mapping",
      Struct: &armrecoveryservicessiterecovery.NetworkMapping{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationNetworkMappingsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationNetworkMappingsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationNetworkMappingsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationNetworkMappingsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings",
		},
		{
      Name: "network",
      Struct: &armrecoveryservicessiterecovery.Network{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationNetworksClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationNetworksClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationNetworksClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationNetworksClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworks",
		},
		{
      Name: "policy",
      Struct: &armrecoveryservicessiterecovery.Policy{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationPoliciesClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationPoliciesClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationPoliciesClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies",
		},
		{
      Name: "replication_protected_item",
      Struct: &armrecoveryservicessiterecovery.ReplicationProtectedItem{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationProtectedItemsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationProtectedItemsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationProtectedItemsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectedItems",
		},
		{
      Name: "protection_container_mapping",
      Struct: &armrecoveryservicessiterecovery.ProtectionContainerMapping{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationProtectionContainerMappingsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationProtectionContainerMappingsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationProtectionContainerMappingsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationProtectionContainerMappingsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings",
		},
		{
      Name: "protection_container",
      Struct: &armrecoveryservicessiterecovery.ProtectionContainer{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationProtectionContainersClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationProtectionContainersClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationProtectionContainersClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationProtectionContainersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers",
		},
		{
      Name: "replication_protection_intent",
      Struct: &armrecoveryservicessiterecovery.ReplicationProtectionIntent{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationProtectionIntentsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationProtectionIntentsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionIntents",
		},
		{
      Name: "recovery_plan",
      Struct: &armrecoveryservicessiterecovery.RecoveryPlan{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationRecoveryPlansClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationRecoveryPlansClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans",
		},
		{
      Name: "recovery_services_provider",
      Struct: &armrecoveryservicessiterecovery.RecoveryServicesProvider{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationRecoveryServicesProvidersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders",
		},
		{
      Name: "storage_classification_mapping",
      Struct: &armrecoveryservicessiterecovery.StorageClassificationMapping{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationStorageClassificationMappingsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationStorageClassificationMappingsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationStorageClassificationMappingsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationStorageClassificationMappingsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings",
		},
		{
      Name: "storage_classification",
      Struct: &armrecoveryservicessiterecovery.StorageClassification{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationStorageClassificationsClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationStorageClassificationsClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationStorageClassificationsClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationStorageClassificationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications",
		},
		{
      Name: "vault_setting",
      Struct: &armrecoveryservicessiterecovery.VaultSetting{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationVaultSettingClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationVaultSettingClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationVaultSettingClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationVaultSettingClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings",
		},
		{
      Name: "v_center",
      Struct: &armrecoveryservicessiterecovery.VCenter{},
      ResponseStruct: &armrecoveryservicessiterecovery.ReplicationvCentersClientListResponse{},
      Client: &armrecoveryservicessiterecovery.ReplicationvCentersClient{},
      ListFunc: (&armrecoveryservicessiterecovery.ReplicationvCentersClient{}).NewListPager,
			NewFunc: armrecoveryservicessiterecovery.NewReplicationvCentersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationvCenters",
		},
	}

	for i := range tables {
		tables[i].Service = "armrecoveryservicessiterecovery"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armrecoveryservicessiterecovery()...)
}