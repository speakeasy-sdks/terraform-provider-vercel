// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EdgeConfigItem - The EdgeConfig.
type EdgeConfigItem struct {
	CreatedAt    int64                `json:"createdAt"`
	EdgeConfigID string               `json:"edgeConfigId"`
	Key          string               `json:"key"`
	UpdatedAt    int64                `json:"updatedAt"`
	Value        *EdgeConfigItemValue `json:"value"`
}
