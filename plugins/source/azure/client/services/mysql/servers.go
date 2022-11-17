// Code generated by codegen; DO NOT EDIT.
package mysql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

type (
	RuntimePagerArmmysqlServersClientListByResourceGroupResponse = runtime.Pager[armmysql.ServersClientListByResourceGroupResponse]
	RuntimePagerArmmysqlServersClientListResponse                = runtime.Pager[armmysql.ServersClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/mysql/servers.go -source=servers.go ServersClient
type ServersClient interface {
	Get(context.Context, string, string, *armmysql.ServersClientGetOptions) (armmysql.ServersClientGetResponse, error)
	NewListByResourceGroupPager(string, *armmysql.ServersClientListByResourceGroupOptions) *RuntimePagerArmmysqlServersClientListByResourceGroupResponse
	NewListPager(*armmysql.ServersClientListOptions) *RuntimePagerArmmysqlServersClientListResponse
}
