package users

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildUsersMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockUsersAPIClient(ctrl)
	services := client.DatadogServices{
		UsersAPI: m,
	}

	var users datadogV2.UsersResponse
	err := faker.FakeObject(&users)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUsers(gomock.Any()).Return(users, nil, nil)

	var permissions datadogV2.PermissionsResponse
	err = faker.FakeObject(&permissions)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPermissions(gomock.Any(), gomock.Any()).Return(permissions, nil, nil)

	var orgs datadogV2.UserResponse
	err = faker.FakeObject(&orgs)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserOrganizations(gomock.Any(), gomock.Any()).Return(orgs, nil, nil)

	return services
}

func TestUsers(t *testing.T) {
	client.DatadogMockTestHelper(t, Users(), buildUsersMock, client.TestOptions{})
}
