// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateDeploymentCrons struct {
	Path     types.String `tfsdk:"path"`
	Schedule types.String `tfsdk:"schedule"`
}