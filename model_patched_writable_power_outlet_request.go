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

// checks if the PatchedWritablePowerOutletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritablePowerOutletRequest{}

// PatchedWritablePowerOutletRequest Adds support for custom fields and tags.
type PatchedWritablePowerOutletRequest struct {
	Device *int32        `json:"device,omitempty"`
	Module NullableInt32 `json:"module,omitempty"`
	Name   *string       `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Physical port type  * `iec-60320-c5` - C5 * `iec-60320-c7` - C7 * `iec-60320-c13` - C13 * `iec-60320-c15` - C15 * `iec-60320-c19` - C19 * `iec-60320-c21` - C21 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15r` - NEMA 1-15R * `nema-5-15r` - NEMA 5-15R * `nema-5-20r` - NEMA 5-20R * `nema-5-30r` - NEMA 5-30R * `nema-5-50r` - NEMA 5-50R * `nema-6-15r` - NEMA 6-15R * `nema-6-20r` - NEMA 6-20R * `nema-6-30r` - NEMA 6-30R * `nema-6-50r` - NEMA 6-50R * `nema-10-30r` - NEMA 10-30R * `nema-10-50r` - NEMA 10-50R * `nema-14-20r` - NEMA 14-20R * `nema-14-30r` - NEMA 14-30R * `nema-14-50r` - NEMA 14-50R * `nema-14-60r` - NEMA 14-60R * `nema-15-15r` - NEMA 15-15R * `nema-15-20r` - NEMA 15-20R * `nema-15-30r` - NEMA 15-30R * `nema-15-50r` - NEMA 15-50R * `nema-15-60r` - NEMA 15-60R * `nema-l1-15r` - NEMA L1-15R * `nema-l5-15r` - NEMA L5-15R * `nema-l5-20r` - NEMA L5-20R * `nema-l5-30r` - NEMA L5-30R * `nema-l5-50r` - NEMA L5-50R * `nema-l6-15r` - NEMA L6-15R * `nema-l6-20r` - NEMA L6-20R * `nema-l6-30r` - NEMA L6-30R * `nema-l6-50r` - NEMA L6-50R * `nema-l10-30r` - NEMA L10-30R * `nema-l14-20r` - NEMA L14-20R * `nema-l14-30r` - NEMA L14-30R * `nema-l14-50r` - NEMA L14-50R * `nema-l14-60r` - NEMA L14-60R * `nema-l15-20r` - NEMA L15-20R * `nema-l15-30r` - NEMA L15-30R * `nema-l15-50r` - NEMA L15-50R * `nema-l15-60r` - NEMA L15-60R * `nema-l21-20r` - NEMA L21-20R * `nema-l21-30r` - NEMA L21-30R * `nema-l22-30r` - NEMA L22-30R * `CS6360C` - CS6360C * `CS6364C` - CS6364C * `CS8164C` - CS8164C * `CS8264C` - CS8264C * `CS8364C` - CS8364C * `CS8464C` - CS8464C * `ita-e` - ITA Type E (CEE 7/5) * `ita-f` - ITA Type F (CEE 7/3) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `ita-multistandard` - ITA Multistandard * `usb-a` - USB Type A * `usb-micro-b` - USB Micro B * `usb-c` - USB Type C * `dc-terminal` - DC Terminal * `hdot-cx` - HDOT Cx * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20a` - Neutrik powerCON (20A) * `neutrik-powercon-32a` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
	Type      *string       `json:"type,omitempty"`
	PowerPort NullableInt32 `json:"power_port,omitempty"`
	// Phase (for three-phase feeds)  * `A` - A * `B` - B * `C` - C
	FeedLeg     *string `json:"feed_leg,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritablePowerOutletRequest PatchedWritablePowerOutletRequest

// NewPatchedWritablePowerOutletRequest instantiates a new PatchedWritablePowerOutletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerOutletRequest() *PatchedWritablePowerOutletRequest {
	this := PatchedWritablePowerOutletRequest{}
	return &this
}

// NewPatchedWritablePowerOutletRequestWithDefaults instantiates a new PatchedWritablePowerOutletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerOutletRequestWithDefaults() *PatchedWritablePowerOutletRequest {
	this := PatchedWritablePowerOutletRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetDevice() int32 {
	if o == nil || IsNil(o.Device) {
		var ret int32
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetDeviceOk() (*int32, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given int32 and assigns it to the Device field.
func (o *PatchedWritablePowerOutletRequest) SetDevice(v int32) {
	o.Device = &v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerOutletRequest) GetModule() int32 {
	if o == nil || IsNil(o.Module.Get()) {
		var ret int32
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerOutletRequest) GetModuleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableInt32 and assigns it to the Module field.
func (o *PatchedWritablePowerOutletRequest) SetModule(v int32) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *PatchedWritablePowerOutletRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PatchedWritablePowerOutletRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerOutletRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritablePowerOutletRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritablePowerOutletRequest) SetType(v string) {
	o.Type = &v
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerOutletRequest) GetPowerPort() int32 {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret int32
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerOutletRequest) GetPowerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableInt32 and assigns it to the PowerPort field.
func (o *PatchedWritablePowerOutletRequest) SetPowerPort(v int32) {
	o.PowerPort.Set(&v)
}

// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *PatchedWritablePowerOutletRequest) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *PatchedWritablePowerOutletRequest) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetFeedLeg() string {
	if o == nil || IsNil(o.FeedLeg) {
		var ret string
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetFeedLegOk() (*string, bool) {
	if o == nil || IsNil(o.FeedLeg) {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasFeedLeg() bool {
	if o != nil && !IsNil(o.FeedLeg) {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given string and assigns it to the FeedLeg field.
func (o *PatchedWritablePowerOutletRequest) SetFeedLeg(v string) {
	o.FeedLeg = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritablePowerOutletRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PatchedWritablePowerOutletRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritablePowerOutletRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePowerOutletRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritablePowerOutletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if !IsNil(o.FeedLeg) {
		toSerialize["feed_leg"] = o.FeedLeg
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

func (o *PatchedWritablePowerOutletRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritablePowerOutletRequest := _PatchedWritablePowerOutletRequest{}

	err = json.Unmarshal(bytes, &varPatchedWritablePowerOutletRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritablePowerOutletRequest(varPatchedWritablePowerOutletRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritablePowerOutletRequest struct {
	value *PatchedWritablePowerOutletRequest
	isSet bool
}

func (v NullablePatchedWritablePowerOutletRequest) Get() *PatchedWritablePowerOutletRequest {
	return v.value
}

func (v *NullablePatchedWritablePowerOutletRequest) Set(val *PatchedWritablePowerOutletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerOutletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerOutletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerOutletRequest(val *PatchedWritablePowerOutletRequest) *NullablePatchedWritablePowerOutletRequest {
	return &NullablePatchedWritablePowerOutletRequest{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerOutletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
