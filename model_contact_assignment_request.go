/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the ContactAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactAssignmentRequest{}

// ContactAssignmentRequest Adds support for custom fields and tags.
type ContactAssignmentRequest struct {
	ObjectType           string                          `json:"object_type"`
	ObjectId             int64                           `json:"object_id"`
	Contact              ContactRequest                  `json:"contact"`
	Role                 NullableContactRoleRequest      `json:"role,omitempty"`
	Priority             *ContactAssignmentPriorityValue `json:"priority,omitempty"`
	Tags                 []NestedTagRequest              `json:"tags,omitempty"`
	CustomFields         map[string]interface{}          `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContactAssignmentRequest ContactAssignmentRequest

// NewContactAssignmentRequest instantiates a new ContactAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactAssignmentRequest(objectType string, objectId int64, contact ContactRequest) *ContactAssignmentRequest {
	this := ContactAssignmentRequest{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Contact = contact
	return &this
}

// NewContactAssignmentRequestWithDefaults instantiates a new ContactAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactAssignmentRequestWithDefaults() *ContactAssignmentRequest {
	this := ContactAssignmentRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *ContactAssignmentRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ContactAssignmentRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ContactAssignmentRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *ContactAssignmentRequest) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ContactAssignmentRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ContactAssignmentRequest) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetContact returns the Contact field value
func (o *ContactAssignmentRequest) GetContact() ContactRequest {
	if o == nil {
		var ret ContactRequest
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ContactAssignmentRequest) GetContactOk() (*ContactRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ContactAssignmentRequest) SetContact(v ContactRequest) {
	o.Contact = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactAssignmentRequest) GetRole() ContactRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret ContactRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactAssignmentRequest) GetRoleOk() (*ContactRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *ContactAssignmentRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableContactRoleRequest and assigns it to the Role field.
func (o *ContactAssignmentRequest) SetRole(v ContactRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *ContactAssignmentRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *ContactAssignmentRequest) UnsetRole() {
	o.Role.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ContactAssignmentRequest) GetPriority() ContactAssignmentPriorityValue {
	if o == nil || IsNil(o.Priority) {
		var ret ContactAssignmentPriorityValue
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignmentRequest) GetPriorityOk() (*ContactAssignmentPriorityValue, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ContactAssignmentRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given ContactAssignmentPriorityValue and assigns it to the Priority field.
func (o *ContactAssignmentRequest) SetPriority(v ContactAssignmentPriorityValue) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContactAssignmentRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignmentRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContactAssignmentRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ContactAssignmentRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ContactAssignmentRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactAssignmentRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ContactAssignmentRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ContactAssignmentRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ContactAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["contact"] = o.Contact
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
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

func (o *ContactAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
		"contact",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContactAssignmentRequest := _ContactAssignmentRequest{}

	err = json.Unmarshal(data, &varContactAssignmentRequest)

	if err != nil {
		return err
	}

	*o = ContactAssignmentRequest(varContactAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "role")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactAssignmentRequest struct {
	value *ContactAssignmentRequest
	isSet bool
}

func (v NullableContactAssignmentRequest) Get() *ContactAssignmentRequest {
	return v.value
}

func (v *NullableContactAssignmentRequest) Set(val *ContactAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactAssignmentRequest(val *ContactAssignmentRequest) *NullableContactAssignmentRequest {
	return &NullableContactAssignmentRequest{value: val, isSet: true}
}

func (v NullableContactAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
