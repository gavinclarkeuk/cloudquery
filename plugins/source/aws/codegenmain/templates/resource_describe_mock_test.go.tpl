// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"testing"

  "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

{{range .Imports}}	"{{.}}"
{{end}}
)

func build{{.AWSService}}{{.AWSSubService | ToCamel}}(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMock{{.AWSService}}Client(ctrl)

	var list types.{{.ItemName}}Summary
	if err := faker.FakeData(&list); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().{{.ListFunctionName}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.{{.ListFunctionName}}Input{},
		gomock.Any(),
	).Return(
		&{{.AWSService | ToLower}}.{{.ListFunctionName}}Output{
		  {{.ItemName}}SummaryList: []types.{{.ItemName}}Summary{list},
    },
		nil,
	)

	var detail types.{{.ItemName}}Detail
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}
	detail.{{.DescribeFieldName}} = list.{{.DescribeFieldName}}
	mock.EXPECT().{{.DescribeFunctionName}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.{{.DescribeFunctionName}}Input{
		  {{.DescribeFieldName}}: list.{{.DescribeFieldName}},
		},
		gomock.Any(),
	).Return(
		&{{.AWSService | ToLower}}.{{.DescribeFunctionName}}Output{
		  {{.ItemName}}: &detail,
    },
		nil,
	)

{{if .HasTags}}
	mock.EXPECT().ListTagsFor{{.ItemName}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Input{
		  {{.DescribeFieldName}}: detail.{{.DescribeFieldName}},
    },
	).Return(
		&{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Output{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
{{end}}
	return client.Services{
	  {{.AWSService}}: mock,
  }
}

func {{.TestFuncName}}(t *testing.T) {
	client.AwsMockTestHelper(t, {{.TableFuncName}}(), {{.MockFuncName}}, client.TestOptions{})
}
