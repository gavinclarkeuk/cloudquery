// Code generated by codegen; DO NOT EDIT.

package armapplicationinsights

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Annotation() *schema.Table {
	return &schema.Table{
		Name:      "azure_armapplicationinsights_annotation",
		Resolver:  fetchAnnotation,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "annotation_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AnnotationName"),
			},
			{
				Name:     "category",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Category"),
			},
			{
				Name:     "event_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EventTime"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "related_annotation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RelatedAnnotation"),
			},
		},
	}
}

func fetchAnnotation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmapplicationinsightsAnnotation
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
