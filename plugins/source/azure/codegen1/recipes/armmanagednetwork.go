// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"

func Armmanagednetwork() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmanagednetwork.NewGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
			URL: "",
		},
		{
			NewFunc: armmanagednetwork.NewManagedNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
			URL: "",
		},
		{
			NewFunc: armmanagednetwork.NewPeeringPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
			URL: "",
		},
		{
			NewFunc: armmanagednetwork.NewScopeAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
			URL: "/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmanagednetwork())
}