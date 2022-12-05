// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"

func Armelasticsan() []Table {
	tables := []Table{
		{
      Name: "sku_information",
      Struct: &armelasticsan.SKUInformation{},
      ResponseStruct: &armelasticsan.SKUsClientListResponse{},
      Client: &armelasticsan.SKUsClient{},
      ListFunc: (&armelasticsan.SKUsClient{}).NewListPager,
			NewFunc: armelasticsan.NewSKUsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ElasticSan/skus",
		},
	}

	for i := range tables {
		tables[i].Service = "armelasticsan"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armelasticsan()...)
}