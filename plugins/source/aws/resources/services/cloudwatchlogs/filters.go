package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudwatchlogsFilters() *schema.Table {
	return &schema.Table{
		Name:          "aws_cloudwatchlogs_filters",
		Description:   "Metric filters express how CloudWatch Logs would extract metric observations from ingested log events and transform them into metric data in a CloudWatch metric.",
		Resolver:      fetchCloudwatchlogsFilters,
		Multiplex:     client.ServiceAccountRegionMultiplexer("logs"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the metric filter, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
				Type:        schema.TypeInt,
			},
			{
				Name:            "name",
				Description:     "The name of the metric filter.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("FilterName"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "pattern",
				Description: "A symbolic description of how CloudWatch Logs should interpret the data in each log event.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FilterPattern"),
			},
			{
				Name:            "log_group_name",
				Description:     "The name of the log group.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "metric_transformations",
				Description: "Indicates how to transform ingested log events to metric data in a CloudWatch metric.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("MetricTransformations"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchCloudwatchlogsFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudwatchlogs.DescribeMetricFiltersInput
	c := meta.(*client.Client)
	svc := c.Services().CloudwatchLogs
	for {
		response, err := svc.DescribeMetricFilters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.MetricFilters
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
