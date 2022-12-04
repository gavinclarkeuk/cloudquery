// Code generated by codegen; DO NOT EDIT.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Usage() *schema.Table {
	return &schema.Table{
		Name:      "azure_armnetwork_usage",
		Resolver:  fetchUsage,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "current_value",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CurrentValue"),
			},
			{
				Name:     "limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Limit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
		},
	}
}

func fetchUsage(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewUsagesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
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
