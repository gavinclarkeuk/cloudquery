// Code generated by codegen; DO NOT EDIT.

package mysql

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/mysql"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mysql"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildServers(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockServersClient := mocks.NewMockServersClient(ctrl)

	var response api.ServersClientListResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockServersClient.EXPECT().NewListPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	mysqlClient := &service.MysqlClient{
		ServersClient: mockServersClient,
	}

	c := &client.Services{Mysql: mysqlClient}

	buildConfigurations(t, ctrl, c)

	return c
}

func TestServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), buildServers)
}
