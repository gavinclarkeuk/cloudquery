// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"

func Armsubscription() []Table {
	tables := []Table{
		{
      Name: "subscription",
      Struct: &armsubscription.Subscription{},
      ResponseStruct: &armsubscription.SubscriptionsClientListResponse{},
      Client: &armsubscription.SubscriptionsClient{},
      ListFunc: (&armsubscription.SubscriptionsClient{}).NewListPager,
			NewFunc: armsubscription.NewSubscriptionsClient,
		},
		{
      Name: "tenant_id_description",
      Struct: &armsubscription.TenantIDDescription{},
      ResponseStruct: &armsubscription.TenantsClientListResponse{},
      Client: &armsubscription.TenantsClient{},
      ListFunc: (&armsubscription.TenantsClient{}).NewListPager,
			NewFunc: armsubscription.NewTenantsClient,
		},
		{
      Name: "operation",
      Struct: &armsubscription.Operation{},
      ResponseStruct: &armsubscription.OperationsClientListResponse{},
      Client: &armsubscription.OperationsClient{},
      ListFunc: (&armsubscription.OperationsClient{}).NewListPager,
			NewFunc: armsubscription.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armsubscription"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armsubscription()...)
}