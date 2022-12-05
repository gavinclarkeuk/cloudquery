// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"

func Armresourcemover() []Table {
	tables := []Table{
		{
      Name: "move_resource",
      Struct: &armresourcemover.MoveResource{},
      ResponseStruct: &armresourcemover.MoveResourcesClientListResponse{},
      Client: &armresourcemover.MoveResourcesClient{},
      ListFunc: (&armresourcemover.MoveResourcesClient{}).NewListPager,
			NewFunc: armresourcemover.NewMoveResourcesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources",
		},
	}

	for i := range tables {
		tables[i].Service = "armresourcemover"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armresourcemover()...)
}