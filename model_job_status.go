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

// checks if the JobStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobStatus{}

// JobStatus struct for JobStatus
type JobStatus struct {
	Value                *JobStatusValue `json:"value,omitempty"`
	Label                *JobStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JobStatus JobStatus

// NewJobStatus instantiates a new JobStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStatus() *JobStatus {
	this := JobStatus{}
	return &this
}

// NewJobStatusWithDefaults instantiates a new JobStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStatusWithDefaults() *JobStatus {
	this := JobStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *JobStatus) GetValue() JobStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret JobStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatus) GetValueOk() (*JobStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *JobStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given JobStatusValue and assigns it to the Value field.
func (o *JobStatus) SetValue(v JobStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *JobStatus) GetLabel() JobStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret JobStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatus) GetLabelOk() (*JobStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *JobStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given JobStatusLabel and assigns it to the Label field.
func (o *JobStatus) SetLabel(v JobStatusLabel) {
	o.Label = &v
}

func (o JobStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobStatus) ToMap() (map[string]interface{}, error) {
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

func (o *JobStatus) UnmarshalJSON(data []byte) (err error) {
	varJobStatus := _JobStatus{}

	err = json.Unmarshal(data, &varJobStatus)

	if err != nil {
		return err
	}

	*o = JobStatus(varJobStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobStatus struct {
	value *JobStatus
	isSet bool
}

func (v NullableJobStatus) Get() *JobStatus {
	return v.value
}

func (v *NullableJobStatus) Set(val *JobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStatus(val *JobStatus) *NullableJobStatus {
	return &NullableJobStatus{value: val, isSet: true}
}

func (v NullableJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
