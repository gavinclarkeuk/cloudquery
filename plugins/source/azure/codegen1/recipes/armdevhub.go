// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"

func Armdevhub() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdevhub.NewWorkflowClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub",
		},
		{
			NewFunc: armdevhub.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub",
		},
		{
			NewFunc: armdevhub.NewDeveloperHubServiceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdevhub())
}