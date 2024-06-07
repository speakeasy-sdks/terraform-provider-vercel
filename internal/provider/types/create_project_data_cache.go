// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateProjectDataCache struct {
	StorageSizeBytes types.Number `tfsdk:"storage_size_bytes"`
	Unlimited        types.Bool   `tfsdk:"unlimited"`
	UserDisabled     types.Bool   `tfsdk:"user_disabled"`
}
