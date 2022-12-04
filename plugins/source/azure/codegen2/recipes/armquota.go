// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"

func Armquota() []Table {
	tables := []Table{
		{
      Name: "current_usages_base",
      Struct: &armquota.CurrentUsagesBase{},
      ResponseStruct: &armquota.UsagesClientListResponse{},
      Client: &armquota.UsagesClient{},
      ListFunc: (&armquota.UsagesClient{}).NewListPager,
			NewFunc: armquota.NewUsagesClient,
		},
		{
      Name: "operation_response",
      Struct: &armquota.OperationResponse{},
      ResponseStruct: &armquota.OperationClientListResponse{},
      Client: &armquota.OperationClient{},
      ListFunc: (&armquota.OperationClient{}).NewListPager,
			NewFunc: armquota.NewOperationClient,
		},
		{
      Name: "request_details",
      Struct: &armquota.RequestDetails{},
      ResponseStruct: &armquota.RequestStatusClientListResponse{},
      Client: &armquota.RequestStatusClient{},
      ListFunc: (&armquota.RequestStatusClient{}).NewListPager,
			NewFunc: armquota.NewRequestStatusClient,
		},
		{
      Name: "current_quota_limit_base",
      Struct: &armquota.CurrentQuotaLimitBase{},
      ResponseStruct: &armquota.ClientListResponse{},
      Client: &armquota.Client{},
      ListFunc: (&armquota.Client{}).NewListPager,
			NewFunc: armquota.NewClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armquota"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armquota()...)
}