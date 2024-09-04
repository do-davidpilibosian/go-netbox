/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the CableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableRequest{}

// CableRequest Adds support for custom fields and tags.
type CableRequest struct {
	Label                *string `json:"label,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CableRequest CableRequest

// NewCableRequest instantiates a new CableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableRequest() *CableRequest {
	this := CableRequest{}
	return &this
}

// NewCableRequestWithDefaults instantiates a new CableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableRequestWithDefaults() *CableRequest {
	this := CableRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CableRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CableRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CableRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CableRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CableRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CableRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CableRequest) UnmarshalJSON(data []byte) (err error) {
	varCableRequest := _CableRequest{}

	err = json.Unmarshal(data, &varCableRequest)

	if err != nil {
		return err
	}

	*o = CableRequest(varCableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCableRequest struct {
	value *CableRequest
	isSet bool
}

func (v NullableCableRequest) Get() *CableRequest {
	return v.value
}

func (v *NullableCableRequest) Set(val *CableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableRequest(val *CableRequest) *NullableCableRequest {
	return &NullableCableRequest{value: val, isSet: true}
}

func (v NullableCableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}