// Code generated by codegen; DO NOT EDIT.

package armdigitaltwins

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createDescription() (*client.Services, error) {
	var item armdigitaltwins.ClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
	emptyStr := ""
	item.NextLink = &emptyStr
	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	cloud.AzurePublic.Services[cloud.ResourceManager] = cloud.ServiceConfiguration{
		Endpoint: ts.URL,
		Audience: "test",
	}
	svc, err := armdigitaltwins.NewClient(client.TestSubscription, &client.MockCreds{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: ts.Client(),
		},
	})
	if err != nil {
		return nil, err
	}
	return &client.Services{
		ArmdigitaltwinsDescription: svc,
	}, nil
}

func TestDescription(t *testing.T) {
	client.MockTestHelper(t, Description(), createDescription)
}
