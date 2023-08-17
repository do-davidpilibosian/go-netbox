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

// checks if the NestedVirtualChassis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVirtualChassis{}

// NestedVirtualChassis Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVirtualChassis struct {
	Id                   int32        `json:"id"`
	Url                  string       `json:"url"`
	Display              string       `json:"display"`
	Name                 string       `json:"name"`
	Master               NestedDevice `json:"master"`
	AdditionalProperties map[string]interface{}
}

type _NestedVirtualChassis NestedVirtualChassis

// NewNestedVirtualChassis instantiates a new NestedVirtualChassis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVirtualChassis(id int32, url string, display string, name string, master NestedDevice) *NestedVirtualChassis {
	this := NestedVirtualChassis{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Master = master
	return &this
}

// NewNestedVirtualChassisWithDefaults instantiates a new NestedVirtualChassis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVirtualChassisWithDefaults() *NestedVirtualChassis {
	this := NestedVirtualChassis{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVirtualChassis) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassis) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVirtualChassis) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedVirtualChassis) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassis) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVirtualChassis) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedVirtualChassis) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassis) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedVirtualChassis) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedVirtualChassis) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassis) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVirtualChassis) SetName(v string) {
	o.Name = v
}

// GetMaster returns the Master field value
func (o *NestedVirtualChassis) GetMaster() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Master
}

// GetMasterOk returns a tuple with the Master field value
// and a boolean to check if the value has been set.
func (o *NestedVirtualChassis) GetMasterOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Master, true
}

// SetMaster sets field value
func (o *NestedVirtualChassis) SetMaster(v NestedDevice) {
	o.Master = v
}

func (o NestedVirtualChassis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVirtualChassis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["name"] = o.Name
	toSerialize["master"] = o.Master

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVirtualChassis) UnmarshalJSON(bytes []byte) (err error) {
	varNestedVirtualChassis := _NestedVirtualChassis{}

	if err = json.Unmarshal(bytes, &varNestedVirtualChassis); err == nil {
		*o = NestedVirtualChassis(varNestedVirtualChassis)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "master")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVirtualChassis struct {
	value *NestedVirtualChassis
	isSet bool
}

func (v NullableNestedVirtualChassis) Get() *NestedVirtualChassis {
	return v.value
}

func (v *NullableNestedVirtualChassis) Set(val *NestedVirtualChassis) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVirtualChassis) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVirtualChassis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVirtualChassis(val *NestedVirtualChassis) *NullableNestedVirtualChassis {
	return &NullableNestedVirtualChassis{value: val, isSet: true}
}

func (v NullableNestedVirtualChassis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVirtualChassis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
