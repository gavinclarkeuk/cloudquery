// Code generated by codegen; DO NOT EDIT.

package mysql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	mysql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Mysql

	server := parent.Item.(*mysql.Server)
	id, err := arm.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}

	pager := svc.ConfigurationsClient.NewListByServerPager(id.ResourceGroupName, *server.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
