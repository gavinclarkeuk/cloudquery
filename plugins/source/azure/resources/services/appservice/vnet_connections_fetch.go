// Code generated by codegen; DO NOT EDIT.

package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	appservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchVnetConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Appservice

	site := parent.Item.(*appservice.Site)
	id, err := arm.ParseResourceID(*site.ID)
	if err != nil {
		return err
	}

	data, err := svc.WebAppsClient.ListVnetConnections(ctx, id.ResourceGroupName, *site.Name, nil)
	if err != nil {
		return err
	}
	res <- data.VnetInfoResourceArray

	return nil
}
