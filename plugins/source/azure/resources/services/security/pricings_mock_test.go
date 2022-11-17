// Code generated by codegen; DO NOT EDIT.

package security

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/security"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/security"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildPricings(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockPricingsClient := mocks.NewMockPricingsClient(ctrl)

	var response api.PricingsClientListResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockPricingsClient.EXPECT().List(gomock.Any(), gomock.Any()).
		Return(response, nil).MinTimes(1)

	securityClient := &service.SecurityClient{
		PricingsClient: mockPricingsClient,
	}

	c := &client.Services{Security: securityClient}

	return c
}

func TestPricings(t *testing.T) {
	client.MockTestHelper(t, Pricings(), buildPricings)
}
