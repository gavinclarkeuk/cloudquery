// Code generated by codegen; DO NOT EDIT.

package network

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_security_groups",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#SecurityGroup`,
		Resolver:    fetchSecurityGroups,
		Multiplex:   client.SubscriptionMultiplex,
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
				Name:     "flush_connection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.FlushConnection"),
			},
			{
				Name:     "security_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.SecurityRules"),
			},
			{
				Name:     "default_security_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.DefaultSecurityRules"),
			},
			{
				Name:     "flow_logs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.FlowLogs"),
			},
			{
				Name:     "network_interfaces",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.NetworkInterfaces"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ResourceGUID"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Subnets"),
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
		},
	}
}
