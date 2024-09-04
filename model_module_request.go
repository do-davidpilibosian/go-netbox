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

// checks if the ModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleRequest{}

// ModuleRequest Adds support for custom fields and tags.
type ModuleRequest struct {
	Device               DeviceRequest          `json:"device"`
	ModuleBay            NestedModuleBayRequest `json:"module_bay"`
	AdditionalProperties map[string]interface{}
}

type _ModuleRequest ModuleRequest

// NewModuleRequest instantiates a new ModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleRequest(device DeviceRequest, moduleBay NestedModuleBayRequest) *ModuleRequest {
	this := ModuleRequest{}
	this.Device = device
	this.ModuleBay = moduleBay
	return &this
}

// NewModuleRequestWithDefaults instantiates a new ModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleRequestWithDefaults() *ModuleRequest {
	this := ModuleRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *ModuleRequest) GetDevice() DeviceRequest {
	if o == nil {
		var ret DeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetDeviceOk() (*DeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ModuleRequest) SetDevice(v DeviceRequest) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *ModuleRequest) GetModuleBay() NestedModuleBayRequest {
	if o == nil {
		var ret NestedModuleBayRequest
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *ModuleRequest) SetModuleBay(v NestedModuleBayRequest) {
	o.ModuleBay = v
}

func (o ModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["module_bay"] = o.ModuleBay

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"module_bay",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModuleRequest := _ModuleRequest{}

	err = json.Unmarshal(data, &varModuleRequest)

	if err != nil {
		return err
	}

	*o = ModuleRequest(varModuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleRequest struct {
	value *ModuleRequest
	isSet bool
}

func (v NullableModuleRequest) Get() *ModuleRequest {
	return v.value
}

func (v *NullableModuleRequest) Set(val *ModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleRequest(val *ModuleRequest) *NullableModuleRequest {
	return &NullableModuleRequest{value: val, isSet: true}
}

func (v NullableModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}