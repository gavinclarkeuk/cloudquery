// Code generated by codegen; DO NOT EDIT.

package armcompute

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RunCommandDocumentBase() *schema.Table {
	return &schema.Table{
		Name:      "azure_armcompute_run_command_document_base",
		Resolver:  fetchRunCommandDocumentBase,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Label"),
			},
			{
				Name:     "os_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OSType"),
			},
			{
				Name:     "$schema",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Schema"),
			},
		},
	}
}

func fetchRunCommandDocumentBase(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmcomputeRunCommandDocumentBase
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
