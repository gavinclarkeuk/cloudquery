// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstancePatches() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_instance_patches",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html`,
		Resolver:    fetchSsmInstancePatches,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
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
				Name:     "kb_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KBId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "classification",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Classification"),
			},
			{
				Name:     "installed_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("InstalledTime"),
			},
			{
				Name:     "severity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Severity"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "cve_ids",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CVEIds"),
			},
		},
	}
}
