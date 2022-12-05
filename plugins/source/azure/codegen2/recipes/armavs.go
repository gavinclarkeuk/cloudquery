// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"

func Armavs() []Table {
	tables := []Table{
		{
      Name: "addon",
      Struct: &armavs.Addon{},
      ResponseStruct: &armavs.AddonsClientListResponse{},
      Client: &armavs.AddonsClient{},
      ListFunc: (&armavs.AddonsClient{}).NewListPager,
			NewFunc: armavs.NewAddonsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/addons",
		},
		{
      Name: "express_route_authorization",
      Struct: &armavs.ExpressRouteAuthorization{},
      ResponseStruct: &armavs.AuthorizationsClientListResponse{},
      Client: &armavs.AuthorizationsClient{},
      ListFunc: (&armavs.AuthorizationsClient{}).NewListPager,
			NewFunc: armavs.NewAuthorizationsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations",
		},
		{
      Name: "cloud_link",
      Struct: &armavs.CloudLink{},
      ResponseStruct: &armavs.CloudLinksClientListResponse{},
      Client: &armavs.CloudLinksClient{},
      ListFunc: (&armavs.CloudLinksClient{}).NewListPager,
			NewFunc: armavs.NewCloudLinksClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/cloudLinks",
		},
		{
      Name: "cluster",
      Struct: &armavs.Cluster{},
      ResponseStruct: &armavs.ClustersClientListResponse{},
      Client: &armavs.ClustersClient{},
      ListFunc: (&armavs.ClustersClient{}).NewListPager,
			NewFunc: armavs.NewClustersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters",
		},
		{
      Name: "datastore",
      Struct: &armavs.Datastore{},
      ResponseStruct: &armavs.DatastoresClientListResponse{},
      Client: &armavs.DatastoresClient{},
      ListFunc: (&armavs.DatastoresClient{}).NewListPager,
			NewFunc: armavs.NewDatastoresClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/datastores",
		},
		{
      Name: "global_reach_connection",
      Struct: &armavs.GlobalReachConnection{},
      ResponseStruct: &armavs.GlobalReachConnectionsClientListResponse{},
      Client: &armavs.GlobalReachConnectionsClient{},
      ListFunc: (&armavs.GlobalReachConnectionsClient{}).NewListPager,
			NewFunc: armavs.NewGlobalReachConnectionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/globalReachConnections",
		},
		{
      Name: "hcx_enterprise_site",
      Struct: &armavs.HcxEnterpriseSite{},
      ResponseStruct: &armavs.HcxEnterpriseSitesClientListResponse{},
      Client: &armavs.HcxEnterpriseSitesClient{},
      ListFunc: (&armavs.HcxEnterpriseSitesClient{}).NewListPager,
			NewFunc: armavs.NewHcxEnterpriseSitesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites",
		},
		{
      Name: "placement_policy",
      Struct: &armavs.PlacementPolicy{},
      ResponseStruct: &armavs.PlacementPoliciesClientListResponse{},
      Client: &armavs.PlacementPoliciesClient{},
      ListFunc: (&armavs.PlacementPoliciesClient{}).NewListPager,
			NewFunc: armavs.NewPlacementPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies",
		},
		{
      Name: "private_cloud",
      Struct: &armavs.PrivateCloud{},
      ResponseStruct: &armavs.PrivateCloudsClientListResponse{},
      Client: &armavs.PrivateCloudsClient{},
      ListFunc: (&armavs.PrivateCloudsClient{}).NewListPager,
			NewFunc: armavs.NewPrivateCloudsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds",
		},
		{
      Name: "script_cmdlet",
      Struct: &armavs.ScriptCmdlet{},
      ResponseStruct: &armavs.ScriptCmdletsClientListResponse{},
      Client: &armavs.ScriptCmdletsClient{},
      ListFunc: (&armavs.ScriptCmdletsClient{}).NewListPager,
			NewFunc: armavs.NewScriptCmdletsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages/{scriptPackageName}/scriptCmdlets",
		},
		{
      Name: "script_execution",
      Struct: &armavs.ScriptExecution{},
      ResponseStruct: &armavs.ScriptExecutionsClientListResponse{},
      Client: &armavs.ScriptExecutionsClient{},
      ListFunc: (&armavs.ScriptExecutionsClient{}).NewListPager,
			NewFunc: armavs.NewScriptExecutionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions",
		},
		{
      Name: "script_package",
      Struct: &armavs.ScriptPackage{},
      ResponseStruct: &armavs.ScriptPackagesClientListResponse{},
      Client: &armavs.ScriptPackagesClient{},
      ListFunc: (&armavs.ScriptPackagesClient{}).NewListPager,
			NewFunc: armavs.NewScriptPackagesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages",
		},
		{
      Name: "virtual_machine",
      Struct: &armavs.VirtualMachine{},
      ResponseStruct: &armavs.VirtualMachinesClientListResponse{},
      Client: &armavs.VirtualMachinesClient{},
      ListFunc: (&armavs.VirtualMachinesClient{}).NewListPager,
			NewFunc: armavs.NewVirtualMachinesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines",
		},
		{
      Name: "workload_network",
      Struct: &armavs.WorkloadNetwork{},
      ResponseStruct: &armavs.WorkloadNetworksClientListResponse{},
      Client: &armavs.WorkloadNetworksClient{},
      ListFunc: (&armavs.WorkloadNetworksClient{}).NewListPager,
			NewFunc: armavs.NewWorkloadNetworksClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/workloadNetworks",
		},
	}

	for i := range tables {
		tables[i].Service = "armavs"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armavs()...)
}