/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// ClassicKeyTargetInfo struct for ClassicKeyTargetInfo
type ClassicKeyTargetInfo struct {
	ExternalKmsId *ExternalKMSKeyId `json:"external_kms_id,omitempty"`
	KeyPurpose []string `json:"key_purpose,omitempty"`
	KeyStatus *ClassicKeyStatusInfo `json:"key_status,omitempty"`
	TargetAssocId *string `json:"target_assoc_id,omitempty"`
	TargetType *string `json:"target_type,omitempty"`
}

// NewClassicKeyTargetInfo instantiates a new ClassicKeyTargetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassicKeyTargetInfo() *ClassicKeyTargetInfo {
	this := ClassicKeyTargetInfo{}
	return &this
}

// NewClassicKeyTargetInfoWithDefaults instantiates a new ClassicKeyTargetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassicKeyTargetInfoWithDefaults() *ClassicKeyTargetInfo {
	this := ClassicKeyTargetInfo{}
	return &this
}

// GetExternalKmsId returns the ExternalKmsId field value if set, zero value otherwise.
func (o *ClassicKeyTargetInfo) GetExternalKmsId() ExternalKMSKeyId {
	if o == nil || o.ExternalKmsId == nil {
		var ret ExternalKMSKeyId
		return ret
	}
	return *o.ExternalKmsId
}

// GetExternalKmsIdOk returns a tuple with the ExternalKmsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicKeyTargetInfo) GetExternalKmsIdOk() (*ExternalKMSKeyId, bool) {
	if o == nil || o.ExternalKmsId == nil {
		return nil, false
	}
	return o.ExternalKmsId, true
}

// HasExternalKmsId returns a boolean if a field has been set.
func (o *ClassicKeyTargetInfo) HasExternalKmsId() bool {
	if o != nil && o.ExternalKmsId != nil {
		return true
	}

	return false
}

// SetExternalKmsId gets a reference to the given ExternalKMSKeyId and assigns it to the ExternalKmsId field.
func (o *ClassicKeyTargetInfo) SetExternalKmsId(v ExternalKMSKeyId) {
	o.ExternalKmsId = &v
}

// GetKeyPurpose returns the KeyPurpose field value if set, zero value otherwise.
func (o *ClassicKeyTargetInfo) GetKeyPurpose() []string {
	if o == nil || o.KeyPurpose == nil {
		var ret []string
		return ret
	}
	return o.KeyPurpose
}

// GetKeyPurposeOk returns a tuple with the KeyPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicKeyTargetInfo) GetKeyPurposeOk() ([]string, bool) {
	if o == nil || o.KeyPurpose == nil {
		return nil, false
	}
	return o.KeyPurpose, true
}

// HasKeyPurpose returns a boolean if a field has been set.
func (o *ClassicKeyTargetInfo) HasKeyPurpose() bool {
	if o != nil && o.KeyPurpose != nil {
		return true
	}

	return false
}

// SetKeyPurpose gets a reference to the given []string and assigns it to the KeyPurpose field.
func (o *ClassicKeyTargetInfo) SetKeyPurpose(v []string) {
	o.KeyPurpose = v
}

// GetKeyStatus returns the KeyStatus field value if set, zero value otherwise.
func (o *ClassicKeyTargetInfo) GetKeyStatus() ClassicKeyStatusInfo {
	if o == nil || o.KeyStatus == nil {
		var ret ClassicKeyStatusInfo
		return ret
	}
	return *o.KeyStatus
}

// GetKeyStatusOk returns a tuple with the KeyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicKeyTargetInfo) GetKeyStatusOk() (*ClassicKeyStatusInfo, bool) {
	if o == nil || o.KeyStatus == nil {
		return nil, false
	}
	return o.KeyStatus, true
}

// HasKeyStatus returns a boolean if a field has been set.
func (o *ClassicKeyTargetInfo) HasKeyStatus() bool {
	if o != nil && o.KeyStatus != nil {
		return true
	}

	return false
}

// SetKeyStatus gets a reference to the given ClassicKeyStatusInfo and assigns it to the KeyStatus field.
func (o *ClassicKeyTargetInfo) SetKeyStatus(v ClassicKeyStatusInfo) {
	o.KeyStatus = &v
}

// GetTargetAssocId returns the TargetAssocId field value if set, zero value otherwise.
func (o *ClassicKeyTargetInfo) GetTargetAssocId() string {
	if o == nil || o.TargetAssocId == nil {
		var ret string
		return ret
	}
	return *o.TargetAssocId
}

// GetTargetAssocIdOk returns a tuple with the TargetAssocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicKeyTargetInfo) GetTargetAssocIdOk() (*string, bool) {
	if o == nil || o.TargetAssocId == nil {
		return nil, false
	}
	return o.TargetAssocId, true
}

// HasTargetAssocId returns a boolean if a field has been set.
func (o *ClassicKeyTargetInfo) HasTargetAssocId() bool {
	if o != nil && o.TargetAssocId != nil {
		return true
	}

	return false
}

// SetTargetAssocId gets a reference to the given string and assigns it to the TargetAssocId field.
func (o *ClassicKeyTargetInfo) SetTargetAssocId(v string) {
	o.TargetAssocId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *ClassicKeyTargetInfo) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicKeyTargetInfo) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *ClassicKeyTargetInfo) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *ClassicKeyTargetInfo) SetTargetType(v string) {
	o.TargetType = &v
}

func (o ClassicKeyTargetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalKmsId != nil {
		toSerialize["external_kms_id"] = o.ExternalKmsId
	}
	if o.KeyPurpose != nil {
		toSerialize["key_purpose"] = o.KeyPurpose
	}
	if o.KeyStatus != nil {
		toSerialize["key_status"] = o.KeyStatus
	}
	if o.TargetAssocId != nil {
		toSerialize["target_assoc_id"] = o.TargetAssocId
	}
	if o.TargetType != nil {
		toSerialize["target_type"] = o.TargetType
	}
	return json.Marshal(toSerialize)
}

type NullableClassicKeyTargetInfo struct {
	value *ClassicKeyTargetInfo
	isSet bool
}

func (v NullableClassicKeyTargetInfo) Get() *ClassicKeyTargetInfo {
	return v.value
}

func (v *NullableClassicKeyTargetInfo) Set(val *ClassicKeyTargetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClassicKeyTargetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClassicKeyTargetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassicKeyTargetInfo(val *ClassicKeyTargetInfo) *NullableClassicKeyTargetInfo {
	return &NullableClassicKeyTargetInfo{value: val, isSet: true}
}

func (v NullableClassicKeyTargetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassicKeyTargetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

