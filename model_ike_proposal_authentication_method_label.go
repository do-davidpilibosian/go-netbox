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

// IKEProposalAuthenticationMethodLabel the model 'IKEProposalAuthenticationMethodLabel'
type IKEProposalAuthenticationMethodLabel string

// List of IKEProposal_authentication_method_label
const (
	IKEPROPOSALAUTHENTICATIONMETHODLABEL_PRE_SHARED_KEYS IKEProposalAuthenticationMethodLabel = "Pre-shared keys"
	IKEPROPOSALAUTHENTICATIONMETHODLABEL_CERTIFICATES    IKEProposalAuthenticationMethodLabel = "Certificates"
	IKEPROPOSALAUTHENTICATIONMETHODLABEL_RSA_SIGNATURES  IKEProposalAuthenticationMethodLabel = "RSA signatures"
	IKEPROPOSALAUTHENTICATIONMETHODLABEL_DSA_SIGNATURES  IKEProposalAuthenticationMethodLabel = "DSA signatures"
)

// All allowed values of IKEProposalAuthenticationMethodLabel enum
var AllowedIKEProposalAuthenticationMethodLabelEnumValues = []IKEProposalAuthenticationMethodLabel{
	"Pre-shared keys",
	"Certificates",
	"RSA signatures",
	"DSA signatures",
}

func (v *IKEProposalAuthenticationMethodLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalAuthenticationMethodLabel(value)
	for _, existing := range AllowedIKEProposalAuthenticationMethodLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalAuthenticationMethodLabel", value)
}

// NewIKEProposalAuthenticationMethodLabelFromValue returns a pointer to a valid IKEProposalAuthenticationMethodLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalAuthenticationMethodLabelFromValue(v string) (*IKEProposalAuthenticationMethodLabel, error) {
	ev := IKEProposalAuthenticationMethodLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalAuthenticationMethodLabel: valid values are %v", v, AllowedIKEProposalAuthenticationMethodLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalAuthenticationMethodLabel) IsValid() bool {
	for _, existing := range AllowedIKEProposalAuthenticationMethodLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_authentication_method_label value
func (v IKEProposalAuthenticationMethodLabel) Ptr() *IKEProposalAuthenticationMethodLabel {
	return &v
}

type NullableIKEProposalAuthenticationMethodLabel struct {
	value *IKEProposalAuthenticationMethodLabel
	isSet bool
}

func (v NullableIKEProposalAuthenticationMethodLabel) Get() *IKEProposalAuthenticationMethodLabel {
	return v.value
}

func (v *NullableIKEProposalAuthenticationMethodLabel) Set(val *IKEProposalAuthenticationMethodLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalAuthenticationMethodLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalAuthenticationMethodLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalAuthenticationMethodLabel(val *IKEProposalAuthenticationMethodLabel) *NullableIKEProposalAuthenticationMethodLabel {
	return &NullableIKEProposalAuthenticationMethodLabel{value: val, isSet: true}
}

func (v NullableIKEProposalAuthenticationMethodLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalAuthenticationMethodLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
