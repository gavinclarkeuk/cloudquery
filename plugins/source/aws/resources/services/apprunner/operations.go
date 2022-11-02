// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Operations() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_operations",
		Description: "https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html",
		Resolver:    fetchApprunnerOperations,
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
				Name:     "ended_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "started_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartedAt"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "target_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetArn"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}