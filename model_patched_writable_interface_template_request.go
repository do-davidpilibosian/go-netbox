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

// checks if the PatchedWritableInterfaceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableInterfaceTemplateRequest{}

// PatchedWritableInterfaceTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableInterfaceTemplateRequest struct {
	DeviceType NullableInt32 `json:"device_type,omitempty"`
	ModuleType NullableInt32 `json:"module_type,omitempty"`
	//          {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// * `virtual` - Virtual * `bridge` - Bridge * `lag` - Link Aggregation Group (LAG) * `100base-fx` - 100BASE-FX (10/100ME FIBER) * `100base-lfx` - 100BASE-LFX (10/100ME FIBER) * `100base-tx` - 100BASE-TX (10/100ME) * `100base-t1` - 100BASE-T1 (10/100ME Single Pair) * `1000base-t` - 1000BASE-T (1GE) * `2.5gbase-t` - 2.5GBASE-T (2.5GE) * `5gbase-t` - 5GBASE-T (5GE) * `10gbase-t` - 10GBASE-T (10GE) * `10gbase-cx4` - 10GBASE-CX4 (10GE) * `1000base-x-gbic` - GBIC (1GE) * `1000base-x-sfp` - SFP (1GE) * `10gbase-x-sfpp` - SFP+ (10GE) * `10gbase-x-xfp` - XFP (10GE) * `10gbase-x-xenpak` - XENPAK (10GE) * `10gbase-x-x2` - X2 (10GE) * `25gbase-x-sfp28` - SFP28 (25GE) * `50gbase-x-sfp56` - SFP56 (50GE) * `40gbase-x-qsfpp` - QSFP+ (40GE) * `50gbase-x-sfp28` - QSFP28 (50GE) * `100gbase-x-cfp` - CFP (100GE) * `100gbase-x-cfp2` - CFP2 (100GE) * `200gbase-x-cfp2` - CFP2 (200GE) * `400gbase-x-cfp2` - CFP2 (400GE) * `100gbase-x-cfp4` - CFP4 (100GE) * `100gbase-x-cxp` - CXP (100GE) * `100gbase-x-cpak` - Cisco CPAK (100GE) * `100gbase-x-dsfp` - DSFP (100GE) * `100gbase-x-sfpdd` - SFP-DD (100GE) * `100gbase-x-qsfp28` - QSFP28 (100GE) * `100gbase-x-qsfpdd` - QSFP-DD (100GE) * `200gbase-x-qsfp56` - QSFP56 (200GE) * `200gbase-x-qsfpdd` - QSFP-DD (200GE) * `400gbase-x-qsfpdd` - QSFP-DD (400GE) * `400gbase-x-osfp` - OSFP (400GE) * `400gbase-x-cdfp` - CDFP (400GE) * `400gbase-x-cfp8` - CPF8 (400GE) * `800gbase-x-qsfpdd` - QSFP-DD (800GE) * `800gbase-x-osfp` - OSFP (800GE) * `1000base-kx` - 1000BASE-KX (1GE) * `10gbase-kr` - 10GBASE-KR (10GE) * `10gbase-kx4` - 10GBASE-KX4 (10GE) * `25gbase-kr` - 25GBASE-KR (25GE) * `40gbase-kr4` - 40GBASE-KR4 (40GE) * `50gbase-kr` - 50GBASE-KR (50GE) * `100gbase-kp4` - 100GBASE-KP4 (100GE) * `100gbase-kr2` - 100GBASE-KR2 (100GE) * `100gbase-kr4` - 100GBASE-KR4 (100GE) * `ieee802.11a` - IEEE 802.11a * `ieee802.11g` - IEEE 802.11b/g * `ieee802.11n` - IEEE 802.11n * `ieee802.11ac` - IEEE 802.11ac * `ieee802.11ad` - IEEE 802.11ad * `ieee802.11ax` - IEEE 802.11ax * `ieee802.11ay` - IEEE 802.11ay * `ieee802.15.1` - IEEE 802.15.1 (Bluetooth) * `other-wireless` - Other (Wireless) * `gsm` - GSM * `cdma` - CDMA * `lte` - LTE * `sonet-oc3` - OC-3/STM-1 * `sonet-oc12` - OC-12/STM-4 * `sonet-oc48` - OC-48/STM-16 * `sonet-oc192` - OC-192/STM-64 * `sonet-oc768` - OC-768/STM-256 * `sonet-oc1920` - OC-1920/STM-640 * `sonet-oc3840` - OC-3840/STM-1234 * `1gfc-sfp` - SFP (1GFC) * `2gfc-sfp` - SFP (2GFC) * `4gfc-sfp` - SFP (4GFC) * `8gfc-sfpp` - SFP+ (8GFC) * `16gfc-sfpp` - SFP+ (16GFC) * `32gfc-sfp28` - SFP28 (32GFC) * `64gfc-qsfpp` - QSFP+ (64GFC) * `128gfc-qsfp28` - QSFP28 (128GFC) * `infiniband-sdr` - SDR (2 Gbps) * `infiniband-ddr` - DDR (4 Gbps) * `infiniband-qdr` - QDR (8 Gbps) * `infiniband-fdr10` - FDR10 (10 Gbps) * `infiniband-fdr` - FDR (13.5 Gbps) * `infiniband-edr` - EDR (25 Gbps) * `infiniband-hdr` - HDR (50 Gbps) * `infiniband-ndr` - NDR (100 Gbps) * `infiniband-xdr` - XDR (250 Gbps) * `t1` - T1 (1.544 Mbps) * `e1` - E1 (2.048 Mbps) * `t3` - T3 (45 Mbps) * `e3` - E3 (34 Mbps) * `xdsl` - xDSL * `docsis` - DOCSIS * `gpon` - GPON (2.5 Gbps / 1.25 Gps) * `xg-pon` - XG-PON (10 Gbps / 2.5 Gbps) * `xgs-pon` - XGS-PON (10 Gbps) * `ng-pon2` - NG-PON2 (TWDM-PON) (4x10 Gbps) * `epon` - EPON (1 Gbps) * `10g-epon` - 10G-EPON (10 Gbps) * `cisco-stackwise` - Cisco StackWise * `cisco-stackwise-plus` - Cisco StackWise Plus * `cisco-flexstack` - Cisco FlexStack * `cisco-flexstack-plus` - Cisco FlexStack Plus * `cisco-stackwise-80` - Cisco StackWise-80 * `cisco-stackwise-160` - Cisco StackWise-160 * `cisco-stackwise-320` - Cisco StackWise-320 * `cisco-stackwise-480` - Cisco StackWise-480 * `cisco-stackwise-1t` - Cisco StackWise-1T * `juniper-vcp` - Juniper VCP * `extreme-summitstack` - Extreme SummitStack * `extreme-summitstack-128` - Extreme SummitStack-128 * `extreme-summitstack-256` - Extreme SummitStack-256 * `extreme-summitstack-512` - Extreme SummitStack-512 * `other` - Other
	Type        *string       `json:"type,omitempty"`
	Enabled     *bool         `json:"enabled,omitempty"`
	MgmtOnly    *bool         `json:"mgmt_only,omitempty"`
	Description *string       `json:"description,omitempty"`
	Bridge      NullableInt32 `json:"bridge,omitempty"`
	// * `pd` - PD * `pse` - PSE
	PoeMode *string `json:"poe_mode,omitempty"`
	// * `type1-ieee802.3af` - 802.3af (Type 1) * `type2-ieee802.3at` - 802.3at (Type 2) * `type3-ieee802.3bt` - 802.3bt (Type 3) * `type4-ieee802.3bt` - 802.3bt (Type 4) * `passive-24v-2pair` - Passive 24V (2-pair) * `passive-24v-4pair` - Passive 24V (4-pair) * `passive-48v-2pair` - Passive 48V (2-pair) * `passive-48v-4pair` - Passive 48V (4-pair)
	PoeType              *string `json:"poe_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableInterfaceTemplateRequest PatchedWritableInterfaceTemplateRequest

// NewPatchedWritableInterfaceTemplateRequest instantiates a new PatchedWritableInterfaceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableInterfaceTemplateRequest() *PatchedWritableInterfaceTemplateRequest {
	this := PatchedWritableInterfaceTemplateRequest{}
	return &this
}

// NewPatchedWritableInterfaceTemplateRequestWithDefaults instantiates a new PatchedWritableInterfaceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableInterfaceTemplateRequestWithDefaults() *PatchedWritableInterfaceTemplateRequest {
	this := PatchedWritableInterfaceTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceTemplateRequest) GetDeviceType() int32 {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret int32
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableInt32 and assigns it to the DeviceType field.
func (o *PatchedWritableInterfaceTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceTemplateRequest) GetModuleType() int32 {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret int32
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceTemplateRequest) GetModuleTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableInt32 and assigns it to the ModuleType field.
func (o *PatchedWritableInterfaceTemplateRequest) SetModuleType(v int32) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableInterfaceTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableInterfaceTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritableInterfaceTemplateRequest) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedWritableInterfaceTemplateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetMgmtOnly() bool {
	if o == nil || IsNil(o.MgmtOnly) {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MgmtOnly) {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasMgmtOnly() bool {
	if o != nil && !IsNil(o.MgmtOnly) {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *PatchedWritableInterfaceTemplateRequest) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableInterfaceTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceTemplateRequest) GetBridge() int32 {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret int32
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceTemplateRequest) GetBridgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableInt32 and assigns it to the Bridge field.
func (o *PatchedWritableInterfaceTemplateRequest) SetBridge(v int32) {
	o.Bridge.Set(&v)
}

// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetPoeMode returns the PoeMode field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeMode() string {
	if o == nil || IsNil(o.PoeMode) {
		var ret string
		return ret
	}
	return *o.PoeMode
}

// GetPoeModeOk returns a tuple with the PoeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeModeOk() (*string, bool) {
	if o == nil || IsNil(o.PoeMode) {
		return nil, false
	}
	return o.PoeMode, true
}

// HasPoeMode returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasPoeMode() bool {
	if o != nil && !IsNil(o.PoeMode) {
		return true
	}

	return false
}

// SetPoeMode gets a reference to the given string and assigns it to the PoeMode field.
func (o *PatchedWritableInterfaceTemplateRequest) SetPoeMode(v string) {
	o.PoeMode = &v
}

// GetPoeType returns the PoeType field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeType() string {
	if o == nil || IsNil(o.PoeType) {
		var ret string
		return ret
	}
	return *o.PoeType
}

// GetPoeTypeOk returns a tuple with the PoeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PoeType) {
		return nil, false
	}
	return o.PoeType, true
}

// HasPoeType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasPoeType() bool {
	if o != nil && !IsNil(o.PoeType) {
		return true
	}

	return false
}

// SetPoeType gets a reference to the given string and assigns it to the PoeType field.
func (o *PatchedWritableInterfaceTemplateRequest) SetPoeType(v string) {
	o.PoeType = &v
}

func (o PatchedWritableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableInterfaceTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
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
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MgmtOnly) {
		toSerialize["mgmt_only"] = o.MgmtOnly
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if !IsNil(o.PoeMode) {
		toSerialize["poe_mode"] = o.PoeMode
	}
	if !IsNil(o.PoeType) {
		toSerialize["poe_type"] = o.PoeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableInterfaceTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableInterfaceTemplateRequest := _PatchedWritableInterfaceTemplateRequest{}

	if err = json.Unmarshal(bytes, &varPatchedWritableInterfaceTemplateRequest); err == nil {
		*o = PatchedWritableInterfaceTemplateRequest(varPatchedWritableInterfaceTemplateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "mgmt_only")
		delete(additionalProperties, "description")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "poe_mode")
		delete(additionalProperties, "poe_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableInterfaceTemplateRequest struct {
	value *PatchedWritableInterfaceTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableInterfaceTemplateRequest) Get() *PatchedWritableInterfaceTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableInterfaceTemplateRequest) Set(val *PatchedWritableInterfaceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableInterfaceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableInterfaceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableInterfaceTemplateRequest(val *PatchedWritableInterfaceTemplateRequest) *NullablePatchedWritableInterfaceTemplateRequest {
	return &NullablePatchedWritableInterfaceTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableInterfaceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
