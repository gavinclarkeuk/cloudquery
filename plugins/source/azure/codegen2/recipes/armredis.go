// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"

func Armredis() []Table {
	tables := []Table{
		{
      Name: "firewall_rule",
      Struct: &armredis.FirewallRule{},
      ResponseStruct: &armredis.FirewallRulesClientListResponse{},
      Client: &armredis.FirewallRulesClient{},
      ListFunc: (&armredis.FirewallRulesClient{}).NewListPager,
			NewFunc: armredis.NewFirewallRulesClient,
		},
		{
      Name: "private_endpoint_connection",
      Struct: &armredis.PrivateEndpointConnection{},
      ResponseStruct: &armredis.PrivateEndpointConnectionsClientListResponse{},
      Client: &armredis.PrivateEndpointConnectionsClient{},
      ListFunc: (&armredis.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armredis.NewPrivateEndpointConnectionsClient,
		},
		{
      Name: "operation",
      Struct: &armredis.Operation{},
      ResponseStruct: &armredis.OperationsClientListResponse{},
      Client: &armredis.OperationsClient{},
      ListFunc: (&armredis.OperationsClient{}).NewListPager,
			NewFunc: armredis.NewOperationsClient,
		},
		{
      Name: "linked_server_with_properties",
      Struct: &armredis.LinkedServerWithProperties{},
      ResponseStruct: &armredis.LinkedServerClientListResponse{},
      Client: &armredis.LinkedServerClient{},
      ListFunc: (&armredis.LinkedServerClient{}).NewListPager,
			NewFunc: armredis.NewLinkedServerClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armredis"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armredis()...)
}