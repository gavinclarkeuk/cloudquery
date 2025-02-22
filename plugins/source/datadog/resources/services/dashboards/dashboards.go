// Code generated by codegen; DO NOT EDIT.

package dashboards

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Dashboards() *schema.Table {
	return &schema.Table{
		Name:      "datadog_dashboards",
		Resolver:  fetchDashboards,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "author_handle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorHandle"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "description",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "is_read_only",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsReadOnly"),
			},
			{
				Name:     "layout_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LayoutType"),
			},
			{
				Name:     "modified_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ModifiedAt"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
			},
		},
	}
}
