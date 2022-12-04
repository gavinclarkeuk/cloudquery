// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"

func Armappplatform() []Table {
	tables := []Table{
		{
      Name: "app_resource",
      Struct: &armappplatform.AppResource{},
      ResponseStruct: &armappplatform.AppsClientListResponse{},
      Client: &armappplatform.AppsClient{},
      ListFunc: (&armappplatform.AppsClient{}).NewListPager,
			NewFunc: armappplatform.NewAppsClient,
		},
		{
      Name: "build_service_agent_pool_resource",
      Struct: &armappplatform.BuildServiceAgentPoolResource{},
      ResponseStruct: &armappplatform.BuildServiceAgentPoolClientListResponse{},
      Client: &armappplatform.BuildServiceAgentPoolClient{},
      ListFunc: (&armappplatform.BuildServiceAgentPoolClient{}).NewListPager,
			NewFunc: armappplatform.NewBuildServiceAgentPoolClient,
		},
		{
      Name: "service_registry_resource",
      Struct: &armappplatform.ServiceRegistryResource{},
      ResponseStruct: &armappplatform.ServiceRegistriesClientListResponse{},
      Client: &armappplatform.ServiceRegistriesClient{},
      ListFunc: (&armappplatform.ServiceRegistriesClient{}).NewListPager,
			NewFunc: armappplatform.NewServiceRegistriesClient,
		},
		{
      Name: "buildpack_binding_resource",
      Struct: &armappplatform.BuildpackBindingResource{},
      ResponseStruct: &armappplatform.BuildpackBindingClientListResponse{},
      Client: &armappplatform.BuildpackBindingClient{},
      ListFunc: (&armappplatform.BuildpackBindingClient{}).NewListPager,
			NewFunc: armappplatform.NewBuildpackBindingClient,
		},
		{
      Name: "deployment_resource",
      Struct: &armappplatform.DeploymentResource{},
      ResponseStruct: &armappplatform.DeploymentsClientListResponse{},
      Client: &armappplatform.DeploymentsClient{},
      ListFunc: (&armappplatform.DeploymentsClient{}).NewListPager,
			NewFunc: armappplatform.NewDeploymentsClient,
		},
		{
      Name: "operation_detail",
      Struct: &armappplatform.OperationDetail{},
      ResponseStruct: &armappplatform.OperationsClientListResponse{},
      Client: &armappplatform.OperationsClient{},
      ListFunc: (&armappplatform.OperationsClient{}).NewListPager,
			NewFunc: armappplatform.NewOperationsClient,
		},
		{
      Name: "service_resource",
      Struct: &armappplatform.ServiceResource{},
      ResponseStruct: &armappplatform.ServicesClientListResponse{},
      Client: &armappplatform.ServicesClient{},
      ListFunc: (&armappplatform.ServicesClient{}).NewListPager,
			NewFunc: armappplatform.NewServicesClient,
		},
		{
      Name: "binding_resource",
      Struct: &armappplatform.BindingResource{},
      ResponseStruct: &armappplatform.BindingsClientListResponse{},
      Client: &armappplatform.BindingsClient{},
      ListFunc: (&armappplatform.BindingsClient{}).NewListPager,
			NewFunc: armappplatform.NewBindingsClient,
		},
		{
      Name: "builder_resource",
      Struct: &armappplatform.BuilderResource{},
      ResponseStruct: &armappplatform.BuildServiceBuilderClientListResponse{},
      Client: &armappplatform.BuildServiceBuilderClient{},
      ListFunc: (&armappplatform.BuildServiceBuilderClient{}).NewListPager,
			NewFunc: armappplatform.NewBuildServiceBuilderClient,
		},
		{
      Name: "certificate_resource",
      Struct: &armappplatform.CertificateResource{},
      ResponseStruct: &armappplatform.CertificatesClientListResponse{},
      Client: &armappplatform.CertificatesClient{},
      ListFunc: (&armappplatform.CertificatesClient{}).NewListPager,
			NewFunc: armappplatform.NewCertificatesClient,
		},
		{
      Name: "configuration_service_resource",
      Struct: &armappplatform.ConfigurationServiceResource{},
      ResponseStruct: &armappplatform.ConfigurationServicesClientListResponse{},
      Client: &armappplatform.ConfigurationServicesClient{},
      ListFunc: (&armappplatform.ConfigurationServicesClient{}).NewListPager,
			NewFunc: armappplatform.NewConfigurationServicesClient,
		},
		{
      Name: "custom_domain_resource",
      Struct: &armappplatform.CustomDomainResource{},
      ResponseStruct: &armappplatform.CustomDomainsClientListResponse{},
      Client: &armappplatform.CustomDomainsClient{},
      ListFunc: (&armappplatform.CustomDomainsClient{}).NewListPager,
			NewFunc: armappplatform.NewCustomDomainsClient,
		},
		{
      Name: "resource_sku",
      Struct: &armappplatform.ResourceSKU{},
      ResponseStruct: &armappplatform.SKUsClientListResponse{},
      Client: &armappplatform.SKUsClient{},
      ListFunc: (&armappplatform.SKUsClient{}).NewListPager,
			NewFunc: armappplatform.NewSKUsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armappplatform"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armappplatform()...)
}