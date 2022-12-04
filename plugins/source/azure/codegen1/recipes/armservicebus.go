// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"

func Armservicebus() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armservicebus.NewDisasterRecoveryConfigsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewNamespacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewQueuesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewSubscriptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewTopicsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
		{
			NewFunc: armservicebus.NewMigrationConfigsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armservicebus())
}