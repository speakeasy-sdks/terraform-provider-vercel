// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/internal/utils"
	"net/http"
)

type UploadFileRequest struct {
	// The file size in bytes
	ContentLength *int64 `header:"style=simple,explode=false,name=Content-Length"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The file SHA1 used to check the integrity
	XNowDigest *string `header:"style=simple,explode=false,name=x-now-digest"`
	// The file size as an alternative to `Content-Length`
	XNowSize *int64 `header:"style=simple,explode=false,name=x-now-size"`
	// The file SHA1 used to check the integrity
	XVercelDigest *string `header:"style=simple,explode=false,name=x-vercel-digest"`
}

func (o *UploadFileRequest) GetContentLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ContentLength
}

func (o *UploadFileRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *UploadFileRequest) GetXNowDigest() *string {
	if o == nil {
		return nil
	}
	return o.XNowDigest
}

func (o *UploadFileRequest) GetXNowSize() *int64 {
	if o == nil {
		return nil
	}
	return o.XNowSize
}

func (o *UploadFileRequest) GetXVercelDigest() *string {
	if o == nil {
		return nil
	}
	return o.XVercelDigest
}

type UploadFile2 struct {
}

type UploadFile1 struct {
	// Array of URLs where the file was updated
	Urls []string `json:"urls"`
}

func (o *UploadFile1) GetUrls() []string {
	if o == nil {
		return []string{}
	}
	return o.Urls
}

type UploadFileResponseBodyType string

const (
	UploadFileResponseBodyTypeUploadFile1 UploadFileResponseBodyType = "uploadFile_1"
	UploadFileResponseBodyTypeUploadFile2 UploadFileResponseBodyType = "uploadFile_2"
)

// UploadFileResponseBody - File already uploaded
// File successfully uploaded
type UploadFileResponseBody struct {
	UploadFile1 *UploadFile1
	UploadFile2 *UploadFile2

	Type UploadFileResponseBodyType
}

func CreateUploadFileResponseBodyUploadFile1(uploadFile1 UploadFile1) UploadFileResponseBody {
	typ := UploadFileResponseBodyTypeUploadFile1

	return UploadFileResponseBody{
		UploadFile1: &uploadFile1,
		Type:        typ,
	}
}

func CreateUploadFileResponseBodyUploadFile2(uploadFile2 UploadFile2) UploadFileResponseBody {
	typ := UploadFileResponseBodyTypeUploadFile2

	return UploadFileResponseBody{
		UploadFile2: &uploadFile2,
		Type:        typ,
	}
}

func (u *UploadFileResponseBody) UnmarshalJSON(data []byte) error {

	uploadFile2 := UploadFile2{}
	if err := utils.UnmarshalJSON(data, &uploadFile2, "", true, true); err == nil {
		u.UploadFile2 = &uploadFile2
		u.Type = UploadFileResponseBodyTypeUploadFile2
		return nil
	}

	uploadFile1 := UploadFile1{}
	if err := utils.UnmarshalJSON(data, &uploadFile1, "", true, true); err == nil {
		u.UploadFile1 = &uploadFile1
		u.Type = UploadFileResponseBodyTypeUploadFile1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UploadFileResponseBody) MarshalJSON() ([]byte, error) {
	if u.UploadFile1 != nil {
		return utils.MarshalJSON(u.UploadFile1, "", true)
	}

	if u.UploadFile2 != nil {
		return utils.MarshalJSON(u.UploadFile2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UploadFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// File already uploaded
	// File successfully uploaded
	OneOf *UploadFileResponseBody
}

func (o *UploadFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadFileResponse) GetOneOf() *UploadFileResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
