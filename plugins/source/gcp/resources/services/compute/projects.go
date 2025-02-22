// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_projects",
		Resolver:  fetchProjects,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "common_instance_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CommonInstanceMetadata"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "default_network_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultNetworkTier"),
			},
			{
				Name:     "default_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultServiceAccount"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enabled_features",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledFeatures"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "quotas",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Quotas"),
			},
			{
				Name:     "usage_export_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UsageExportLocation"),
			},
			{
				Name:     "xpn_project_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("XpnProjectStatus"),
			},
		},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.GetProjectRequest{
		Project: c.ProjectId,
	}
	resp, err := c.Services.ComputeProjectsClient.Get(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- resp
	return nil
}
