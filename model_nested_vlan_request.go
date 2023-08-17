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

// checks if the NestedVLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVLANRequest{}

// NestedVLANRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVLANRequest struct {
	// Numeric VLAN ID (1-4094)
	Vid                  int32  `json:"vid"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedVLANRequest NestedVLANRequest

// NewNestedVLANRequest instantiates a new NestedVLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVLANRequest(vid int32, name string) *NestedVLANRequest {
	this := NestedVLANRequest{}
	this.Vid = vid
	this.Name = name
	return &this
}

// NewNestedVLANRequestWithDefaults instantiates a new NestedVLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVLANRequestWithDefaults() *NestedVLANRequest {
	this := NestedVLANRequest{}
	return &this
}

// GetVid returns the Vid field value
func (o *NestedVLANRequest) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *NestedVLANRequest) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *NestedVLANRequest) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *NestedVLANRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVLANRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVLANRequest) SetName(v string) {
	o.Name = v
}

func (o NestedVLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vid"] = o.Vid
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVLANRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedVLANRequest := _NestedVLANRequest{}

	if err = json.Unmarshal(bytes, &varNestedVLANRequest); err == nil {
		*o = NestedVLANRequest(varNestedVLANRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVLANRequest struct {
	value *NestedVLANRequest
	isSet bool
}

func (v NullableNestedVLANRequest) Get() *NestedVLANRequest {
	return v.value
}

func (v *NullableNestedVLANRequest) Set(val *NestedVLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVLANRequest(val *NestedVLANRequest) *NullableNestedVLANRequest {
	return &NullableNestedVLANRequest{value: val, isSet: true}
}

func (v NullableNestedVLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
