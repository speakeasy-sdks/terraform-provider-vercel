// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Attribution struct {
	LandingPage              types.String `tfsdk:"landing_page"`
	PageBeforeConversionPage types.String `tfsdk:"page_before_conversion_page"`
	SessionReferrer          types.String `tfsdk:"session_referrer"`
	Utm                      *Utm         `tfsdk:"utm"`
}
