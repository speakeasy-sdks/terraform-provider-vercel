// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateRecord2 struct {
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	TTL     types.Number `tfsdk:"ttl"`
	Type    types.String `tfsdk:"type"`
	UID     types.String `tfsdk:"uid"`
	Value   types.String `tfsdk:"value"`
}
