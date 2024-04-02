// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zchee/terraform-provider-vercel/internal/sdk/internal/utils"
	"net/http"
)

type GetProjectEnvRequest struct {
	// The unique ID for the environment variable to get the decrypted value.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier or slug to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
}

func (o *GetProjectEnvRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetProjectEnvRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *GetProjectEnvRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

type GetProjectEnvProjectsResponseType string

const (
	GetProjectEnvProjectsResponseTypePostgresDatabase GetProjectEnvProjectsResponseType = "postgres-database"
)

func (e GetProjectEnvProjectsResponseType) ToPointer() *GetProjectEnvProjectsResponseType {
	return &e
}

func (e *GetProjectEnvProjectsResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-database":
		*e = GetProjectEnvProjectsResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponseType: %v", v)
	}
}

type GetProjectEnv12 struct {
	StoreID string                            `json:"storeId"`
	Type    GetProjectEnvProjectsResponseType `json:"type"`
}

func (o *GetProjectEnv12) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv12) GetType() GetProjectEnvProjectsResponseType {
	if o == nil {
		return GetProjectEnvProjectsResponseType("")
	}
	return o.Type
}

type GetProjectEnvProjectsType string

const (
	GetProjectEnvProjectsTypePostgresPassword GetProjectEnvProjectsType = "postgres-password"
)

func (e GetProjectEnvProjectsType) ToPointer() *GetProjectEnvProjectsType {
	return &e
}

func (e *GetProjectEnvProjectsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-password":
		*e = GetProjectEnvProjectsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsType: %v", v)
	}
}

type GetProjectEnv11 struct {
	StoreID string                    `json:"storeId"`
	Type    GetProjectEnvProjectsType `json:"type"`
}

func (o *GetProjectEnv11) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv11) GetType() GetProjectEnvProjectsType {
	if o == nil {
		return GetProjectEnvProjectsType("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10TypePostgresHost GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type = "postgres-host"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-host":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type: %v", v)
	}
}

type GetProjectEnv10 struct {
	StoreID string                                                                       `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type `json:"type"`
}

func (o *GetProjectEnv10) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv10) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint10Type("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9TypePostgresUser GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type = "postgres-user"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-user":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type: %v", v)
	}
}

type GetProjectEnv9 struct {
	StoreID string                                                                      `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type `json:"type"`
}

func (o *GetProjectEnv9) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv9) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint9Type("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8TypePostgresPrismaURL GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type = "postgres-prisma-url"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-prisma-url":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type: %v", v)
	}
}

type GetProjectEnv8 struct {
	StoreID string                                                                      `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type `json:"type"`
}

func (o *GetProjectEnv8) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv8) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint8Type("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7TypePostgresURLNonPooling GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type = "postgres-url-non-pooling"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-url-non-pooling":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type: %v", v)
	}
}

type GetProjectEnv7 struct {
	StoreID string                                                                      `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type `json:"type"`
}

func (o *GetProjectEnv7) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv7) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint7Type("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6TypePostgresURL GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type = "postgres-url"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres-url":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type: %v", v)
	}
}

type GetProjectEnv6 struct {
	StoreID string                                                                      `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type `json:"type"`
}

func (o *GetProjectEnv6) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv6) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint6Type("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5TypeBlobReadWriteToken GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type = "blob-read-write-token"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "blob-read-write-token":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type: %v", v)
	}
}

type GetProjectEnv5 struct {
	StoreID string                                                                      `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type `json:"type"`
}

func (o *GetProjectEnv5) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv5) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHint5Type("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintTypeRedisRestAPIReadOnlyToken GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType = "redis-rest-api-read-only-token"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redis-rest-api-read-only-token":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType: %v", v)
	}
}

type GetProjectEnv4 struct {
	StoreID string                                                                     `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType `json:"type"`
}

func (o *GetProjectEnv4) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv4) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyContentHintType("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyTypeRedisRestAPIToken GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType = "redis-rest-api-token"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redis-rest-api-token":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType: %v", v)
	}
}

type GetProjectEnv3 struct {
	StoreID string                                                          `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType `json:"type"`
}

