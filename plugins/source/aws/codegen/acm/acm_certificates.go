// Code generated by codegen using template resource_list_describe.go.tpl; DO NOT EDIT.

package acm

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
)

func ACMCertificates() *schema.Table {
	return &schema.Table{
		Name:      "aws_acm_certificates",
		Resolver:  fetchACMCertificates,
		Multiplex: client.ServiceAccountRegionMultiplexer("acm"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "certificate_authority_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateAuthorityArn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "domain_validation_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainValidationOptions"),
			},
			{
				Name:     "extended_key_usages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedKeyUsages"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "imported_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ImportedAt"),
			},
			{
				Name:     "in_use_by",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("InUseBy"),
			},
			{
				Name:     "issued_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("IssuedAt"),
			},
			{
				Name:     "issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issuer"),
			},
			{
				Name:     "key_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyAlgorithm"),
			},
			{
				Name:     "key_usages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KeyUsages.Name"),
			},
			{
				Name:     "not_after",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotAfter"),
			},
			{
				Name:     "not_before",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotBefore"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "renewal_eligibility",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RenewalEligibility"),
			},
			{
				Name:     "renewal_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RenewalSummary"),
			},
			{
				Name:     "revocation_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RevocationReason"),
			},
			{
				Name:     "revoked_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RevokedAt"),
			},
			{
				Name:     "serial",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Serial"),
			},
			{
				Name:     "signature_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SignatureAlgorithm"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subject",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject"),
			},
			{
				Name:     "subject_alternative_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubjectAlternativeNames"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Description: `The tags that have been applied to the ACM certificate`,
			},
		},
	}
}

func fetchACMCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().ACM

	input := acm.ListCertificatesInput{}
	paginator := acm.NewListCertificatesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {

			return errors.WithStack(err)
		}

		for _, item := range output.CertificateSummaryList {

			do, err := svc.DescribeCertificate(ctx, &acm.DescribeCertificateInput{

				CertificateArn: item.CertificateArn,
			})
			if err != nil {

				if cl.IsNotFoundError(err) {
					continue
				}
				return errors.WithStack(err)
			}
			res <- do.Certificate
		}
	}
	return nil
}

func resolveACMCertificatesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(*types.CertificateDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().ACM
	out, err := svc.ListTagsForCertificate(ctx, &acm.ListTagsForCertificateInput{

		CertificateArn: item.CertificateArn,
	})
	if err != nil {

		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, client.TagsToMap(out.Tags)))
}
