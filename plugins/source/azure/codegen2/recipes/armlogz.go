// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"

func Armlogz() []Table {
	tables := []Table{
		{
      Name: "operation_result",
      Struct: &armlogz.OperationResult{},
      ResponseStruct: &armlogz.OperationsClientListResponse{},
      Client: &armlogz.OperationsClient{},
      ListFunc: (&armlogz.OperationsClient{}).NewListPager,
			NewFunc: armlogz.NewOperationsClient,
		},
		{
      Name: "monitoring_tag_rules",
      Struct: &armlogz.MonitoringTagRules{},
      ResponseStruct: &armlogz.SubAccountTagRulesClientListResponse{},
      Client: &armlogz.SubAccountTagRulesClient{},
      ListFunc: (&armlogz.SubAccountTagRulesClient{}).NewListPager,
			NewFunc: armlogz.NewSubAccountTagRulesClient,
		},
		{
      Name: "single_sign_on_resource",
      Struct: &armlogz.SingleSignOnResource{},
      ResponseStruct: &armlogz.SingleSignOnClientListResponse{},
      Client: &armlogz.SingleSignOnClient{},
      ListFunc: (&armlogz.SingleSignOnClient{}).NewListPager,
			NewFunc: armlogz.NewSingleSignOnClient,
		},
		{
      Name: "monitor_resource",
      Struct: &armlogz.MonitorResource{},
      ResponseStruct: &armlogz.SubAccountClientListResponse{},
      Client: &armlogz.SubAccountClient{},
      ListFunc: (&armlogz.SubAccountClient{}).NewListPager,
			NewFunc: armlogz.NewSubAccountClient,
		},
		{
      Name: "monitoring_tag_rules",
      Struct: &armlogz.MonitoringTagRules{},
      ResponseStruct: &armlogz.TagRulesClientListResponse{},
      Client: &armlogz.TagRulesClient{},
      ListFunc: (&armlogz.TagRulesClient{}).NewListPager,
			NewFunc: armlogz.NewTagRulesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armlogz"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armlogz()...)
}