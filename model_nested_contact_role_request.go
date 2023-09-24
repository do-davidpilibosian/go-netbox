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

// checks if the NestedContactRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedContactRoleRequest{}

// NestedContactRoleRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedContactRoleRequest struct {
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedContactRoleRequest NestedContactRoleRequest

// NewNestedContactRoleRequest instantiates a new NestedContactRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedContactRoleRequest(name string, slug string) *NestedContactRoleRequest {
	this := NestedContactRoleRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedContactRoleRequestWithDefaults instantiates a new NestedContactRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedContactRoleRequestWithDefaults() *NestedContactRoleRequest {
	this := NestedContactRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedContactRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedContactRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedContactRoleRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedContactRoleRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedContactRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedContactRoleRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedContactRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedContactRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedContactRoleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedContactRoleRequest := _NestedContactRoleRequest{}

	err = json.Unmarshal(bytes, &varNestedContactRoleRequest)

	if err != nil {
		return err
	}

	*o = NestedContactRoleRequest(varNestedContactRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedContactRoleRequest struct {
	value *NestedContactRoleRequest
	isSet bool
}

func (v NullableNestedContactRoleRequest) Get() *NestedContactRoleRequest {
	return v.value
}

func (v *NullableNestedContactRoleRequest) Set(val *NestedContactRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedContactRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedContactRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedContactRoleRequest(val *NestedContactRoleRequest) *NullableNestedContactRoleRequest {
	return &NullableNestedContactRoleRequest{value: val, isSet: true}
}

func (v NullableNestedContactRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedContactRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
