// Code generated by codegen; DO NOT EDIT.

package storage

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/storage/v1"
	// k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createCsiDrivers(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.CSIDriver{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockCSIDriverInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.CSIDriverList{Items: []resource.CSIDriver{r}}, nil,
	)

	serviceClient := resourcemock.NewMockStorageV1Interface(ctrl)

	serviceClient.EXPECT().CSIDrivers().Return(resourceClient)

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().StorageV1().Return(serviceClient)

	return client
}

func TestCsiDrivers(t *testing.T) {
	client.K8sMockTestHelper(t, CsiDrivers(), createCsiDrivers)
}
