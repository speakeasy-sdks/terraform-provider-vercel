// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/vercel/terraform-provider-vercel/internal/provider/types"
	"github.com/vercel/terraform-provider-vercel/internal/sdk/models/operations"
	"math/big"
)

func (r *EdgeConfigDataSourceModel) RefreshFromOperationsGetEdgeConfigResponseBody(resp *operations.GetEdgeConfigResponseBody) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.NumberValue(big.NewFloat(float64(*resp.CreatedAt)))
		} else {
			r.CreatedAt = types.NumberNull()
		}
		r.Digest = types.StringPointerValue(resp.Digest)
		r.ID = types.StringPointerValue(resp.ID)
		r.ItemCount = types.NumberValue(big.NewFloat(float64(resp.ItemCount)))
		r.OwnerID = types.StringPointerValue(resp.OwnerID)
		if resp.Schema == nil {
			r.Schema = nil
		} else {
			r.Schema = &tfTypes.GetProjectLastRollbackTarget{}
		}
		r.SizeInBytes = types.NumberValue(big.NewFloat(float64(resp.SizeInBytes)))
		r.Slug = types.StringPointerValue(resp.Slug)
		if resp.Transfer == nil {
			r.Transfer = nil
		} else {
			r.Transfer = &tfTypes.GetEdgeConfigTransfer{}
			if resp.Transfer.DoneAt != nil {
				r.Transfer.DoneAt = types.NumberValue(big.NewFloat(float64(*resp.Transfer.DoneAt)))
			} else {
				r.Transfer.DoneAt = types.NumberNull()
			}
			r.Transfer.FromAccountID = types.StringValue(resp.Transfer.FromAccountID)
			r.Transfer.StartedAt = types.NumberValue(big.NewFloat(float64(resp.Transfer.StartedAt)))
		}
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.NumberValue(big.NewFloat(float64(*resp.UpdatedAt)))
		} else {
			r.UpdatedAt = types.NumberNull()
		}
	}
}
