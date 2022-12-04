// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"

func Armdataboxedge() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdataboxedge.NewJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewAvailableSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewBandwidthSchedulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewSharesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewDeviceCapacityInfoClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewNodesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewAddonsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewDeviceCapacityCheckClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewOperationsStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewTriggersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewUsersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewDiagnosticSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewRolesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewStorageAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewContainersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewDevicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewMonitoringConfigClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewOrdersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewStorageAccountCredentialsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
		{
			NewFunc: armdataboxedge.NewSupportPackagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdataboxedge())
}