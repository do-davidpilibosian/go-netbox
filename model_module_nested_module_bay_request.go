/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.2 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the ModuleNestedModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleNestedModuleBayRequest{}

// ModuleNestedModuleBayRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type ModuleNestedModuleBayRequest struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ModuleNestedModuleBayRequest ModuleNestedModuleBayRequest

// NewModuleNestedModuleBayRequest instantiates a new ModuleNestedModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleNestedModuleBayRequest(name string) *ModuleNestedModuleBayRequest {
	this := ModuleNestedModuleBayRequest{}
	this.Name = name
	return &this
}

// NewModuleNestedModuleBayRequestWithDefaults instantiates a new ModuleNestedModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleNestedModuleBayRequestWithDefaults() *ModuleNestedModuleBayRequest {
	this := ModuleNestedModuleBayRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ModuleNestedModuleBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModuleNestedModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModuleNestedModuleBayRequest) SetName(v string) {
	o.Name = v
}

func (o ModuleNestedModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleNestedModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleNestedModuleBayRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModuleNestedModuleBayRequest := _ModuleNestedModuleBayRequest{}

	err = json.Unmarshal(bytes, &varModuleNestedModuleBayRequest)

	if err != nil {
		return err
	}

	*o = ModuleNestedModuleBayRequest(varModuleNestedModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleNestedModuleBayRequest struct {
	value *ModuleNestedModuleBayRequest
	isSet bool
}

func (v NullableModuleNestedModuleBayRequest) Get() *ModuleNestedModuleBayRequest {
	return v.value
}

func (v *NullableModuleNestedModuleBayRequest) Set(val *ModuleNestedModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleNestedModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleNestedModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleNestedModuleBayRequest(val *ModuleNestedModuleBayRequest) *NullableModuleNestedModuleBayRequest {
	return &NullableModuleNestedModuleBayRequest{value: val, isSet: true}
}

func (v NullableModuleNestedModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleNestedModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
