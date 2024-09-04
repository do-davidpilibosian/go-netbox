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

// checks if the PatchedPowerPanelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedPowerPanelRequest{}

// PatchedPowerPanelRequest Adds support for custom fields and tags.
type PatchedPowerPanelRequest struct {
	Site                 *SiteRequest            `json:"site,omitempty"`
	Location             NullableLocationRequest `json:"location,omitempty"`
	Name                 *string                 `json:"name,omitempty"`
	Description          *string                 `json:"description,omitempty"`
	Comments             *string                 `json:"comments,omitempty"`
	Tags                 []NestedTagRequest      `json:"tags,omitempty"`
	CustomFields         map[string]interface{}  `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedPowerPanelRequest PatchedPowerPanelRequest

// NewPatchedPowerPanelRequest instantiates a new PatchedPowerPanelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPowerPanelRequest() *PatchedPowerPanelRequest {
	this := PatchedPowerPanelRequest{}
	return &this
}

// NewPatchedPowerPanelRequestWithDefaults instantiates a new PatchedPowerPanelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPowerPanelRequestWithDefaults() *PatchedPowerPanelRequest {
	this := PatchedPowerPanelRequest{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *PatchedPowerPanelRequest) GetSite() SiteRequest {
	if o == nil || IsNil(o.Site) {
		var ret SiteRequest
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPowerPanelRequest) GetSiteOk() (*SiteRequest, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given SiteRequest and assigns it to the Site field.
func (o *PatchedPowerPanelRequest) SetSite(v SiteRequest) {
	o.Site = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPowerPanelRequest) GetLocation() LocationRequest {
	if o == nil || IsNil(o.Location.Get()) {
		var ret LocationRequest
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPowerPanelRequest) GetLocationOk() (*LocationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableLocationRequest and assigns it to the Location field.
func (o *PatchedPowerPanelRequest) SetLocation(v LocationRequest) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *PatchedPowerPanelRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *PatchedPowerPanelRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPowerPanelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPowerPanelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPowerPanelRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedPowerPanelRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPowerPanelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedPowerPanelRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedPowerPanelRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPowerPanelRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedPowerPanelRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedPowerPanelRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPowerPanelRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedPowerPanelRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedPowerPanelRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPowerPanelRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedPowerPanelRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedPowerPanelRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedPowerPanelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedPowerPanelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

func (o *PatchedPowerPanelRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedPowerPanelRequest := _PatchedPowerPanelRequest{}

	err = json.Unmarshal(data, &varPatchedPowerPanelRequest)

	if err != nil {
		return err
	}

	*o = PatchedPowerPanelRequest(varPatchedPowerPanelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedPowerPanelRequest struct {
	value *PatchedPowerPanelRequest
	isSet bool
}

func (v NullablePatchedPowerPanelRequest) Get() *PatchedPowerPanelRequest {
	return v.value
}

func (v *NullablePatchedPowerPanelRequest) Set(val *PatchedPowerPanelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPowerPanelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPowerPanelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPowerPanelRequest(val *PatchedPowerPanelRequest) *NullablePatchedPowerPanelRequest {
	return &NullablePatchedPowerPanelRequest{value: val, isSet: true}
}

func (v NullablePatchedPowerPanelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPowerPanelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}