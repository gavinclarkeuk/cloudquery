// Code generated by codegen; DO NOT EDIT.

package armoperationalinsights

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UsageMetric() *schema.Table {
	return &schema.Table{
		Name:      "azure_armoperationalinsights_usage_metric",
		Resolver:  fetchUsageMetric,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "current_value",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("CurrentValue"),
			},
			{
				Name:     "limit",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Limit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "next_reset_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NextResetTime"),
			},
			{
				Name:     "quota_period",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QuotaPeriod"),
			},
			{
				Name:     "unit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Unit"),
			},
		},
	}
}

func fetchUsageMetric(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmoperationalinsightsUsageMetric
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
