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

// checks if the IPAddressStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPAddressStatus{}

// IPAddressStatus struct for IPAddressStatus
type IPAddressStatus struct {
	// * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated * `dhcp` - DHCP * `slaac` - SLAAC
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPAddressStatus IPAddressStatus

// NewIPAddressStatus instantiates a new IPAddressStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAddressStatus() *IPAddressStatus {
	this := IPAddressStatus{}
	return &this
}

// NewIPAddressStatusWithDefaults instantiates a new IPAddressStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAddressStatusWithDefaults() *IPAddressStatus {
	this := IPAddressStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IPAddressStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IPAddressStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IPAddressStatus) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IPAddressStatus) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressStatus) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IPAddressStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IPAddressStatus) SetLabel(v string) {
	o.Label = &v
}

func (o IPAddressStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPAddressStatus) ToMap() (map[string]interface{}, error) {
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

func (o *IPAddressStatus) UnmarshalJSON(bytes []byte) (err error) {
	varIPAddressStatus := _IPAddressStatus{}

	if err = json.Unmarshal(bytes, &varIPAddressStatus); err == nil {
		*o = IPAddressStatus(varIPAddressStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPAddressStatus struct {
	value *IPAddressStatus
	isSet bool
}

func (v NullableIPAddressStatus) Get() *IPAddressStatus {
	return v.value
}

func (v *NullableIPAddressStatus) Set(val *IPAddressStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressStatus(val *IPAddressStatus) *NullableIPAddressStatus {
	return &NullableIPAddressStatus{value: val, isSet: true}
}

func (v NullableIPAddressStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
