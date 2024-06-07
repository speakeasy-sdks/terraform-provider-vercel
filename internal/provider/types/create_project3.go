// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateProject3 struct {
	CreatedAt        types.Number  `tfsdk:"created_at"`
	DeployHooks      []DeployHooks `tfsdk:"deploy_hooks"`
	GitCredentialID  types.String  `tfsdk:"git_credential_id"`
	Name             types.String  `tfsdk:"name"`
	Owner            types.String  `tfsdk:"owner"`
	ProductionBranch types.String  `tfsdk:"production_branch"`
	Slug             types.String  `tfsdk:"slug"`
	Sourceless       types.Bool    `tfsdk:"sourceless"`
	Type             types.String  `tfsdk:"type"`
	UpdatedAt        types.Number  `tfsdk:"updated_at"`
	UUID             types.String  `tfsdk:"uuid"`
	WorkspaceUUID    types.String  `tfsdk:"workspace_uuid"`
}
