// Code generated by codegen; DO NOT EDIT.

package organizations

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Members() *schema.Table {
	return &schema.Table{
		Name:     "github_organization_members",
		Resolver: fetchMembers,
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "membership",
				Type:     schema.TypeJSON,
				Resolver: resolveMembership,
			},
			{
				Name:     "login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Login"),
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
			{
				Name:     "avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvatarURL"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GravatarID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Company"),
			},
			{
				Name:     "blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Blog"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Hireable"),
			},
			{
				Name:     "bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Bio"),
			},
			{
				Name:     "twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TwitterUsername"),
			},
			{
				Name:     "public_repos",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PublicRepos"),
			},
			{
				Name:     "public_gists",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PublicGists"),
			},
			{
				Name:     "followers",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Followers"),
			},
			{
				Name:     "following",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Following"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "suspended_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SuspendedAt"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SiteAdmin"),
			},
			{
				Name:     "total_private_repos",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalPrivateRepos"),
			},
			{
				Name:     "owned_private_repos",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("OwnedPrivateRepos"),
			},
			{
				Name:     "private_gists",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PrivateGists"),
			},
			{
				Name:     "disk_usage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskUsage"),
			},
			{
				Name:     "collaborators",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Collaborators"),
			},
			{
				Name:     "two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TwoFactorAuthentication"),
			},
			{
				Name:     "plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Plan"),
			},
			{
				Name:     "ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LdapDn"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventsURL"),
			},
			{
				Name:     "following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FollowingURL"),
			},
			{
				Name:     "followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FollowersURL"),
			},
			{
				Name:     "gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GistsURL"),
			},
			{
				Name:     "organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OrganizationsURL"),
			},
			{
				Name:     "received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReceivedEventsURL"),
			},
			{
				Name:     "repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReposURL"),
			},
			{
				Name:     "starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StarredURL"),
			},
			{
				Name:     "subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionsURL"),
			},
			{
				Name:     "text_matches",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TextMatches"),
			},
			{
				Name:     "permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Permissions"),
			},
			{
				Name:     "role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleName"),
			},
		},
	}
}
