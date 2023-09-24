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

// checks if the InterfaceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceType{}

// InterfaceType struct for InterfaceType
type InterfaceType struct {
	// * `virtual` - Virtual * `bridge` - Bridge * `lag` - Link Aggregation Group (LAG) * `100base-fx` - 100BASE-FX (10/100ME FIBER) * `100base-lfx` - 100BASE-LFX (10/100ME FIBER) * `100base-tx` - 100BASE-TX (10/100ME) * `100base-t1` - 100BASE-T1 (10/100ME Single Pair) * `1000base-t` - 1000BASE-T (1GE) * `2.5gbase-t` - 2.5GBASE-T (2.5GE) * `5gbase-t` - 5GBASE-T (5GE) * `10gbase-t` - 10GBASE-T (10GE) * `10gbase-cx4` - 10GBASE-CX4 (10GE) * `1000base-x-gbic` - GBIC (1GE) * `1000base-x-sfp` - SFP (1GE) * `10gbase-x-sfpp` - SFP+ (10GE) * `10gbase-x-xfp` - XFP (10GE) * `10gbase-x-xenpak` - XENPAK (10GE) * `10gbase-x-x2` - X2 (10GE) * `25gbase-x-sfp28` - SFP28 (25GE) * `50gbase-x-sfp56` - SFP56 (50GE) * `40gbase-x-qsfpp` - QSFP+ (40GE) * `50gbase-x-sfp28` - QSFP28 (50GE) * `100gbase-x-cfp` - CFP (100GE) * `100gbase-x-cfp2` - CFP2 (100GE) * `200gbase-x-cfp2` - CFP2 (200GE) * `400gbase-x-cfp2` - CFP2 (400GE) * `100gbase-x-cfp4` - CFP4 (100GE) * `100gbase-x-cxp` - CXP (100GE) * `100gbase-x-cpak` - Cisco CPAK (100GE) * `100gbase-x-dsfp` - DSFP (100GE) * `100gbase-x-sfpdd` - SFP-DD (100GE) * `100gbase-x-qsfp28` - QSFP28 (100GE) * `100gbase-x-qsfpdd` - QSFP-DD (100GE) * `200gbase-x-qsfp56` - QSFP56 (200GE) * `200gbase-x-qsfpdd` - QSFP-DD (200GE) * `400gbase-x-qsfp112` - QSFP112 (400GE) * `400gbase-x-qsfpdd` - QSFP-DD (400GE) * `400gbase-x-osfp` - OSFP (400GE) * `400gbase-x-osfp-rhs` - OSFP-RHS (400GE) * `400gbase-x-cdfp` - CDFP (400GE) * `400gbase-x-cfp8` - CPF8 (400GE) * `800gbase-x-qsfpdd` - QSFP-DD (800GE) * `800gbase-x-osfp` - OSFP (800GE) * `1000base-kx` - 1000BASE-KX (1GE) * `10gbase-kr` - 10GBASE-KR (10GE) * `10gbase-kx4` - 10GBASE-KX4 (10GE) * `25gbase-kr` - 25GBASE-KR (25GE) * `40gbase-kr4` - 40GBASE-KR4 (40GE) * `50gbase-kr` - 50GBASE-KR (50GE) * `100gbase-kp4` - 100GBASE-KP4 (100GE) * `100gbase-kr2` - 100GBASE-KR2 (100GE) * `100gbase-kr4` - 100GBASE-KR4 (100GE) * `ieee802.11a` - IEEE 802.11a * `ieee802.11g` - IEEE 802.11b/g * `ieee802.11n` - IEEE 802.11n * `ieee802.11ac` - IEEE 802.11ac * `ieee802.11ad` - IEEE 802.11ad * `ieee802.11ax` - IEEE 802.11ax * `ieee802.11ay` - IEEE 802.11ay * `ieee802.15.1` - IEEE 802.15.1 (Bluetooth) * `other-wireless` - Other (Wireless) * `gsm` - GSM * `cdma` - CDMA * `lte` - LTE * `sonet-oc3` - OC-3/STM-1 * `sonet-oc12` - OC-12/STM-4 * `sonet-oc48` - OC-48/STM-16 * `sonet-oc192` - OC-192/STM-64 * `sonet-oc768` - OC-768/STM-256 * `sonet-oc1920` - OC-1920/STM-640 * `sonet-oc3840` - OC-3840/STM-1234 * `1gfc-sfp` - SFP (1GFC) * `2gfc-sfp` - SFP (2GFC) * `4gfc-sfp` - SFP (4GFC) * `8gfc-sfpp` - SFP+ (8GFC) * `16gfc-sfpp` - SFP+ (16GFC) * `32gfc-sfp28` - SFP28 (32GFC) * `64gfc-qsfpp` - QSFP+ (64GFC) * `128gfc-qsfp28` - QSFP28 (128GFC) * `infiniband-sdr` - SDR (2 Gbps) * `infiniband-ddr` - DDR (4 Gbps) * `infiniband-qdr` - QDR (8 Gbps) * `infiniband-fdr10` - FDR10 (10 Gbps) * `infiniband-fdr` - FDR (13.5 Gbps) * `infiniband-edr` - EDR (25 Gbps) * `infiniband-hdr` - HDR (50 Gbps) * `infiniband-ndr` - NDR (100 Gbps) * `infiniband-xdr` - XDR (250 Gbps) * `t1` - T1 (1.544 Mbps) * `e1` - E1 (2.048 Mbps) * `t3` - T3 (45 Mbps) * `e3` - E3 (34 Mbps) * `xdsl` - xDSL * `docsis` - DOCSIS * `gpon` - GPON (2.5 Gbps / 1.25 Gps) * `xg-pon` - XG-PON (10 Gbps / 2.5 Gbps) * `xgs-pon` - XGS-PON (10 Gbps) * `ng-pon2` - NG-PON2 (TWDM-PON) (4x10 Gbps) * `epon` - EPON (1 Gbps) * `10g-epon` - 10G-EPON (10 Gbps) * `cisco-stackwise` - Cisco StackWise * `cisco-stackwise-plus` - Cisco StackWise Plus * `cisco-flexstack` - Cisco FlexStack * `cisco-flexstack-plus` - Cisco FlexStack Plus * `cisco-stackwise-80` - Cisco StackWise-80 * `cisco-stackwise-160` - Cisco StackWise-160 * `cisco-stackwise-320` - Cisco StackWise-320 * `cisco-stackwise-480` - Cisco StackWise-480 * `cisco-stackwise-1t` - Cisco StackWise-1T * `juniper-vcp` - Juniper VCP * `extreme-summitstack` - Extreme SummitStack * `extreme-summitstack-128` - Extreme SummitStack-128 * `extreme-summitstack-256` - Extreme SummitStack-256 * `extreme-summitstack-512` - Extreme SummitStack-512 * `other` - Other
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceType InterfaceType

// NewInterfaceType instantiates a new InterfaceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceType() *InterfaceType {
	this := InterfaceType{}
	return &this
}

// NewInterfaceTypeWithDefaults instantiates a new InterfaceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTypeWithDefaults() *InterfaceType {
	this := InterfaceType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InterfaceType) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceType) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceType) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceType) SetLabel(v string) {
	o.Label = &v
}

func (o InterfaceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceType) UnmarshalJSON(bytes []byte) (err error) {
	varInterfaceType := _InterfaceType{}

	err = json.Unmarshal(bytes, &varInterfaceType)

	if err != nil {
		return err
	}

	*o = InterfaceType(varInterfaceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceType struct {
	value *InterfaceType
	isSet bool
}

func (v NullableInterfaceType) Get() *InterfaceType {
	return v.value
}

func (v *NullableInterfaceType) Set(val *InterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceType(val *InterfaceType) *NullableInterfaceType {
	return &NullableInterfaceType{value: val, isSet: true}
}

func (v NullableInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
