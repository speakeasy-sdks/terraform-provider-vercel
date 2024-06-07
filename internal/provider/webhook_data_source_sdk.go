// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/operations"
	"math/big"
)

func (r *WebhookDataSourceModel) RefreshFromOperationsGetWebhookResponseBody(resp *operations.GetWebhookResponseBody) {
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
