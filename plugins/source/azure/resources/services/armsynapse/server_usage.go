// Code generated by codegen; DO NOT EDIT.

package armsynapse

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServerUsage() *schema.Table {
	return &schema.Table{
		Name:      "azure_armsynapse_server_usage",
		Resolver:  fetchServerUsage,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "current_value",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("CurrentValue"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "limit",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Limit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "next_reset_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NextResetTime"),
			},
			{
				Name:     "resource_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceName"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
		},
	}
}

func fetchServerUsage(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armsynapse.NewWorkspaceManagedSQLServerUsagesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
