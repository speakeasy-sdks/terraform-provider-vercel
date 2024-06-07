// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Pagination - This object contains information related to the pagination of the current request, including the necessary parameters to get the next or previous page of data.
type Pagination struct {
	// Amount of items in the current page.
	Count float64 `json:"count"`
	// Timestamp that must be used to request the next page.
	Next *float64 `json:"next"`
	// Timestamp that must be used to request the previous page.
	Prev *float64 `json:"prev"`
}

func (o *Pagination) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *Pagination) GetNext() *float64 {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *Pagination) GetPrev() *float64 {
	if o == nil {
		return nil
	}
	return o.Prev
}
