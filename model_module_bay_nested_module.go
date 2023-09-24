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

// checks if the ModuleBayNestedModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleBayNestedModule{}

// ModuleBayNestedModule Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type ModuleBayNestedModule struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Serial               *string `json:"serial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleBayNestedModule ModuleBayNestedModule

// NewModuleBayNestedModule instantiates a new ModuleBayNestedModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleBayNestedModule(id int32, url string, display string) *ModuleBayNestedModule {
	this := ModuleBayNestedModule{}
	this.Id = id
	this.Url = url
	this.Display = display
	return &this
}

// NewModuleBayNestedModuleWithDefaults instantiates a new ModuleBayNestedModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleBayNestedModuleWithDefaults() *ModuleBayNestedModule {
	this := ModuleBayNestedModule{}
	return &this
}

// GetId returns the Id field value
func (o *ModuleBayNestedModule) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModuleBayNestedModule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModuleBayNestedModule) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ModuleBayNestedModule) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ModuleBayNestedModule) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ModuleBayNestedModule) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ModuleBayNestedModule) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ModuleBayNestedModule) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ModuleBayNestedModule) SetDisplay(v string) {
	o.Display = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *ModuleBayNestedModule) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayNestedModule) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *ModuleBayNestedModule) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *ModuleBayNestedModule) SetSerial(v string) {
	o.Serial = &v
}

func (o ModuleBayNestedModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleBayNestedModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleBayNestedModule) UnmarshalJSON(bytes []byte) (err error) {
	varModuleBayNestedModule := _ModuleBayNestedModule{}

	err = json.Unmarshal(bytes, &varModuleBayNestedModule)

	if err != nil {
		return err
	}

	*o = ModuleBayNestedModule(varModuleBayNestedModule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "serial")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleBayNestedModule struct {
	value *ModuleBayNestedModule
	isSet bool
}

func (v NullableModuleBayNestedModule) Get() *ModuleBayNestedModule {
	return v.value
}

func (v *NullableModuleBayNestedModule) Set(val *ModuleBayNestedModule) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleBayNestedModule) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleBayNestedModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleBayNestedModule(val *ModuleBayNestedModule) *NullableModuleBayNestedModule {
	return &NullableModuleBayNestedModule{value: val, isSet: true}
}

func (v NullableModuleBayNestedModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleBayNestedModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
