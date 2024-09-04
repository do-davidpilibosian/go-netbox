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

// checks if the PatchedWritableConsoleServerPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableConsoleServerPortTemplateRequest{}

// PatchedWritableConsoleServerPortTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableConsoleServerPortTemplateRequest struct {
	DeviceType NullableDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label                *string               `json:"label,omitempty"`
	Type                 *ConsolePortTypeValue `json:"type,omitempty"`
	Description          *string               `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableConsoleServerPortTemplateRequest PatchedWritableConsoleServerPortTemplateRequest

// NewPatchedWritableConsoleServerPortTemplateRequest instantiates a new PatchedWritableConsoleServerPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableConsoleServerPortTemplateRequest() *PatchedWritableConsoleServerPortTemplateRequest {
	this := PatchedWritableConsoleServerPortTemplateRequest{}
	return &this
}

// NewPatchedWritableConsoleServerPortTemplateRequestWithDefaults instantiates a new PatchedWritableConsoleServerPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableConsoleServerPortTemplateRequestWithDefaults() *PatchedWritableConsoleServerPortTemplateRequest {
	this := PatchedWritableConsoleServerPortTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetDeviceType() DeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret DeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetDeviceTypeOk() (*DeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableDeviceTypeRequest and assigns it to the DeviceType field.
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetDeviceType(v DeviceTypeRequest) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PatchedWritableConsoleServerPortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetModuleType() ModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret ModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetModuleTypeOk() (*ModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableModuleTypeRequest and assigns it to the ModuleType field.
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetModuleType(v ModuleTypeRequest) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PatchedWritableConsoleServerPortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetType() ConsolePortTypeValue {
	if o == nil || IsNil(o.Type) {
		var ret ConsolePortTypeValue
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetTypeOk() (*ConsolePortTypeValue, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortTypeValue and assigns it to the Type field.
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetType(v ConsolePortTypeValue) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableConsoleServerPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableConsoleServerPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedWritableConsoleServerPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableConsoleServerPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableConsoleServerPortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableConsoleServerPortTemplateRequest := _PatchedWritableConsoleServerPortTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedWritableConsoleServerPortTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableConsoleServerPortTemplateRequest(varPatchedWritableConsoleServerPortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableConsoleServerPortTemplateRequest struct {
	value *PatchedWritableConsoleServerPortTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableConsoleServerPortTemplateRequest) Get() *PatchedWritableConsoleServerPortTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableConsoleServerPortTemplateRequest) Set(val *PatchedWritableConsoleServerPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConsoleServerPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConsoleServerPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConsoleServerPortTemplateRequest(val *PatchedWritableConsoleServerPortTemplateRequest) *NullablePatchedWritableConsoleServerPortTemplateRequest {
	return &NullablePatchedWritableConsoleServerPortTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableConsoleServerPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConsoleServerPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
