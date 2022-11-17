// Code generated by codegen; DO NOT EDIT.
package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

type (
	RuntimePagerArmsqlServerSecurityAlertPoliciesClientListByServerResponse = runtime.Pager[armsql.ServerSecurityAlertPoliciesClientListByServerResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/sql/server_security_alert_policies.go -source=server_security_alert_policies.go ServerSecurityAlertPoliciesClient
type ServerSecurityAlertPoliciesClient interface {
	Get(context.Context, string, string, armsql.SecurityAlertPolicyName, *armsql.ServerSecurityAlertPoliciesClientGetOptions) (armsql.ServerSecurityAlertPoliciesClientGetResponse, error)
	NewListByServerPager(string, string, *armsql.ServerSecurityAlertPoliciesClientListByServerOptions) *RuntimePagerArmsqlServerSecurityAlertPoliciesClientListByServerResponse
}
