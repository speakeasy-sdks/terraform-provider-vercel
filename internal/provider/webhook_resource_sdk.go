// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/operations"
	"math/big"
)

func (r *WebhookResourceModel) ToOperationsCreateWebhookRequestBody() *operations.CreateWebhookRequestBody {
	url := r.URL.ValueString()
	var events []operations.Events = []operations.Events{}
	for _, eventsItem := range r.Events {
		events = append(events, operations.Events(eventsItem.ValueString()))
	}
	var projectIds []string = []string{}
	for _, projectIdsItem := range r.ProjectIds {
		projectIds = append(projectIds, projectIdsItem.ValueString())
	}
	out := operations.CreateWebhookRequestBody{
		URL:        url,
		Events:     events,
		ProjectIds: projectIds,
	}
	return &out
}

func (r *WebhookResourceModel) RefreshFromOperationsCreateWebhookResponseBody(resp *operations.CreateWebhookResponseBody) {
	if resp != nil {
		r.CreatedAt = types.NumberValue(big.NewFloat(float64(resp.CreatedAt)))
		r.Events = []types.String{}
		for _, v := range resp.Events {
			r.Events = append(r.Events, types.StringValue(string(v)))
		}
		r.ID = types.StringValue(resp.ID)
		r.OwnerID = types.StringValue(resp.OwnerID)
		r.ProjectIds = []types.String{}
		for _, v := range resp.ProjectIds {
			r.ProjectIds = append(r.ProjectIds, types.StringValue(v))
		}
		r.Secret = types.StringValue(resp.Secret)
		r.UpdatedAt = types.NumberValue(big.NewFloat(float64(resp.UpdatedAt)))
		r.URL = types.StringValue(resp.URL)
	}
}

func (r *WebhookResourceModel) RefreshFromOperationsGetWebhookResponseBody(resp *operations.GetWebhookResponseBody) {
	if resp != nil {
		r.CreatedAt = types.NumberValue(big.NewFloat(float64(resp.CreatedAt)))
		r.Events = []types.String{}
		for _, v := range resp.Events {
			r.Events = append(r.Events, types.StringValue(string(v)))
		}
		r.ID = types.StringValue(resp.ID)
		r.OwnerID = types.StringValue(resp.OwnerID)
		r.ProjectIds = []types.String{}
		for _, v := range resp.ProjectIds {
			r.ProjectIds = append(r.ProjectIds, types.StringValue(v))
		}
		r.UpdatedAt = types.NumberValue(big.NewFloat(float64(resp.UpdatedAt)))
		r.URL = types.StringValue(resp.URL)
	}
}
