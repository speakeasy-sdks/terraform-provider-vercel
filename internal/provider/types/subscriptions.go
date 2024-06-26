// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Subscriptions struct {
	Discount  *Discount    `tfsdk:"discount"`
	Frequency Frequency    `tfsdk:"frequency"`
	ID        types.String `tfsdk:"id"`
	Items     []Items      `tfsdk:"items"`
	Period    Contract     `tfsdk:"period"`
	Trial     *Contract    `tfsdk:"trial"`
}
