// Code generated by codegen; DO NOT EDIT.

package armrecoveryservicessiterecovery

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReplicationAppliance() *schema.Table {
	return &schema.Table{
		Name:      "azure_armrecoveryservicessiterecovery_replication_appliance",
		Resolver:  fetchReplicationAppliance,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
		},
	}
}

func fetchReplicationAppliance(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmrecoveryservicessiterecoveryReplicationAppliance
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
