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

// checks if the NestedConfigTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedConfigTemplate{}

// NestedConfigTemplate Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedConfigTemplate struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedConfigTemplate NestedConfigTemplate

// NewNestedConfigTemplate instantiates a new NestedConfigTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedConfigTemplate(id int32, url string, display string, name string) *NestedConfigTemplate {
	this := NestedConfigTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedConfigTemplateWithDefaults instantiates a new NestedConfigTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedConfigTemplateWithDefaults() *NestedConfigTemplate {
	this := NestedConfigTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *NestedConfigTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedConfigTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedConfigTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedConfigTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedConfigTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedConfigTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedConfigTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedConfigTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedConfigTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedConfigTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedConfigTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedConfigTemplate) SetName(v string) {
	o.Name = v
}

func (o NestedConfigTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedConfigTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedConfigTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varNestedConfigTemplate := _NestedConfigTemplate{}

	err = json.Unmarshal(bytes, &varNestedConfigTemplate)

	if err != nil {
		return err
	}

	*o = NestedConfigTemplate(varNestedConfigTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedConfigTemplate struct {
	value *NestedConfigTemplate
	isSet bool
}

func (v NullableNestedConfigTemplate) Get() *NestedConfigTemplate {
	return v.value
}

func (v *NullableNestedConfigTemplate) Set(val *NestedConfigTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedConfigTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedConfigTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedConfigTemplate(val *NestedConfigTemplate) *NullableNestedConfigTemplate {
	return &NullableNestedConfigTemplate{value: val, isSet: true}
}

func (v NullableNestedConfigTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedConfigTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
