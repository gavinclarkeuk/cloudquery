// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"

func Armdns() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdns.NewRecordSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns",
			URL: "",
		},
		{
			NewFunc: armdns.NewResourceReferenceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns",
			URL: "",
		},
		{
			NewFunc: armdns.NewZonesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdns())
}