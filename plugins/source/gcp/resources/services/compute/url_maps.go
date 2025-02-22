// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func UrlMaps() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_url_maps",
		Resolver:  fetchUrlMaps,
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
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "default_route_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultRouteAction"),
			},
			{
				Name:     "default_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultService"),
			},
			{
				Name:     "default_url_redirect",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultUrlRedirect"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "header_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HeaderAction"),
			},
			{
				Name:     "host_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostRules"),
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
				Name:     "path_matchers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PathMatchers"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "tests",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tests"),
			},
		},
	}
}

func fetchUrlMaps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListUrlMapsRequest{
		Project: c.ProjectId,
	}
	it := c.Services.ComputeUrlMapsClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.UrlMaps

	}
	return nil
}
