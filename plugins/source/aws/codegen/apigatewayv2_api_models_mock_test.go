// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func buildApigatewayv2ApiModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockApigatewayv2Client(ctrl)

	item := types.Model{}
	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetModelsOutput{
			Items: []types.Model{item},
		}, nil)
	return client.Services{
		Apigatewayv2: mock,
	}
}

func TestApigatewayv2ApiModels(t *testing.T) {
	client.AwsMockTestHelper(t, Apigatewayv2ApiModels(), buildApigatewayv2ApiModels, client.TestOptions{})
}
