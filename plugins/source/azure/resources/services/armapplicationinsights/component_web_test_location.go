// Code generated by codegen; DO NOT EDIT.

package armapplicationinsights

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ComponentWebTestLocation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armapplicationinsights_component_web_test_location",
		Resolver:  fetchComponentWebTestLocation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "tag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tag"),
			},
		},
	}
}

func fetchComponentWebTestLocation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmapplicationinsightsComponentWebTestLocation
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
