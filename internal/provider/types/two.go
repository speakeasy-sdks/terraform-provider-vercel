// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Two struct {
	Org  types.String `tfsdk:"org"`
	PrID types.Number `tfsdk:"pr_id"`
	Ref  types.String `tfsdk:"ref"`
	Repo types.String `tfsdk:"repo"`
	Sha  types.String `tfsdk:"sha"`
	Type types.String `tfsdk:"type"`
}
