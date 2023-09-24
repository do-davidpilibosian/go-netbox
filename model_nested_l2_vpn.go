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

// checks if the NestedL2VPN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedL2VPN{}

// NestedL2VPN Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedL2VPN struct {
	Id         int32         `json:"id"`
	Url        string        `json:"url"`
	Display    string        `json:"display"`
	Identifier NullableInt64 `json:"identifier,omitempty"`
	Name       string        `json:"name"`
	Slug       string        `json:"slug"`
	// * `vpws` - VPWS * `vpls` - VPLS * `vxlan` - VXLAN * `vxlan-evpn` - VXLAN-EVPN * `mpls-evpn` - MPLS EVPN * `pbb-evpn` - PBB EVPN * `epl` - EPL * `evpl` - EVPL * `ep-lan` - Ethernet Private LAN * `evp-lan` - Ethernet Virtual Private LAN * `ep-tree` - Ethernet Private Tree * `evp-tree` - Ethernet Virtual Private Tree
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _NestedL2VPN NestedL2VPN

// NewNestedL2VPN instantiates a new NestedL2VPN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedL2VPN(id int32, url string, display string, name string, slug string, type_ string) *NestedL2VPN {
	this := NestedL2VPN{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Type = type_
	return &this
}

// NewNestedL2VPNWithDefaults instantiates a new NestedL2VPN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedL2VPNWithDefaults() *NestedL2VPN {
	this := NestedL2VPN{}
	return &this
}

// GetId returns the Id field value
func (o *NestedL2VPN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedL2VPN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedL2VPN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedL2VPN) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedL2VPN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedL2VPN) SetDisplay(v string) {
	o.Display = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedL2VPN) GetIdentifier() int64 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int64
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedL2VPN) GetIdentifierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *NestedL2VPN) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt64 and assigns it to the Identifier field.
func (o *NestedL2VPN) SetIdentifier(v int64) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *NestedL2VPN) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *NestedL2VPN) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetName returns the Name field value
func (o *NestedL2VPN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedL2VPN) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedL2VPN) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPN) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedL2VPN) SetSlug(v string) {
	o.Slug = v
}

// GetType returns the Type field value
func (o *NestedL2VPN) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPN) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NestedL2VPN) SetType(v string) {
	o.Type = v
}

func (o NestedL2VPN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedL2VPN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedL2VPN) UnmarshalJSON(bytes []byte) (err error) {
	varNestedL2VPN := _NestedL2VPN{}

	err = json.Unmarshal(bytes, &varNestedL2VPN)

	if err != nil {
		return err
	}

	*o = NestedL2VPN(varNestedL2VPN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedL2VPN struct {
	value *NestedL2VPN
	isSet bool
}

func (v NullableNestedL2VPN) Get() *NestedL2VPN {
	return v.value
}

func (v *NullableNestedL2VPN) Set(val *NestedL2VPN) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedL2VPN) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedL2VPN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedL2VPN(val *NestedL2VPN) *NullableNestedL2VPN {
	return &NullableNestedL2VPN{value: val, isSet: true}
}

func (v NullableNestedL2VPN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedL2VPN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
