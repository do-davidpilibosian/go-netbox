/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritableCustomFieldRequestType The type of data this custom field holds  * `text` - Text * `longtext` - Text (long) * `integer` - Integer * `decimal` - Decimal * `boolean` - Boolean (true/false) * `date` - Date * `datetime` - Date & time * `url` - URL * `json` - JSON * `select` - Selection * `multiselect` - Multiple selection * `object` - Object * `multiobject` - Multiple objects
type PatchedWritableCustomFieldRequestType string

// List of PatchedWritableCustomFieldRequest_type
const (
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_TEXT        PatchedWritableCustomFieldRequestType = "text"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_LONGTEXT    PatchedWritableCustomFieldRequestType = "longtext"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_INTEGER     PatchedWritableCustomFieldRequestType = "integer"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_DECIMAL     PatchedWritableCustomFieldRequestType = "decimal"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_BOOLEAN     PatchedWritableCustomFieldRequestType = "boolean"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_DATE        PatchedWritableCustomFieldRequestType = "date"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_DATETIME    PatchedWritableCustomFieldRequestType = "datetime"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_URL         PatchedWritableCustomFieldRequestType = "url"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_JSON        PatchedWritableCustomFieldRequestType = "json"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_SELECT      PatchedWritableCustomFieldRequestType = "select"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_MULTISELECT PatchedWritableCustomFieldRequestType = "multiselect"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_OBJECT      PatchedWritableCustomFieldRequestType = "object"
	PATCHEDWRITABLECUSTOMFIELDREQUESTTYPE_MULTIOBJECT PatchedWritableCustomFieldRequestType = "multiobject"
)

// All allowed values of PatchedWritableCustomFieldRequestType enum
var AllowedPatchedWritableCustomFieldRequestTypeEnumValues = []PatchedWritableCustomFieldRequestType{
	"text",
	"longtext",
	"integer",
	"decimal",
	"boolean",
	"date",
	"datetime",
	"url",
	"json",
	"select",
	"multiselect",
	"object",
	"multiobject",
}

func (v *PatchedWritableCustomFieldRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableCustomFieldRequestType(value)
	for _, existing := range AllowedPatchedWritableCustomFieldRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableCustomFieldRequestType", value)
}

// NewPatchedWritableCustomFieldRequestTypeFromValue returns a pointer to a valid PatchedWritableCustomFieldRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableCustomFieldRequestTypeFromValue(v string) (*PatchedWritableCustomFieldRequestType, error) {
	ev := PatchedWritableCustomFieldRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableCustomFieldRequestType: valid values are %v", v, AllowedPatchedWritableCustomFieldRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableCustomFieldRequestType) IsValid() bool {
	for _, existing := range AllowedPatchedWritableCustomFieldRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableCustomFieldRequest_type value
func (v PatchedWritableCustomFieldRequestType) Ptr() *PatchedWritableCustomFieldRequestType {
	return &v
}

type NullablePatchedWritableCustomFieldRequestType struct {
	value *PatchedWritableCustomFieldRequestType
	isSet bool
}

func (v NullablePatchedWritableCustomFieldRequestType) Get() *PatchedWritableCustomFieldRequestType {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldRequestType) Set(val *PatchedWritableCustomFieldRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldRequestType(val *PatchedWritableCustomFieldRequestType) *NullablePatchedWritableCustomFieldRequestType {
	return &NullablePatchedWritableCustomFieldRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
