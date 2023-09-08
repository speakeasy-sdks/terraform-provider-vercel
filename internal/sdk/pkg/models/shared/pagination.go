// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Pagination - This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
type Pagination struct {
	// Amount of items in the current page.
	Count int64 `json:"count"`
	// Timestamp that must be used to request the next page.
	Next *int64 `json:"next"`
	// Timestamp that must be used to request the previous page.
	Prev *int64 `json:"prev"`
}
