// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

func AccessAnalyzerAccessanalyzers() *schema.Table {
	return &schema.Table{
		Name:      "aws_accessanalyzer_accessanalyzers",
		Resolver:  fetchAccessAnalyzerAccessanalyzers,
		Multiplex: client.ServiceAccountRegionMultiplexer("accessanalyzer"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "last_resource_analyzed",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastResourceAnalyzed"),
			},
			{
				Name:     "last_resource_analyzed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastResourceAnalyzedAt"),
			},
			{
				Name:     "status_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StatusReason"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchAccessAnalyzerAccessanalyzers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().AccessAnalyzer

	input := accessanalyzer.ListAnalyzersInput{}

	for {
		response, err := svc.ListAnalyzers(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}
		res <- response.Analyzers
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
