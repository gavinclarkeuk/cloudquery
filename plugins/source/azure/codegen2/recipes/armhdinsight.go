// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"

func Armhdinsight() []Table {
	tables := []Table{
		{
      Name: "cluster",
      Struct: &armhdinsight.Cluster{},
      ResponseStruct: &armhdinsight.ClustersClientListResponse{},
      Client: &armhdinsight.ClustersClient{},
      ListFunc: (&armhdinsight.ClustersClient{}).NewListPager,
			NewFunc: armhdinsight.NewClustersClient,
		},
		{
      Name: "operation",
      Struct: &armhdinsight.Operation{},
      ResponseStruct: &armhdinsight.OperationsClientListResponse{},
      Client: &armhdinsight.OperationsClient{},
      ListFunc: (&armhdinsight.OperationsClient{}).NewListPager,
			NewFunc: armhdinsight.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armhdinsight"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armhdinsight()...)
}