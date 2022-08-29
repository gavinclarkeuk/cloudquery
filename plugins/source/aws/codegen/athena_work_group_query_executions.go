// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

func AthenaWorkGroupQueryExecutions() *schema.Table {
	return &schema.Table{
		Name:      "aws_athena_work_group_query_executions",
		Resolver:  fetchAthenaWorkGroupQueryExecutions,
		Multiplex: client.ServiceAccountRegionMultiplexer("athena"),
		Columns: []schema.Column{
			{
				Name:     "engine_version",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "execution_parameters",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ExecutionParameters"),
			},
			{
				Name:     "query",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Query"),
			},
			{
				Name:     "query_execution_context",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("QueryExecutionContext"),
			},
			{
				Name:     "query_execution_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryExecutionId"),
			},
			{
				Name:     "result_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultConfiguration"),
			},
			{
				Name:     "statement_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatementType"),
			},
			{
				Name:     "statistics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Statistics"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "work_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkGroup"),
			},
		},
	}
}

func fetchAthenaWorkGroupQueryExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena

	r1 := parent.Item.(types.WorkGroup)

	input := athena.ListQueryExecutionsInput{
		WorkGroup: r1.Name,
	}
	paginator := athena.NewListQueryExecutionsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range output.QueryExecutionIds {
			do, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{

				QueryExecutionId: &item,
			})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- do.QueryExecution
		}
	}
	return nil
}
