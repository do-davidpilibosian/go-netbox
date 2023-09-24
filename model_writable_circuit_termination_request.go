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

// checks if the WritableCircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableCircuitTerminationRequest{}

// WritableCircuitTerminationRequest Adds support for custom fields and tags.
type WritableCircuitTerminationRequest struct {
	Circuit int32 `json:"circuit"`
	// * `A` - A * `Z` - Z
	TermSide        string        `json:"term_side"`
	Site            NullableInt32 `json:"site,omitempty"`
	ProviderNetwork NullableInt32 `json:"provider_network,omitempty"`
	// Physical circuit speed
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId *string `json:"xconnect_id,omitempty"`
	// Patch panel ID and port number(s)
	PpInfo      *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableCircuitTerminationRequest WritableCircuitTerminationRequest

// NewWritableCircuitTerminationRequest instantiates a new WritableCircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableCircuitTerminationRequest(circuit int32, termSide string) *WritableCircuitTerminationRequest {
	this := WritableCircuitTerminationRequest{}
	this.Circuit = circuit
	this.TermSide = termSide
	return &this
}

// NewWritableCircuitTerminationRequestWithDefaults instantiates a new WritableCircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableCircuitTerminationRequestWithDefaults() *WritableCircuitTerminationRequest {
	this := WritableCircuitTerminationRequest{}
	return &this
}

// GetCircuit returns the Circuit field value
func (o *WritableCircuitTerminationRequest) GetCircuit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetCircuitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Circuit, true
}

// SetCircuit sets field value
func (o *WritableCircuitTerminationRequest) SetCircuit(v int32) {
	o.Circuit = v
}

// GetTermSide returns the TermSide field value
func (o *WritableCircuitTerminationRequest) GetTermSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TermSide
}

// GetTermSideOk returns a tuple with the TermSide field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetTermSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermSide, true
}

// SetTermSide sets field value
func (o *WritableCircuitTerminationRequest) SetTermSide(v string) {
	o.TermSide = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitTerminationRequest) GetSite() int32 {
	if o == nil || IsNil(o.Site.Get()) {
		var ret int32
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitTerminationRequest) GetSiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableInt32 and assigns it to the Site field.
func (o *WritableCircuitTerminationRequest) SetSite(v int32) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *WritableCircuitTerminationRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *WritableCircuitTerminationRequest) UnsetSite() {
	o.Site.Unset()
}

// GetProviderNetwork returns the ProviderNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitTerminationRequest) GetProviderNetwork() int32 {
	if o == nil || IsNil(o.ProviderNetwork.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitTerminationRequest) GetProviderNetworkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// HasProviderNetwork returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasProviderNetwork() bool {
	if o != nil && o.ProviderNetwork.IsSet() {
		return true
	}

	return false
}

// SetProviderNetwork gets a reference to the given NullableInt32 and assigns it to the ProviderNetwork field.
func (o *WritableCircuitTerminationRequest) SetProviderNetwork(v int32) {
	o.ProviderNetwork.Set(&v)
}

// SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil
func (o *WritableCircuitTerminationRequest) SetProviderNetworkNil() {
	o.ProviderNetwork.Set(nil)
}

// UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
func (o *WritableCircuitTerminationRequest) UnsetProviderNetwork() {
	o.ProviderNetwork.Unset()
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitTerminationRequest) GetPortSpeed() int32 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *WritableCircuitTerminationRequest) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}

// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *WritableCircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *WritableCircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitTerminationRequest) GetUpstreamSpeed() int32 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *WritableCircuitTerminationRequest) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}

// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *WritableCircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *WritableCircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *WritableCircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *WritableCircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *WritableCircuitTerminationRequest) GetPpInfo() string {
	if o == nil || IsNil(o.PpInfo) {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetPpInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PpInfo) {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasPpInfo() bool {
	if o != nil && !IsNil(o.PpInfo) {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *WritableCircuitTerminationRequest) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableCircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableCircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *WritableCircuitTerminationRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *WritableCircuitTerminationRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableCircuitTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableCircuitTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableCircuitTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableCircuitTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableCircuitTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableCircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["circuit"] = o.Circuit
	toSerialize["term_side"] = o.TermSide
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.ProviderNetwork.IsSet() {
		toSerialize["provider_network"] = o.ProviderNetwork.Get()
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.PpInfo) {
		toSerialize["pp_info"] = o.PpInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *WritableCircuitTerminationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWritableCircuitTerminationRequest := _WritableCircuitTerminationRequest{}

	err = json.Unmarshal(bytes, &varWritableCircuitTerminationRequest)

	if err != nil {
		return err
	}

	*o = WritableCircuitTerminationRequest(varWritableCircuitTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "term_side")
		delete(additionalProperties, "site")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "port_speed")
		delete(additionalProperties, "upstream_speed")
		delete(additionalProperties, "xconnect_id")
		delete(additionalProperties, "pp_info")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableCircuitTerminationRequest struct {
	value *WritableCircuitTerminationRequest
	isSet bool
}

func (v NullableWritableCircuitTerminationRequest) Get() *WritableCircuitTerminationRequest {
	return v.value
}

func (v *NullableWritableCircuitTerminationRequest) Set(val *WritableCircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableCircuitTerminationRequest(val *WritableCircuitTerminationRequest) *NullableWritableCircuitTerminationRequest {
	return &NullableWritableCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullableWritableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
