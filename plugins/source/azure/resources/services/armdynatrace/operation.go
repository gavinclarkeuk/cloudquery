// Code generated by codegen; DO NOT EDIT.

package armdynatrace

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Operation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armdynatrace_operation",
		Resolver:  fetchOperation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "display",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Display"),
			},
			{
				Name:     "action_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActionType"),
			},
			{
				Name:     "is_data_action",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDataAction"),
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
	svc := meta.(*client.Client).Services().ArmdynatraceOperation
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
