// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"

func Armdatabricks() []Table {
	tables := []Table{
		{
      Name: "group_id_information",
      Struct: &armdatabricks.GroupIDInformation{},
      ResponseStruct: &armdatabricks.PrivateLinkResourcesClientListResponse{},
      Client: &armdatabricks.PrivateLinkResourcesClient{},
      ListFunc: (&armdatabricks.PrivateLinkResourcesClient{}).NewListPager,
			NewFunc: armdatabricks.NewPrivateLinkResourcesClient,
		},
		{
      Name: "operation",
      Struct: &armdatabricks.Operation{},
      ResponseStruct: &armdatabricks.OperationsClientListResponse{},
      Client: &armdatabricks.OperationsClient{},
      ListFunc: (&armdatabricks.OperationsClient{}).NewListPager,
			NewFunc: armdatabricks.NewOperationsClient,
		},
		{
      Name: "private_endpoint_connection",
      Struct: &armdatabricks.PrivateEndpointConnection{},
      ResponseStruct: &armdatabricks.PrivateEndpointConnectionsClientListResponse{},
      Client: &armdatabricks.PrivateEndpointConnectionsClient{},
      ListFunc: (&armdatabricks.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armdatabricks.NewPrivateEndpointConnectionsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdatabricks"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdatabricks()...)
}