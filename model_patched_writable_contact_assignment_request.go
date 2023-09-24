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

// checks if the PatchedWritableContactAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableContactAssignmentRequest{}

// PatchedWritableContactAssignmentRequest Adds support for custom fields and tags.
type PatchedWritableContactAssignmentRequest struct {
	ContentType *string `json:"content_type,omitempty"`
	ObjectId    *int64  `json:"object_id,omitempty"`
	Contact     *int32  `json:"contact,omitempty"`
	Role        *int32  `json:"role,omitempty"`
	// * `primary` - Primary * `secondary` - Secondary * `tertiary` - Tertiary * `inactive` - Inactive
	Priority             *string            `json:"priority,omitempty"`
	Tags                 []NestedTagRequest `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableContactAssignmentRequest PatchedWritableContactAssignmentRequest

// NewPatchedWritableContactAssignmentRequest instantiates a new PatchedWritableContactAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableContactAssignmentRequest() *PatchedWritableContactAssignmentRequest {
	this := PatchedWritableContactAssignmentRequest{}
	return &this
}

// NewPatchedWritableContactAssignmentRequestWithDefaults instantiates a new PatchedWritableContactAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableContactAssignmentRequestWithDefaults() *PatchedWritableContactAssignmentRequest {
	this := PatchedWritableContactAssignmentRequest{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PatchedWritableContactAssignmentRequest) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableContactAssignmentRequest) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PatchedWritableContactAssignmentRequest) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PatchedWritableContactAssignmentRequest) SetContentType(v string) {
	o.ContentType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *PatchedWritableContactAssignmentRequest) GetObjectId() int64 {
	if o == nil || IsNil(o.ObjectId) {
		var ret int64
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableContactAssignmentRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *PatchedWritableContactAssignmentRequest) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int64 and assigns it to the ObjectId field.
func (o *PatchedWritableContactAssignmentRequest) SetObjectId(v int64) {
	o.ObjectId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *PatchedWritableContactAssignmentRequest) GetContact() int32 {
	if o == nil || IsNil(o.Contact) {
		var ret int32
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableContactAssignmentRequest) GetContactOk() (*int32, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *PatchedWritableContactAssignmentRequest) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given int32 and assigns it to the Contact field.
func (o *PatchedWritableContactAssignmentRequest) SetContact(v int32) {
	o.Contact = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PatchedWritableContactAssignmentRequest) GetRole() int32 {
	if o == nil || IsNil(o.Role) {
		var ret int32
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableContactAssignmentRequest) GetRoleOk() (*int32, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedWritableContactAssignmentRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given int32 and assigns it to the Role field.
func (o *PatchedWritableContactAssignmentRequest) SetRole(v int32) {
	o.Role = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PatchedWritableContactAssignmentRequest) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableContactAssignmentRequest) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PatchedWritableContactAssignmentRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *PatchedWritableContactAssignmentRequest) SetPriority(v string) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableContactAssignmentRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableContactAssignmentRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableContactAssignmentRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableContactAssignmentRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o PatchedWritableContactAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableContactAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.ObjectId) {
		toSerialize["object_id"] = o.ObjectId
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableContactAssignmentRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableContactAssignmentRequest := _PatchedWritableContactAssignmentRequest{}

	err = json.Unmarshal(bytes, &varPatchedWritableContactAssignmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableContactAssignmentRequest(varPatchedWritableContactAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "role")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableContactAssignmentRequest struct {
	value *PatchedWritableContactAssignmentRequest
	isSet bool
}

func (v NullablePatchedWritableContactAssignmentRequest) Get() *PatchedWritableContactAssignmentRequest {
	return v.value
}

func (v *NullablePatchedWritableContactAssignmentRequest) Set(val *PatchedWritableContactAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableContactAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableContactAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableContactAssignmentRequest(val *PatchedWritableContactAssignmentRequest) *NullablePatchedWritableContactAssignmentRequest {
	return &NullablePatchedWritableContactAssignmentRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableContactAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableContactAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
