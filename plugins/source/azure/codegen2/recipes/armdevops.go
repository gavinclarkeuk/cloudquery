// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"

func Armdevops() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armdevops.Operation{},
      ResponseStruct: &armdevops.OperationsClientListResponse{},
      Client: &armdevops.OperationsClient{},
      ListFunc: (&armdevops.OperationsClient{}).NewListPager,
			NewFunc: armdevops.NewOperationsClient,
		},
		{
      Name: "pipeline_template_definition",
      Struct: &armdevops.PipelineTemplateDefinition{},
      ResponseStruct: &armdevops.PipelineTemplateDefinitionsClientListResponse{},
      Client: &armdevops.PipelineTemplateDefinitionsClient{},
      ListFunc: (&armdevops.PipelineTemplateDefinitionsClient{}).NewListPager,
			NewFunc: armdevops.NewPipelineTemplateDefinitionsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdevops"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdevops()...)
}