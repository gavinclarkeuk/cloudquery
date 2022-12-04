// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"

func Armkubernetesconfiguration() []Table {
	tables := []Table{
		{
      Name: "source_control_configuration",
      Struct: &armkubernetesconfiguration.SourceControlConfiguration{},
      ResponseStruct: &armkubernetesconfiguration.SourceControlConfigurationsClientListResponse{},
      Client: &armkubernetesconfiguration.SourceControlConfigurationsClient{},
      ListFunc: (&armkubernetesconfiguration.SourceControlConfigurationsClient{}).NewListPager,
			NewFunc: armkubernetesconfiguration.NewSourceControlConfigurationsClient,
		},
		{
      Name: "extension",
      Struct: &armkubernetesconfiguration.Extension{},
      ResponseStruct: &armkubernetesconfiguration.ExtensionsClientListResponse{},
      Client: &armkubernetesconfiguration.ExtensionsClient{},
      ListFunc: (&armkubernetesconfiguration.ExtensionsClient{}).NewListPager,
			NewFunc: armkubernetesconfiguration.NewExtensionsClient,
		},
		{
      Name: "resource_provider_operation",
      Struct: &armkubernetesconfiguration.ResourceProviderOperation{},
      ResponseStruct: &armkubernetesconfiguration.OperationsClientListResponse{},
      Client: &armkubernetesconfiguration.OperationsClient{},
      ListFunc: (&armkubernetesconfiguration.OperationsClient{}).NewListPager,
			NewFunc: armkubernetesconfiguration.NewOperationsClient,
		},
		{
      Name: "operation_status_result",
      Struct: &armkubernetesconfiguration.OperationStatusResult{},
      ResponseStruct: &armkubernetesconfiguration.OperationStatusClientListResponse{},
      Client: &armkubernetesconfiguration.OperationStatusClient{},
      ListFunc: (&armkubernetesconfiguration.OperationStatusClient{}).NewListPager,
			NewFunc: armkubernetesconfiguration.NewOperationStatusClient,
		},
		{
      Name: "flux_configuration",
      Struct: &armkubernetesconfiguration.FluxConfiguration{},
      ResponseStruct: &armkubernetesconfiguration.FluxConfigurationsClientListResponse{},
      Client: &armkubernetesconfiguration.FluxConfigurationsClient{},
      ListFunc: (&armkubernetesconfiguration.FluxConfigurationsClient{}).NewListPager,
			NewFunc: armkubernetesconfiguration.NewFluxConfigurationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armkubernetesconfiguration"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armkubernetesconfiguration()...)
}