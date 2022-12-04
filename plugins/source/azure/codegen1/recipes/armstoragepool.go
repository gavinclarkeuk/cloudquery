// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"

func Armstoragepool() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armstoragepool.NewDiskPoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool",
		},
		{
			NewFunc: armstoragepool.NewDiskPoolZonesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool",
		},
		{
			NewFunc: armstoragepool.NewIscsiTargetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool",
		},
		{
			NewFunc: armstoragepool.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool",
		},
		{
			NewFunc: armstoragepool.NewResourceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armstoragepool())
}