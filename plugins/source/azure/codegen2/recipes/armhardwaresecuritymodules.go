// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"

func Armhardwaresecuritymodules() []Table {
	tables := []Table{
		{
      Name: "dedicated_hsm_operation",
      Struct: &armhardwaresecuritymodules.DedicatedHsmOperation{},
      ResponseStruct: &armhardwaresecuritymodules.OperationsClientListResponse{},
      Client: &armhardwaresecuritymodules.OperationsClient{},
      ListFunc: (&armhardwaresecuritymodules.OperationsClient{}).NewListPager,
			NewFunc: armhardwaresecuritymodules.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armhardwaresecuritymodules"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armhardwaresecuritymodules()...)
}