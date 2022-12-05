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
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules",
		},
		{
      Name: "linked_server_with_properties",
      Struct: &armredis.LinkedServerWithProperties{},
      ResponseStruct: &armredis.LinkedServerClientListResponse{},
      Client: &armredis.LinkedServerClient{},
      ListFunc: (&armredis.LinkedServerClient{}).NewListPager,
			NewFunc: armredis.NewLinkedServerClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers",
		},
		{
      Name: "private_endpoint_connection",
      Struct: &armredis.PrivateEndpointConnection{},
      ResponseStruct: &armredis.PrivateEndpointConnectionsClientListResponse{},
      Client: &armredis.PrivateEndpointConnectionsClient{},
      ListFunc: (&armredis.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armredis.NewPrivateEndpointConnectionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/privateEndpointConnections",
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