/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.2 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the TokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRequest{}

// TokenRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type TokenRequest struct {
	User     NestedUserRequest `json:"user"`
	Expires  NullableTime      `json:"expires,omitempty"`
	LastUsed NullableTime      `json:"last_used,omitempty"`
	Key      *string           `json:"key,omitempty"`
	// Permit create/update/delete operations using this key
	WriteEnabled         *bool   `json:"write_enabled,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenRequest TokenRequest

// NewTokenRequest instantiates a new TokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequest(user NestedUserRequest) *TokenRequest {
	this := TokenRequest{}
	this.User = user
	return &this
}

// NewTokenRequestWithDefaults instantiates a new TokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestWithDefaults() *TokenRequest {
	this := TokenRequest{}
	return &this
}

// GetUser returns the User field value
func (o *TokenRequest) GetUser() NestedUserRequest {
	if o == nil {
		var ret NestedUserRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetUserOk() (*NestedUserRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TokenRequest) SetUser(v NestedUserRequest) {
	o.User = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenRequest) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *TokenRequest) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *TokenRequest) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *TokenRequest) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *TokenRequest) UnsetExpires() {
	o.Expires.Unset()
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenRequest) GetLastUsed() time.Time {
	if o == nil || IsNil(o.LastUsed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenRequest) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// HasLastUsed returns a boolean if a field has been set.
func (o *TokenRequest) HasLastUsed() bool {
	if o != nil && o.LastUsed.IsSet() {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given NullableTime and assigns it to the LastUsed field.
func (o *TokenRequest) SetLastUsed(v time.Time) {
	o.LastUsed.Set(&v)
}

// SetLastUsedNil sets the value for LastUsed to be an explicit nil
func (o *TokenRequest) SetLastUsedNil() {
	o.LastUsed.Set(nil)
}

// UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
func (o *TokenRequest) UnsetLastUsed() {
	o.LastUsed.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TokenRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TokenRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TokenRequest) SetKey(v string) {
	o.Key = &v
}

// GetWriteEnabled returns the WriteEnabled field value if set, zero value otherwise.
func (o *TokenRequest) GetWriteEnabled() bool {
	if o == nil || IsNil(o.WriteEnabled) {
		var ret bool
		return ret
	}
	return *o.WriteEnabled
}

// GetWriteEnabledOk returns a tuple with the WriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetWriteEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WriteEnabled) {
		return nil, false
	}
	return o.WriteEnabled, true
}

// HasWriteEnabled returns a boolean if a field has been set.
func (o *TokenRequest) HasWriteEnabled() bool {
	if o != nil && !IsNil(o.WriteEnabled) {
		return true
	}

	return false
}

// SetWriteEnabled gets a reference to the given bool and assigns it to the WriteEnabled field.
func (o *TokenRequest) SetWriteEnabled(v bool) {
	o.WriteEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TokenRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TokenRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TokenRequest) SetDescription(v string) {
	o.Description = &v
}

func (o TokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.LastUsed.IsSet() {
		toSerialize["last_used"] = o.LastUsed.Get()
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.WriteEnabled) {
		toSerialize["write_enabled"] = o.WriteEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTokenRequest := _TokenRequest{}

	err = json.Unmarshal(bytes, &varTokenRequest)

	if err != nil {
		return err
	}

	*o = TokenRequest(varTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "expires")
		delete(additionalProperties, "last_used")
		delete(additionalProperties, "key")
		delete(additionalProperties, "write_enabled")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenRequest struct {
	value *TokenRequest
	isSet bool
}

func (v NullableTokenRequest) Get() *TokenRequest {
	return v.value
}

func (v *NullableTokenRequest) Set(val *TokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequest(val *TokenRequest) *NullableTokenRequest {
	return &NullableTokenRequest{value: val, isSet: true}
}

func (v NullableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
