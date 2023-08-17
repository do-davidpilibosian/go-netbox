/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.8 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the ImageAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageAttachment{}

// ImageAttachment Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ImageAttachment struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	ContentType          string                 `json:"content_type"`
	ObjectId             int64                  `json:"object_id"`
	Parent               map[string]interface{} `json:"parent"`
	Name                 *string                `json:"name,omitempty"`
	Image                string                 `json:"image"`
	ImageHeight          int32                  `json:"image_height"`
	ImageWidth           int32                  `json:"image_width"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ImageAttachment ImageAttachment

// NewImageAttachment instantiates a new ImageAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAttachment(id int32, url string, display string, contentType string, objectId int64, parent map[string]interface{}, image string, imageHeight int32, imageWidth int32, created NullableTime, lastUpdated NullableTime) *ImageAttachment {
	this := ImageAttachment{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ContentType = contentType
	this.ObjectId = objectId
	this.Parent = parent
	this.Image = image
	this.ImageHeight = imageHeight
	this.ImageWidth = imageWidth
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewImageAttachmentWithDefaults instantiates a new ImageAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAttachmentWithDefaults() *ImageAttachment {
	this := ImageAttachment{}
	return &this
}

// GetId returns the Id field value
func (o *ImageAttachment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageAttachment) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ImageAttachment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ImageAttachment) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ImageAttachment) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ImageAttachment) SetDisplay(v string) {
	o.Display = v
}

// GetContentType returns the ContentType field value
func (o *ImageAttachment) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ImageAttachment) SetContentType(v string) {
	o.ContentType = v
}

// GetObjectId returns the ObjectId field value
func (o *ImageAttachment) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ImageAttachment) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetParent returns the Parent field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ImageAttachment) GetParent() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageAttachment) GetParentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parent) {
		return map[string]interface{}{}, false
	}
	return o.Parent, true
}

// SetParent sets field value
func (o *ImageAttachment) SetParent(v map[string]interface{}) {
	o.Parent = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageAttachment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageAttachment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageAttachment) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value
func (o *ImageAttachment) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ImageAttachment) SetImage(v string) {
	o.Image = v
}

// GetImageHeight returns the ImageHeight field value
func (o *ImageAttachment) GetImageHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetImageHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageHeight, true
}

// SetImageHeight sets field value
func (o *ImageAttachment) SetImageHeight(v int32) {
	o.ImageHeight = v
}

// GetImageWidth returns the ImageWidth field value
func (o *ImageAttachment) GetImageWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetImageWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageWidth, true
}

// SetImageWidth sets field value
func (o *ImageAttachment) SetImageWidth(v int32) {
	o.ImageWidth = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ImageAttachment) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageAttachment) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ImageAttachment) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ImageAttachment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageAttachment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ImageAttachment) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ImageAttachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["content_type"] = o.ContentType
	toSerialize["object_id"] = o.ObjectId
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["image"] = o.Image
	toSerialize["image_height"] = o.ImageHeight
	toSerialize["image_width"] = o.ImageWidth
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageAttachment) UnmarshalJSON(bytes []byte) (err error) {
	varImageAttachment := _ImageAttachment{}

	if err = json.Unmarshal(bytes, &varImageAttachment); err == nil {
		*o = ImageAttachment(varImageAttachment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "image")
		delete(additionalProperties, "image_height")
		delete(additionalProperties, "image_width")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageAttachment struct {
	value *ImageAttachment
	isSet bool
}

func (v NullableImageAttachment) Get() *ImageAttachment {
	return v.value
}

func (v *NullableImageAttachment) Set(val *ImageAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAttachment(val *ImageAttachment) *NullableImageAttachment {
	return &NullableImageAttachment{value: val, isSet: true}
}

func (v NullableImageAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
