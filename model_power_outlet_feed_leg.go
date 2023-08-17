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

// checks if the PowerOutletFeedLeg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerOutletFeedLeg{}

// PowerOutletFeedLeg struct for PowerOutletFeedLeg
type PowerOutletFeedLeg struct {
	// * `A` - A * `B` - B * `C` - C
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerOutletFeedLeg PowerOutletFeedLeg

// NewPowerOutletFeedLeg instantiates a new PowerOutletFeedLeg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOutletFeedLeg() *PowerOutletFeedLeg {
	this := PowerOutletFeedLeg{}
	return &this
}

// NewPowerOutletFeedLegWithDefaults instantiates a new PowerOutletFeedLeg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOutletFeedLegWithDefaults() *PowerOutletFeedLeg {
	this := PowerOutletFeedLeg{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerOutletFeedLeg) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletFeedLeg) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerOutletFeedLeg) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PowerOutletFeedLeg) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerOutletFeedLeg) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletFeedLeg) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerOutletFeedLeg) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerOutletFeedLeg) SetLabel(v string) {
	o.Label = &v
}

func (o PowerOutletFeedLeg) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerOutletFeedLeg) ToMap() (map[string]interface{}, error) {
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

func (o *PowerOutletFeedLeg) UnmarshalJSON(bytes []byte) (err error) {
	varPowerOutletFeedLeg := _PowerOutletFeedLeg{}

	if err = json.Unmarshal(bytes, &varPowerOutletFeedLeg); err == nil {
		*o = PowerOutletFeedLeg(varPowerOutletFeedLeg)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerOutletFeedLeg struct {
	value *PowerOutletFeedLeg
	isSet bool
}

func (v NullablePowerOutletFeedLeg) Get() *PowerOutletFeedLeg {
	return v.value
}

func (v *NullablePowerOutletFeedLeg) Set(val *PowerOutletFeedLeg) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletFeedLeg) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletFeedLeg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletFeedLeg(val *PowerOutletFeedLeg) *NullablePowerOutletFeedLeg {
	return &NullablePowerOutletFeedLeg{value: val, isSet: true}
}

func (v NullablePowerOutletFeedLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletFeedLeg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