func (o *GetProjectEnv3) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv3) GetType() GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONResponseBodyType("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200ApplicationJSONType string

const (
	GetProjectEnvProjectsResponse200ApplicationJSONTypeRedisRestAPIURL GetProjectEnvProjectsResponse200ApplicationJSONType = "redis-rest-api-url"
)

func (e GetProjectEnvProjectsResponse200ApplicationJSONType) ToPointer() *GetProjectEnvProjectsResponse200ApplicationJSONType {
	return &e
}

func (e *GetProjectEnvProjectsResponse200ApplicationJSONType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redis-rest-api-url":
		*e = GetProjectEnvProjectsResponse200ApplicationJSONType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200ApplicationJSONType: %v", v)
	}
}

type GetProjectEnv2 struct {
	StoreID string                                              `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200ApplicationJSONType `json:"type"`
}

func (o *GetProjectEnv2) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv2) GetType() GetProjectEnvProjectsResponse200ApplicationJSONType {
	if o == nil {
		return GetProjectEnvProjectsResponse200ApplicationJSONType("")
	}
	return o.Type
}

type GetProjectEnvProjectsResponse200Type string

const (
	GetProjectEnvProjectsResponse200TypeRedisURL GetProjectEnvProjectsResponse200Type = "redis-url"
)

func (e GetProjectEnvProjectsResponse200Type) ToPointer() *GetProjectEnvProjectsResponse200Type {
	return &e
}

func (e *GetProjectEnvProjectsResponse200Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redis-url":
		*e = GetProjectEnvProjectsResponse200Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjectsResponse200Type: %v", v)
	}
}

type GetProjectEnv1 struct {
	StoreID string                               `json:"storeId"`
	Type    GetProjectEnvProjectsResponse200Type `json:"type"`
}

func (o *GetProjectEnv1) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *GetProjectEnv1) GetType() GetProjectEnvProjectsResponse200Type {
	if o == nil {
		return GetProjectEnvProjectsResponse200Type("")
	}
	return o.Type
}

type GetProjectEnvContentHintType string

const (
	GetProjectEnvContentHintTypeGetProjectEnv1  GetProjectEnvContentHintType = "getProjectEnv_1"
	GetProjectEnvContentHintTypeGetProjectEnv2  GetProjectEnvContentHintType = "getProjectEnv_2"
	GetProjectEnvContentHintTypeGetProjectEnv3  GetProjectEnvContentHintType = "getProjectEnv_3"
	GetProjectEnvContentHintTypeGetProjectEnv4  GetProjectEnvContentHintType = "getProjectEnv_4"
	GetProjectEnvContentHintTypeGetProjectEnv5  GetProjectEnvContentHintType = "getProjectEnv_5"
	GetProjectEnvContentHintTypeGetProjectEnv6  GetProjectEnvContentHintType = "getProjectEnv_6"
	GetProjectEnvContentHintTypeGetProjectEnv7  GetProjectEnvContentHintType = "getProjectEnv_7"
	GetProjectEnvContentHintTypeGetProjectEnv8  GetProjectEnvContentHintType = "getProjectEnv_8"
	GetProjectEnvContentHintTypeGetProjectEnv9  GetProjectEnvContentHintType = "getProjectEnv_9"
	GetProjectEnvContentHintTypeGetProjectEnv10 GetProjectEnvContentHintType = "getProjectEnv_10"
	GetProjectEnvContentHintTypeGetProjectEnv11 GetProjectEnvContentHintType = "getProjectEnv_11"
	GetProjectEnvContentHintTypeGetProjectEnv12 GetProjectEnvContentHintType = "getProjectEnv_12"
)

type GetProjectEnvContentHint struct {
	GetProjectEnv1  *GetProjectEnv1
	GetProjectEnv2  *GetProjectEnv2
	GetProjectEnv3  *GetProjectEnv3
	GetProjectEnv4  *GetProjectEnv4
	GetProjectEnv5  *GetProjectEnv5
	GetProjectEnv6  *GetProjectEnv6
	GetProjectEnv7  *GetProjectEnv7
	GetProjectEnv8  *GetProjectEnv8
	GetProjectEnv9  *GetProjectEnv9
	GetProjectEnv10 *GetProjectEnv10
	GetProjectEnv11 *GetProjectEnv11
	GetProjectEnv12 *GetProjectEnv12

	Type GetProjectEnvContentHintType
}

func CreateGetProjectEnvContentHintGetProjectEnv1(getProjectEnv1 GetProjectEnv1) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv1

	return GetProjectEnvContentHint{
		GetProjectEnv1: &getProjectEnv1,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv2(getProjectEnv2 GetProjectEnv2) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv2

	return GetProjectEnvContentHint{
		GetProjectEnv2: &getProjectEnv2,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv3(getProjectEnv3 GetProjectEnv3) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv3

	return GetProjectEnvContentHint{
		GetProjectEnv3: &getProjectEnv3,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv4(getProjectEnv4 GetProjectEnv4) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv4

	return GetProjectEnvContentHint{
		GetProjectEnv4: &getProjectEnv4,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv5(getProjectEnv5 GetProjectEnv5) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv5

	return GetProjectEnvContentHint{
		GetProjectEnv5: &getProjectEnv5,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv6(getProjectEnv6 GetProjectEnv6) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv6

	return GetProjectEnvContentHint{
		GetProjectEnv6: &getProjectEnv6,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv7(getProjectEnv7 GetProjectEnv7) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv7

	return GetProjectEnvContentHint{
		GetProjectEnv7: &getProjectEnv7,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv8(getProjectEnv8 GetProjectEnv8) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv8

	return GetProjectEnvContentHint{
		GetProjectEnv8: &getProjectEnv8,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv9(getProjectEnv9 GetProjectEnv9) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv9

	return GetProjectEnvContentHint{
		GetProjectEnv9: &getProjectEnv9,
		Type:           typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv10(getProjectEnv10 GetProjectEnv10) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv10

	return GetProjectEnvContentHint{
		GetProjectEnv10: &getProjectEnv10,
		Type:            typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv11(getProjectEnv11 GetProjectEnv11) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv11

	return GetProjectEnvContentHint{
		GetProjectEnv11: &getProjectEnv11,
		Type:            typ,
	}
}

func CreateGetProjectEnvContentHintGetProjectEnv12(getProjectEnv12 GetProjectEnv12) GetProjectEnvContentHint {
	typ := GetProjectEnvContentHintTypeGetProjectEnv12

	return GetProjectEnvContentHint{
		GetProjectEnv12: &getProjectEnv12,
		Type:            typ,
	}
}

func (u *GetProjectEnvContentHint) UnmarshalJSON(data []byte) error {

	getProjectEnv1 := GetProjectEnv1{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv1, "", true, true); err == nil {
		u.GetProjectEnv1 = &getProjectEnv1
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv1
		return nil
	}

	getProjectEnv2 := GetProjectEnv2{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv2, "", true, true); err == nil {
		u.GetProjectEnv2 = &getProjectEnv2
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv2
		return nil
	}

	getProjectEnv3 := GetProjectEnv3{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv3, "", true, true); err == nil {
		u.GetProjectEnv3 = &getProjectEnv3
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv3
		return nil
	}

	getProjectEnv4 := GetProjectEnv4{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv4, "", true, true); err == nil {
		u.GetProjectEnv4 = &getProjectEnv4
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv4
		return nil
	}

	getProjectEnv5 := GetProjectEnv5{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv5, "", true, true); err == nil {
		u.GetProjectEnv5 = &getProjectEnv5
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv5
		return nil
	}

	getProjectEnv6 := GetProjectEnv6{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv6, "", true, true); err == nil {
		u.GetProjectEnv6 = &getProjectEnv6
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv6
		return nil
	}

	getProjectEnv7 := GetProjectEnv7{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv7, "", true, true); err == nil {
		u.GetProjectEnv7 = &getProjectEnv7
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv7
		return nil
	}

	getProjectEnv8 := GetProjectEnv8{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv8, "", true, true); err == nil {
		u.GetProjectEnv8 = &getProjectEnv8
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv8
		return nil
	}

	getProjectEnv9 := GetProjectEnv9{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv9, "", true, true); err == nil {
		u.GetProjectEnv9 = &getProjectEnv9
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv9
		return nil
	}

	getProjectEnv10 := GetProjectEnv10{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv10, "", true, true); err == nil {
		u.GetProjectEnv10 = &getProjectEnv10
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv10
		return nil
	}

	getProjectEnv11 := GetProjectEnv11{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv11, "", true, true); err == nil {
		u.GetProjectEnv11 = &getProjectEnv11
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv11
		return nil
	}

	getProjectEnv12 := GetProjectEnv12{}
	if err := utils.UnmarshalJSON(data, &getProjectEnv12, "", true, true); err == nil {
		u.GetProjectEnv12 = &getProjectEnv12
		u.Type = GetProjectEnvContentHintTypeGetProjectEnv12
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetProjectEnvContentHint) MarshalJSON() ([]byte, error) {
	if u.GetProjectEnv1 != nil {
		return utils.MarshalJSON(u.GetProjectEnv1, "", true)
	}

	if u.GetProjectEnv2 != nil {
		return utils.MarshalJSON(u.GetProjectEnv2, "", true)
	}

	if u.GetProjectEnv3 != nil {
		return utils.MarshalJSON(u.GetProjectEnv3, "", true)
	}

	if u.GetProjectEnv4 != nil {
		return utils.MarshalJSON(u.GetProjectEnv4, "", true)
	}

	if u.GetProjectEnv5 != nil {
		return utils.MarshalJSON(u.GetProjectEnv5, "", true)
	}

	if u.GetProjectEnv6 != nil {
		return utils.MarshalJSON(u.GetProjectEnv6, "", true)
	}

	if u.GetProjectEnv7 != nil {
		return utils.MarshalJSON(u.GetProjectEnv7, "", true)
	}

	if u.GetProjectEnv8 != nil {
		return utils.MarshalJSON(u.GetProjectEnv8, "", true)
	}

	if u.GetProjectEnv9 != nil {
		return utils.MarshalJSON(u.GetProjectEnv9, "", true)
	}

	if u.GetProjectEnv10 != nil {
		return utils.MarshalJSON(u.GetProjectEnv10, "", true)
	}

	if u.GetProjectEnv11 != nil {
		return utils.MarshalJSON(u.GetProjectEnv11, "", true)
	}

	if u.GetProjectEnv12 != nil {
		return utils.MarshalJSON(u.GetProjectEnv12, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetProjectEnvProjects2 string

const (
	GetProjectEnvProjects2Production  GetProjectEnvProjects2 = "production"
	GetProjectEnvProjects2Preview     GetProjectEnvProjects2 = "preview"
	GetProjectEnvProjects2Development GetProjectEnvProjects2 = "development"
)

func (e GetProjectEnvProjects2) ToPointer() *GetProjectEnvProjects2 {
	return &e
}

func (e *GetProjectEnvProjects2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		fallthrough
	case "development":
		*e = GetProjectEnvProjects2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjects2: %v", v)
	}
}

type GetProjectEnvProjects1 string

const (
	GetProjectEnvProjects1Production  GetProjectEnvProjects1 = "production"
	GetProjectEnvProjects1Preview     GetProjectEnvProjects1 = "preview"
	GetProjectEnvProjects1Development GetProjectEnvProjects1 = "development"
)

func (e GetProjectEnvProjects1) ToPointer() *GetProjectEnvProjects1 {
	return &e
}

func (e *GetProjectEnvProjects1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		fallthrough
	case "development":
		*e = GetProjectEnvProjects1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvProjects1: %v", v)
	}
}

type GetProjectEnvTargetType string

const (
	GetProjectEnvTargetTypeArrayOfgetProjectEnvProjects1 GetProjectEnvTargetType = "arrayOfgetProjectEnv_projects_1"
	GetProjectEnvTargetTypeGetProjectEnvProjects2        GetProjectEnvTargetType = "getProjectEnv_projects_2"
)

type GetProjectEnvTarget struct {
	ArrayOfgetProjectEnvProjects1 []GetProjectEnvProjects1
	GetProjectEnvProjects2        *GetProjectEnvProjects2

	Type GetProjectEnvTargetType
}

func CreateGetProjectEnvTargetArrayOfgetProjectEnvProjects1(arrayOfgetProjectEnvProjects1 []GetProjectEnvProjects1) GetProjectEnvTarget {
	typ := GetProjectEnvTargetTypeArrayOfgetProjectEnvProjects1

	return GetProjectEnvTarget{
		ArrayOfgetProjectEnvProjects1: arrayOfgetProjectEnvProjects1,
		Type:                          typ,
	}
}

func CreateGetProjectEnvTargetGetProjectEnvProjects2(getProjectEnvProjects2 GetProjectEnvProjects2) GetProjectEnvTarget {
	typ := GetProjectEnvTargetTypeGetProjectEnvProjects2

	return GetProjectEnvTarget{
		GetProjectEnvProjects2: &getProjectEnvProjects2,
		Type:                   typ,
	}
}

func (u *GetProjectEnvTarget) UnmarshalJSON(data []byte) error {

	arrayOfgetProjectEnvProjects1 := []GetProjectEnvProjects1{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetProjectEnvProjects1, "", true, true); err == nil {
		u.ArrayOfgetProjectEnvProjects1 = arrayOfgetProjectEnvProjects1
		u.Type = GetProjectEnvTargetTypeArrayOfgetProjectEnvProjects1
		return nil
	}

	getProjectEnvProjects2 := GetProjectEnvProjects2("")
	if err := utils.UnmarshalJSON(data, &getProjectEnvProjects2, "", true, true); err == nil {
		u.GetProjectEnvProjects2 = &getProjectEnvProjects2
		u.Type = GetProjectEnvTargetTypeGetProjectEnvProjects2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetProjectEnvTarget) MarshalJSON() ([]byte, error) {
	if u.ArrayOfgetProjectEnvProjects1 != nil {
		return utils.MarshalJSON(u.ArrayOfgetProjectEnvProjects1, "", true)
	}

	if u.GetProjectEnvProjects2 != nil {
		return utils.MarshalJSON(u.GetProjectEnvProjects2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetProjectEnvType string

const (
	GetProjectEnvTypeSecret    GetProjectEnvType = "secret"
	GetProjectEnvTypeSystem    GetProjectEnvType = "system"
	GetProjectEnvTypeEncrypted GetProjectEnvType = "encrypted"
	GetProjectEnvTypePlain     GetProjectEnvType = "plain"
	GetProjectEnvTypeSensitive GetProjectEnvType = "sensitive"
)

func (e GetProjectEnvType) ToPointer() *GetProjectEnvType {
	return &e
}

func (e *GetProjectEnvType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "system":
		fallthrough
	case "encrypted":
		fallthrough
	case "plain":
		fallthrough
	case "sensitive":
		*e = GetProjectEnvType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProjectEnvType: %v", v)
	}
}

type GetProjectEnvResponseBody struct {
	ConfigurationID *string                   `json:"configurationId,omitempty"`
	ContentHint     *GetProjectEnvContentHint `json:"contentHint,omitempty"`
	CreatedAt       *int64                    `json:"createdAt,omitempty"`
	CreatedBy       *string                   `json:"createdBy,omitempty"`
	// Whether `value` is decrypted.
	Decrypted         *bool                `json:"decrypted,omitempty"`
	EdgeConfigID      *string              `json:"edgeConfigId,omitempty"`
	EdgeConfigTokenID *string              `json:"edgeConfigTokenId,omitempty"`
	GitBranch         *string              `json:"gitBranch,omitempty"`
	ID                *string              `json:"id,omitempty"`
	Key               string               `json:"key"`
	Target            *GetProjectEnvTarget `json:"target,omitempty"`
	Type              GetProjectEnvType    `json:"type"`
	UpdatedAt         *int64               `json:"updatedAt,omitempty"`
	UpdatedBy         *string              `json:"updatedBy,omitempty"`
	Value             string               `json:"value"`
}

func (o *GetProjectEnvResponseBody) GetConfigurationID() *string {
	if o == nil {
		return nil
	}
	return o.ConfigurationID
}

func (o *GetProjectEnvResponseBody) GetContentHint() *GetProjectEnvContentHint {
	if o == nil {
		return nil
	}
	return o.ContentHint
}

func (o *GetProjectEnvResponseBody) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetProjectEnvResponseBody) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *GetProjectEnvResponseBody) GetDecrypted() *bool {
	if o == nil {
		return nil
	}
	return o.Decrypted
}

func (o *GetProjectEnvResponseBody) GetEdgeConfigID() *string {
	if o == nil {
		return nil
	}
	return o.EdgeConfigID
}

func (o *GetProjectEnvResponseBody) GetEdgeConfigTokenID() *string {
	if o == nil {
		return nil
	}
	return o.EdgeConfigTokenID
}

func (o *GetProjectEnvResponseBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *GetProjectEnvResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectEnvResponseBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *GetProjectEnvResponseBody) GetTarget() *GetProjectEnvTarget {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *GetProjectEnvResponseBody) GetType() GetProjectEnvType {
	if o == nil {
		return GetProjectEnvType("")
	}
	return o.Type
}

func (o *GetProjectEnvResponseBody) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetProjectEnvResponseBody) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *GetProjectEnvResponseBody) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type GetProjectEnvResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Object      *GetProjectEnvResponseBody
}

func (o *GetProjectEnvResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProjectEnvResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProjectEnvResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProjectEnvResponse) GetObject() *GetProjectEnvResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}