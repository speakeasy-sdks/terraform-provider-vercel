// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateProjectTrustedIps1 struct {
	Addresses      []TrustedIpsAddresses `tfsdk:"addresses"`
	DeploymentType types.String          `tfsdk:"deployment_type"`
	ProtectionMode types.String          `tfsdk:"protection_mode"`
}