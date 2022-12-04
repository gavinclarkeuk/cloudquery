// Code generated by codegen; DO NOT EDIT.

package armdatabox

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Operation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armdatabox_operation",
		Resolver:  fetchOperation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "is_data_action",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDataAction"),
			},
			{
				Name:     "display",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Display"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Origin"),
			},
		},
	}
}

func fetchOperation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armdatabox.NewOperationsClient(cl.Creds, cl.Options)
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
