// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateProjectProjectsTarget struct {
	ArrayOfOne []types.String `tfsdk:"array_of_one" tfPlanOnly:"true"`
	Two        types.String   `tfsdk:"two" tfPlanOnly:"true"`
}