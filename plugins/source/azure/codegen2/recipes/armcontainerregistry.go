// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

func Armcontainerregistry() []Table {
	tables := []Table{
		{
      Name: "pipeline_run",
      Struct: &armcontainerregistry.PipelineRun{},
      ResponseStruct: &armcontainerregistry.PipelineRunsClientListResponse{},
      Client: &armcontainerregistry.PipelineRunsClient{},
      ListFunc: (&armcontainerregistry.PipelineRunsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewPipelineRunsClient,
		},
		{
      Name: "token",
      Struct: &armcontainerregistry.Token{},
      ResponseStruct: &armcontainerregistry.TokensClientListResponse{},
      Client: &armcontainerregistry.TokensClient{},
      ListFunc: (&armcontainerregistry.TokensClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewTokensClient,
		},
		{
      Name: "connected_registry",
      Struct: &armcontainerregistry.ConnectedRegistry{},
      ResponseStruct: &armcontainerregistry.ConnectedRegistriesClientListResponse{},
      Client: &armcontainerregistry.ConnectedRegistriesClient{},
      ListFunc: (&armcontainerregistry.ConnectedRegistriesClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewConnectedRegistriesClient,
		},
		{
      Name: "import_pipeline",
      Struct: &armcontainerregistry.ImportPipeline{},
      ResponseStruct: &armcontainerregistry.ImportPipelinesClientListResponse{},
      Client: &armcontainerregistry.ImportPipelinesClient{},
      ListFunc: (&armcontainerregistry.ImportPipelinesClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewImportPipelinesClient,
		},
		{
      Name: "task_run",
      Struct: &armcontainerregistry.TaskRun{},
      ResponseStruct: &armcontainerregistry.TaskRunsClientListResponse{},
      Client: &armcontainerregistry.TaskRunsClient{},
      ListFunc: (&armcontainerregistry.TaskRunsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewTaskRunsClient,
		},
		{
      Name: "replication",
      Struct: &armcontainerregistry.Replication{},
      ResponseStruct: &armcontainerregistry.ReplicationsClientListResponse{},
      Client: &armcontainerregistry.ReplicationsClient{},
      ListFunc: (&armcontainerregistry.ReplicationsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewReplicationsClient,
		},
		{
      Name: "run",
      Struct: &armcontainerregistry.Run{},
      ResponseStruct: &armcontainerregistry.RunsClientListResponse{},
      Client: &armcontainerregistry.RunsClient{},
      ListFunc: (&armcontainerregistry.RunsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewRunsClient,
		},
		{
      Name: "scope_map",
      Struct: &armcontainerregistry.ScopeMap{},
      ResponseStruct: &armcontainerregistry.ScopeMapsClientListResponse{},
      Client: &armcontainerregistry.ScopeMapsClient{},
      ListFunc: (&armcontainerregistry.ScopeMapsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewScopeMapsClient,
		},
		{
      Name: "registry",
      Struct: &armcontainerregistry.Registry{},
      ResponseStruct: &armcontainerregistry.RegistriesClientListResponse{},
      Client: &armcontainerregistry.RegistriesClient{},
      ListFunc: (&armcontainerregistry.RegistriesClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewRegistriesClient,
		},
		{
      Name: "task",
      Struct: &armcontainerregistry.Task{},
      ResponseStruct: &armcontainerregistry.TasksClientListResponse{},
      Client: &armcontainerregistry.TasksClient{},
      ListFunc: (&armcontainerregistry.TasksClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewTasksClient,
		},
		{
      Name: "webhook",
      Struct: &armcontainerregistry.Webhook{},
      ResponseStruct: &armcontainerregistry.WebhooksClientListResponse{},
      Client: &armcontainerregistry.WebhooksClient{},
      ListFunc: (&armcontainerregistry.WebhooksClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewWebhooksClient,
		},
		{
      Name: "agent_pool",
      Struct: &armcontainerregistry.AgentPool{},
      ResponseStruct: &armcontainerregistry.AgentPoolsClientListResponse{},
      Client: &armcontainerregistry.AgentPoolsClient{},
      ListFunc: (&armcontainerregistry.AgentPoolsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewAgentPoolsClient,
		},
		{
      Name: "export_pipeline",
      Struct: &armcontainerregistry.ExportPipeline{},
      ResponseStruct: &armcontainerregistry.ExportPipelinesClientListResponse{},
      Client: &armcontainerregistry.ExportPipelinesClient{},
      ListFunc: (&armcontainerregistry.ExportPipelinesClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewExportPipelinesClient,
		},
		{
      Name: "operation_definition",
      Struct: &armcontainerregistry.OperationDefinition{},
      ResponseStruct: &armcontainerregistry.OperationsClientListResponse{},
      Client: &armcontainerregistry.OperationsClient{},
      ListFunc: (&armcontainerregistry.OperationsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewOperationsClient,
		},
		{
      Name: "private_endpoint_connection",
      Struct: &armcontainerregistry.PrivateEndpointConnection{},
      ResponseStruct: &armcontainerregistry.PrivateEndpointConnectionsClientListResponse{},
      Client: &armcontainerregistry.PrivateEndpointConnectionsClient{},
      ListFunc: (&armcontainerregistry.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewPrivateEndpointConnectionsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armcontainerregistry"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armcontainerregistry()...)
}