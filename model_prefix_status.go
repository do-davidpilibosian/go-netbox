/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.8 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PrefixStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrefixStatus{}

// PrefixStatus struct for PrefixStatus
type PrefixStatus struct {
	// * `container` - Container * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrefixStatus PrefixStatus

// NewPrefixStatus instantiates a new PrefixStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrefixStatus() *PrefixStatus {
	this := PrefixStatus{}
	return &this
}

// NewPrefixStatusWithDefaults instantiates a new PrefixStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrefixStatusWithDefaults() *PrefixStatus {
	this := PrefixStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PrefixStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PrefixStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PrefixStatus) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PrefixStatus) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefixStatus) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PrefixStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PrefixStatus) SetLabel(v string) {
	o.Label = &v
}

func (o PrefixStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrefixStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrefixStatus) UnmarshalJSON(bytes []byte) (err error) {
	varPrefixStatus := _PrefixStatus{}

	if err = json.Unmarshal(bytes, &varPrefixStatus); err == nil {
		*o = PrefixStatus(varPrefixStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrefixStatus struct {
	value *PrefixStatus
	isSet bool
}

func (v NullablePrefixStatus) Get() *PrefixStatus {
	return v.value
}

func (v *NullablePrefixStatus) Set(val *PrefixStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefixStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefixStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefixStatus(val *PrefixStatus) *NullablePrefixStatus {
	return &NullablePrefixStatus{value: val, isSet: true}
}

func (v NullablePrefixStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefixStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
