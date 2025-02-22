// Code generated by codegen; DO NOT EDIT.

package ram

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceShareInvitations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_invitations",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html`,
		Resolver:    fetchRamResourceShareInvitations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
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
				Name:     "invitation_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("InvitationTimestamp"),
			},
			{
				Name:     "receiver_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReceiverAccountId"),
			},
			{
				Name:     "receiver_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReceiverArn"),
			},
			{
				Name:     "resource_share_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareArn"),
			},
			{
				Name:     "resource_share_associations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceShareAssociations"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareInvitationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "resource_share_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareName"),
			},
			{
				Name:     "sender_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SenderAccountId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
