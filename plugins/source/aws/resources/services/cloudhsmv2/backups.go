// Code generated by codegen; DO NOT EDIT.

package cloudhsmv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Backups() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudhsmv2_backups",
		Description: `https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Backup.html`,
		Resolver:    fetchCloudhsmv2Backups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudhsmv2"),
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
				Resolver: resolveBackupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
			{
				Name:     "backup_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupId"),
			},
			{
				Name:     "backup_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupState"),
			},
			{
				Name:     "cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterId"),
			},
			{
				Name:     "copy_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CopyTimestamp"),
			},
			{
				Name:     "create_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTimestamp"),
			},
			{
				Name:     "delete_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeleteTimestamp"),
			},
			{
				Name:     "never_expires",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("NeverExpires"),
			},
			{
				Name:     "source_backup",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceBackup"),
			},
			{
				Name:     "source_cluster",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceCluster"),
			},
			{
				Name:     "source_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceRegion"),
			},
		},
	}
}
