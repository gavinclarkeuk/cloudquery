// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"

func Armdelegatednetwork() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armdelegatednetwork.Operation{},
      ResponseStruct: &armdelegatednetwork.OperationsClientListResponse{},
      Client: &armdelegatednetwork.OperationsClient{},
      ListFunc: (&armdelegatednetwork.OperationsClient{}).NewListPager,
			NewFunc: armdelegatednetwork.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdelegatednetwork"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdelegatednetwork()...)
}