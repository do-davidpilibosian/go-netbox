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

// checks if the ModuleBayTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleBayTemplateRequest{}

// ModuleBayTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ModuleBayTemplateRequest struct {
	DeviceType NestedDeviceTypeRequest `json:"device_type"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Identifier to reference when renaming installed components
	Position             *string `json:"position,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleBayTemplateRequest ModuleBayTemplateRequest

// NewModuleBayTemplateRequest instantiates a new ModuleBayTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleBayTemplateRequest(deviceType NestedDeviceTypeRequest, name string) *ModuleBayTemplateRequest {
	this := ModuleBayTemplateRequest{}
	this.DeviceType = deviceType
	this.Name = name
	return &this
}

// NewModuleBayTemplateRequestWithDefaults instantiates a new ModuleBayTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleBayTemplateRequestWithDefaults() *ModuleBayTemplateRequest {
	this := ModuleBayTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value
func (o *ModuleBayTemplateRequest) GetDeviceType() NestedDeviceTypeRequest {
	if o == nil {
		var ret NestedDeviceTypeRequest
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *ModuleBayTemplateRequest) GetDeviceTypeOk() (*NestedDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *ModuleBayTemplateRequest) SetDeviceType(v NestedDeviceTypeRequest) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *ModuleBayTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModuleBayTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModuleBayTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModuleBayTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModuleBayTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ModuleBayTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ModuleBayTemplateRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayTemplateRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ModuleBayTemplateRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ModuleBayTemplateRequest) SetPosition(v string) {
	o.Position = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModuleBayTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModuleBayTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModuleBayTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ModuleBayTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleBayTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_type"] = o.DeviceType
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleBayTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModuleBayTemplateRequest := _ModuleBayTemplateRequest{}

	err = json.Unmarshal(bytes, &varModuleBayTemplateRequest)

	if err != nil {
		return err
	}

	*o = ModuleBayTemplateRequest(varModuleBayTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "position")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleBayTemplateRequest struct {
	value *ModuleBayTemplateRequest
	isSet bool
}

func (v NullableModuleBayTemplateRequest) Get() *ModuleBayTemplateRequest {
	return v.value
}

func (v *NullableModuleBayTemplateRequest) Set(val *ModuleBayTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleBayTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleBayTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleBayTemplateRequest(val *ModuleBayTemplateRequest) *NullableModuleBayTemplateRequest {
	return &NullableModuleBayTemplateRequest{value: val, isSet: true}
}

func (v NullableModuleBayTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleBayTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
