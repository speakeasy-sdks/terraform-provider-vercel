// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/zchee/terraform-provider-vercel/internal/provider/types"
	"github.com/zchee/terraform-provider-vercel/internal/sdk"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &EdgeConfigDataSource{}
var _ datasource.DataSourceWithConfigure = &EdgeConfigDataSource{}

func NewEdgeConfigDataSource() datasource.DataSource {
	return &EdgeConfigDataSource{}
}

// EdgeConfigDataSource is the data source implementation.
type EdgeConfigDataSource struct {
	client *sdk.Vercel
}

// EdgeConfigDataSourceModel describes the data model.
type EdgeConfigDataSourceModel struct {
	CreatedAt    types.Number      `tfsdk:"created_at"`
	Digest       types.String      `tfsdk:"digest"`
	EdgeConfigID types.String      `tfsdk:"edge_config_id"`
	ID           types.String      `tfsdk:"id"`
	ItemCount    types.Number      `tfsdk:"item_count"`
	OwnerID      types.String      `tfsdk:"owner_id"`
	Schema       *tfTypes.Schema   `tfsdk:"schema"`
	SizeInBytes  types.Number      `tfsdk:"size_in_bytes"`
	Slug         types.String      `tfsdk:"slug"`
	TeamID       types.String      `tfsdk:"team_id"`
	Transfer     *tfTypes.Transfer `tfsdk:"transfer"`
	UpdatedAt    types.Number      `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *EdgeConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_edge_config"
}

// Schema defines the schema for the data source.
func (r *EdgeConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "EdgeConfig DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.NumberAttribute{
				Computed: true,
			},
			"digest": schema.StringAttribute{
				Computed: true,
			},
			"edge_config_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"item_count": schema.NumberAttribute{
				Computed: true,
			},
			"owner_id": schema.StringAttribute{
				Computed: true,
			},
			"schema": schema.SingleNestedAttribute{
				Computed:   true,
				Attributes: map[string]schema.Attribute{},
			},
			"size_in_bytes": schema.NumberAttribute{
				Computed: true,
			},
			"slug": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The Team slug to perform the request on behalf of.`,
			},
			"team_id": schema.StringAttribute{
				Optional:    true,
				Description: `The Team identifier to perform the request on behalf of.`,
			},
			"transfer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"done_at": schema.NumberAttribute{
						Computed: true,
					},
					"from_account_id": schema.StringAttribute{
						Computed: true,
					},
					"started_at": schema.NumberAttribute{
						Computed: true,
					},
				},
				Description: `Keeps track of the current state of the Edge Config while it gets transferred.`,
			},
			"updated_at": schema.NumberAttribute{
				Computed: true,
			},
		},
	}
}

func (r *EdgeConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Vercel)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Vercel, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EdgeConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *EdgeConfigDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	edgeConfigID := data.EdgeConfigID.ValueString()
	slug := new(string)
	if !data.Slug.IsUnknown() && !data.Slug.IsNull() {
		*slug = data.Slug.ValueString()
	} else {
		slug = nil
	}
	teamID := new(string)
	if !data.TeamID.IsUnknown() && !data.TeamID.IsNull() {
		*teamID = data.TeamID.ValueString()
	} else {
		teamID = nil
	}
	request := operations.GetEdgeConfigRequest{
		EdgeConfigID: edgeConfigID,
		Slug:         slug,
		TeamID:       teamID,
	}
	res, err := r.client.EdgeConfig.Get(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Object == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromOperationsGetEdgeConfigResponseBody(res.Object)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
