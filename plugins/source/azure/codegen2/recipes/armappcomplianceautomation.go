// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"

func Armappcomplianceautomation() []Table {
	tables := []Table{
		{
      Name: "report_resource",
      Struct: &armappcomplianceautomation.ReportResource{},
      ResponseStruct: &armappcomplianceautomation.ReportsClientListResponse{},
      Client: &armappcomplianceautomation.ReportsClient{},
      ListFunc: (&armappcomplianceautomation.ReportsClient{}).NewListPager,
			NewFunc: armappcomplianceautomation.NewReportsClient,
			URL: "/providers/Microsoft.AppComplianceAutomation/reports",
		},
		{
      Name: "snapshot_resource",
      Struct: &armappcomplianceautomation.SnapshotResource{},
      ResponseStruct: &armappcomplianceautomation.SnapshotsClientListResponse{},
      Client: &armappcomplianceautomation.SnapshotsClient{},
      ListFunc: (&armappcomplianceautomation.SnapshotsClient{}).NewListPager,
			NewFunc: armappcomplianceautomation.NewSnapshotsClient,
			URL: "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/snapshots",
		},
	}

	for i := range tables {
		tables[i].Service = "armappcomplianceautomation"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armappcomplianceautomation()...)
}