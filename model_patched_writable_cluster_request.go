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

// checks if the PatchedWritableClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableClusterRequest{}

// PatchedWritableClusterRequest Adds support for custom fields and tags.
type PatchedWritableClusterRequest struct {
	Name  *string       `json:"name,omitempty"`
	Type  *int32        `json:"type,omitempty"`
	Group NullableInt32 `json:"group,omitempty"`
	// * `planned` - Planned * `staging` - Staging * `active` - Active * `decommissioning` - Decommissioning * `offline` - Offline
	Status               *string                `json:"status,omitempty"`
	Tenant               NullableInt32          `json:"tenant,omitempty"`
	Site                 NullableInt32          `json:"site,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableClusterRequest PatchedWritableClusterRequest

// NewPatchedWritableClusterRequest instantiates a new PatchedWritableClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableClusterRequest() *PatchedWritableClusterRequest {
	this := PatchedWritableClusterRequest{}
	return &this
}

// NewPatchedWritableClusterRequestWithDefaults instantiates a new PatchedWritableClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableClusterRequestWithDefaults() *PatchedWritableClusterRequest {
	this := PatchedWritableClusterRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableClusterRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *PatchedWritableClusterRequest) SetType(v int32) {
	o.Type = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableClusterRequest) GetGroup() int32 {
	if o == nil || IsNil(o.Group.Get()) {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableClusterRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *PatchedWritableClusterRequest) SetGroup(v int32) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PatchedWritableClusterRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PatchedWritableClusterRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchedWritableClusterRequest) SetStatus(v string) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableClusterRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableClusterRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableClusterRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableClusterRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableClusterRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableClusterRequest) GetSite() int32 {
	if o == nil || IsNil(o.Site.Get()) {
		var ret int32
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableClusterRequest) GetSiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableInt32 and assigns it to the Site field.
func (o *PatchedWritableClusterRequest) SetSite(v int32) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *PatchedWritableClusterRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *PatchedWritableClusterRequest) UnsetSite() {
	o.Site.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableClusterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableClusterRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableClusterRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableClusterRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableClusterRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableClusterRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableClusterRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
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

func (o *PatchedWritableClusterRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableClusterRequest := _PatchedWritableClusterRequest{}

	err = json.Unmarshal(bytes, &varPatchedWritableClusterRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableClusterRequest(varPatchedWritableClusterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "site")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableClusterRequest struct {
	value *PatchedWritableClusterRequest
	isSet bool
}

func (v NullablePatchedWritableClusterRequest) Get() *PatchedWritableClusterRequest {
	return v.value
}

func (v *NullablePatchedWritableClusterRequest) Set(val *PatchedWritableClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableClusterRequest(val *PatchedWritableClusterRequest) *NullablePatchedWritableClusterRequest {
	return &NullablePatchedWritableClusterRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
