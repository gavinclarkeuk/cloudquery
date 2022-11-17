// Code generated by codegen; DO NOT EDIT.

package keyvault

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/keyvault"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/keyvault"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildManagedHsms(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockManagedHsmsClient := mocks.NewMockManagedHsmsClient(ctrl)

	var response api.ManagedHsmsClientListBySubscriptionResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockManagedHsmsClient.EXPECT().NewListBySubscriptionPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	keyvaultClient := &service.KeyvaultClient{
		ManagedHsmsClient: mockManagedHsmsClient,
	}

	c := &client.Services{Keyvault: keyvaultClient}

	return c
}

func TestManagedHsms(t *testing.T) {
	client.MockTestHelper(t, ManagedHsms(), buildManagedHsms)
}
