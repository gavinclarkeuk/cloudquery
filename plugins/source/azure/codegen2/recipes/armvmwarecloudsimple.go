// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"

func Armvmwarecloudsimple() []Table {
	tables := []Table{
		{
      Name: "private_cloud",
      Struct: &armvmwarecloudsimple.PrivateCloud{},
      ResponseStruct: &armvmwarecloudsimple.PrivateCloudsClientListResponse{},
      Client: &armvmwarecloudsimple.PrivateCloudsClient{},
      ListFunc: (&armvmwarecloudsimple.PrivateCloudsClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewPrivateCloudsClient,
		},
		{
      Name: "customization_policy",
      Struct: &armvmwarecloudsimple.CustomizationPolicy{},
      ResponseStruct: &armvmwarecloudsimple.CustomizationPoliciesClientListResponse{},
      Client: &armvmwarecloudsimple.CustomizationPoliciesClient{},
      ListFunc: (&armvmwarecloudsimple.CustomizationPoliciesClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewCustomizationPoliciesClient,
		},
		{
      Name: "sku_availability",
      Struct: &armvmwarecloudsimple.SKUAvailability{},
      ResponseStruct: &armvmwarecloudsimple.SKUsAvailabilityClientListResponse{},
      Client: &armvmwarecloudsimple.SKUsAvailabilityClient{},
      ListFunc: (&armvmwarecloudsimple.SKUsAvailabilityClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewSKUsAvailabilityClient,
		},
		{
      Name: "virtual_network",
      Struct: &armvmwarecloudsimple.VirtualNetwork{},
      ResponseStruct: &armvmwarecloudsimple.VirtualNetworksClientListResponse{},
      Client: &armvmwarecloudsimple.VirtualNetworksClient{},
      ListFunc: (&armvmwarecloudsimple.VirtualNetworksClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewVirtualNetworksClient,
		},
		{
      Name: "resource_pool",
      Struct: &armvmwarecloudsimple.ResourcePool{},
      ResponseStruct: &armvmwarecloudsimple.ResourcePoolsClientListResponse{},
      Client: &armvmwarecloudsimple.ResourcePoolsClient{},
      ListFunc: (&armvmwarecloudsimple.ResourcePoolsClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewResourcePoolsClient,
		},
		{
      Name: "virtual_machine_template",
      Struct: &armvmwarecloudsimple.VirtualMachineTemplate{},
      ResponseStruct: &armvmwarecloudsimple.VirtualMachineTemplatesClientListResponse{},
      Client: &armvmwarecloudsimple.VirtualMachineTemplatesClient{},
      ListFunc: (&armvmwarecloudsimple.VirtualMachineTemplatesClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewVirtualMachineTemplatesClient,
		},
		{
      Name: "available_operation",
      Struct: &armvmwarecloudsimple.AvailableOperation{},
      ResponseStruct: &armvmwarecloudsimple.OperationsClientListResponse{},
      Client: &armvmwarecloudsimple.OperationsClient{},
      ListFunc: (&armvmwarecloudsimple.OperationsClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewOperationsClient,
		},
		{
      Name: "usage",
      Struct: &armvmwarecloudsimple.Usage{},
      ResponseStruct: &armvmwarecloudsimple.UsagesClientListResponse{},
      Client: &armvmwarecloudsimple.UsagesClient{},
      ListFunc: (&armvmwarecloudsimple.UsagesClient{}).NewListPager,
			NewFunc: armvmwarecloudsimple.NewUsagesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armvmwarecloudsimple"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armvmwarecloudsimple()...)
}