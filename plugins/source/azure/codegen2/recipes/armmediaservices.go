// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"

func Armmediaservices() []Table {
	tables := []Table{
		{
      Name: "account_filter",
      Struct: &armmediaservices.AccountFilter{},
      ResponseStruct: &armmediaservices.AccountFiltersClientListResponse{},
      Client: &armmediaservices.AccountFiltersClient{},
      ListFunc: (&armmediaservices.AccountFiltersClient{}).NewListPager,
			NewFunc: armmediaservices.NewAccountFiltersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters",
		},
		{
      Name: "asset_filter",
      Struct: &armmediaservices.AssetFilter{},
      ResponseStruct: &armmediaservices.AssetFiltersClientListResponse{},
      Client: &armmediaservices.AssetFiltersClient{},
      ListFunc: (&armmediaservices.AssetFiltersClient{}).NewListPager,
			NewFunc: armmediaservices.NewAssetFiltersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters",
		},
		{
      Name: "asset",
      Struct: &armmediaservices.Asset{},
      ResponseStruct: &armmediaservices.AssetsClientListResponse{},
      Client: &armmediaservices.AssetsClient{},
      ListFunc: (&armmediaservices.AssetsClient{}).NewListPager,
			NewFunc: armmediaservices.NewAssetsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets",
		},
		{
      Name: "media_service",
      Struct: &armmediaservices.MediaService{},
      ResponseStruct: &armmediaservices.ClientListResponse{},
      Client: &armmediaservices.Client{},
      ListFunc: (&armmediaservices.Client{}).NewListPager,
			NewFunc: armmediaservices.NewClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices",
		},
		{
      Name: "content_key_policy",
      Struct: &armmediaservices.ContentKeyPolicy{},
      ResponseStruct: &armmediaservices.ContentKeyPoliciesClientListResponse{},
      Client: &armmediaservices.ContentKeyPoliciesClient{},
      ListFunc: (&armmediaservices.ContentKeyPoliciesClient{}).NewListPager,
			NewFunc: armmediaservices.NewContentKeyPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies",
		},
		{
      Name: "job",
      Struct: &armmediaservices.Job{},
      ResponseStruct: &armmediaservices.JobsClientListResponse{},
      Client: &armmediaservices.JobsClient{},
      ListFunc: (&armmediaservices.JobsClient{}).NewListPager,
			NewFunc: armmediaservices.NewJobsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs",
		},
		{
      Name: "live_event",
      Struct: &armmediaservices.LiveEvent{},
      ResponseStruct: &armmediaservices.LiveEventsClientListResponse{},
      Client: &armmediaservices.LiveEventsClient{},
      ListFunc: (&armmediaservices.LiveEventsClient{}).NewListPager,
			NewFunc: armmediaservices.NewLiveEventsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents",
		},
		{
      Name: "live_output",
      Struct: &armmediaservices.LiveOutput{},
      ResponseStruct: &armmediaservices.LiveOutputsClientListResponse{},
      Client: &armmediaservices.LiveOutputsClient{},
      ListFunc: (&armmediaservices.LiveOutputsClient{}).NewListPager,
			NewFunc: armmediaservices.NewLiveOutputsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs",
		},
		{
      Name: "streaming_endpoint",
      Struct: &armmediaservices.StreamingEndpoint{},
      ResponseStruct: &armmediaservices.StreamingEndpointsClientListResponse{},
      Client: &armmediaservices.StreamingEndpointsClient{},
      ListFunc: (&armmediaservices.StreamingEndpointsClient{}).NewListPager,
			NewFunc: armmediaservices.NewStreamingEndpointsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints",
		},
		{
      Name: "streaming_locator",
      Struct: &armmediaservices.StreamingLocator{},
      ResponseStruct: &armmediaservices.StreamingLocatorsClientListResponse{},
      Client: &armmediaservices.StreamingLocatorsClient{},
      ListFunc: (&armmediaservices.StreamingLocatorsClient{}).NewListPager,
			NewFunc: armmediaservices.NewStreamingLocatorsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators",
		},
		{
      Name: "streaming_policy",
      Struct: &armmediaservices.StreamingPolicy{},
      ResponseStruct: &armmediaservices.StreamingPoliciesClientListResponse{},
      Client: &armmediaservices.StreamingPoliciesClient{},
      ListFunc: (&armmediaservices.StreamingPoliciesClient{}).NewListPager,
			NewFunc: armmediaservices.NewStreamingPoliciesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies",
		},
		{
      Name: "asset_track",
      Struct: &armmediaservices.AssetTrack{},
      ResponseStruct: &armmediaservices.TracksClientListResponse{},
      Client: &armmediaservices.TracksClient{},
      ListFunc: (&armmediaservices.TracksClient{}).NewListPager,
			NewFunc: armmediaservices.NewTracksClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks",
		},
		{
      Name: "transform",
      Struct: &armmediaservices.Transform{},
      ResponseStruct: &armmediaservices.TransformsClientListResponse{},
      Client: &armmediaservices.TransformsClient{},
      ListFunc: (&armmediaservices.TransformsClient{}).NewListPager,
			NewFunc: armmediaservices.NewTransformsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms",
		},
	}

	for i := range tables {
		tables[i].Service = "armmediaservices"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmediaservices()...)
}