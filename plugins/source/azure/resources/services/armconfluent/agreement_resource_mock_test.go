// Code generated by codegen; DO NOT EDIT.

package armconfluent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createAgreementResource() (*client.Services, error) {
	var item armconfluent.MarketplaceAgreementsClientListResponse
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
	svc, err := armconfluent.NewMarketplaceAgreementsClient(client.TestSubscription, &client.MockCreds{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: ts.Client(),
		},
	})
	if err != nil {
		return nil, err
	}
	return &client.Services{
		ArmconfluentAgreementResource: svc,
	}, nil
}

func TestAgreementResource(t *testing.T) {
	client.MockTestHelper(t, AgreementResource(), createAgreementResource)
}
