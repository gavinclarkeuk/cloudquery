// Code generated by codegen; DO NOT EDIT.

package mysql

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:        "azure_mysql_servers",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql#Server`,
		Resolver:    fetchServers,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "administrator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.AdministratorLogin"),
			},
			{
				Name:     "earliest_restore_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Properties.EarliestRestoreDate"),
			},
			{
				Name:     "fully_qualified_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.FullyQualifiedDomainName"),
			},
			{
				Name:     "infrastructure_encryption",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.InfrastructureEncryption"),
			},
			{
				Name:     "master_server_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.MasterServerID"),
			},
			{
				Name:     "minimal_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.MinimalTLSVersion"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.PublicNetworkAccess"),
			},
			{
				Name:     "replica_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Properties.ReplicaCapacity"),
			},
			{
				Name:     "replication_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ReplicationRole"),
			},
			{
				Name:     "ssl_enforcement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.SSLEnforcement"),
			},
			{
				Name:     "storage_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.StorageProfile"),
			},
			{
				Name:     "user_visible_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.UserVisibleState"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Version"),
			},
			{
				Name:     "byok_enforcement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ByokEnforcement"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.PrivateEndpointConnections"),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			configurations(),
		},
	}
}
