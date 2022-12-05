// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"

func Armstoragepool() []Table {
	tables := []Table{
		{
      Name: "disk_pool_zone_info",
      Struct: &armstoragepool.DiskPoolZoneInfo{},
      ResponseStruct: &armstoragepool.DiskPoolZonesClientListResponse{},
      Client: &armstoragepool.DiskPoolZonesClient{},
      ListFunc: (&armstoragepool.DiskPoolZonesClient{}).NewListPager,
			NewFunc: armstoragepool.NewDiskPoolZonesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.StoragePool/locations/{location}/diskPoolZones",
		},
		{
      Name: "resource_sku_info",
      Struct: &armstoragepool.ResourceSKUInfo{},
      ResponseStruct: &armstoragepool.ResourceSKUsClientListResponse{},
      Client: &armstoragepool.ResourceSKUsClient{},
      ListFunc: (&armstoragepool.ResourceSKUsClient{}).NewListPager,
			NewFunc: armstoragepool.NewResourceSKUsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.StoragePool/locations/{location}/skus",
		},
	}

	for i := range tables {
		tables[i].Service = "armstoragepool"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armstoragepool()...)
}