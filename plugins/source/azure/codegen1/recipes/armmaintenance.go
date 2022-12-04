// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"

func Armmaintenance() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmaintenance.NewUpdatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewConfigurationsForResourceGroupClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewApplyUpdateForResourceGroupClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewPublicMaintenanceConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewApplyUpdatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
		{
			NewFunc: armmaintenance.NewConfigurationAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmaintenance())
}