// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Transfer struct {
	DoneAt        types.Number `tfsdk:"done_at"`
	FromAccountID types.String `tfsdk:"from_account_id"`
	StartedAt     types.Number `tfsdk:"started_at"`
}
