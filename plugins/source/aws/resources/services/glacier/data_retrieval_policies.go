// Code generated by codegen; DO NOT EDIT.

package glacier

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataRetrievalPolicies() *schema.Table {
	return &schema.Table{
		Name:      "aws_glacier_data_retrieval_policies",
		Resolver:  fetchGlacierDataRetrievalPolicies,
		Multiplex: client.ServiceAccountRegionMultiplexer("glacier"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rules"),
			},
		},
	}
}
