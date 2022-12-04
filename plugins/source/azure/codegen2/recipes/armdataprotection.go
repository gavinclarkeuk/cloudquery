// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"

func Armdataprotection() []Table {
	tables := []Table{
		{
      Name: "azure_backup_recovery_point_resource",
      Struct: &armdataprotection.AzureBackupRecoveryPointResource{},
      ResponseStruct: &armdataprotection.RecoveryPointsClientListResponse{},
      Client: &armdataprotection.RecoveryPointsClient{},
      ListFunc: (&armdataprotection.RecoveryPointsClient{}).NewListPager,
			NewFunc: armdataprotection.NewRecoveryPointsClient,
		},
		{
      Name: "base_backup_policy_resource",
      Struct: &armdataprotection.BaseBackupPolicyResource{},
      ResponseStruct: &armdataprotection.BackupPoliciesClientListResponse{},
      Client: &armdataprotection.BackupPoliciesClient{},
      ListFunc: (&armdataprotection.BackupPoliciesClient{}).NewListPager,
			NewFunc: armdataprotection.NewBackupPoliciesClient,
		},
		{
      Name: "azure_backup_job_resource",
      Struct: &armdataprotection.AzureBackupJobResource{},
      ResponseStruct: &armdataprotection.JobsClientListResponse{},
      Client: &armdataprotection.JobsClient{},
      ListFunc: (&armdataprotection.JobsClient{}).NewListPager,
			NewFunc: armdataprotection.NewJobsClient,
		},
		{
      Name: "client_discovery_value_for_single_api",
      Struct: &armdataprotection.ClientDiscoveryValueForSingleAPI{},
      ResponseStruct: &armdataprotection.OperationsClientListResponse{},
      Client: &armdataprotection.OperationsClient{},
      ListFunc: (&armdataprotection.OperationsClient{}).NewListPager,
			NewFunc: armdataprotection.NewOperationsClient,
		},
		{
      Name: "backup_instance_resource",
      Struct: &armdataprotection.BackupInstanceResource{},
      ResponseStruct: &armdataprotection.BackupInstancesClientListResponse{},
      Client: &armdataprotection.BackupInstancesClient{},
      ListFunc: (&armdataprotection.BackupInstancesClient{}).NewListPager,
			NewFunc: armdataprotection.NewBackupInstancesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdataprotection"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdataprotection()...)
}