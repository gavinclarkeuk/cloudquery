// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"

func Armblueprint() []Table {
	tables := []Table{
		{
      Name: "assignment_operation",
      Struct: &armblueprint.AssignmentOperation{},
      ResponseStruct: &armblueprint.AssignmentOperationsClientListResponse{},
      Client: &armblueprint.AssignmentOperationsClient{},
      ListFunc: (&armblueprint.AssignmentOperationsClient{}).NewListPager,
			NewFunc: armblueprint.NewAssignmentOperationsClient,
			URL: "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations",
		},
		{
      Name: "assignment",
      Struct: &armblueprint.Assignment{},
      ResponseStruct: &armblueprint.AssignmentsClientListResponse{},
      Client: &armblueprint.AssignmentsClient{},
      ListFunc: (&armblueprint.AssignmentsClient{}).NewListPager,
			NewFunc: armblueprint.NewAssignmentsClient,
			URL: "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments",
		},
		{
      Name: "blueprint",
      Struct: &armblueprint.Blueprint{},
      ResponseStruct: &armblueprint.BlueprintsClientListResponse{},
      Client: &armblueprint.BlueprintsClient{},
      ListFunc: (&armblueprint.BlueprintsClient{}).NewListPager,
			NewFunc: armblueprint.NewBlueprintsClient,
			URL: "/{resourceScope}/providers/Microsoft.Blueprint/blueprints",
		},
		{
      Name: "published_blueprint",
      Struct: &armblueprint.PublishedBlueprint{},
      ResponseStruct: &armblueprint.PublishedBlueprintsClientListResponse{},
      Client: &armblueprint.PublishedBlueprintsClient{},
      ListFunc: (&armblueprint.PublishedBlueprintsClient{}).NewListPager,
			NewFunc: armblueprint.NewPublishedBlueprintsClient,
			URL: "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions",
		},
	}

	for i := range tables {
		tables[i].Service = "armblueprint"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armblueprint()...)
}