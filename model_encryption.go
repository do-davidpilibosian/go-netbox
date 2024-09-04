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
)

// Encryption * `aes-128-cbc` - 128-bit AES (CBC) * `aes-128-gcm` - 128-bit AES (GCM) * `aes-192-cbc` - 192-bit AES (CBC) * `aes-192-gcm` - 192-bit AES (GCM) * `aes-256-cbc` - 256-bit AES (CBC) * `aes-256-gcm` - 256-bit AES (GCM) * `3des-cbc` - 3DES * `des-cbc` - DES
type Encryption string

// List of Encryption
const (
	ENCRYPTION_AES_128_CBC Encryption = "aes-128-cbc"
	ENCRYPTION_AES_128_GCM Encryption = "aes-128-gcm"
	ENCRYPTION_AES_192_CBC Encryption = "aes-192-cbc"
	ENCRYPTION_AES_192_GCM Encryption = "aes-192-gcm"
	ENCRYPTION_AES_256_CBC Encryption = "aes-256-cbc"
	ENCRYPTION_AES_256_GCM Encryption = "aes-256-gcm"
	ENCRYPTION__3DES_CBC   Encryption = "3des-cbc"
	ENCRYPTION_DES_CBC     Encryption = "des-cbc"
	ENCRYPTION_EMPTY       Encryption = ""
)

// All allowed values of Encryption enum
var AllowedEncryptionEnumValues = []Encryption{
	"aes-128-cbc",
	"aes-128-gcm",
	"aes-192-cbc",
	"aes-192-gcm",
	"aes-256-cbc",
	"aes-256-gcm",
	"3des-cbc",
	"des-cbc",
	"",
}

func (v *Encryption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Encryption(value)
	for _, existing := range AllowedEncryptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Encryption", value)
}

// NewEncryptionFromValue returns a pointer to a valid Encryption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncryptionFromValue(v string) (*Encryption, error) {
	ev := Encryption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Encryption: valid values are %v", v, AllowedEncryptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Encryption) IsValid() bool {
	for _, existing := range AllowedEncryptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Encryption value
func (v Encryption) Ptr() *Encryption {
	return &v
}

type NullableEncryption struct {
	value *Encryption
	isSet bool
}

func (v NullableEncryption) Get() *Encryption {
	return v.value
}

func (v *NullableEncryption) Set(val *Encryption) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryption(val *Encryption) *NullableEncryption {
	return &NullableEncryption{value: val, isSet: true}
}

func (v NullableEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
