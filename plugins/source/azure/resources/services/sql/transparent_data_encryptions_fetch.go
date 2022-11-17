// Code generated by codegen; DO NOT EDIT.

package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	sql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTransparentDataEncryptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sql

	database := parent.Item.(*sql.Database)
	server := parent.Parent.Item.(*sql.Server)
	id, err := arm.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}

	pager := svc.TransparentDataEncryptionsClient.NewListByDatabasePager(id.ResourceGroupName, *server.Name, *database.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
