// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"

func Armresourcehealth() []Table {
	tables := []Table{
		{
      Name: "availability_status",
      Struct: &armresourcehealth.AvailabilityStatus{},
      ResponseStruct: &armresourcehealth.AvailabilityStatusesClientListResponse{},
      Client: &armresourcehealth.AvailabilityStatusesClient{},
      ListFunc: (&armresourcehealth.AvailabilityStatusesClient{}).NewListPager,
			NewFunc: armresourcehealth.NewAvailabilityStatusesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armresourcehealth"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armresourcehealth()...)
}