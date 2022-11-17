// Code generated by codegen; DO NOT EDIT.
package logic

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

type (
	RuntimePagerArmlogicWorkflowsClientListByResourceGroupResponse = runtime.Pager[armlogic.WorkflowsClientListByResourceGroupResponse]
	RuntimePagerArmlogicWorkflowsClientListBySubscriptionResponse  = runtime.Pager[armlogic.WorkflowsClientListBySubscriptionResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/logic/workflows.go -source=workflows.go WorkflowsClient
type WorkflowsClient interface {
	Get(context.Context, string, string, *armlogic.WorkflowsClientGetOptions) (armlogic.WorkflowsClientGetResponse, error)
	ListCallbackURL(context.Context, string, string, armlogic.GetCallbackURLParameters, *armlogic.WorkflowsClientListCallbackURLOptions) (armlogic.WorkflowsClientListCallbackURLResponse, error)
	ListSwagger(context.Context, string, string, *armlogic.WorkflowsClientListSwaggerOptions) (armlogic.WorkflowsClientListSwaggerResponse, error)
	NewListByResourceGroupPager(string, *armlogic.WorkflowsClientListByResourceGroupOptions) *RuntimePagerArmlogicWorkflowsClientListByResourceGroupResponse
	NewListBySubscriptionPager(*armlogic.WorkflowsClientListBySubscriptionOptions) *RuntimePagerArmlogicWorkflowsClientListBySubscriptionResponse
}
