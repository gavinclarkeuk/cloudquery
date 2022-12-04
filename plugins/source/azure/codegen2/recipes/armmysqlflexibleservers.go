// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"

func Armmysqlflexibleservers() []Table {
	tables := []Table{
		{
      Name: "server",
      Struct: &armmysqlflexibleservers.Server{},
      ResponseStruct: &armmysqlflexibleservers.ServersClientListResponse{},
      Client: &armmysqlflexibleservers.ServersClient{},
      ListFunc: (&armmysqlflexibleservers.ServersClient{}).NewListPager,
			NewFunc: armmysqlflexibleservers.NewServersClient,
		},
		{
      Name: "operation",
      Struct: &armmysqlflexibleservers.Operation{},
      ResponseStruct: &armmysqlflexibleservers.OperationsClientListResponse{},
      Client: &armmysqlflexibleservers.OperationsClient{},
      ListFunc: (&armmysqlflexibleservers.OperationsClient{}).NewListPager,
			NewFunc: armmysqlflexibleservers.NewOperationsClient,
		},
		{
      Name: "capability_properties",
      Struct: &armmysqlflexibleservers.CapabilityProperties{},
      ResponseStruct: &armmysqlflexibleservers.LocationBasedCapabilitiesClientListResponse{},
      Client: &armmysqlflexibleservers.LocationBasedCapabilitiesClient{},
      ListFunc: (&armmysqlflexibleservers.LocationBasedCapabilitiesClient{}).NewListPager,
			NewFunc: armmysqlflexibleservers.NewLocationBasedCapabilitiesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armmysqlflexibleservers"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmysqlflexibleservers()...)
}