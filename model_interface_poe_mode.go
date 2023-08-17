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

// checks if the InterfacePoeMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfacePoeMode{}

// InterfacePoeMode struct for InterfacePoeMode
type InterfacePoeMode struct {
	// * `pd` - PD * `pse` - PSE
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfacePoeMode InterfacePoeMode

// NewInterfacePoeMode instantiates a new InterfacePoeMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfacePoeMode() *InterfacePoeMode {
	this := InterfacePoeMode{}
	return &this
}

// NewInterfacePoeModeWithDefaults instantiates a new InterfacePoeMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfacePoeModeWithDefaults() *InterfacePoeMode {
	this := InterfacePoeMode{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfacePoeMode) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfacePoeMode) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfacePoeMode) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InterfacePoeMode) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfacePoeMode) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfacePoeMode) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfacePoeMode) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfacePoeMode) SetLabel(v string) {
	o.Label = &v
}

func (o InterfacePoeMode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfacePoeMode) ToMap() (map[string]interface{}, error) {
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

func (o *InterfacePoeMode) UnmarshalJSON(bytes []byte) (err error) {
	varInterfacePoeMode := _InterfacePoeMode{}

	if err = json.Unmarshal(bytes, &varInterfacePoeMode); err == nil {
		*o = InterfacePoeMode(varInterfacePoeMode)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfacePoeMode struct {
	value *InterfacePoeMode
	isSet bool
}

func (v NullableInterfacePoeMode) Get() *InterfacePoeMode {
	return v.value
}

func (v *NullableInterfacePoeMode) Set(val *InterfacePoeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfacePoeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfacePoeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfacePoeMode(val *InterfacePoeMode) *NullableInterfacePoeMode {
	return &NullableInterfacePoeMode{value: val, isSet: true}
}

func (v NullableInterfacePoeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfacePoeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
