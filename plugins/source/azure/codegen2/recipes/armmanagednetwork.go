// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"

func Armmanagednetwork() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armmanagednetwork.Operation{},
      ResponseStruct: &armmanagednetwork.OperationsClientListResponse{},
      Client: &armmanagednetwork.OperationsClient{},
      ListFunc: (&armmanagednetwork.OperationsClient{}).NewListPager,
			NewFunc: armmanagednetwork.NewOperationsClient,
		},
		{
      Name: "scope_assignment",
      Struct: &armmanagednetwork.ScopeAssignment{},
      ResponseStruct: &armmanagednetwork.ScopeAssignmentsClientListResponse{},
      Client: &armmanagednetwork.ScopeAssignmentsClient{},
      ListFunc: (&armmanagednetwork.ScopeAssignmentsClient{}).NewListPager,
			NewFunc: armmanagednetwork.NewScopeAssignmentsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armmanagednetwork"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmanagednetwork()...)
}