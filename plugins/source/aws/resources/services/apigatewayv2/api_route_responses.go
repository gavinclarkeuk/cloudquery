// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiRouteResponses() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_api_route_responses",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html`,
		Resolver:    fetchApigatewayv2ApiRouteResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Name:     "api_route_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "route_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("route_id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiRouteResponseArn(),
			},
			{
				Name:     "route_response_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteResponseKey"),
			},
			{
				Name:     "model_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelSelectionExpression"),
			},
			{
				Name:     "response_models",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponseModels"),
			},
			{
				Name:     "response_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponseParameters"),
			},
			{
				Name:     "route_response_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteResponseId"),
			},
		},
	}
}
