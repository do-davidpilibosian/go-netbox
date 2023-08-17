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

// checks if the PatchedWritableConfigTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableConfigTemplateRequest{}

// PatchedWritableConfigTemplateRequest Introduces support for Tag assignment. Adds `tags` serialization, and handles tag assignment on create() and update().
type PatchedWritableConfigTemplateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Any <a href=\"https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment\">additional parameters</a> to pass when constructing the Jinja2 environment.
	EnvironmentParams map[string]interface{} `json:"environment_params,omitempty"`
	// Jinja2 template code.
	TemplateCode *string `json:"template_code,omitempty"`
	// Remote data source
	DataSource           NullableInt32      `json:"data_source,omitempty"`
	Tags                 []NestedTagRequest `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableConfigTemplateRequest PatchedWritableConfigTemplateRequest

// NewPatchedWritableConfigTemplateRequest instantiates a new PatchedWritableConfigTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableConfigTemplateRequest() *PatchedWritableConfigTemplateRequest {
	this := PatchedWritableConfigTemplateRequest{}
	return &this
}

// NewPatchedWritableConfigTemplateRequestWithDefaults instantiates a new PatchedWritableConfigTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableConfigTemplateRequestWithDefaults() *PatchedWritableConfigTemplateRequest {
	this := PatchedWritableConfigTemplateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableConfigTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableConfigTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableConfigTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableConfigTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableConfigTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableConfigTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironmentParams returns the EnvironmentParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConfigTemplateRequest) GetEnvironmentParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EnvironmentParams
}

// GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConfigTemplateRequest) GetEnvironmentParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentParams) {
		return map[string]interface{}{}, false
	}
	return o.EnvironmentParams, true
}

// HasEnvironmentParams returns a boolean if a field has been set.
func (o *PatchedWritableConfigTemplateRequest) HasEnvironmentParams() bool {
	if o != nil && IsNil(o.EnvironmentParams) {
		return true
	}

	return false
}

// SetEnvironmentParams gets a reference to the given map[string]interface{} and assigns it to the EnvironmentParams field.
func (o *PatchedWritableConfigTemplateRequest) SetEnvironmentParams(v map[string]interface{}) {
	o.EnvironmentParams = v
}

// GetTemplateCode returns the TemplateCode field value if set, zero value otherwise.
func (o *PatchedWritableConfigTemplateRequest) GetTemplateCode() string {
	if o == nil || IsNil(o.TemplateCode) {
		var ret string
		return ret
	}
	return *o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigTemplateRequest) GetTemplateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateCode) {
		return nil, false
	}
	return o.TemplateCode, true
}

// HasTemplateCode returns a boolean if a field has been set.
func (o *PatchedWritableConfigTemplateRequest) HasTemplateCode() bool {
	if o != nil && !IsNil(o.TemplateCode) {
		return true
	}

	return false
}

// SetTemplateCode gets a reference to the given string and assigns it to the TemplateCode field.
func (o *PatchedWritableConfigTemplateRequest) SetTemplateCode(v string) {
	o.TemplateCode = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConfigTemplateRequest) GetDataSource() int32 {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret int32
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConfigTemplateRequest) GetDataSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *PatchedWritableConfigTemplateRequest) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableInt32 and assigns it to the DataSource field.
func (o *PatchedWritableConfigTemplateRequest) SetDataSource(v int32) {
	o.DataSource.Set(&v)
}

// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *PatchedWritableConfigTemplateRequest) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *PatchedWritableConfigTemplateRequest) UnsetDataSource() {
	o.DataSource.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableConfigTemplateRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigTemplateRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableConfigTemplateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableConfigTemplateRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o PatchedWritableConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableConfigTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.EnvironmentParams != nil {
		toSerialize["environment_params"] = o.EnvironmentParams
	}
	if !IsNil(o.TemplateCode) {
		toSerialize["template_code"] = o.TemplateCode
	}
	if o.DataSource.IsSet() {
		toSerialize["data_source"] = o.DataSource.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableConfigTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableConfigTemplateRequest := _PatchedWritableConfigTemplateRequest{}

	if err = json.Unmarshal(bytes, &varPatchedWritableConfigTemplateRequest); err == nil {
		*o = PatchedWritableConfigTemplateRequest(varPatchedWritableConfigTemplateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "environment_params")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableConfigTemplateRequest struct {
	value *PatchedWritableConfigTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableConfigTemplateRequest) Get() *PatchedWritableConfigTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableConfigTemplateRequest) Set(val *PatchedWritableConfigTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConfigTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConfigTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConfigTemplateRequest(val *PatchedWritableConfigTemplateRequest) *NullablePatchedWritableConfigTemplateRequest {
	return &NullablePatchedWritableConfigTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConfigTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
