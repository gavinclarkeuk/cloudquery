package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/appservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func AppService() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armappservice.Site),
			Resolver: &resource.FuncParams{
				Func: appservice.WebAppsClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					SubService: "functions",
					Struct:     new(armappservice.FunctionEnvelope),
					Resolver: &resource.FuncParams{
						Func: appservice.WebAppsClient.NewListFunctionsPager,
					},
				},
				{
					Struct: new(armappservice.SiteAuthSettings),
					Resolver: &resource.FuncParams{
						Func: appservice.WebAppsClient.GetAuthSettings,
					},
				},
				{
					SubService: "site_auth_settings_v2",
					Struct:     new(armappservice.SiteAuthSettingsV2),
					Resolver: &resource.FuncParams{
						Func: appservice.WebAppsClient.GetAuthSettingsV2,
					},
				},
				{
					SubService: "vnet_connections",
					Struct:     new(armappservice.VnetInfoResource),
					Resolver: &resource.FuncParams{
						Func: appservice.WebAppsClient.ListVnetConnections,
					},
				},
			},
		},
	}
}
