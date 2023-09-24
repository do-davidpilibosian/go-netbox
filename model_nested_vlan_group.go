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

// checks if the NestedVLANGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVLANGroup{}

// NestedVLANGroup Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVLANGroup struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedVLANGroup NestedVLANGroup

// NewNestedVLANGroup instantiates a new NestedVLANGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVLANGroup(id int32, url string, display string, name string, slug string) *NestedVLANGroup {
	this := NestedVLANGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedVLANGroupWithDefaults instantiates a new NestedVLANGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVLANGroupWithDefaults() *NestedVLANGroup {
	this := NestedVLANGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVLANGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVLANGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedVLANGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVLANGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedVLANGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedVLANGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedVLANGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVLANGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedVLANGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedVLANGroup) SetSlug(v string) {
	o.Slug = v
}

func (o NestedVLANGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVLANGroup) ToMap() (map[string]interface{}, error) {
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

func (o *NestedVLANGroup) UnmarshalJSON(bytes []byte) (err error) {
	varNestedVLANGroup := _NestedVLANGroup{}

	err = json.Unmarshal(bytes, &varNestedVLANGroup)

	if err != nil {
		return err
	}

	*o = NestedVLANGroup(varNestedVLANGroup)

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

type NullableNestedVLANGroup struct {
	value *NestedVLANGroup
	isSet bool
}

func (v NullableNestedVLANGroup) Get() *NestedVLANGroup {
	return v.value
}

func (v *NullableNestedVLANGroup) Set(val *NestedVLANGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVLANGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVLANGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVLANGroup(val *NestedVLANGroup) *NullableNestedVLANGroup {
	return &NullableNestedVLANGroup{value: val, isSet: true}
}

func (v NullableNestedVLANGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVLANGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
