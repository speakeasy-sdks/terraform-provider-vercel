// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type GitMetadata struct {
	CommitAuthorName types.String `tfsdk:"commit_author_name"`
	CommitMessage    types.String `tfsdk:"commit_message"`
	CommitRef        types.String `tfsdk:"commit_ref"`
	CommitSha        types.String `tfsdk:"commit_sha"`
	Dirty            types.Bool   `tfsdk:"dirty"`
	RemoteURL        types.String `tfsdk:"remote_url"`
}
