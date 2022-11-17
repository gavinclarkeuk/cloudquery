package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func CDN() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armcdn.Profile),
			Resolver: &resource.FuncParams{
				Func: cdn.ProfilesClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armcdn.Endpoint),
					Resolver: &resource.FuncParams{
						Func: cdn.EndpointsClient.NewListByProfilePager,
					},
					Children: []*resource.Resource{
						{
							Struct: new(armcdn.CustomDomain),
							Resolver: &resource.FuncParams{
								Func: cdn.CustomDomainsClient.NewListByEndpointPager,
							},
						},
						{
							Struct: new(armcdn.Route),
							Resolver: &resource.FuncParams{
								Func: cdn.RoutesClient.NewListByEndpointPager,
							},
						},
					},
				},
				{
					Struct: new(armcdn.RuleSet),
					Resolver: &resource.FuncParams{
						Func: cdn.RuleSetsClient.NewListByProfilePager,
					},
					Children: []*resource.Resource{
						{
							Struct: new(armcdn.Rule),
							Resolver: &resource.FuncParams{
								Func: cdn.RulesClient.NewListByRuleSetPager,
							},
						},
					},
				},
				{
					Struct: new(armcdn.SecurityPolicy),
					Resolver: &resource.FuncParams{
						Func: cdn.SecurityPoliciesClient.NewListByProfilePager,
					},
				},
			},
		},
	}
}
