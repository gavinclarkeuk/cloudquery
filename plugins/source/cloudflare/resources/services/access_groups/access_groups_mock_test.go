package access_groups

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAccessGroups(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var accessGroup cloudflare.AccessGroup
	if err := faker.FakeObject(&accessGroup); err != nil {
		t.Fatal(err)
	}
	accessGroup.Include = []interface{}{"a"}
	accessGroup.Exclude = []interface{}{"b"}
	accessGroup.Require = []interface{}{"c"}
	mock.EXPECT().ZoneLevelAccessGroups(
		gomock.Any(),
		client.TestZoneID,
		cloudflare.PaginationOptions{
			Page:    1,
			PerPage: 200,
		},
	).Return(
		[]cloudflare.AccessGroup{accessGroup},
		cloudflare.ResultInfo{
			Page:    1,
			PerPage: 1,
			Count:   1,
			Total:   1,
		},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestAccessGroups(t *testing.T) {
	client.MockTestHelper(t, AccessGroups(), buildAccessGroups)
}
