// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateProjectProjects1 struct {
	CreatedAt        types.Number  `tfsdk:"created_at"`
	DeployHooks      []DeployHooks `tfsdk:"deploy_hooks"`
	GitCredentialID  types.String  `tfsdk:"git_credential_id"`
	Org              types.String  `tfsdk:"org"`
	ProductionBranch types.String  `tfsdk:"production_branch"`
	Repo             types.String  `tfsdk:"repo"`
	RepoID           types.Number  `tfsdk:"repo_id"`
	Sourceless       types.Bool    `tfsdk:"sourceless"`
	Type             types.String  `tfsdk:"type"`
	UpdatedAt        types.Number  `tfsdk:"updated_at"`
}
