// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"

func Armcosmos() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcosmos.NewMongoDBResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableSQLResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewSQLResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewDatabaseAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewNotebookWorkspacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableDatabaseAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewDatabaseClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCassandraDataCentersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPartitionKeyRangeIDClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableMongodbDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCassandraClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPartitionKeyRangeIDRegionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableMongodbResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCollectionRegionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPercentileSourceTargetClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPercentileTargetClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableMongodbCollectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableSQLDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewTableResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCollectionPartitionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewRestorableSQLContainersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCollectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewGremlinResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewPercentileClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCollectionPartitionRegionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewDatabaseAccountRegionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
		{
			NewFunc: armcosmos.NewCassandraResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcosmos())
}