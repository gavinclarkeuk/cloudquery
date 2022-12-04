// Code generated by codegen; DO NOT EDIT.

package armlogic

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IntegrationServiceEnvironmentSkuDefinition() *schema.Table {
	return &schema.Table{
		Name:      "azure_armlogic_integration_service_environment_sku_definition",
		Resolver:  fetchIntegrationServiceEnvironmentSkuDefinition,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "capacity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capacity"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SKU"),
			},
		},
	}
}

func fetchIntegrationServiceEnvironmentSkuDefinition(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmlogicIntegrationServiceEnvironmentSkuDefinition
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
