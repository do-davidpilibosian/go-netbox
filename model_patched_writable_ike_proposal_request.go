/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableIKEProposalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableIKEProposalRequest{}

// PatchedWritableIKEProposalRequest Adds support for custom fields and tags.
type PatchedWritableIKEProposalRequest struct {
	Name                    *string                                                   `json:"name,omitempty"`
	Description             *string                                                   `json:"description,omitempty"`
	AuthenticationMethod    *IKEProposalAuthenticationMethodValue                     `json:"authentication_method,omitempty"`
	EncryptionAlgorithm     *IKEProposalEncryptionAlgorithmValue                      `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *PatchedWritableIKEProposalRequestAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`
	Group                   *PatchedWritableIKEProposalRequestGroup                   `json:"group,omitempty"`
	// Security association lifetime (in seconds)
	SaLifetime           NullableInt32          `json:"sa_lifetime,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableIKEProposalRequest PatchedWritableIKEProposalRequest

// NewPatchedWritableIKEProposalRequest instantiates a new PatchedWritableIKEProposalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableIKEProposalRequest() *PatchedWritableIKEProposalRequest {
	this := PatchedWritableIKEProposalRequest{}
	return &this
}

// NewPatchedWritableIKEProposalRequestWithDefaults instantiates a new PatchedWritableIKEProposalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableIKEProposalRequestWithDefaults() *PatchedWritableIKEProposalRequest {
	this := PatchedWritableIKEProposalRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableIKEProposalRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableIKEProposalRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetAuthenticationMethod() IKEProposalAuthenticationMethodValue {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret IKEProposalAuthenticationMethodValue
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetAuthenticationMethodOk() (*IKEProposalAuthenticationMethodValue, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given IKEProposalAuthenticationMethodValue and assigns it to the AuthenticationMethod field.
func (o *PatchedWritableIKEProposalRequest) SetAuthenticationMethod(v IKEProposalAuthenticationMethodValue) {
	o.AuthenticationMethod = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithmValue {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		var ret IKEProposalEncryptionAlgorithmValue
		return ret
	}
	return *o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithmValue, bool) {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		return nil, false
	}
	return o.EncryptionAlgorithm, true
}

// HasEncryptionAlgorithm returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.EncryptionAlgorithm) {
		return true
	}

	return false
}

// SetEncryptionAlgorithm gets a reference to the given IKEProposalEncryptionAlgorithmValue and assigns it to the EncryptionAlgorithm field.
func (o *PatchedWritableIKEProposalRequest) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithmValue) {
	o.EncryptionAlgorithm = &v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetAuthenticationAlgorithm() PatchedWritableIKEProposalRequestAuthenticationAlgorithm {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		var ret PatchedWritableIKEProposalRequestAuthenticationAlgorithm
		return ret
	}
	return *o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetAuthenticationAlgorithmOk() (*PatchedWritableIKEProposalRequestAuthenticationAlgorithm, bool) {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		return nil, false
	}
	return o.AuthenticationAlgorithm, true
}

// HasAuthenticationAlgorithm returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasAuthenticationAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationAlgorithm gets a reference to the given PatchedWritableIKEProposalRequestAuthenticationAlgorithm and assigns it to the AuthenticationAlgorithm field.
func (o *PatchedWritableIKEProposalRequest) SetAuthenticationAlgorithm(v PatchedWritableIKEProposalRequestAuthenticationAlgorithm) {
	o.AuthenticationAlgorithm = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetGroup() PatchedWritableIKEProposalRequestGroup {
	if o == nil || IsNil(o.Group) {
		var ret PatchedWritableIKEProposalRequestGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetGroupOk() (*PatchedWritableIKEProposalRequestGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given PatchedWritableIKEProposalRequestGroup and assigns it to the Group field.
func (o *PatchedWritableIKEProposalRequest) SetGroup(v PatchedWritableIKEProposalRequestGroup) {
	o.Group = &v
}

// GetSaLifetime returns the SaLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableIKEProposalRequest) GetSaLifetime() int32 {
	if o == nil || IsNil(o.SaLifetime.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetime.Get()
}

// GetSaLifetimeOk returns a tuple with the SaLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableIKEProposalRequest) GetSaLifetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetime.Get(), o.SaLifetime.IsSet()
}

// HasSaLifetime returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasSaLifetime() bool {
	if o != nil && o.SaLifetime.IsSet() {
		return true
	}

	return false
}

// SetSaLifetime gets a reference to the given NullableInt32 and assigns it to the SaLifetime field.
func (o *PatchedWritableIKEProposalRequest) SetSaLifetime(v int32) {
	o.SaLifetime.Set(&v)
}

// SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil
func (o *PatchedWritableIKEProposalRequest) SetSaLifetimeNil() {
	o.SaLifetime.Set(nil)
}

// UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
func (o *PatchedWritableIKEProposalRequest) UnsetSaLifetime() {
	o.SaLifetime.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableIKEProposalRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableIKEProposalRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableIKEProposalRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIKEProposalRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableIKEProposalRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableIKEProposalRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableIKEProposalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableIKEProposalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authentication_method"] = o.AuthenticationMethod
	}
	if !IsNil(o.EncryptionAlgorithm) {
		toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	}
	if !IsNil(o.AuthenticationAlgorithm) {
		toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if o.SaLifetime.IsSet() {
		toSerialize["sa_lifetime"] = o.SaLifetime.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *PatchedWritableIKEProposalRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableIKEProposalRequest := _PatchedWritableIKEProposalRequest{}

	err = json.Unmarshal(data, &varPatchedWritableIKEProposalRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableIKEProposalRequest(varPatchedWritableIKEProposalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "authentication_method")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "group")
		delete(additionalProperties, "sa_lifetime")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableIKEProposalRequest struct {
	value *PatchedWritableIKEProposalRequest
	isSet bool
}

func (v NullablePatchedWritableIKEProposalRequest) Get() *PatchedWritableIKEProposalRequest {
	return v.value
}

func (v *NullablePatchedWritableIKEProposalRequest) Set(val *PatchedWritableIKEProposalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIKEProposalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIKEProposalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIKEProposalRequest(val *PatchedWritableIKEProposalRequest) *NullablePatchedWritableIKEProposalRequest {
	return &NullablePatchedWritableIKEProposalRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableIKEProposalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIKEProposalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
