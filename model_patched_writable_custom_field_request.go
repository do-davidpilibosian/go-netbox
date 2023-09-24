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

// checks if the PatchedWritableCustomFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableCustomFieldRequest{}

// PatchedWritableCustomFieldRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableCustomFieldRequest struct {
	ContentTypes []string `json:"content_types,omitempty"`
	// The type of data this custom field holds  * `text` - Text * `longtext` - Text (long) * `integer` - Integer * `decimal` - Decimal * `boolean` - Boolean (true/false) * `date` - Date * `datetime` - Date & time * `url` - URL * `json` - JSON * `select` - Selection * `multiselect` - Multiple selection * `object` - Object * `multiobject` - Multiple objects
	Type       *string `json:"type,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	// Internal field name
	Name *string `json:"name,omitempty"`
	// Name of the field as displayed to users (if not provided, 'the field's name will be used)
	Label *string `json:"label,omitempty"`
	// Custom fields within the same group will be displayed together
	GroupName   *string `json:"group_name,omitempty"`
	Description *string `json:"description,omitempty"`
	// If true, this field is required when creating new objects or editing an existing object.
	Required *bool `json:"required,omitempty"`
	// Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored.
	SearchWeight *int32 `json:"search_weight,omitempty"`
	// Loose matches any instance of a given string; exact matches the entire field.  * `disabled` - Disabled * `loose` - Loose * `exact` - Exact
	FilterLogic *string `json:"filter_logic,omitempty"`
	// Specifies the visibility of custom field in the UI  * `read-write` - Read/write * `read-only` - Read-only * `hidden` - Hidden * `hidden-ifunset` - Hidden (if unset)
	UiVisibility *string `json:"ui_visibility,omitempty"`
	// Replicate this value when cloning objects
	IsCloneable *bool `json:"is_cloneable,omitempty"`
	// Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \"Foo\").
	Default map[string]interface{} `json:"default,omitempty"`
	// Fields with higher weights appear lower in a form.
	Weight *int32 `json:"weight,omitempty"`
	// Minimum allowed value (for numeric fields)
	ValidationMinimum NullableInt32 `json:"validation_minimum,omitempty"`
	// Maximum allowed value (for numeric fields)
	ValidationMaximum NullableInt32 `json:"validation_maximum,omitempty"`
	// Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters.
	ValidationRegex      *string       `json:"validation_regex,omitempty"`
	ChoiceSet            NullableInt32 `json:"choice_set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableCustomFieldRequest PatchedWritableCustomFieldRequest

// NewPatchedWritableCustomFieldRequest instantiates a new PatchedWritableCustomFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCustomFieldRequest() *PatchedWritableCustomFieldRequest {
	this := PatchedWritableCustomFieldRequest{}
	return &this
}

// NewPatchedWritableCustomFieldRequestWithDefaults instantiates a new PatchedWritableCustomFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCustomFieldRequestWithDefaults() *PatchedWritableCustomFieldRequest {
	this := PatchedWritableCustomFieldRequest{}
	return &this
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *PatchedWritableCustomFieldRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritableCustomFieldRequest) SetType(v string) {
	o.Type = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *PatchedWritableCustomFieldRequest) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableCustomFieldRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableCustomFieldRequest) SetLabel(v string) {
	o.Label = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *PatchedWritableCustomFieldRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCustomFieldRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PatchedWritableCustomFieldRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetSearchWeight returns the SearchWeight field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetSearchWeight() int32 {
	if o == nil || IsNil(o.SearchWeight) {
		var ret int32
		return ret
	}
	return *o.SearchWeight
}

// GetSearchWeightOk returns a tuple with the SearchWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetSearchWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SearchWeight) {
		return nil, false
	}
	return o.SearchWeight, true
}

// HasSearchWeight returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasSearchWeight() bool {
	if o != nil && !IsNil(o.SearchWeight) {
		return true
	}

	return false
}

// SetSearchWeight gets a reference to the given int32 and assigns it to the SearchWeight field.
func (o *PatchedWritableCustomFieldRequest) SetSearchWeight(v int32) {
	o.SearchWeight = &v
}

// GetFilterLogic returns the FilterLogic field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetFilterLogic() string {
	if o == nil || IsNil(o.FilterLogic) {
		var ret string
		return ret
	}
	return *o.FilterLogic
}

// GetFilterLogicOk returns a tuple with the FilterLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetFilterLogicOk() (*string, bool) {
	if o == nil || IsNil(o.FilterLogic) {
		return nil, false
	}
	return o.FilterLogic, true
}

