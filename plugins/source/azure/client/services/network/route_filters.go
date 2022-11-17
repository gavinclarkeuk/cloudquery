// Code generated by codegen; DO NOT EDIT.
package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

type (
	RuntimePagerArmnetworkRouteFiltersClientListByResourceGroupResponse = runtime.Pager[armnetwork.RouteFiltersClientListByResourceGroupResponse]
	RuntimePagerArmnetworkRouteFiltersClientListResponse                = runtime.Pager[armnetwork.RouteFiltersClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/network/route_filters.go -source=route_filters.go RouteFiltersClient
type RouteFiltersClient interface {
	Get(context.Context, string, string, *armnetwork.RouteFiltersClientGetOptions) (armnetwork.RouteFiltersClientGetResponse, error)
	NewListByResourceGroupPager(string, *armnetwork.RouteFiltersClientListByResourceGroupOptions) *RuntimePagerArmnetworkRouteFiltersClientListByResourceGroupResponse
	NewListPager(*armnetwork.RouteFiltersClientListOptions) *RuntimePagerArmnetworkRouteFiltersClientListResponse
}
