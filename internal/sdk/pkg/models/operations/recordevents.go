// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Event - One of `HIT` or `MISS`. `HIT` specifies that a cached artifact for `hash` was found in the cache. `MISS` specifies that a cached artifact with `hash` was not found.
type Event string

const (
	EventHit  Event = "HIT"
	EventMiss Event = "MISS"
)

func (e Event) ToPointer() *Event {
	return &e
}

func (e *Event) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HIT":
		fallthrough
	case "MISS":
		*e = Event(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Event: %v", v)
	}
}

// Source - One of `LOCAL` or `REMOTE`. `LOCAL` specifies that the cache event was from the user's filesystem cache. `REMOTE` specifies that the cache event is from a remote cache.
type Source string

const (
	SourceLocal  Source = "LOCAL"
	SourceRemote Source = "REMOTE"
)

func (e Source) ToPointer() *Source {
	return &e
}

func (e *Source) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOCAL":
		fallthrough
	case "REMOTE":
		*e = Source(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Source: %v", v)
	}
}

type RequestBody struct {
	// The time taken to generate the artifact. This should be sent as a body parameter on `HIT` events.
	Duration *int64 `json:"duration,omitempty"`
	// One of `HIT` or `MISS`. `HIT` specifies that a cached artifact for `hash` was found in the cache. `MISS` specifies that a cached artifact with `hash` was not found.
	Event Event `json:"event"`
	// The artifact hash
	Hash string `json:"hash"`
	// A UUID (universally unique identifer) for the session that generated this event.
	SessionID string `json:"sessionId"`
	// One of `LOCAL` or `REMOTE`. `LOCAL` specifies that the cache event was from the user's filesystem cache. `REMOTE` specifies that the cache event is from a remote cache.
	Source Source `json:"source"`
}

func (o *RequestBody) GetDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *RequestBody) GetEvent() Event {
	if o == nil {
		return Event("")
	}
	return o.Event
}

func (o *RequestBody) GetHash() string {
	if o == nil {
		return ""
	}
	return o.Hash
}

func (o *RequestBody) GetSessionID() string {
	if o == nil {
		return ""
	}
	return o.SessionID
}

func (o *RequestBody) GetSource() Source {
	if o == nil {
		return Source("")
	}
	return o.Source
}

type RecordEventsRequest struct {
	RequestBody []RequestBody `request:"mediaType=application/json"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The continuous integration or delivery environment where this artifact is downloaded.
	XArtifactClientCi *string `header:"style=simple,explode=false,name=x-artifact-client-ci"`
	// 1 if the client is an interactive shell. Otherwise 0
	XArtifactClientInteractive *int64 `header:"style=simple,explode=false,name=x-artifact-client-interactive"`
}

func (o *RecordEventsRequest) GetRequestBody() []RequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *RecordEventsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *RecordEventsRequest) GetXArtifactClientCi() *string {
	if o == nil {
		return nil
	}
	return o.XArtifactClientCi
}

func (o *RecordEventsRequest) GetXArtifactClientInteractive() *int64 {
	if o == nil {
		return nil
	}
	return o.XArtifactClientInteractive
}

type RecordEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RecordEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RecordEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RecordEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
