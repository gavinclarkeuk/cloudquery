// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"

func Armsecurityinsights() []Table {
	tables := []Table{
		{
      Name: "watchlist",
      Struct: &armsecurityinsights.Watchlist{},
      ResponseStruct: &armsecurityinsights.WatchlistsClientListResponse{},
      Client: &armsecurityinsights.WatchlistsClient{},
      ListFunc: (&armsecurityinsights.WatchlistsClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewWatchlistsClient,
		},
		// {
    //   Name: "data_connector_classification",
    //   Struct: &armsecurityinsights.DataConnectorClassification{},
    //   ResponseStruct: &armsecurityinsights.DataConnectorsClientListResponse{},
    //   Client: &armsecurityinsights.DataConnectorsClient{},
    //   ListFunc: (&armsecurityinsights.DataConnectorsClient{}).NewListPager,
		// 	NewFunc: armsecurityinsights.NewDataConnectorsClient,
		// },
		{
      Name: "relation",
      Struct: &armsecurityinsights.Relation{},
      ResponseStruct: &armsecurityinsights.IncidentRelationsClientListResponse{},
      Client: &armsecurityinsights.IncidentRelationsClient{},
      ListFunc: (&armsecurityinsights.IncidentRelationsClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewIncidentRelationsClient,
		},
		// {
    //   Name: "threat_intelligence_information_classification",
    //   Struct: &armsecurityinsights.ThreatIntelligenceInformationClassification{},
    //   ResponseStruct: &armsecurityinsights.ThreatIntelligenceIndicatorsClientListResponse{},
    //   Client: &armsecurityinsights.ThreatIntelligenceIndicatorsClient{},
    //   ListFunc: (&armsecurityinsights.ThreatIntelligenceIndicatorsClient{}).NewListPager,
		// 	NewFunc: armsecurityinsights.NewThreatIntelligenceIndicatorsClient,
		// },
		// {
    //   Name: "alert_rule_classification",
    //   Struct: &armsecurityinsights.AlertRuleClassification{},
    //   ResponseStruct: &armsecurityinsights.AlertRulesClientListResponse{},
    //   Client: &armsecurityinsights.AlertRulesClient{},
    //   ListFunc: (&armsecurityinsights.AlertRulesClient{}).NewListPager,
		// 	NewFunc: armsecurityinsights.NewAlertRulesClient,
		// },
		// {
    //   Name: "alert_rule_template_classification",
    //   Struct: &armsecurityinsights.AlertRuleTemplateClassification{},
    //   ResponseStruct: &armsecurityinsights.AlertRuleTemplatesClientListResponse{},
    //   Client: &armsecurityinsights.AlertRuleTemplatesClient{},
    //   ListFunc: (&armsecurityinsights.AlertRuleTemplatesClient{}).NewListPager,
		// 	NewFunc: armsecurityinsights.NewAlertRuleTemplatesClient,
		// },
		{
      Name: "incident",
      Struct: &armsecurityinsights.Incident{},
      ResponseStruct: &armsecurityinsights.IncidentsClientListResponse{},
      Client: &armsecurityinsights.IncidentsClient{},
      ListFunc: (&armsecurityinsights.IncidentsClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewIncidentsClient,
		},
		{
      Name: "watchlist_item",
      Struct: &armsecurityinsights.WatchlistItem{},
      ResponseStruct: &armsecurityinsights.WatchlistItemsClientListResponse{},
      Client: &armsecurityinsights.WatchlistItemsClient{},
      ListFunc: (&armsecurityinsights.WatchlistItemsClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewWatchlistItemsClient,
		},
		{
      Name: "automation_rule",
      Struct: &armsecurityinsights.AutomationRule{},
      ResponseStruct: &armsecurityinsights.AutomationRulesClientListResponse{},
      Client: &armsecurityinsights.AutomationRulesClient{},
      ListFunc: (&armsecurityinsights.AutomationRulesClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewAutomationRulesClient,
		},
		{
      Name: "bookmark",
      Struct: &armsecurityinsights.Bookmark{},
      ResponseStruct: &armsecurityinsights.BookmarksClientListResponse{},
      Client: &armsecurityinsights.BookmarksClient{},
      ListFunc: (&armsecurityinsights.BookmarksClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewBookmarksClient,
		},
		{
      Name: "incident_comment",
      Struct: &armsecurityinsights.IncidentComment{},
      ResponseStruct: &armsecurityinsights.IncidentCommentsClientListResponse{},
      Client: &armsecurityinsights.IncidentCommentsClient{},
      ListFunc: (&armsecurityinsights.IncidentCommentsClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewIncidentCommentsClient,
		},
		{
      Name: "operation",
      Struct: &armsecurityinsights.Operation{},
      ResponseStruct: &armsecurityinsights.OperationsClientListResponse{},
      Client: &armsecurityinsights.OperationsClient{},
      ListFunc: (&armsecurityinsights.OperationsClient{}).NewListPager,
			NewFunc: armsecurityinsights.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armsecurityinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armsecurityinsights()...)
}