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

// checks if the WritableDeviceRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableDeviceRoleRequest{}

// WritableDeviceRoleRequest Adds support for custom fields and tags.
type WritableDeviceRoleRequest struct {
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Color *string `json:"color,omitempty"`
	// Virtual machines may be assigned to this role
	VmRole               *bool                  `json:"vm_role,omitempty"`
	ConfigTemplate       NullableInt32          `json:"config_template,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableDeviceRoleRequest WritableDeviceRoleRequest

// NewWritableDeviceRoleRequest instantiates a new WritableDeviceRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableDeviceRoleRequest(name string, slug string) *WritableDeviceRoleRequest {
	this := WritableDeviceRoleRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewWritableDeviceRoleRequestWithDefaults instantiates a new WritableDeviceRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableDeviceRoleRequestWithDefaults() *WritableDeviceRoleRequest {
	this := WritableDeviceRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableDeviceRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableDeviceRoleRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritableDeviceRoleRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritableDeviceRoleRequest) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *WritableDeviceRoleRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *WritableDeviceRoleRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *WritableDeviceRoleRequest) SetColor(v string) {
	o.Color = &v
}

// GetVmRole returns the VmRole field value if set, zero value otherwise.
func (o *WritableDeviceRoleRequest) GetVmRole() bool {
	if o == nil || IsNil(o.VmRole) {
		var ret bool
		return ret
	}
	return *o.VmRole
}

// GetVmRoleOk returns a tuple with the VmRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetVmRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.VmRole) {
		return nil, false
	}
	return o.VmRole, true
}

// HasVmRole returns a boolean if a field has been set.
func (o *WritableDeviceRoleRequest) HasVmRole() bool {
	if o != nil && !IsNil(o.VmRole) {
		return true
	}

	return false
}

// SetVmRole gets a reference to the given bool and assigns it to the VmRole field.
func (o *WritableDeviceRoleRequest) SetVmRole(v bool) {
	o.VmRole = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableDeviceRoleRequest) GetConfigTemplate() int32 {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret int32
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableDeviceRoleRequest) GetConfigTemplateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *WritableDeviceRoleRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableInt32 and assigns it to the ConfigTemplate field.
func (o *WritableDeviceRoleRequest) SetConfigTemplate(v int32) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *WritableDeviceRoleRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *WritableDeviceRoleRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableDeviceRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableDeviceRoleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableDeviceRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableDeviceRoleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableDeviceRoleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableDeviceRoleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableDeviceRoleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceRoleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableDeviceRoleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableDeviceRoleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableDeviceRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableDeviceRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.VmRole) {
		toSerialize["vm_role"] = o.VmRole
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableDeviceRoleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWritableDeviceRoleRequest := _WritableDeviceRoleRequest{}

	err = json.Unmarshal(bytes, &varWritableDeviceRoleRequest)

	if err != nil {
		return err
	}

	*o = WritableDeviceRoleRequest(varWritableDeviceRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "vm_role")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableDeviceRoleRequest struct {
	value *WritableDeviceRoleRequest
	isSet bool
}

func (v NullableWritableDeviceRoleRequest) Get() *WritableDeviceRoleRequest {
	return v.value
}

func (v *NullableWritableDeviceRoleRequest) Set(val *WritableDeviceRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableDeviceRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableDeviceRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableDeviceRoleRequest(val *WritableDeviceRoleRequest) *NullableWritableDeviceRoleRequest {
	return &NullableWritableDeviceRoleRequest{value: val, isSet: true}
}

func (v NullableWritableDeviceRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableDeviceRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
