// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"

func Armfeatures() []Table {
	tables := []Table{
		{
      Name: "feature_result",
      Struct: &armfeatures.FeatureResult{},
      ResponseStruct: &armfeatures.ClientListResponse{},
      Client: &armfeatures.Client{},
      ListFunc: (&armfeatures.Client{}).NewListPager,
			NewFunc: armfeatures.NewClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features",
		},
	}

	for i := range tables {
		tables[i].Service = "armfeatures"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armfeatures()...)
}