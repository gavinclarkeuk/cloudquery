package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IamPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_policies",
		Description: "Contains information about a managed policy, including the policy's ARN, versions, and the number of principal entities (users, groups, and roles) that the policy is attached to.",
		Resolver:    fetchIamPolicies,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN). ARNs are unique identifiers for AWS resources. For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the AWS General Reference. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "attachment_count",
				Description: "The number of principal entities (users, groups, and roles) that the policy is attached to. ",
				Type:        schema.TypeInt,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy was created. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_version_id",
				Description: "The identifier for the version of the policy that is set as the default (operative) version. For more information about policy versions, see Versioning for managed policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html) in the IAM User Guide. ",
				Type:        schema.TypeString,
			},
			{
				Name:          "description",
				Description:   "A friendly description of the policy. ",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "is_attachable",
				Description: "Specifies whether the policy can be attached to an IAM user, group, or role. ",
				Type:        schema.TypeBool,
			},
			{
				Name:        "path",
				Description: "The path to the policy. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "permissions_boundary_usage_count",
				Description: "The number of entities (users and roles) for which the policy is used as the permissions boundary. For more information about permissions boundaries, see Permissions boundaries for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the IAM User Guide. ",
				Type:        schema.TypeInt,
			},
			{
				Name:            "id",
				Description:     "The stable and unique string identifying the policy. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. ",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("PolicyId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The friendly name (not ARN) identifying the policy. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyName"),
			},
			{
				Name:        "update_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the policy was last updated. When a policy has only one version, this field contains the date and time when the policy was created. When a policy has more than one version, this field contains the date and time when the most recent policy version was created. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "tags",
				Description: "A list of tags that are attached to the role. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamPolicyTags,
			},
			{
				Name:        "policy_version_list",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.GetAccountAuthorizationDetailsInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.GetAccountAuthorizationDetails(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Policies
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}


func resolveIamPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ManagedPolicyDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().IAM
	response, err := svc.ListPolicyTags(ctx, &iam.ListPolicyTagsInput{PolicyArn: r.Arn})
	if err != nil {
		if cl.IsNotFoundError(err) {
			meta.Logger().Debug().Err(err).Msg("ListPolicyTags: Policy does not exist")
			return nil
		}
		return err
	}
	return resource.Set("tags", client.TagsToMap(response.Tags))
}
