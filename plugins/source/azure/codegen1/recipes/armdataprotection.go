// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"

func Armdataprotection() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdataprotection.NewBackupInstancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances",
		},
		{
			NewFunc: armdataprotection.NewBackupPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies",
		},
		{
			NewFunc: armdataprotection.NewBackupVaultOperationResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewBackupVaultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewExportJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewExportJobsOperationResultClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupJobs",
		},
		{
			NewFunc: armdataprotection.NewOperationResultClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewOperationStatusBackupVaultContextClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewOperationStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewOperationStatusResourceGroupContextClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewRecoveryPointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/recoveryPoints",
		},
		{
			NewFunc: armdataprotection.NewResourceGuardsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
		{
			NewFunc: armdataprotection.NewRestorableTimeRangesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdataprotection())
}