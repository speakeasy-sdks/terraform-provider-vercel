// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type GetDeploymentGitSourceDeployments3 struct {
	Type      types.String `tfsdk:"type"`
	ProjectID RepoID       `tfsdk:"project_id"`
	Ref       types.String `tfsdk:"ref"`
	Sha       types.String `tfsdk:"sha"`
	PrID      types.Number `tfsdk:"pr_id"`
}