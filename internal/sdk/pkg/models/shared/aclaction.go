// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ACLAction - Enum containing the actions that can be performed against a resource. Group operations are included.
type ACLAction string

const (
	ACLActionCreate ACLAction = "create"
	ACLActionDelete ACLAction = "delete"
	ACLActionRead   ACLAction = "read"
	ACLActionUpdate ACLAction = "update"
	ACLActionList   ACLAction = "list"
)

func (e ACLAction) ToPointer() *ACLAction {
	return &e
}

func (e *ACLAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "create":
		fallthrough
	case "delete":
		fallthrough
	case "read":
		fallthrough
	case "update":
		fallthrough
	case "list":
		*e = ACLAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ACLAction: %v", v)
	}
}
