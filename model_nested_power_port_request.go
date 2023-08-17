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

// checks if the NestedPowerPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedPowerPortRequest{}

// NestedPowerPortRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedPowerPortRequest struct {
	Name                 string        `json:"name"`
	Cable                NullableInt32 `json:"cable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedPowerPortRequest NestedPowerPortRequest

// NewNestedPowerPortRequest instantiates a new NestedPowerPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedPowerPortRequest(name string) *NestedPowerPortRequest {
	this := NestedPowerPortRequest{}
	this.Name = name
	return &this
}

// NewNestedPowerPortRequestWithDefaults instantiates a new NestedPowerPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedPowerPortRequestWithDefaults() *NestedPowerPortRequest {
	this := NestedPowerPortRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedPowerPortRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPortRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedPowerPortRequest) SetName(v string) {
	o.Name = v
}

// GetCable returns the Cable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedPowerPortRequest) GetCable() int32 {
	if o == nil || IsNil(o.Cable.Get()) {
		var ret int32
		return ret
	}
	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedPowerPortRequest) GetCableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// HasCable returns a boolean if a field has been set.
func (o *NestedPowerPortRequest) HasCable() bool {
	if o != nil && o.Cable.IsSet() {
		return true
	}

	return false
}

// SetCable gets a reference to the given NullableInt32 and assigns it to the Cable field.
func (o *NestedPowerPortRequest) SetCable(v int32) {
	o.Cable.Set(&v)
}

// SetCableNil sets the value for Cable to be an explicit nil
func (o *NestedPowerPortRequest) SetCableNil() {
	o.Cable.Set(nil)
}

// UnsetCable ensures that no value is present for Cable, not even an explicit nil
func (o *NestedPowerPortRequest) UnsetCable() {
	o.Cable.Unset()
}

func (o NestedPowerPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedPowerPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Cable.IsSet() {
		toSerialize["cable"] = o.Cable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedPowerPortRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedPowerPortRequest := _NestedPowerPortRequest{}

	if err = json.Unmarshal(bytes, &varNestedPowerPortRequest); err == nil {
		*o = NestedPowerPortRequest(varNestedPowerPortRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "cable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedPowerPortRequest struct {
	value *NestedPowerPortRequest
	isSet bool
}

func (v NullableNestedPowerPortRequest) Get() *NestedPowerPortRequest {
	return v.value
}

func (v *NullableNestedPowerPortRequest) Set(val *NestedPowerPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedPowerPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedPowerPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedPowerPortRequest(val *NestedPowerPortRequest) *NullableNestedPowerPortRequest {
	return &NullableNestedPowerPortRequest{value: val, isSet: true}
}

func (v NullableNestedPowerPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedPowerPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
