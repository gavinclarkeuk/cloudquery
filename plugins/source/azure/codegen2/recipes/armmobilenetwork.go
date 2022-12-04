// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"

func Armmobilenetwork() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armmobilenetwork.Operation{},
      ResponseStruct: &armmobilenetwork.OperationsClientListResponse{},
      Client: &armmobilenetwork.OperationsClient{},
      ListFunc: (&armmobilenetwork.OperationsClient{}).NewListPager,
			NewFunc: armmobilenetwork.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armmobilenetwork"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmobilenetwork()...)
}