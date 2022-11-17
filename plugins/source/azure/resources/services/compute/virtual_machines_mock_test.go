// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/compute"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/compute"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildVirtualMachines(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockVirtualMachinesClient := mocks.NewMockVirtualMachinesClient(ctrl)

	var response api.VirtualMachinesClientListAllResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockVirtualMachinesClient.EXPECT().NewListAllPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	computeClient := &service.ComputeClient{
		VirtualMachinesClient: mockVirtualMachinesClient,
	}

	c := &client.Services{Compute: computeClient}

	buildVirtualMachineScaleSets(t, ctrl, c)

	return c
}

func TestVirtualMachines(t *testing.T) {
	client.MockTestHelper(t, VirtualMachines(), buildVirtualMachines)
}
