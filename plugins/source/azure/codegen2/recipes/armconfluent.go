// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"

func Armconfluent() []Table {
	tables := []Table{
		{
      Name: "agreement_resource",
      Struct: &armconfluent.AgreementResource{},
      ResponseStruct: &armconfluent.MarketplaceAgreementsClientListResponse{},
      Client: &armconfluent.MarketplaceAgreementsClient{},
      ListFunc: (&armconfluent.MarketplaceAgreementsClient{}).NewListPager,
			NewFunc: armconfluent.NewMarketplaceAgreementsClient,
		},
		{
      Name: "operation_result",
      Struct: &armconfluent.OperationResult{},
      ResponseStruct: &armconfluent.OrganizationOperationsClientListResponse{},
      Client: &armconfluent.OrganizationOperationsClient{},
      ListFunc: (&armconfluent.OrganizationOperationsClient{}).NewListPager,
			NewFunc: armconfluent.NewOrganizationOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armconfluent"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armconfluent()...)
}