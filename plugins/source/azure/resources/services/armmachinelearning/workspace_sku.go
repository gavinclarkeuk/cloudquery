// Code generated by codegen; DO NOT EDIT.

package armmachinelearning

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkspaceSku() *schema.Table {
	return &schema.Table{
		Name:      "azure_armmachinelearning_workspace_sku",
		Resolver:  fetchWorkspaceSku,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "restrictions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Restrictions"),
			},
			{
				Name:     "capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capabilities"),
			},
			{
				Name:     "location_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocationInfo"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
		},
	}
}

func fetchWorkspaceSku(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmmachinelearningWorkspaceSku
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
