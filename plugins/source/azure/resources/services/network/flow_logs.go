// Code generated by codegen; DO NOT EDIT.

package network

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func flowLogs() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_flow_logs",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#FlowLog`,
		Resolver:    fetchFlowLogs,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "storage_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.StorageID"),
			},
			{
				Name:     "target_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.TargetResourceID"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.Enabled"),
			},
			{
				Name:     "flow_analytics_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.FlowAnalyticsConfiguration"),
			},
			{
				Name:     "format",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Format"),
			},
			{
				Name:     "retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.RetentionPolicy"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "target_resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.TargetResourceGUID"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "watcher_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
