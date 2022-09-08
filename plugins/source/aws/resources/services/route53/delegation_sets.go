package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Route53ReusableDelegationSets() *schema.Table {
	return &schema.Table{
		Name:          "aws_route53_reusable_delegation_sets",
		Resolver:      fetchRoute53DelegationSets,
		Multiplex:     client.AccountMultiplex,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"delegationset", *resource.Item.(types.DelegationSet).Id}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "name_servers",
				Type: schema.TypeStringArray,
			},
			{
				Name: "caller_reference",
				Type: schema.TypeString,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRoute53DelegationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config route53.ListReusableDelegationSetsInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListReusableDelegationSets(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DelegationSets
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
