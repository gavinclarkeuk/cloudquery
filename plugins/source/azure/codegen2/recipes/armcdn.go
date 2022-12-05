// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func Armcdn() []Table {
	tables := []Table{
		{
      Name: "edge_node",
      Struct: &armcdn.EdgeNode{},
      ResponseStruct: &armcdn.EdgeNodesClientListResponse{},
      Client: &armcdn.EdgeNodesClient{},
      ListFunc: (&armcdn.EdgeNodesClient{}).NewListPager,
			NewFunc: armcdn.NewEdgeNodesClient,
			URL: "/providers/Microsoft.Cdn/edgenodes",
		},
		{
      Name: "managed_rule_set_definition",
      Struct: &armcdn.ManagedRuleSetDefinition{},
      ResponseStruct: &armcdn.ManagedRuleSetsClientListResponse{},
      Client: &armcdn.ManagedRuleSetsClient{},
      ListFunc: (&armcdn.ManagedRuleSetsClient{}).NewListPager,
			NewFunc: armcdn.NewManagedRuleSetsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets",
		},
		{
      Name: "web_application_firewall_policy",
      Struct: &armcdn.WebApplicationFirewallPolicy{},
      ResponseStruct: &armcdn.PoliciesClientListResponse{},
      Client: &armcdn.PoliciesClient{},
      ListFunc: (&armcdn.PoliciesClient{}).NewListPager,
			NewFunc: armcdn.NewPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/cdnWebApplicationFirewallPolicies",
		},
		{
      Name: "profile",
      Struct: &armcdn.Profile{},
      ResponseStruct: &armcdn.ProfilesClientListResponse{},
      Client: &armcdn.ProfilesClient{},
      ListFunc: (&armcdn.ProfilesClient{}).NewListPager,
			NewFunc: armcdn.NewProfilesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles",
		},
		{
      Name: "resource_usage",
      Struct: &armcdn.ResourceUsage{},
      ResponseStruct: &armcdn.ResourceUsageClientListResponse{},
      Client: &armcdn.ResourceUsageClient{},
      ListFunc: (&armcdn.ResourceUsageClient{}).NewListPager,
			NewFunc: armcdn.NewResourceUsageClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/checkResourceUsage",
		},
	}

	for i := range tables {
		tables[i].Service = "armcdn"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armcdn()...)
}