// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/zchee/terraform-provider-vercel/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetEdgeConfigItemRequest struct {
	EdgeConfigID      string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	EdgeConfigItemKey string `pathParam:"style=simple,explode=false,name=edgeConfigItemKey"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetEdgeConfigItemRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *GetEdgeConfigItemRequest) GetEdgeConfigItemKey() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigItemKey
}

func (o *GetEdgeConfigItemRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetEdgeConfigItemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The EdgeConfig.
	EdgeConfigItem *shared.EdgeConfigItem
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEdgeConfigItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEdgeConfigItemResponse) GetEdgeConfigItem() *shared.EdgeConfigItem {
	if o == nil {
		return nil
	}
	return o.EdgeConfigItem
}

func (o *GetEdgeConfigItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEdgeConfigItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
