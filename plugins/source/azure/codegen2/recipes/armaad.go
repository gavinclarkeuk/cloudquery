// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"

func Armaad() []Table {
	tables := []Table{
		{
      Name: "private_link_policy",
      Struct: &armaad.PrivateLinkPolicy{},
      ResponseStruct: &armaad.PrivateLinkForAzureAdClientListResponse{},
      Client: &armaad.PrivateLinkForAzureAdClient{},
      ListFunc: (&armaad.PrivateLinkForAzureAdClient{}).NewListPager,
			NewFunc: armaad.NewPrivateLinkForAzureAdClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd",
		},
	}

	for i := range tables {
		tables[i].Service = "armaad"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armaad()...)
}