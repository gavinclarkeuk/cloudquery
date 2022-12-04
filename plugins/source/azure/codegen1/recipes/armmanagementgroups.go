// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"

func Armmanagementgroups() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmanagementgroups.NewAPIClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups",
		},
		{
			NewFunc: armmanagementgroups.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups",
		},
		{
			NewFunc: armmanagementgroups.NewManagementGroupSubscriptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups",
		},
		{
			NewFunc: armmanagementgroups.NewEntitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups",
		},
		{
			NewFunc: armmanagementgroups.NewHierarchySettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups",
		},
		{
			NewFunc: armmanagementgroups.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmanagementgroups())
}