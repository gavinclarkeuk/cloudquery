// Code generated by codegen; DO NOT EDIT.
package policy

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

type (
	RuntimePagerArmpolicyAssignmentsClientListForManagementGroupResponse = runtime.Pager[armpolicy.AssignmentsClientListForManagementGroupResponse]
	RuntimePagerArmpolicyAssignmentsClientListForResourceGroupResponse   = runtime.Pager[armpolicy.AssignmentsClientListForResourceGroupResponse]
	RuntimePagerArmpolicyAssignmentsClientListForResourceResponse        = runtime.Pager[armpolicy.AssignmentsClientListForResourceResponse]
	RuntimePagerArmpolicyAssignmentsClientListResponse                   = runtime.Pager[armpolicy.AssignmentsClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/policy/assignments.go -source=assignments.go AssignmentsClient
type AssignmentsClient interface {
	Get(context.Context, string, string, *armpolicy.AssignmentsClientGetOptions) (armpolicy.AssignmentsClientGetResponse, error)
	GetByID(context.Context, string, *armpolicy.AssignmentsClientGetByIDOptions) (armpolicy.AssignmentsClientGetByIDResponse, error)
	NewListForManagementGroupPager(string, *armpolicy.AssignmentsClientListForManagementGroupOptions) *RuntimePagerArmpolicyAssignmentsClientListForManagementGroupResponse
	NewListForResourceGroupPager(string, *armpolicy.AssignmentsClientListForResourceGroupOptions) *RuntimePagerArmpolicyAssignmentsClientListForResourceGroupResponse
	NewListForResourcePager(string, string, string, string, string, *armpolicy.AssignmentsClientListForResourceOptions) *RuntimePagerArmpolicyAssignmentsClientListForResourceResponse
	NewListPager(*armpolicy.AssignmentsClientListOptions) *RuntimePagerArmpolicyAssignmentsClientListResponse
}
