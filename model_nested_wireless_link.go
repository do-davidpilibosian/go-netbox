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

// checks if the NestedWirelessLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedWirelessLink{}

// NestedWirelessLink Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedWirelessLink struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Ssid                 *string `json:"ssid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedWirelessLink NestedWirelessLink

// NewNestedWirelessLink instantiates a new NestedWirelessLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedWirelessLink(id int32, url string, display string) *NestedWirelessLink {
	this := NestedWirelessLink{}
	this.Id = id
	this.Url = url
	this.Display = display
	return &this
}

// NewNestedWirelessLinkWithDefaults instantiates a new NestedWirelessLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedWirelessLinkWithDefaults() *NestedWirelessLink {
	this := NestedWirelessLink{}
	return &this
}

// GetId returns the Id field value
func (o *NestedWirelessLink) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLink) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedWirelessLink) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedWirelessLink) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedWirelessLink) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedWirelessLink) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLink) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedWirelessLink) SetDisplay(v string) {
	o.Display = v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NestedWirelessLink) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedWirelessLink) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NestedWirelessLink) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *NestedWirelessLink) SetSsid(v string) {
	o.Ssid = &v
}

func (o NestedWirelessLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedWirelessLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedWirelessLink) UnmarshalJSON(bytes []byte) (err error) {
	varNestedWirelessLink := _NestedWirelessLink{}

	err = json.Unmarshal(bytes, &varNestedWirelessLink)

	if err != nil {
		return err
	}

	*o = NestedWirelessLink(varNestedWirelessLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "ssid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedWirelessLink struct {
	value *NestedWirelessLink
	isSet bool
}

func (v NullableNestedWirelessLink) Get() *NestedWirelessLink {
	return v.value
}

func (v *NullableNestedWirelessLink) Set(val *NestedWirelessLink) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedWirelessLink) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedWirelessLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedWirelessLink(val *NestedWirelessLink) *NullableNestedWirelessLink {
	return &NullableNestedWirelessLink{value: val, isSet: true}
}

func (v NullableNestedWirelessLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedWirelessLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
