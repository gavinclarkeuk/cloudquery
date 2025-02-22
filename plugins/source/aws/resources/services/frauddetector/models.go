// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Models() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_models",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_Model.html`,
		Resolver:    fetchFrauddetectorModels,
		Multiplex:   client.ServiceAccountRegionMultiplexer("frauddetector"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "event_type_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventTypeName"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "model_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelId"),
			},
			{
				Name:     "model_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelType"),
			},
		},

		Relations: []*schema.Table{
			ModelVersions(),
		},
	}
}
