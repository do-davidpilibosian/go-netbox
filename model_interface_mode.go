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

// checks if the InterfaceMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceMode{}

// InterfaceMode struct for InterfaceMode
type InterfaceMode struct {
	// * `access` - Access * `tagged` - Tagged * `tagged-all` - Tagged (All)
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceMode InterfaceMode

// NewInterfaceMode instantiates a new InterfaceMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceMode() *InterfaceMode {
	this := InterfaceMode{}
	return &this
}

// NewInterfaceModeWithDefaults instantiates a new InterfaceMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceModeWithDefaults() *InterfaceMode {
	this := InterfaceMode{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceMode) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceMode) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceMode) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InterfaceMode) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceMode) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceMode) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceMode) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceMode) SetLabel(v string) {
	o.Label = &v
}

func (o InterfaceMode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceMode) ToMap() (map[string]interface{}, error) {
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

func (o *InterfaceMode) UnmarshalJSON(bytes []byte) (err error) {
	varInterfaceMode := _InterfaceMode{}

	if err = json.Unmarshal(bytes, &varInterfaceMode); err == nil {
		*o = InterfaceMode(varInterfaceMode)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceMode struct {
	value *InterfaceMode
	isSet bool
}

func (v NullableInterfaceMode) Get() *InterfaceMode {
	return v.value
}

func (v *NullableInterfaceMode) Set(val *InterfaceMode) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceMode) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceMode(val *InterfaceMode) *NullableInterfaceMode {
	return &NullableInterfaceMode{value: val, isSet: true}
}

func (v NullableInterfaceMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
