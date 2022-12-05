// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"

func Armsecuritydevops() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsecuritydevops.NewAzureDevOpsConnectorClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "",
		},
		{
			NewFunc: armsecuritydevops.NewAzureDevOpsConnectorStatsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "",
		},
		{
			NewFunc: armsecuritydevops.NewAzureDevOpsOrgClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs",
		},
		{
			NewFunc: armsecuritydevops.NewAzureDevOpsProjectClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs/{azureDevOpsOrgName}/projects",
		},
		{
			NewFunc: armsecuritydevops.NewAzureDevOpsRepoClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs/{azureDevOpsOrgName}/projects/{azureDevOpsProjectName}/repos",
		},
		{
			NewFunc: armsecuritydevops.NewGitHubConnectorClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "",
		},
		{
			NewFunc: armsecuritydevops.NewGitHubConnectorStatsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "",
		},
		{
			NewFunc: armsecuritydevops.NewGitHubOwnerClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}/owners",
		},
		{
			NewFunc: armsecuritydevops.NewGitHubRepoClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}/owners/{gitHubOwnerName}/repos",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsecuritydevops())
}