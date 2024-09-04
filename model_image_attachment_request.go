/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
	"os"
)

// checks if the ImageAttachmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageAttachmentRequest{}

// ImageAttachmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ImageAttachmentRequest struct {
	ObjectType           string   `json:"object_type"`
	ObjectId             int64    `json:"object_id"`
	Name                 *string  `json:"name,omitempty"`
	Image                *os.File `json:"image"`
	ImageHeight          int32    `json:"image_height"`
	ImageWidth           int32    `json:"image_width"`
	AdditionalProperties map[string]interface{}
}

type _ImageAttachmentRequest ImageAttachmentRequest

// NewImageAttachmentRequest instantiates a new ImageAttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAttachmentRequest(objectType string, objectId int64, image *os.File, imageHeight int32, imageWidth int32) *ImageAttachmentRequest {
	this := ImageAttachmentRequest{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Image = image
	this.ImageHeight = imageHeight
	this.ImageWidth = imageWidth
	return &this
}

// NewImageAttachmentRequestWithDefaults instantiates a new ImageAttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAttachmentRequestWithDefaults() *ImageAttachmentRequest {
	this := ImageAttachmentRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *ImageAttachmentRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ImageAttachmentRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ImageAttachmentRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *ImageAttachmentRequest) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ImageAttachmentRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ImageAttachmentRequest) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageAttachmentRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAttachmentRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageAttachmentRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageAttachmentRequest) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value
func (o *ImageAttachmentRequest) GetImage() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ImageAttachmentRequest) GetImageOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ImageAttachmentRequest) SetImage(v *os.File) {
	o.Image = v
}

// GetImageHeight returns the ImageHeight field value
func (o *ImageAttachmentRequest) GetImageHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value
// and a boolean to check if the value has been set.
func (o *ImageAttachmentRequest) GetImageHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageHeight, true
}

// SetImageHeight sets field value
func (o *ImageAttachmentRequest) SetImageHeight(v int32) {
	o.ImageHeight = v
}

// GetImageWidth returns the ImageWidth field value
func (o *ImageAttachmentRequest) GetImageWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value
// and a boolean to check if the value has been set.
func (o *ImageAttachmentRequest) GetImageWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageWidth, true
}

// SetImageWidth sets field value
func (o *ImageAttachmentRequest) SetImageWidth(v int32) {
	o.ImageWidth = v
}

func (o ImageAttachmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageAttachmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["image"] = o.Image
	toSerialize["image_height"] = o.ImageHeight
	toSerialize["image_width"] = o.ImageWidth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageAttachmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
		"image",
		"image_height",
		"image_width",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varImageAttachmentRequest := _ImageAttachmentRequest{}

	err = json.Unmarshal(data, &varImageAttachmentRequest)

	if err != nil {
		return err
	}

	*o = ImageAttachmentRequest(varImageAttachmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "image")
		delete(additionalProperties, "image_height")
		delete(additionalProperties, "image_width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageAttachmentRequest struct {
	value *ImageAttachmentRequest
	isSet bool
}

func (v NullableImageAttachmentRequest) Get() *ImageAttachmentRequest {
	return v.value
}

func (v *NullableImageAttachmentRequest) Set(val *ImageAttachmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAttachmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAttachmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAttachmentRequest(val *ImageAttachmentRequest) *NullableImageAttachmentRequest {
	return &NullableImageAttachmentRequest{value: val, isSet: true}
}

func (v NullableImageAttachmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAttachmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
