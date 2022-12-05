// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"

func Armapplicationinsights() []Table {
	tables := []Table{
		{
      Name: "component_api_key",
      Struct: &armapplicationinsights.ComponentAPIKey{},
      ResponseStruct: &armapplicationinsights.APIKeysClientListResponse{},
      Client: &armapplicationinsights.APIKeysClient{},
      ListFunc: (&armapplicationinsights.APIKeysClient{}).NewListPager,
			NewFunc: armapplicationinsights.NewAPIKeysClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys",
		},
		{
      Name: "annotation",
      Struct: &armapplicationinsights.Annotation{},
      ResponseStruct: &armapplicationinsights.AnnotationsClientListResponse{},
      Client: &armapplicationinsights.AnnotationsClient{},
      ListFunc: (&armapplicationinsights.AnnotationsClient{}).NewListPager,
			NewFunc: armapplicationinsights.NewAnnotationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations",
		},
		{
      Name: "component",
      Struct: &armapplicationinsights.Component{},
      ResponseStruct: &armapplicationinsights.ComponentsClientListResponse{},
      Client: &armapplicationinsights.ComponentsClient{},
      ListFunc: (&armapplicationinsights.ComponentsClient{}).NewListPager,
			NewFunc: armapplicationinsights.NewComponentsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components",
		},
		{
      Name: "component_web_test_location",
      Struct: &armapplicationinsights.ComponentWebTestLocation{},
      ResponseStruct: &armapplicationinsights.WebTestLocationsClientListResponse{},
      Client: &armapplicationinsights.WebTestLocationsClient{},
      ListFunc: (&armapplicationinsights.WebTestLocationsClient{}).NewListPager,
			NewFunc: armapplicationinsights.NewWebTestLocationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/syntheticmonitorlocations",
		},
		{
      Name: "web_test",
      Struct: &armapplicationinsights.WebTest{},
      ResponseStruct: &armapplicationinsights.WebTestsClientListResponse{},
      Client: &armapplicationinsights.WebTestsClient{},
      ListFunc: (&armapplicationinsights.WebTestsClient{}).NewListPager,
			NewFunc: armapplicationinsights.NewWebTestsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests",
		},
		{
      Name: "work_item_configuration",
      Struct: &armapplicationinsights.WorkItemConfiguration{},
      ResponseStruct: &armapplicationinsights.WorkItemConfigurationsClientListResponse{},
      Client: &armapplicationinsights.WorkItemConfigurationsClient{},
      ListFunc: (&armapplicationinsights.WorkItemConfigurationsClient{}).NewListPager,
			NewFunc: armapplicationinsights.NewWorkItemConfigurationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs",
		},
	}

	for i := range tables {
		tables[i].Service = "armapplicationinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armapplicationinsights()...)
}