// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"

func Armnetwork() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armnetwork.NewAdminRuleCollectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections",
		},
		{
			NewFunc: armnetwork.NewAdminRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules",
		},
		{
			NewFunc: armnetwork.NewApplicationGatewayPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections",
		},
		{
			NewFunc: armnetwork.NewApplicationGatewayPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateLinkResources",
		},
		{
			NewFunc: armnetwork.NewApplicationGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways",
		},
		{
			NewFunc: armnetwork.NewApplicationSecurityGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups",
		},
		{
			NewFunc: armnetwork.NewAvailableDelegationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations",
		},
		{
			NewFunc: armnetwork.NewAvailableEndpointServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices",
		},
		{
			NewFunc: armnetwork.NewAvailablePrivateEndpointTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes",
		},
		{
			NewFunc: armnetwork.NewAvailableResourceGroupDelegationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations",
		},
		{
			NewFunc: armnetwork.NewAvailableServiceAliasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases",
		},
		{
			NewFunc: armnetwork.NewAzureFirewallFqdnTagsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewAzureFirewallsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls",
		},
		{
			NewFunc: armnetwork.NewBastionHostsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bastionHosts",
		},
		{
			NewFunc: armnetwork.NewBgpServiceCommunitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities",
		},
		{
			NewFunc: armnetwork.NewConfigurationPolicyGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewConnectionMonitorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors",
		},
		{
			NewFunc: armnetwork.NewConnectivityConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/connectivityConfigurations",
		},
		{
			NewFunc: armnetwork.NewCustomIPPrefixesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/customIpPrefixes",
		},
		{
			NewFunc: armnetwork.NewDdosCustomPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewDdosProtectionPlansClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans",
		},
		{
			NewFunc: armnetwork.NewDefaultSecurityRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules",
		},
		{
			NewFunc: armnetwork.NewDscpConfigurationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations",
		},
		{
			NewFunc: armnetwork.NewExpressRouteCircuitAuthorizationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations",
		},
		{
			NewFunc: armnetwork.NewExpressRouteCircuitConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections",
		},
		{
			NewFunc: armnetwork.NewExpressRouteCircuitPeeringsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings",
		},
		{
			NewFunc: armnetwork.NewExpressRouteCircuitsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits",
		},
		{
			NewFunc: armnetwork.NewExpressRouteConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections",
		},
		{
			NewFunc: armnetwork.NewExpressRouteCrossConnectionPeeringsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings",
		},
		{
			NewFunc: armnetwork.NewExpressRouteCrossConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections",
		},
		{
			NewFunc: armnetwork.NewExpressRouteGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewExpressRouteLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links",
		},
		{
			NewFunc: armnetwork.NewExpressRoutePortAuthorizationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations",
		},
		{
			NewFunc: armnetwork.NewExpressRoutePortsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts",
		},
		{
			NewFunc: armnetwork.NewExpressRoutePortsLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations",
		},
		{
			NewFunc: armnetwork.NewExpressRouteProviderPortsLocationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteProviderPorts",
		},
		{
			NewFunc: armnetwork.NewExpressRouteServiceProvidersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders",
		},
		{
			NewFunc: armnetwork.NewFirewallPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies",
		},
		{
			NewFunc: armnetwork.NewFirewallPolicyIdpsSignaturesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/listIdpsSignatures",
		},
		{
			NewFunc: armnetwork.NewFirewallPolicyIdpsSignaturesFilterValuesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/listIdpsFilterOptions",
		},
		{
			NewFunc: armnetwork.NewFirewallPolicyIdpsSignaturesOverridesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides",
		},
		{
			NewFunc: armnetwork.NewFirewallPolicyRuleCollectionGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups",
		},
		{
			NewFunc: armnetwork.NewFlowLogsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs",
		},
		{
			NewFunc: armnetwork.NewGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups",
		},
		{
			NewFunc: armnetwork.NewHubRouteTablesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables",
		},
		{
			NewFunc: armnetwork.NewHubVirtualNetworkConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections",
		},
		{
			NewFunc: armnetwork.NewIPAllocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/IpAllocations",
		},
		{
			NewFunc: armnetwork.NewIPGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ipGroups",
		},
		{
			NewFunc: armnetwork.NewInboundNatRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules",
		},
		{
			NewFunc: armnetwork.NewInboundSecurityRuleClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewInterfaceIPConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations",
		},
		{
			NewFunc: armnetwork.NewInterfaceLoadBalancersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers",
		},
		{
			NewFunc: armnetwork.NewInterfaceTapConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations",
		},
		{
			NewFunc: armnetwork.NewInterfacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces",
		},
		{
			NewFunc: armnetwork.NewLoadBalancerBackendAddressPoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools",
		},
		{
			NewFunc: armnetwork.NewLoadBalancerFrontendIPConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations",
		},
		{
			NewFunc: armnetwork.NewLoadBalancerLoadBalancingRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules",
		},
		{
			NewFunc: armnetwork.NewLoadBalancerNetworkInterfacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces",
		},
		{
			NewFunc: armnetwork.NewLoadBalancerOutboundRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules",
		},
		{
			NewFunc: armnetwork.NewLoadBalancerProbesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes",
		},
		{
			NewFunc: armnetwork.NewLoadBalancersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers",
		},
		{
			NewFunc: armnetwork.NewLocalNetworkGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways",
		},
		{
			NewFunc: armnetwork.NewManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewManagementGroupNetworkManagerConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Network/networkManagerConnections",
		},
		{
			NewFunc: armnetwork.NewManagerCommitsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewManagerDeploymentStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/listDeploymentStatus",
		},
		{
			NewFunc: armnetwork.NewManagersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers",
		},
		{
			NewFunc: armnetwork.NewNatGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways",
		},
		{
			NewFunc: armnetwork.NewNatRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewP2SVPNGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways",
		},
		{
			NewFunc: armnetwork.NewPacketCapturesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures",
		},
		{
			NewFunc: armnetwork.NewPeerExpressRouteCircuitConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections",
		},
		{
			NewFunc: armnetwork.NewPrivateDNSZoneGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups",
		},
		{
			NewFunc: armnetwork.NewPrivateEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints",
		},
		{
			NewFunc: armnetwork.NewPrivateLinkServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices",
		},
		{
			NewFunc: armnetwork.NewProfilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles",
		},
		{
			NewFunc: armnetwork.NewPublicIPAddressesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses",
		},
		{
			NewFunc: armnetwork.NewPublicIPPrefixesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes",
		},
		{
			NewFunc: armnetwork.NewResourceNavigationLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks",
		},
		{
			NewFunc: armnetwork.NewRouteFilterRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewRouteFiltersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters",
		},
		{
			NewFunc: armnetwork.NewRouteTablesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables",
		},
		{
			NewFunc: armnetwork.NewRoutesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes",
		},
		{
			NewFunc: armnetwork.NewRoutingIntentClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent",
		},
		{
			NewFunc: armnetwork.NewScopeConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections",
		},
		{
			NewFunc: armnetwork.NewSecurityAdminConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations",
		},
		{
			NewFunc: armnetwork.NewSecurityGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups",
		},
		{
			NewFunc: armnetwork.NewSecurityPartnerProvidersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/securityPartnerProviders",
		},
		{
			NewFunc: armnetwork.NewSecurityRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules",
		},
		{
			NewFunc: armnetwork.NewServiceAssociationLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks",
		},
		{
			NewFunc: armnetwork.NewServiceEndpointPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies",
		},
		{
			NewFunc: armnetwork.NewServiceEndpointPolicyDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewServiceTagInformationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTagDetails",
		},
		{
			NewFunc: armnetwork.NewServiceTagsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTags",
		},
		{
			NewFunc: armnetwork.NewStaticMembersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups/{networkGroupName}/staticMembers",
		},
		{
			NewFunc: armnetwork.NewSubnetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets",
		},
		{
			NewFunc: armnetwork.NewSubscriptionNetworkManagerConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkManagerConnections",
		},
		{
			NewFunc: armnetwork.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages",
		},
		{
			NewFunc: armnetwork.NewVPNConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVPNGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways",
		},
		{
			NewFunc: armnetwork.NewVPNLinkConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVPNServerConfigurationsAssociatedWithVirtualWanClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnServerConfigurations",
		},
		{
			NewFunc: armnetwork.NewVPNServerConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations",
		},
		{
			NewFunc: armnetwork.NewVPNSiteLinkConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVPNSiteLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVPNSitesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites",
		},
		{
			NewFunc: armnetwork.NewVPNSitesConfigurationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVirtualApplianceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualApplianceSkus",
		},
		{
			NewFunc: armnetwork.NewVirtualApplianceSitesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites",
		},
		{
			NewFunc: armnetwork.NewVirtualAppliancesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances",
		},
		{
			NewFunc: armnetwork.NewVirtualHubBgpConnectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVirtualHubBgpConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections",
		},
		{
			NewFunc: armnetwork.NewVirtualHubIPConfigurationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations",
		},
		{
			NewFunc: armnetwork.NewVirtualHubRouteTableV2SClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables",
		},
		{
			NewFunc: armnetwork.NewVirtualHubsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs",
		},
		{
			NewFunc: armnetwork.NewVirtualNetworkGatewayConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections",
		},
		{
			NewFunc: armnetwork.NewVirtualNetworkGatewayNatRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVirtualNetworkGatewaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways",
		},
		{
			NewFunc: armnetwork.NewVirtualNetworkPeeringsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings",
		},
		{
			NewFunc: armnetwork.NewVirtualNetworkTapsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
		{
			NewFunc: armnetwork.NewVirtualNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks",
		},
		{
			NewFunc: armnetwork.NewVirtualRouterPeeringsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings",
		},
		{
			NewFunc: armnetwork.NewVirtualRoutersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualRouters",
		},
		{
			NewFunc: armnetwork.NewVirtualWansClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans",
		},
		{
			NewFunc: armnetwork.NewWatchersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers",
		},
		{
			NewFunc: armnetwork.NewWebApplicationFirewallPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies",
		},
		{
			NewFunc: armnetwork.NewWebCategoriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armnetwork())
}