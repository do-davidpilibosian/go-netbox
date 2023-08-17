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

// checks if the WritableModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableModuleRequest{}

// WritableModuleRequest Adds support for custom fields and tags.
type WritableModuleRequest struct {
	Device     int32 `json:"device"`
	ModuleBay  int32 `json:"module_bay"`
	ModuleType int32 `json:"module_type"`
	// * `offline` - Offline * `active` - Active * `planned` - Planned * `staged` - Staged * `failed` - Failed * `decommissioning` - Decommissioning
	Status *string `json:"status,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag             NullableString         `json:"asset_tag,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableModuleRequest WritableModuleRequest

// NewWritableModuleRequest instantiates a new WritableModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableModuleRequest(device int32, moduleBay int32, moduleType int32) *WritableModuleRequest {
	this := WritableModuleRequest{}
	this.Device = device
	this.ModuleBay = moduleBay
	this.ModuleType = moduleType
	return &this
}

// NewWritableModuleRequestWithDefaults instantiates a new WritableModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableModuleRequestWithDefaults() *WritableModuleRequest {
	this := WritableModuleRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *WritableModuleRequest) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritableModuleRequest) SetDevice(v int32) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *WritableModuleRequest) GetModuleBay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetModuleBayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *WritableModuleRequest) SetModuleBay(v int32) {
	o.ModuleBay = v
}

// GetModuleType returns the ModuleType field value
func (o *WritableModuleRequest) GetModuleType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetModuleTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleType, true
}

// SetModuleType sets field value
func (o *WritableModuleRequest) SetModuleType(v int32) {
	o.ModuleType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableModuleRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WritableModuleRequest) SetStatus(v string) {
	o.Status = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *WritableModuleRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *WritableModuleRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableModuleRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableModuleRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *WritableModuleRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *WritableModuleRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *WritableModuleRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableModuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableModuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableModuleRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableModuleRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableModuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableModuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableModuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableModuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableModuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["module_bay"] = o.ModuleBay
	toSerialize["module_type"] = o.ModuleType
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *WritableModuleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWritableModuleRequest := _WritableModuleRequest{}

	if err = json.Unmarshal(bytes, &varWritableModuleRequest); err == nil {
		*o = WritableModuleRequest(varWritableModuleRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableModuleRequest struct {
	value *WritableModuleRequest
	isSet bool
}

func (v NullableWritableModuleRequest) Get() *WritableModuleRequest {
	return v.value
}

func (v *NullableWritableModuleRequest) Set(val *WritableModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableModuleRequest(val *WritableModuleRequest) *NullableWritableModuleRequest {
	return &NullableWritableModuleRequest{value: val, isSet: true}
}

func (v NullableWritableModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
