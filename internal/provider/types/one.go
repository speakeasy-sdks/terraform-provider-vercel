// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type One struct {
	PrID   types.Number `tfsdk:"pr_id"`
	Ref    types.String `tfsdk:"ref"`
	RepoID RepoID       `tfsdk:"repo_id"`
	Sha    types.String `tfsdk:"sha"`
	Type   types.String `tfsdk:"type"`
}