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

// SraInfo struct for SraInfo
type SraInfo struct {
	Package *string `json:"package,omitempty"`
	Tier *string `json:"tier,omitempty"`
	UserType *string `json:"user_type,omitempty"`
}

// NewSraInfo instantiates a new SraInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSraInfo() *SraInfo {
	this := SraInfo{}
	return &this
}

// NewSraInfoWithDefaults instantiates a new SraInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSraInfoWithDefaults() *SraInfo {
	this := SraInfo{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *SraInfo) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SraInfo) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *SraInfo) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *SraInfo) SetPackage(v string) {
	o.Package = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *SraInfo) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SraInfo) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *SraInfo) HasTier() bool {
	if o != nil && o.Tier != nil {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *SraInfo) SetTier(v string) {
	o.Tier = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *SraInfo) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SraInfo) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *SraInfo) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *SraInfo) SetUserType(v string) {
	o.UserType = &v
}

func (o SraInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	return json.Marshal(toSerialize)
}

type NullableSraInfo struct {
	value *SraInfo
	isSet bool
}

func (v NullableSraInfo) Get() *SraInfo {
	return v.value
}

func (v *NullableSraInfo) Set(val *SraInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSraInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSraInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSraInfo(val *SraInfo) *NullableSraInfo {
	return &NullableSraInfo{value: val, isSet: true}
}

func (v NullableSraInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSraInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


