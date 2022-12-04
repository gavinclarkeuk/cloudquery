// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"

func Armcontainerservice() []Table {
	tables := []Table{
		{
      Name: "managed_cluster",
      Struct: &armcontainerservice.ManagedCluster{},
      ResponseStruct: &armcontainerservice.ManagedClustersClientListResponse{},
      Client: &armcontainerservice.ManagedClustersClient{},
      ListFunc: (&armcontainerservice.ManagedClustersClient{}).NewListPager,
			NewFunc: armcontainerservice.NewManagedClustersClient,
		},
		{
      Name: "agent_pool",
      Struct: &armcontainerservice.AgentPool{},
      ResponseStruct: &armcontainerservice.AgentPoolsClientListResponse{},
      Client: &armcontainerservice.AgentPoolsClient{},
      ListFunc: (&armcontainerservice.AgentPoolsClient{}).NewListPager,
			NewFunc: armcontainerservice.NewAgentPoolsClient,
		},
		{
      Name: "operation_value",
      Struct: &armcontainerservice.OperationValue{},
      ResponseStruct: &armcontainerservice.OperationsClientListResponse{},
      Client: &armcontainerservice.OperationsClient{},
      ListFunc: (&armcontainerservice.OperationsClient{}).NewListPager,
			NewFunc: armcontainerservice.NewOperationsClient,
		},
		{
      Name: "snapshot",
      Struct: &armcontainerservice.Snapshot{},
      ResponseStruct: &armcontainerservice.SnapshotsClientListResponse{},
      Client: &armcontainerservice.SnapshotsClient{},
      ListFunc: (&armcontainerservice.SnapshotsClient{}).NewListPager,
			NewFunc: armcontainerservice.NewSnapshotsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armcontainerservice"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armcontainerservice()...)
}