// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type GetProjectProtectionBypass struct {
	CreatedAt types.Number `tfsdk:"created_at"`
	CreatedBy types.String `tfsdk:"created_by"`
	Scope     types.String `tfsdk:"scope"`
}