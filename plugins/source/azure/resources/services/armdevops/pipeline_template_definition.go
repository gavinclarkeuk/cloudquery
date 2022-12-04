// Code generated by codegen; DO NOT EDIT.

package armdevops

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PipelineTemplateDefinition() *schema.Table {
	return &schema.Table{
		Name:      "azure_armdevops_pipeline_template_definition",
		Resolver:  fetchPipelineTemplateDefinition,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "inputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Inputs"),
			},
		},
	}
}

func fetchPipelineTemplateDefinition(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armdevops.NewPipelineTemplateDefinitionsClient(cl.Creds, cl.Options)
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
