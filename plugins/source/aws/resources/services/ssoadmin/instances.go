// Code generated by codegen; DO NOT EDIT.

package ssoadmin

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssoadmin_instances",
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html`,
		Resolver:    fetchSsoadminInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("identitystore"),
		Columns: []schema.Column{
			{
				Name:     "identity_store_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityStoreId"),
			},
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceArn"),
			},
		},

		Relations: []*schema.Table{
			PermissionSets(),
		},
	}
}
