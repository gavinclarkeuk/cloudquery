// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/apigatewayv2"
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
			{
				Name:     "model_template",
				Type:     schema.TypeString,
				Resolver: resolvers.ResolveApiModelTemplate,
			},
		},
	}
}

func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	r1 := parent.Item.(types.Api)

	input := apigatewayv2.GetModelsInput{
		ApiId: r1.ApiId,
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
