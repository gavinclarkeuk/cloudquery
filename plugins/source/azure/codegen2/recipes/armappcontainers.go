// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"

func Armappcontainers() []Table {
	tables := []Table{
		{
      Name: "dapr_component",
      Struct: &armappcontainers.DaprComponent{},
      ResponseStruct: &armappcontainers.DaprComponentsClientListResponse{},
      Client: &armappcontainers.DaprComponentsClient{},
      ListFunc: (&armappcontainers.DaprComponentsClient{}).NewListPager,
			NewFunc: armappcontainers.NewDaprComponentsClient,
		},
		{
      Name: "operation_detail",
      Struct: &armappcontainers.OperationDetail{},
      ResponseStruct: &armappcontainers.OperationsClientListResponse{},
      Client: &armappcontainers.OperationsClient{},
      ListFunc: (&armappcontainers.OperationsClient{}).NewListPager,
			NewFunc: armappcontainers.NewOperationsClient,
		},
		{
      Name: "certificate",
      Struct: &armappcontainers.Certificate{},
      ResponseStruct: &armappcontainers.CertificatesClientListResponse{},
      Client: &armappcontainers.CertificatesClient{},
      ListFunc: (&armappcontainers.CertificatesClient{}).NewListPager,
			NewFunc: armappcontainers.NewCertificatesClient,
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