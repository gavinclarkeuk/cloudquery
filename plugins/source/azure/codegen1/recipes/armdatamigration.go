// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"

func Armdatamigration() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatamigration.NewFilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewServiceTasksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewResourceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewTasksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewProjectsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
		{
			NewFunc: armdatamigration.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatamigration())
}