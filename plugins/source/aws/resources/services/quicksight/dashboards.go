// Code generated by codegen; DO NOT EDIT.

package quicksight

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Dashboards() *schema.Table {
	return &schema.Table{
		Name:      "aws_quicksight_dashboards",
		Resolver:  fetchQuicksightDashboards,
		Multiplex: client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "dashboard_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DashboardId"),
			},
			{
				Name:     "last_published_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastPublishedTime"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "published_version_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PublishedVersionNumber"),
			},
		},
	}
}
