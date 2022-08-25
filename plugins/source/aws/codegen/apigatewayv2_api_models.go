// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func Apigatewayv2ApiModels() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigatewayv2_api_models",
		Resolver:  fetchApigatewayv2ApiModels,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "content_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContentType"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "model_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelId"),
			},
			{
				Name:     "schema",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Schema"),
			},
		},
	}
}

func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	r := parent.Item.(types.Api)

	input := apigatewayv2.GetModelsInput{
		ApiId: r.ApiId,
	}

	for {
		response, err := svc.GetModels(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
