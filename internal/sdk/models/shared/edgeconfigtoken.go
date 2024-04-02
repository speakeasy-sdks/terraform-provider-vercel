// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EdgeConfigToken - The EdgeConfig.
type EdgeConfigToken struct {
	CreatedAt    int64  `json:"createdAt"`
	EdgeConfigID string `json:"edgeConfigId"`
	// This is not the token itself, but rather an id to identify the token by
	ID    string `json:"id"`
	Label string `json:"label"`
	Token string `json:"token"`
}

func (o *EdgeConfigToken) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *EdgeConfigToken) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *EdgeConfigToken) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *EdgeConfigToken) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *EdgeConfigToken) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}