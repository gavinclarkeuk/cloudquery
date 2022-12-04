// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"

func Armeducation() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armeducation.NewManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
		{
			NewFunc: armeducation.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
		{
			NewFunc: armeducation.NewStudentLabsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
		{
			NewFunc: armeducation.NewGrantsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
		{
			NewFunc: armeducation.NewJoinRequestsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
		{
			NewFunc: armeducation.NewLabsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
		{
			NewFunc: armeducation.NewStudentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armeducation())
}