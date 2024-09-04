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

// JournalEntryKindLabel the model 'JournalEntryKindLabel'
type JournalEntryKindLabel string

// List of JournalEntry_kind_label
const (
	JOURNALENTRYKINDLABEL_INFO    JournalEntryKindLabel = "Info"
	JOURNALENTRYKINDLABEL_SUCCESS JournalEntryKindLabel = "Success"
	JOURNALENTRYKINDLABEL_WARNING JournalEntryKindLabel = "Warning"
	JOURNALENTRYKINDLABEL_DANGER  JournalEntryKindLabel = "Danger"
)

// All allowed values of JournalEntryKindLabel enum
var AllowedJournalEntryKindLabelEnumValues = []JournalEntryKindLabel{
	"Info",
	"Success",
	"Warning",
	"Danger",
}

func (v *JournalEntryKindLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JournalEntryKindLabel(value)
	for _, existing := range AllowedJournalEntryKindLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JournalEntryKindLabel", value)
}

// NewJournalEntryKindLabelFromValue returns a pointer to a valid JournalEntryKindLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJournalEntryKindLabelFromValue(v string) (*JournalEntryKindLabel, error) {
	ev := JournalEntryKindLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JournalEntryKindLabel: valid values are %v", v, AllowedJournalEntryKindLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JournalEntryKindLabel) IsValid() bool {
	for _, existing := range AllowedJournalEntryKindLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JournalEntry_kind_label value
func (v JournalEntryKindLabel) Ptr() *JournalEntryKindLabel {
	return &v
}

type NullableJournalEntryKindLabel struct {
	value *JournalEntryKindLabel
	isSet bool
}

func (v NullableJournalEntryKindLabel) Get() *JournalEntryKindLabel {
	return v.value
}

func (v *NullableJournalEntryKindLabel) Set(val *JournalEntryKindLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalEntryKindLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalEntryKindLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalEntryKindLabel(val *JournalEntryKindLabel) *NullableJournalEntryKindLabel {
	return &NullableJournalEntryKindLabel{value: val, isSet: true}
}

func (v NullableJournalEntryKindLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalEntryKindLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
