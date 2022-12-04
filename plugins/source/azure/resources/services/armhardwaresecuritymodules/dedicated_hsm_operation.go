// Code generated by codegen; DO NOT EDIT.

package armhardwaresecuritymodules

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DedicatedHsmOperation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armhardwaresecuritymodules_dedicated_hsm_operation",
		Resolver:  fetchDedicatedHsmOperation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
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
				Name:     "is_data_action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IsDataAction"),
			},
		},
	}
}

func fetchDedicatedHsmOperation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmhardwaresecuritymodulesDedicatedHsmOperation
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
