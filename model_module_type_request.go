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

// checks if the ModuleTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleTypeRequest{}

// ModuleTypeRequest Adds support for custom fields and tags.
type ModuleTypeRequest struct {
	Manufacturer         ManufacturerRequest `json:"manufacturer"`
	Model                string              `json:"model"`
	Description          *string             `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleTypeRequest ModuleTypeRequest

// NewModuleTypeRequest instantiates a new ModuleTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleTypeRequest(manufacturer ManufacturerRequest, model string) *ModuleTypeRequest {
	this := ModuleTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	return &this
}

// NewModuleTypeRequestWithDefaults instantiates a new ModuleTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleTypeRequestWithDefaults() *ModuleTypeRequest {
	this := ModuleTypeRequest{}
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *ModuleTypeRequest) GetManufacturer() ManufacturerRequest {
	if o == nil {
		var ret ManufacturerRequest
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *ModuleTypeRequest) GetManufacturerOk() (*ManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *ModuleTypeRequest) SetManufacturer(v ManufacturerRequest) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *ModuleTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ModuleTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ModuleTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModuleTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModuleTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModuleTypeRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ModuleTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"manufacturer",
		"model",
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

	varModuleTypeRequest := _ModuleTypeRequest{}

	err = json.Unmarshal(data, &varModuleTypeRequest)

	if err != nil {
		return err
	}

	*o = ModuleTypeRequest(varModuleTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleTypeRequest struct {
	value *ModuleTypeRequest
	isSet bool
}

func (v NullableModuleTypeRequest) Get() *ModuleTypeRequest {
	return v.value
}

func (v *NullableModuleTypeRequest) Set(val *ModuleTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleTypeRequest(val *ModuleTypeRequest) *NullableModuleTypeRequest {
	return &NullableModuleTypeRequest{value: val, isSet: true}
}

func (v NullableModuleTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
