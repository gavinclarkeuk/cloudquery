// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"

func Armdatashare() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatashare.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewDataSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewShareSubscriptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewEmailRegistrationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewInvitationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewProviderShareSubscriptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewSharesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewSynchronizationSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewTriggersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewConsumerInvitationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewConsumerSourceDataSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
		{
			NewFunc: armdatashare.NewDataSetMappingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatashare())
}