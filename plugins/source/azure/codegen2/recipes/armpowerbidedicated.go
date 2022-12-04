// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"

func Armpowerbidedicated() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armpowerbidedicated.Operation{},
      ResponseStruct: &armpowerbidedicated.OperationsClientListResponse{},
      Client: &armpowerbidedicated.OperationsClient{},
      ListFunc: (&armpowerbidedicated.OperationsClient{}).NewListPager,
			NewFunc: armpowerbidedicated.NewOperationsClient,
		},
		{
      Name: "dedicated_capacity",
      Struct: &armpowerbidedicated.DedicatedCapacity{},
      ResponseStruct: &armpowerbidedicated.CapacitiesClientListResponse{},
      Client: &armpowerbidedicated.CapacitiesClient{},
      ListFunc: (&armpowerbidedicated.CapacitiesClient{}).NewListPager,
			NewFunc: armpowerbidedicated.NewCapacitiesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armpowerbidedicated"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armpowerbidedicated()...)
}