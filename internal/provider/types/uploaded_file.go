// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type UploadedFile struct {
	File types.String `tfsdk:"file"`
	Sha  types.String `tfsdk:"sha"`
	Size types.Int64  `tfsdk:"size"`
}
