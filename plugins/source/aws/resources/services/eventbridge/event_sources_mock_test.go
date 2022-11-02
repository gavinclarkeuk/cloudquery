// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEventbridgeEventSourcesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventBridgeClient(ctrl)
	object := types.EventSource{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEventSources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListEventSourcesOutput{
			EventSources: []types.EventSource{object},
		}, nil)

	return client.Services{
		EventBridge: m,
	}
}

func TestEventbridgeEventSources(t *testing.T) {
	client.AwsMockTestHelper(t, EventSources(), buildEventbridgeEventSourcesMock, client.TestOptions{})
}