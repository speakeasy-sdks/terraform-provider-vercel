// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ProjectSettings struct {
	BuildCommand                    types.String `tfsdk:"build_command"`
	CommandForIgnoringBuildStep     types.String `tfsdk:"command_for_ignoring_build_step"`
	DevCommand                      types.String `tfsdk:"dev_command"`
	Framework                       types.String `tfsdk:"framework"`
	InstallCommand                  types.String `tfsdk:"install_command"`
	NodeVersion                     types.String `tfsdk:"node_version"`
	OutputDirectory                 types.String `tfsdk:"output_directory"`
	RootDirectory                   types.String `tfsdk:"root_directory"`
	ServerlessFunctionRegion        types.String `tfsdk:"serverless_function_region"`
	SkipGitConnectDuringLink        types.Bool   `tfsdk:"skip_git_connect_during_link"`
	SourceFilesOutsideRootDirectory types.Bool   `tfsdk:"source_files_outside_root_directory"`
}
