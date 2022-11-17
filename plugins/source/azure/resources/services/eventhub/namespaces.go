// Code generated by codegen; DO NOT EDIT.

package eventhub

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:        "azure_eventhub_namespaces",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub#EHNamespace`,
		Resolver:    fetchNamespaces,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "alternate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.AlternateName"),
			},
			{
				Name:     "cluster_arm_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ClusterArmID"),
			},
			{
				Name:     "disable_local_auth",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.DisableLocalAuth"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Encryption"),
			},
			{
				Name:     "is_auto_inflate_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.IsAutoInflateEnabled"),
			},
			{
				Name:     "kafka_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.KafkaEnabled"),
			},
			{
				Name:     "maximum_throughput_units",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Properties.MaximumThroughputUnits"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.PrivateEndpointConnections"),
			},
			{
				Name:     "zone_redundant",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.ZoneRedundant"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Properties.CreatedAt"),
			},
			{
				Name:     "metric_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.MetricID"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "service_bus_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ServiceBusEndpoint"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Status"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Properties.UpdatedAt"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SKU"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			networkRuleSets(),
		},
	}
}
