// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"

func Armappcontainers() []Table {
	tables := []Table{
		{
      Name: "certificate",
      Struct: &armappcontainers.Certificate{},
      ResponseStruct: &armappcontainers.CertificatesClientListResponse{},
      Client: &armappcontainers.CertificatesClient{},
      ListFunc: (&armappcontainers.CertificatesClient{}).NewListPager,
			NewFunc: armappcontainers.NewCertificatesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/certificates",
		},
		{
      Name: "dapr_component",
      Struct: &armappcontainers.DaprComponent{},
      ResponseStruct: &armappcontainers.DaprComponentsClientListResponse{},
      Client: &armappcontainers.DaprComponentsClient{},
      ListFunc: (&armappcontainers.DaprComponentsClient{}).NewListPager,
			NewFunc: armappcontainers.NewDaprComponentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents",
		},
	}

	for i := range tables {
		tables[i].Service = "armappcontainers"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armappcontainers()...)
}