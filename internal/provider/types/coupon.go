// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Coupon struct {
	AmountOff        types.Number `tfsdk:"amount_off"`
	Duration         types.String `tfsdk:"duration"`
	DurationInMonths types.Number `tfsdk:"duration_in_months"`
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	PercentageOff    types.Number `tfsdk:"percentage_off"`
}