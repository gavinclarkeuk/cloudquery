// Code generated by codegen; DO NOT EDIT.

package armtestbase

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountUsageData() *schema.Table {
	return &schema.Table{
		Name:      "azure_armtestbase_account_usage_data",
		Resolver:  fetchAccountUsageData,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "current_value",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CurrentValue"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
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
		},
	}
}

func fetchAccountUsageData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armtestbase.NewUsageClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
