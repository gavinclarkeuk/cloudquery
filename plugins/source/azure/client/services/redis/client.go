// Code generated by codegen; DO NOT EDIT.
package redis

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

type (
	RuntimePagerArmredisClientListByResourceGroupResponse      = runtime.Pager[armredis.ClientListByResourceGroupResponse]
	RuntimePagerArmredisClientListBySubscriptionResponse       = runtime.Pager[armredis.ClientListBySubscriptionResponse]
	RuntimePagerArmredisClientListUpgradeNotificationsResponse = runtime.Pager[armredis.ClientListUpgradeNotificationsResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/redis/client.go -source=client.go Client
type Client interface {
	Get(context.Context, string, string, *armredis.ClientGetOptions) (armredis.ClientGetResponse, error)
	ListKeys(context.Context, string, string, *armredis.ClientListKeysOptions) (armredis.ClientListKeysResponse, error)
	NewListByResourceGroupPager(string, *armredis.ClientListByResourceGroupOptions) *RuntimePagerArmredisClientListByResourceGroupResponse
	NewListBySubscriptionPager(*armredis.ClientListBySubscriptionOptions) *RuntimePagerArmredisClientListBySubscriptionResponse
	NewListUpgradeNotificationsPager(string, string, float64, *armredis.ClientListUpgradeNotificationsOptions) *RuntimePagerArmredisClientListUpgradeNotificationsResponse
}
