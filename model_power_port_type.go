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

// checks if the PowerPortType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerPortType{}

// PowerPortType struct for PowerPortType
type PowerPortType struct {
	// * `iec-60320-c6` - C6 * `iec-60320-c8` - C8 * `iec-60320-c14` - C14 * `iec-60320-c16` - C16 * `iec-60320-c20` - C20 * `iec-60320-c22` - C22 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15p` - NEMA 1-15P * `nema-5-15p` - NEMA 5-15P * `nema-5-20p` - NEMA 5-20P * `nema-5-30p` - NEMA 5-30P * `nema-5-50p` - NEMA 5-50P * `nema-6-15p` - NEMA 6-15P * `nema-6-20p` - NEMA 6-20P * `nema-6-30p` - NEMA 6-30P * `nema-6-50p` - NEMA 6-50P * `nema-10-30p` - NEMA 10-30P * `nema-10-50p` - NEMA 10-50P * `nema-14-20p` - NEMA 14-20P * `nema-14-30p` - NEMA 14-30P * `nema-14-50p` - NEMA 14-50P * `nema-14-60p` - NEMA 14-60P * `nema-15-15p` - NEMA 15-15P * `nema-15-20p` - NEMA 15-20P * `nema-15-30p` - NEMA 15-30P * `nema-15-50p` - NEMA 15-50P * `nema-15-60p` - NEMA 15-60P * `nema-l1-15p` - NEMA L1-15P * `nema-l5-15p` - NEMA L5-15P * `nema-l5-20p` - NEMA L5-20P * `nema-l5-30p` - NEMA L5-30P * `nema-l5-50p` - NEMA L5-50P * `nema-l6-15p` - NEMA L6-15P * `nema-l6-20p` - NEMA L6-20P * `nema-l6-30p` - NEMA L6-30P * `nema-l6-50p` - NEMA L6-50P * `nema-l10-30p` - NEMA L10-30P * `nema-l14-20p` - NEMA L14-20P * `nema-l14-30p` - NEMA L14-30P * `nema-l14-50p` - NEMA L14-50P * `nema-l14-60p` - NEMA L14-60P * `nema-l15-20p` - NEMA L15-20P * `nema-l15-30p` - NEMA L15-30P * `nema-l15-50p` - NEMA L15-50P * `nema-l15-60p` - NEMA L15-60P * `nema-l21-20p` - NEMA L21-20P * `nema-l21-30p` - NEMA L21-30P * `nema-l22-30p` - NEMA L22-30P * `cs6361c` - CS6361C * `cs6365c` - CS6365C * `cs8165c` - CS8165C * `cs8265c` - CS8265C * `cs8365c` - CS8365C * `cs8465c` - CS8465C * `ita-c` - ITA Type C (CEE 7/16) * `ita-e` - ITA Type E (CEE 7/6) * `ita-f` - ITA Type F (CEE 7/4) * `ita-ef` - ITA Type E/F (CEE 7/7) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `usb-a` - USB Type A * `usb-b` - USB Type B * `usb-c` - USB Type C * `usb-mini-a` - USB Mini A * `usb-mini-b` - USB Mini B * `usb-micro-a` - USB Micro A * `usb-micro-b` - USB Micro B * `usb-micro-ab` - USB Micro AB * `usb-3-b` - USB 3.0 Type B * `usb-3-micro-b` - USB 3.0 Micro B * `dc-terminal` - DC Terminal * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20` - Neutrik powerCON (20A) * `neutrik-powercon-32` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
	Value                *string `json:"value,omitempty"`
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerPortType PowerPortType

// NewPowerPortType instantiates a new PowerPortType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerPortType() *PowerPortType {
	this := PowerPortType{}
	return &this
}

// NewPowerPortTypeWithDefaults instantiates a new PowerPortType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerPortTypeWithDefaults() *PowerPortType {
	this := PowerPortType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerPortType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPortType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerPortType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PowerPortType) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerPortType) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPortType) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerPortType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerPortType) SetLabel(v string) {
	o.Label = &v
}

func (o PowerPortType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerPortType) ToMap() (map[string]interface{}, error) {
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

func (o *PowerPortType) UnmarshalJSON(bytes []byte) (err error) {
	varPowerPortType := _PowerPortType{}

	err = json.Unmarshal(bytes, &varPowerPortType)

	if err != nil {
		return err
	}

	*o = PowerPortType(varPowerPortType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerPortType struct {
	value *PowerPortType
	isSet bool
}

func (v NullablePowerPortType) Get() *PowerPortType {
	return v.value
}

func (v *NullablePowerPortType) Set(val *PowerPortType) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPortType) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPortType(val *PowerPortType) *NullablePowerPortType {
	return &NullablePowerPortType{value: val, isSet: true}
}

func (v NullablePowerPortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
