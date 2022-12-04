// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"

func Armiothub() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armiothub.Operation{},
      ResponseStruct: &armiothub.OperationsClientListResponse{},
      Client: &armiothub.OperationsClient{},
      ListFunc: (&armiothub.OperationsClient{}).NewListPager,
			NewFunc: armiothub.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armiothub"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armiothub()...)
}