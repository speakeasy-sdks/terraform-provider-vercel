// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"vercel/internal/sdk/pkg/models/shared"
)

type GetEdgeConfigItemSecurity struct {
	BearerToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetEdgeConfigItemRequest struct {
	EdgeConfigID      string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	EdgeConfigItemKey string `pathParam:"style=simple,explode=false,name=edgeConfigItemKey"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

type GetEdgeConfigItemResponse struct {
	ContentType string
	// The EdgeConfig.
	EdgeConfigItem *shared.EdgeConfigItem
	StatusCode     int
	RawResponse    *http.Response
}
