// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"

func Armtimeseriesinsights() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armtimeseriesinsights.NewAccessPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights",
			URL: "",
		},
		{
			NewFunc: armtimeseriesinsights.NewEnvironmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights",
			URL: "",
		},
		{
			NewFunc: armtimeseriesinsights.NewEventSourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights",
			URL: "",
		},
		{
			NewFunc: armtimeseriesinsights.NewReferenceDataSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armtimeseriesinsights())
}