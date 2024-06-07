// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Three struct {
	PrID      types.Number `tfsdk:"pr_id"`
	ProjectID RepoID       `tfsdk:"project_id"`
	Ref       types.String `tfsdk:"ref"`
	Sha       types.String `tfsdk:"sha"`
	Type      types.String `tfsdk:"type"`
}
