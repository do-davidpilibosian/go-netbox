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

// checks if the PatchedWritableTunnelTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableTunnelTerminationRequest{}

// PatchedWritableTunnelTerminationRequest Adds support for custom fields and tags.
type PatchedWritableTunnelTerminationRequest struct {
	Tunnel               *TunnelRequest                               `json:"tunnel,omitempty"`
	Role                 *PatchedWritableTunnelTerminationRequestRole `json:"role,omitempty"`
	TerminationType      *string                                      `json:"termination_type,omitempty"`
	TerminationId        NullableInt64                                `json:"termination_id,omitempty"`
	OutsideIp            NullableIPAddressRequest                     `json:"outside_ip,omitempty"`
	Tags                 []NestedTagRequest                           `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                       `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableTunnelTerminationRequest PatchedWritableTunnelTerminationRequest

// NewPatchedWritableTunnelTerminationRequest instantiates a new PatchedWritableTunnelTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableTunnelTerminationRequest() *PatchedWritableTunnelTerminationRequest {
	this := PatchedWritableTunnelTerminationRequest{}
	return &this
}

// NewPatchedWritableTunnelTerminationRequestWithDefaults instantiates a new PatchedWritableTunnelTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableTunnelTerminationRequestWithDefaults() *PatchedWritableTunnelTerminationRequest {
	this := PatchedWritableTunnelTerminationRequest{}
	return &this
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *PatchedWritableTunnelTerminationRequest) GetTunnel() TunnelRequest {
	if o == nil || IsNil(o.Tunnel) {
		var ret TunnelRequest
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableTunnelTerminationRequest) GetTunnelOk() (*TunnelRequest, bool) {
	if o == nil || IsNil(o.Tunnel) {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasTunnel() bool {
	if o != nil && !IsNil(o.Tunnel) {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given TunnelRequest and assigns it to the Tunnel field.
func (o *PatchedWritableTunnelTerminationRequest) SetTunnel(v TunnelRequest) {
	o.Tunnel = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PatchedWritableTunnelTerminationRequest) GetRole() PatchedWritableTunnelTerminationRequestRole {
	if o == nil || IsNil(o.Role) {
		var ret PatchedWritableTunnelTerminationRequestRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableTunnelTerminationRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given PatchedWritableTunnelTerminationRequestRole and assigns it to the Role field.
func (o *PatchedWritableTunnelTerminationRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole) {
	o.Role = &v
}

// GetTerminationType returns the TerminationType field value if set, zero value otherwise.
func (o *PatchedWritableTunnelTerminationRequest) GetTerminationType() string {
	if o == nil || IsNil(o.TerminationType) {
		var ret string
		return ret
	}
	return *o.TerminationType
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableTunnelTerminationRequest) GetTerminationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TerminationType) {
		return nil, false
	}
	return o.TerminationType, true
}

// HasTerminationType returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasTerminationType() bool {
	if o != nil && !IsNil(o.TerminationType) {
		return true
	}

	return false
}

// SetTerminationType gets a reference to the given string and assigns it to the TerminationType field.
func (o *PatchedWritableTunnelTerminationRequest) SetTerminationType(v string) {
	o.TerminationType = &v
}

// GetTerminationId returns the TerminationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableTunnelTerminationRequest) GetTerminationId() int64 {
	if o == nil || IsNil(o.TerminationId.Get()) {
		var ret int64
		return ret
	}
	return *o.TerminationId.Get()
}

// GetTerminationIdOk returns a tuple with the TerminationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableTunnelTerminationRequest) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationId.Get(), o.TerminationId.IsSet()
}

// HasTerminationId returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasTerminationId() bool {
	if o != nil && o.TerminationId.IsSet() {
		return true
	}

	return false
}

// SetTerminationId gets a reference to the given NullableInt64 and assigns it to the TerminationId field.
func (o *PatchedWritableTunnelTerminationRequest) SetTerminationId(v int64) {
	o.TerminationId.Set(&v)
}

// SetTerminationIdNil sets the value for TerminationId to be an explicit nil
func (o *PatchedWritableTunnelTerminationRequest) SetTerminationIdNil() {
	o.TerminationId.Set(nil)
}

// UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
func (o *PatchedWritableTunnelTerminationRequest) UnsetTerminationId() {
	o.TerminationId.Unset()
}

// GetOutsideIp returns the OutsideIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableTunnelTerminationRequest) GetOutsideIp() IPAddressRequest {
	if o == nil || IsNil(o.OutsideIp.Get()) {
		var ret IPAddressRequest
		return ret
	}
	return *o.OutsideIp.Get()
}

// GetOutsideIpOk returns a tuple with the OutsideIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableTunnelTerminationRequest) GetOutsideIpOk() (*IPAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutsideIp.Get(), o.OutsideIp.IsSet()
}

// HasOutsideIp returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasOutsideIp() bool {
	if o != nil && o.OutsideIp.IsSet() {
		return true
	}

	return false
}

// SetOutsideIp gets a reference to the given NullableIPAddressRequest and assigns it to the OutsideIp field.
func (o *PatchedWritableTunnelTerminationRequest) SetOutsideIp(v IPAddressRequest) {
	o.OutsideIp.Set(&v)
}

// SetOutsideIpNil sets the value for OutsideIp to be an explicit nil
func (o *PatchedWritableTunnelTerminationRequest) SetOutsideIpNil() {
	o.OutsideIp.Set(nil)
}

// UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
func (o *PatchedWritableTunnelTerminationRequest) UnsetOutsideIp() {
	o.OutsideIp.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableTunnelTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableTunnelTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableTunnelTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableTunnelTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableTunnelTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableTunnelTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableTunnelTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableTunnelTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableTunnelTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tunnel) {
		toSerialize["tunnel"] = o.Tunnel
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.TerminationType) {
		toSerialize["termination_type"] = o.TerminationType
	}
	if o.TerminationId.IsSet() {
		toSerialize["termination_id"] = o.TerminationId.Get()
	}
	if o.OutsideIp.IsSet() {
		toSerialize["outside_ip"] = o.OutsideIp.Get()
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

func (o *PatchedWritableTunnelTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableTunnelTerminationRequest := _PatchedWritableTunnelTerminationRequest{}

	err = json.Unmarshal(data, &varPatchedWritableTunnelTerminationRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableTunnelTerminationRequest(varPatchedWritableTunnelTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tunnel")
		delete(additionalProperties, "role")
		delete(additionalProperties, "termination_type")
		delete(additionalProperties, "termination_id")
		delete(additionalProperties, "outside_ip")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableTunnelTerminationRequest struct {
	value *PatchedWritableTunnelTerminationRequest
	isSet bool
}

func (v NullablePatchedWritableTunnelTerminationRequest) Get() *PatchedWritableTunnelTerminationRequest {
	return v.value
}

func (v *NullablePatchedWritableTunnelTerminationRequest) Set(val *PatchedWritableTunnelTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableTunnelTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableTunnelTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableTunnelTerminationRequest(val *PatchedWritableTunnelTerminationRequest) *NullablePatchedWritableTunnelTerminationRequest {
	return &NullablePatchedWritableTunnelTerminationRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableTunnelTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableTunnelTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
