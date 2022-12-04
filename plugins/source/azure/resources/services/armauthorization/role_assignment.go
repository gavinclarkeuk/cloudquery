// Code generated by codegen; DO NOT EDIT.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RoleAssignment() *schema.Table {
	return &schema.Table{
		Name:      "azure_armauthorization_role_assignment",
		Resolver:  fetchRoleAssignment,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchRoleAssignment(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armauthorization.NewRoleAssignmentsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
