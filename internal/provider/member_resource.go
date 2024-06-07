// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-terraform/internal/provider/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MemberResource{}
var _ resource.ResourceWithImportState = &MemberResource{}

func NewMemberResource() resource.Resource {
	return &MemberResource{}
}

// MemberResource defines the resource implementation.
type MemberResource struct {
	client *sdk.SDK
}

// MemberResourceModel describes the resource data model.
type MemberResourceModel struct {
	ID       types.String                          `tfsdk:"id"`
	IDOrName types.String                          `tfsdk:"id_or_name"`
	One      *tfTypes.AddProjectMemberRequestBody1 `tfsdk:"one" tfPlanOnly:"true"`
	Slug     types.String                          `tfsdk:"slug"`
	TeamID   types.String                          `tfsdk:"team_id"`
	Three    *tfTypes.AddProjectMemberRequestBody3 `tfsdk:"three" tfPlanOnly:"true"`
	Two      *tfTypes.AddProjectMemberRequestBody2 `tfsdk:"two" tfPlanOnly:"true"`
	UID      types.String                          `tfsdk:"uid"`
}

func (r *MemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_member"
}

func (r *MemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Member Resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"id_or_name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The ID or name of the Project. Requires replacement if changed. `,
			},
			"one": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `The email of the team member that should be added to this project. Requires replacement if changed. `,
					},
					"role": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Required:    true,
						Description: `The project role of the member that will be added. Requires replacement if changed. ; must be one of ["ADMIN", "PROJECT_DEVELOPER", "PROJECT_VIEWER"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"ADMIN",
								"PROJECT_DEVELOPER",
								"PROJECT_VIEWER",
							),
						},
					},
					"uid": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Required:    true,
						Description: `The ID of the team member that should be added to this project. Requires replacement if changed. `,
					},
					"username": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `The username of the team member that should be added to this project. Requires replacement if changed. `,
					},
				},
				Description: `Requires replacement if changed. `,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("two"),
						path.MatchRelative().AtParent().AtName("three"),
					}...),
				},
			},
			"slug": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The Team slug to perform the request on behalf of. Requires replacement if changed. `,
			},
			"team_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The Team identifier to perform the request on behalf of. Requires replacement if changed. `,
			},
			"three": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Required:    true,
						Description: `The email of the team member that should be added to this project. Requires replacement if changed. `,
					},
					"role": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Required:    true,
						Description: `The project role of the member that will be added. Requires replacement if changed. ; must be one of ["ADMIN", "PROJECT_DEVELOPER", "PROJECT_VIEWER"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"ADMIN",
								"PROJECT_DEVELOPER",
								"PROJECT_VIEWER",
							),
						},
					},
					"uid": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `The ID of the team member that should be added to this project. Requires replacement if changed. `,
					},
					"username": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `The username of the team member that should be added to this project. Requires replacement if changed. `,
					},
				},
				Description: `Requires replacement if changed. `,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("one"),
						path.MatchRelative().AtParent().AtName("two"),
					}...),
				},
			},
			"two": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `The email of the team member that should be added to this project. Requires replacement if changed. `,
					},
					"role": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Required:    true,
						Description: `The project role of the member that will be added. Requires replacement if changed. ; must be one of ["ADMIN", "PROJECT_DEVELOPER", "PROJECT_VIEWER"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"ADMIN",
								"PROJECT_DEVELOPER",
								"PROJECT_VIEWER",
							),
						},
					},
					"uid": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `The ID of the team member that should be added to this project. Requires replacement if changed. `,
					},
					"username": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Required:    true,
						Description: `The username of the team member that should be added to this project. Requires replacement if changed. `,
					},
				},
				Description: `Requires replacement if changed. `,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("one"),
						path.MatchRelative().AtParent().AtName("three"),
					}...),
				},
			},
			"uid": schema.StringAttribute{
				Required:    true,
				Description: `The user ID of the member.`,
			},
		},
	}
}

func (r *MemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MemberResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	idOrName := data.IDOrName.ValueString()
	teamID := new(string)
	if !data.TeamID.IsUnknown() && !data.TeamID.IsNull() {
		*teamID = data.TeamID.ValueString()
	} else {
		teamID = nil
	}
	slug := new(string)
	if !data.Slug.IsUnknown() && !data.Slug.IsNull() {
		*slug = data.Slug.ValueString()
	} else {
		slug = nil
	}
	requestBody := data.ToOperationsAddProjectMemberRequestBody()
	request := operations.AddProjectMemberRequest{
		IDOrName:    idOrName,
		TeamID:      teamID,
		Slug:        slug,
		RequestBody: requestBody,
	}
	res, err := r.client.ProjectMembers.Create(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Object == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromOperationsAddProjectMemberResponseBody(res.Object)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MemberResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MemberResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MemberResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	idOrName := data.IDOrName.ValueString()
	uid := data.UID.ValueString()
	teamID := new(string)
	if !data.TeamID.IsUnknown() && !data.TeamID.IsNull() {
		*teamID = data.TeamID.ValueString()
	} else {
		teamID = nil
	}
	slug := new(string)
	if !data.Slug.IsUnknown() && !data.Slug.IsNull() {
		*slug = data.Slug.ValueString()
	} else {
		slug = nil
	}
	request := operations.RemoveProjectMemberRequest{
		IDOrName: idOrName,
		UID:      uid,
		TeamID:   teamID,
		Slug:     slug,
	}
	res, err := r.client.ProjectMembers.Delete(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *MemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource member.")
}
