// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Erl struct {
	Algo   types.String   `tfsdk:"algo"`
	Keys   []types.String `tfsdk:"keys"`
	Limit  types.Number   `tfsdk:"limit"`
	Window types.Number   `tfsdk:"window"`
}
