// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetEdgeConfigRequest struct {
	EdgeConfigID string `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

// GetEdgeConfig200ApplicationJSONTransfer - Keeps track of the current state of the Edge Config while it gets transferred.
type GetEdgeConfig200ApplicationJSONTransfer struct {
	DoneAt        *int64 `json:"doneAt"`
	FromAccountID string `json:"fromAccountId"`
	StartedAt     int64  `json:"startedAt"`
}

// GetEdgeConfig200ApplicationJSON - The EdgeConfig.
type GetEdgeConfig200ApplicationJSON struct {
	CreatedAt   *int64  `json:"createdAt,omitempty"`
	Digest      *string `json:"digest,omitempty"`
	ID          *string `json:"id,omitempty"`
	ItemCount   int64   `json:"itemCount"`
	OwnerID     *string `json:"ownerId,omitempty"`
	SizeInBytes int64   `json:"sizeInBytes"`
	// Name for the Edge Config Names are not unique. Must start with an alphabetic character and can contain only alphanumeric characters and underscores).
	Slug *string `json:"slug,omitempty"`
	// Keeps track of the current state of the Edge Config while it gets transferred.
	Transfer  *GetEdgeConfig200ApplicationJSONTransfer `json:"transfer,omitempty"`
	UpdatedAt *int64                                   `json:"updatedAt,omitempty"`
}

type GetEdgeConfigResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The EdgeConfig.
	GetEdgeConfig200ApplicationJSONObject *GetEdgeConfig200ApplicationJSON
}
