// Code generated by codegen; DO NOT EDIT.
package frontdoor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

type (
	RuntimePagerArmfrontdoorFrontDoorsClientListByResourceGroupResponse = runtime.Pager[armfrontdoor.FrontDoorsClientListByResourceGroupResponse]
	RuntimePagerArmfrontdoorFrontDoorsClientListResponse                = runtime.Pager[armfrontdoor.FrontDoorsClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/frontdoor/front_doors.go -source=front_doors.go FrontDoorsClient
type FrontDoorsClient interface {
	Get(context.Context, string, string, *armfrontdoor.FrontDoorsClientGetOptions) (armfrontdoor.FrontDoorsClientGetResponse, error)
	NewListByResourceGroupPager(string, *armfrontdoor.FrontDoorsClientListByResourceGroupOptions) *RuntimePagerArmfrontdoorFrontDoorsClientListByResourceGroupResponse
	NewListPager(*armfrontdoor.FrontDoorsClientListOptions) *RuntimePagerArmfrontdoorFrontDoorsClientListResponse
}
