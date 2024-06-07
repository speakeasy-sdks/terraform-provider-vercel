// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateProjectCreator struct {
	Email       types.String `tfsdk:"email"`
	GithubLogin types.String `tfsdk:"github_login"`
	GitlabLogin types.String `tfsdk:"gitlab_login"`
	UID         types.String `tfsdk:"uid"`
	Username    types.String `tfsdk:"username"`
}
