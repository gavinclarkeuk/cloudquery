// Code generated by codegen; DO NOT EDIT.

package network

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ExpressRouteGateways() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_express_route_gateways",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#ExpressRouteGateway`,
		Resolver:    fetchExpressRouteGateways,
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
				Name:     "virtual_hub",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.VirtualHub"),
			},
			{
				Name:     "allow_non_virtual_wan_traffic",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.AllowNonVirtualWanTraffic"),
			},
			{
				Name:     "auto_scale_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.AutoScaleConfiguration"),
			},
			{
				Name:     "express_route_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.ExpressRouteConnections"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
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
