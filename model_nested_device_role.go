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

// checks if the NestedDeviceRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedDeviceRole{}

// NestedDeviceRole Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedDeviceRole struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedDeviceRole NestedDeviceRole

// NewNestedDeviceRole instantiates a new NestedDeviceRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDeviceRole(id int32, url string, display string, name string, slug string) *NestedDeviceRole {
	this := NestedDeviceRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedDeviceRoleWithDefaults instantiates a new NestedDeviceRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeviceRoleWithDefaults() *NestedDeviceRole {
	this := NestedDeviceRole{}
	return &this
}

// GetId returns the Id field value
func (o *NestedDeviceRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedDeviceRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedDeviceRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedDeviceRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedDeviceRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedDeviceRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedDeviceRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedDeviceRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedDeviceRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedDeviceRole) SetSlug(v string) {
	o.Slug = v
}

func (o NestedDeviceRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedDeviceRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedDeviceRole) UnmarshalJSON(bytes []byte) (err error) {
	varNestedDeviceRole := _NestedDeviceRole{}

	err = json.Unmarshal(bytes, &varNestedDeviceRole)

	if err != nil {
		return err
	}

	*o = NestedDeviceRole(varNestedDeviceRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedDeviceRole struct {
	value *NestedDeviceRole
	isSet bool
}

func (v NullableNestedDeviceRole) Get() *NestedDeviceRole {
	return v.value
}

func (v *NullableNestedDeviceRole) Set(val *NestedDeviceRole) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDeviceRole) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDeviceRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDeviceRole(val *NestedDeviceRole) *NullableNestedDeviceRole {
	return &NullableNestedDeviceRole{value: val, isSet: true}
}

func (v NullableNestedDeviceRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDeviceRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
