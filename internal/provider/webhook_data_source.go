// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/zchee/terraform-provider-vercel/internal/sdk"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &WebhookDataSource{}
var _ datasource.DataSourceWithConfigure = &WebhookDataSource{}

func NewWebhookDataSource() datasource.DataSource {
	return &WebhookDataSource{}
}

// WebhookDataSource is the data source implementation.
type WebhookDataSource struct {
	client *sdk.Vercel
}

// WebhookDataSourceModel describes the data model.
type WebhookDataSourceModel struct {
	CreatedAt  types.Number   `tfsdk:"created_at"`
	Events     []types.String `tfsdk:"events"`
	ID         types.String   `tfsdk:"id"`
	OwnerID    types.String   `tfsdk:"owner_id"`
	ProjectIds []types.String `tfsdk:"project_ids"`
	Slug       types.String   `tfsdk:"slug"`
	TeamID     types.String   `tfsdk:"team_id"`
	UpdatedAt  types.Number   `tfsdk:"updated_at"`
	URL        types.String   `tfsdk:"url"`
}

// Metadata returns the data source type name.
func (r *WebhookDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

// Schema defines the schema for the data source.
func (r *WebhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Webhook DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.NumberAttribute{
				Computed:    true,
				Description: `A number containing the date when the webhook was created in in milliseconds`,
			},
			"events": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The webhooks events`,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"owner_id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique ID of the team the webhook belongs to`,
			},
			"project_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The ID of the projects the webhook is associated with`,
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Description: `The Team slug to perform the request on behalf of.`,
			},
			"team_id": schema.StringAttribute{
				Optional:    true,
				Description: `The Team identifier to perform the request on behalf of.`,
			},
			"updated_at": schema.NumberAttribute{
				Computed:    true,
				Description: `A number containing the date when the webhook was updated in in milliseconds`,
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: `A string with the URL of the webhook`,
			},
		},
	}
}

func (r *WebhookDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *WebhookDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WebhookDataSourceModel
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

	id := data.ID.ValueString()
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
	request := operations.GetWebhookRequest{
		ID:     id,
		Slug:   slug,
		TeamID: teamID,
	}
	res, err := r.client.Webhooks.Get(ctx, request)
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
	data.RefreshFromOperationsGetWebhookResponseBody(res.Object)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}