// HasFilterLogic returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasFilterLogic() bool {
	if o != nil && !IsNil(o.FilterLogic) {
		return true
	}

	return false
}

// SetFilterLogic gets a reference to the given string and assigns it to the FilterLogic field.
func (o *PatchedWritableCustomFieldRequest) SetFilterLogic(v string) {
	o.FilterLogic = &v
}

// GetUiVisibility returns the UiVisibility field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetUiVisibility() string {
	if o == nil || IsNil(o.UiVisibility) {
		var ret string
		return ret
	}
	return *o.UiVisibility
}

// GetUiVisibilityOk returns a tuple with the UiVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetUiVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.UiVisibility) {
		return nil, false
	}
	return o.UiVisibility, true
}

// HasUiVisibility returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasUiVisibility() bool {
	if o != nil && !IsNil(o.UiVisibility) {
		return true
	}

	return false
}

// SetUiVisibility gets a reference to the given string and assigns it to the UiVisibility field.
func (o *PatchedWritableCustomFieldRequest) SetUiVisibility(v string) {
	o.UiVisibility = &v
}

// GetIsCloneable returns the IsCloneable field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetIsCloneable() bool {
	if o == nil || IsNil(o.IsCloneable) {
		var ret bool
		return ret
	}
	return *o.IsCloneable
}

// GetIsCloneableOk returns a tuple with the IsCloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetIsCloneableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloneable) {
		return nil, false
	}
	return o.IsCloneable, true
}

// HasIsCloneable returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasIsCloneable() bool {
	if o != nil && !IsNil(o.IsCloneable) {
		return true
	}

	return false
}

// SetIsCloneable gets a reference to the given bool and assigns it to the IsCloneable field.
func (o *PatchedWritableCustomFieldRequest) SetIsCloneable(v bool) {
	o.IsCloneable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCustomFieldRequest) GetDefault() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCustomFieldRequest) GetDefaultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Default) {
		return map[string]interface{}{}, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasDefault() bool {
	if o != nil && IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given map[string]interface{} and assigns it to the Default field.
func (o *PatchedWritableCustomFieldRequest) SetDefault(v map[string]interface{}) {
	o.Default = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedWritableCustomFieldRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetValidationMinimum returns the ValidationMinimum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCustomFieldRequest) GetValidationMinimum() int32 {
	if o == nil || IsNil(o.ValidationMinimum.Get()) {
		var ret int32
		return ret
	}
	return *o.ValidationMinimum.Get()
}

// GetValidationMinimumOk returns a tuple with the ValidationMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCustomFieldRequest) GetValidationMinimumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMinimum.Get(), o.ValidationMinimum.IsSet()
}

// HasValidationMinimum returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasValidationMinimum() bool {
	if o != nil && o.ValidationMinimum.IsSet() {
		return true
	}

	return false
}

// SetValidationMinimum gets a reference to the given NullableInt32 and assigns it to the ValidationMinimum field.
func (o *PatchedWritableCustomFieldRequest) SetValidationMinimum(v int32) {
	o.ValidationMinimum.Set(&v)
}

// SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil
func (o *PatchedWritableCustomFieldRequest) SetValidationMinimumNil() {
	o.ValidationMinimum.Set(nil)
}

// UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
func (o *PatchedWritableCustomFieldRequest) UnsetValidationMinimum() {
	o.ValidationMinimum.Unset()
}

// GetValidationMaximum returns the ValidationMaximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCustomFieldRequest) GetValidationMaximum() int32 {
	if o == nil || IsNil(o.ValidationMaximum.Get()) {
		var ret int32
		return ret
	}
	return *o.ValidationMaximum.Get()
}

// GetValidationMaximumOk returns a tuple with the ValidationMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCustomFieldRequest) GetValidationMaximumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMaximum.Get(), o.ValidationMaximum.IsSet()
}

// HasValidationMaximum returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasValidationMaximum() bool {
	if o != nil && o.ValidationMaximum.IsSet() {
		return true
	}

	return false
}

// SetValidationMaximum gets a reference to the given NullableInt32 and assigns it to the ValidationMaximum field.
func (o *PatchedWritableCustomFieldRequest) SetValidationMaximum(v int32) {
	o.ValidationMaximum.Set(&v)
}

// SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil
func (o *PatchedWritableCustomFieldRequest) SetValidationMaximumNil() {
	o.ValidationMaximum.Set(nil)
}

// UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
func (o *PatchedWritableCustomFieldRequest) UnsetValidationMaximum() {
	o.ValidationMaximum.Unset()
}

// GetValidationRegex returns the ValidationRegex field value if set, zero value otherwise.
func (o *PatchedWritableCustomFieldRequest) GetValidationRegex() string {
	if o == nil || IsNil(o.ValidationRegex) {
		var ret string
		return ret
	}
	return *o.ValidationRegex
}

// GetValidationRegexOk returns a tuple with the ValidationRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCustomFieldRequest) GetValidationRegexOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationRegex) {
		return nil, false
	}
	return o.ValidationRegex, true
}

// HasValidationRegex returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasValidationRegex() bool {
	if o != nil && !IsNil(o.ValidationRegex) {
		return true
	}

	return false
}

// SetValidationRegex gets a reference to the given string and assigns it to the ValidationRegex field.
func (o *PatchedWritableCustomFieldRequest) SetValidationRegex(v string) {
	o.ValidationRegex = &v
}

// GetChoiceSet returns the ChoiceSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCustomFieldRequest) GetChoiceSet() int32 {
	if o == nil || IsNil(o.ChoiceSet.Get()) {
		var ret int32
		return ret
	}
	return *o.ChoiceSet.Get()
}

// GetChoiceSetOk returns a tuple with the ChoiceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCustomFieldRequest) GetChoiceSetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChoiceSet.Get(), o.ChoiceSet.IsSet()
}

// HasChoiceSet returns a boolean if a field has been set.
func (o *PatchedWritableCustomFieldRequest) HasChoiceSet() bool {
	if o != nil && o.ChoiceSet.IsSet() {
		return true
	}

	return false
}

// SetChoiceSet gets a reference to the given NullableInt32 and assigns it to the ChoiceSet field.
func (o *PatchedWritableCustomFieldRequest) SetChoiceSet(v int32) {
	o.ChoiceSet.Set(&v)
}

// SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil
func (o *PatchedWritableCustomFieldRequest) SetChoiceSetNil() {
	o.ChoiceSet.Set(nil)
}

// UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
func (o *PatchedWritableCustomFieldRequest) UnsetChoiceSet() {
	o.ChoiceSet.Unset()
}

func (o PatchedWritableCustomFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableCustomFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentTypes) {
		toSerialize["content_types"] = o.ContentTypes
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.SearchWeight) {
		toSerialize["search_weight"] = o.SearchWeight
	}
	if !IsNil(o.FilterLogic) {
		toSerialize["filter_logic"] = o.FilterLogic
	}
	if !IsNil(o.UiVisibility) {
		toSerialize["ui_visibility"] = o.UiVisibility
	}
	if !IsNil(o.IsCloneable) {
		toSerialize["is_cloneable"] = o.IsCloneable
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if o.ValidationMinimum.IsSet() {
		toSerialize["validation_minimum"] = o.ValidationMinimum.Get()
	}
	if o.ValidationMaximum.IsSet() {
		toSerialize["validation_maximum"] = o.ValidationMaximum.Get()
	}
	if !IsNil(o.ValidationRegex) {
		toSerialize["validation_regex"] = o.ValidationRegex
	}
	if o.ChoiceSet.IsSet() {
		toSerialize["choice_set"] = o.ChoiceSet.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableCustomFieldRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableCustomFieldRequest := _PatchedWritableCustomFieldRequest{}

	err = json.Unmarshal(bytes, &varPatchedWritableCustomFieldRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableCustomFieldRequest(varPatchedWritableCustomFieldRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "type")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "required")
		delete(additionalProperties, "search_weight")
		delete(additionalProperties, "filter_logic")
		delete(additionalProperties, "ui_visibility")
		delete(additionalProperties, "is_cloneable")
		delete(additionalProperties, "default")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "validation_minimum")
		delete(additionalProperties, "validation_maximum")
		delete(additionalProperties, "validation_regex")
		delete(additionalProperties, "choice_set")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableCustomFieldRequest struct {
	value *PatchedWritableCustomFieldRequest
	isSet bool
}

func (v NullablePatchedWritableCustomFieldRequest) Get() *PatchedWritableCustomFieldRequest {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldRequest) Set(val *PatchedWritableCustomFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldRequest(val *PatchedWritableCustomFieldRequest) *NullablePatchedWritableCustomFieldRequest {
	return &NullablePatchedWritableCustomFieldRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
