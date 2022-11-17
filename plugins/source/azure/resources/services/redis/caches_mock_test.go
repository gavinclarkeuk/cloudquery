// Code generated by codegen; DO NOT EDIT.

package redis

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/redis"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/redis"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCaches(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockClient := mocks.NewMockClient(ctrl)

	var response api.ClientListBySubscriptionResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockClient.EXPECT().NewListBySubscriptionPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	redisClient := &service.RedisClient{
		Client: mockClient,
	}

	c := &client.Services{Redis: redisClient}

	return c
}

func TestCaches(t *testing.T) {
	client.MockTestHelper(t, Caches(), buildCaches)
}
