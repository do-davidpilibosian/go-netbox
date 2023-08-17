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

// checks if the L2VPNTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2VPNTerminationRequest{}

// L2VPNTerminationRequest Adds support for custom fields and tags.
type L2VPNTerminationRequest struct {
	L2vpn                NestedL2VPNRequest     `json:"l2vpn"`
	AssignedObjectType   string                 `json:"assigned_object_type"`
	AssignedObjectId     int64                  `json:"assigned_object_id"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _L2VPNTerminationRequest L2VPNTerminationRequest

// NewL2VPNTerminationRequest instantiates a new L2VPNTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2VPNTerminationRequest(l2vpn NestedL2VPNRequest, assignedObjectType string, assignedObjectId int64) *L2VPNTerminationRequest {
	this := L2VPNTerminationRequest{}
	this.L2vpn = l2vpn
	this.AssignedObjectType = assignedObjectType
	this.AssignedObjectId = assignedObjectId
	return &this
}

// NewL2VPNTerminationRequestWithDefaults instantiates a new L2VPNTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2VPNTerminationRequestWithDefaults() *L2VPNTerminationRequest {
	this := L2VPNTerminationRequest{}
	return &this
}

// GetL2vpn returns the L2vpn field value
func (o *L2VPNTerminationRequest) GetL2vpn() NestedL2VPNRequest {
	if o == nil {
		var ret NestedL2VPNRequest
		return ret
	}

	return o.L2vpn
}

// GetL2vpnOk returns a tuple with the L2vpn field value
// and a boolean to check if the value has been set.
func (o *L2VPNTerminationRequest) GetL2vpnOk() (*NestedL2VPNRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2vpn, true
}

// SetL2vpn sets field value
func (o *L2VPNTerminationRequest) SetL2vpn(v NestedL2VPNRequest) {
	o.L2vpn = v
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *L2VPNTerminationRequest) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *L2VPNTerminationRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *L2VPNTerminationRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *L2VPNTerminationRequest) GetAssignedObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *L2VPNTerminationRequest) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *L2VPNTerminationRequest) SetAssignedObjectId(v int64) {
	o.AssignedObjectId = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *L2VPNTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *L2VPNTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *L2VPNTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *L2VPNTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *L2VPNTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *L2VPNTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o L2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2VPNTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["l2vpn"] = o.L2vpn
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object_id"] = o.AssignedObjectId
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

func (o *L2VPNTerminationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varL2VPNTerminationRequest := _L2VPNTerminationRequest{}

	if err = json.Unmarshal(bytes, &varL2VPNTerminationRequest); err == nil {
		*o = L2VPNTerminationRequest(varL2VPNTerminationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "l2vpn")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableL2VPNTerminationRequest struct {
	value *L2VPNTerminationRequest
	isSet bool
}

func (v NullableL2VPNTerminationRequest) Get() *L2VPNTerminationRequest {
	return v.value
}

func (v *NullableL2VPNTerminationRequest) Set(val *L2VPNTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNTerminationRequest(val *L2VPNTerminationRequest) *NullableL2VPNTerminationRequest {
	return &NullableL2VPNTerminationRequest{value: val, isSet: true}
}

func (v NullableL2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
