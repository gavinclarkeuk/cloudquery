// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"

func Armdatafactory() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatafactory.NewActivityRunsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewDataFlowDebugSessionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewDataFlowsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewDatasetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewExposureControlClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewFactoriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories",
		},
		{
			NewFunc: armdatafactory.NewGlobalParametersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewIntegrationRuntimeNodesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewIntegrationRuntimeObjectMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewIntegrationRuntimesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewLinkedServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewManagedPrivateEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewManagedVirtualNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewPipelineRunsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewPipelinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewPrivateEndPointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewPrivateEndpointConnectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewTriggerRunsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
		{
			NewFunc: armdatafactory.NewTriggersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatafactory())
}