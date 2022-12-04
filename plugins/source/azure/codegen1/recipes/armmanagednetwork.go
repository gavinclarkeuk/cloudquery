// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"

func Armmanagednetwork() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmanagednetwork.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
		},
		{
			NewFunc: armmanagednetwork.NewManagedNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
		},
		{
			NewFunc: armmanagednetwork.NewPeeringPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
		},
		{
			NewFunc: armmanagednetwork.NewScopeAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
		},
		{
			NewFunc: armmanagednetwork.NewGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmanagednetwork())
}