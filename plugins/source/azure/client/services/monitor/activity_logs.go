// Code generated by codegen; DO NOT EDIT.
package monitor

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

type (
	RuntimePagerArmmonitorActivityLogsClientListResponse = runtime.Pager[armmonitor.ActivityLogsClientListResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/monitor/activity_logs.go -source=activity_logs.go ActivityLogsClient
type ActivityLogsClient interface {
	NewListPager(string, *armmonitor.ActivityLogsClientListOptions) *RuntimePagerArmmonitorActivityLogsClientListResponse
}
