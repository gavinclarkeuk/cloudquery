// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"

func Armnetapp() []Table {
	tables := []Table{
		{
      Name: "snapshot",
      Struct: &armnetapp.Snapshot{},
      ResponseStruct: &armnetapp.SnapshotsClientListResponse{},
      Client: &armnetapp.SnapshotsClient{},
      ListFunc: (&armnetapp.SnapshotsClient{}).NewListPager,
			NewFunc: armnetapp.NewSnapshotsClient,
		},
		{
      Name: "capacity_pool",
      Struct: &armnetapp.CapacityPool{},
      ResponseStruct: &armnetapp.PoolsClientListResponse{},
      Client: &armnetapp.PoolsClient{},
      ListFunc: (&armnetapp.PoolsClient{}).NewListPager,
			NewFunc: armnetapp.NewPoolsClient,
		},
		{
      Name: "subscription_quota_item",
      Struct: &armnetapp.SubscriptionQuotaItem{},
      ResponseStruct: &armnetapp.ResourceQuotaLimitsClientListResponse{},
      Client: &armnetapp.ResourceQuotaLimitsClient{},
      ListFunc: (&armnetapp.ResourceQuotaLimitsClient{}).NewListPager,
			NewFunc: armnetapp.NewResourceQuotaLimitsClient,
		},
		{
      Name: "backup",
      Struct: &armnetapp.Backup{},
      ResponseStruct: &armnetapp.AccountBackupsClientListResponse{},
      Client: &armnetapp.AccountBackupsClient{},
      ListFunc: (&armnetapp.AccountBackupsClient{}).NewListPager,
			NewFunc: armnetapp.NewAccountBackupsClient,
		},
		{
      Name: "operation",
      Struct: &armnetapp.Operation{},
      ResponseStruct: &armnetapp.OperationsClientListResponse{},
      Client: &armnetapp.OperationsClient{},
      ListFunc: (&armnetapp.OperationsClient{}).NewListPager,
			NewFunc: armnetapp.NewOperationsClient,
		},
		{
      Name: "backup",
      Struct: &armnetapp.Backup{},
      ResponseStruct: &armnetapp.BackupsClientListResponse{},
      Client: &armnetapp.BackupsClient{},
      ListFunc: (&armnetapp.BackupsClient{}).NewListPager,
			NewFunc: armnetapp.NewBackupsClient,
		},
		{
      Name: "vault",
      Struct: &armnetapp.Vault{},
      ResponseStruct: &armnetapp.VaultsClientListResponse{},
      Client: &armnetapp.VaultsClient{},
      ListFunc: (&armnetapp.VaultsClient{}).NewListPager,
			NewFunc: armnetapp.NewVaultsClient,
		},
		{
      Name: "volume",
      Struct: &armnetapp.Volume{},
      ResponseStruct: &armnetapp.VolumesClientListResponse{},
      Client: &armnetapp.VolumesClient{},
      ListFunc: (&armnetapp.VolumesClient{}).NewListPager,
			NewFunc: armnetapp.NewVolumesClient,
		},
		{
      Name: "account",
      Struct: &armnetapp.Account{},
      ResponseStruct: &armnetapp.AccountsClientListResponse{},
      Client: &armnetapp.AccountsClient{},
      ListFunc: (&armnetapp.AccountsClient{}).NewListPager,
			NewFunc: armnetapp.NewAccountsClient,
		},
		{
      Name: "backup_policy",
      Struct: &armnetapp.BackupPolicy{},
      ResponseStruct: &armnetapp.BackupPoliciesClientListResponse{},
      Client: &armnetapp.BackupPoliciesClient{},
      ListFunc: (&armnetapp.BackupPoliciesClient{}).NewListPager,
			NewFunc: armnetapp.NewBackupPoliciesClient,
		},
		{
      Name: "snapshot_policy",
      Struct: &armnetapp.SnapshotPolicy{},
      ResponseStruct: &armnetapp.SnapshotPoliciesClientListResponse{},
      Client: &armnetapp.SnapshotPoliciesClient{},
      ListFunc: (&armnetapp.SnapshotPoliciesClient{}).NewListPager,
			NewFunc: armnetapp.NewSnapshotPoliciesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armnetapp"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armnetapp()...)
}