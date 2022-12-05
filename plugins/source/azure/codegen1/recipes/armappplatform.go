// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"

func Armappplatform() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armappplatform.NewAppsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps",
		},
		{
			NewFunc: armappplatform.NewBindingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings",
		},
		{
			NewFunc: armappplatform.NewBuildServiceAgentPoolClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools",
		},
		{
			NewFunc: armappplatform.NewBuildServiceBuilderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders",
		},
		{
			NewFunc: armappplatform.NewBuildServiceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "",
		},
		{
			NewFunc: armappplatform.NewBuildpackBindingClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}/buildpackBindings",
		},
		{
			NewFunc: armappplatform.NewCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/certificates",
		},
		{
			NewFunc: armappplatform.NewConfigServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "",
		},
		{
			NewFunc: armappplatform.NewConfigurationServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configurationServices",
		},
		{
			NewFunc: armappplatform.NewCustomDomainsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/domains",
		},
		{
			NewFunc: armappplatform.NewDeploymentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments",
		},
		{
			NewFunc: armappplatform.NewMonitoringSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "",
		},
		{
			NewFunc: armappplatform.NewRuntimeVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "",
		},
		{
			NewFunc: armappplatform.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/skus",
		},
		{
			NewFunc: armappplatform.NewServiceRegistriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries",
		},
		{
			NewFunc: armappplatform.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armappplatform())
}