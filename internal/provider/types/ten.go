// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Ten struct {
	Comment types.String     `tfsdk:"comment"`
	HTTPS   RequestBodyHTTPS `tfsdk:"https"`
	Name    types.String     `tfsdk:"name"`
	TTL     types.Number     `tfsdk:"ttl"`
	Type    types.String     `tfsdk:"type"`
}
