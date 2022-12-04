// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"

func Armblockchain() []Table {
	tables := []Table{
		{
      Name: "member",
      Struct: &armblockchain.Member{},
      ResponseStruct: &armblockchain.MembersClientListResponse{},
      Client: &armblockchain.MembersClient{},
      ListFunc: (&armblockchain.MembersClient{}).NewListPager,
			NewFunc: armblockchain.NewMembersClient,
		},
		{
      Name: "resource_provider_operation",
      Struct: &armblockchain.ResourceProviderOperation{},
      ResponseStruct: &armblockchain.OperationsClientListResponse{},
      Client: &armblockchain.OperationsClient{},
      ListFunc: (&armblockchain.OperationsClient{}).NewListPager,
			NewFunc: armblockchain.NewOperationsClient,
		},
		{
      Name: "transaction_node",
      Struct: &armblockchain.TransactionNode{},
      ResponseStruct: &armblockchain.TransactionNodesClientListResponse{},
      Client: &armblockchain.TransactionNodesClient{},
      ListFunc: (&armblockchain.TransactionNodesClient{}).NewListPager,
			NewFunc: armblockchain.NewTransactionNodesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armblockchain"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armblockchain()...)
}