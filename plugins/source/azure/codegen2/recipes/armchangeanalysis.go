// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"

func Armchangeanalysis() []Table {
	tables := []Table{
		{
      Name: "change",
      Struct: &armchangeanalysis.Change{},
      ResponseStruct: &armchangeanalysis.ResourceChangesClientListResponse{},
      Client: &armchangeanalysis.ResourceChangesClient{},
      ListFunc: (&armchangeanalysis.ResourceChangesClient{}).NewListPager,
			NewFunc: armchangeanalysis.NewResourceChangesClient,
		},
		{
      Name: "resource_provider_operation_definition",
      Struct: &armchangeanalysis.ResourceProviderOperationDefinition{},
      ResponseStruct: &armchangeanalysis.OperationsClientListResponse{},
      Client: &armchangeanalysis.OperationsClient{},
      ListFunc: (&armchangeanalysis.OperationsClient{}).NewListPager,
			NewFunc: armchangeanalysis.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armchangeanalysis"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armchangeanalysis()...)
}