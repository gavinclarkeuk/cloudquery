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
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers",
		},
		{
      Name: "transaction_node",
      Struct: &armblockchain.TransactionNode{},
      ResponseStruct: &armblockchain.TransactionNodesClientListResponse{},
      Client: &armblockchain.TransactionNodesClient{},
      ListFunc: (&armblockchain.TransactionNodesClient{}).NewListPager,
			NewFunc: armblockchain.NewTransactionNodesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes",
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