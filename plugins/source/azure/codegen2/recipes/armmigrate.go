// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"

func Armmigrate() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armmigrate.Operation{},
      ResponseStruct: &armmigrate.OperationsClientListResponse{},
      Client: &armmigrate.OperationsClient{},
      ListFunc: (&armmigrate.OperationsClient{}).NewListPager,
			NewFunc: armmigrate.NewOperationsClient,
		},
		{
      Name: "project",
      Struct: &armmigrate.Project{},
      ResponseStruct: &armmigrate.ProjectsClientListResponse{},
      Client: &armmigrate.ProjectsClient{},
      ListFunc: (&armmigrate.ProjectsClient{}).NewListPager,
			NewFunc: armmigrate.NewProjectsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armmigrate"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmigrate()...)
}