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

// PasswordPolicyInfo struct for PasswordPolicyInfo
type PasswordPolicyInfo struct {
	PasswordLength *int64 `json:"password_length,omitempty"`
	UseCapitalLetters *bool `json:"use_capital_letters,omitempty"`
	UseLowerLetters *bool `json:"use_lower_letters,omitempty"`
	UseNumbers *bool `json:"use_numbers,omitempty"`
	UseSpecialCharacters *bool `json:"use_special_characters,omitempty"`
}

// NewPasswordPolicyInfo instantiates a new PasswordPolicyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyInfo() *PasswordPolicyInfo {
	this := PasswordPolicyInfo{}
	return &this
}

// NewPasswordPolicyInfoWithDefaults instantiates a new PasswordPolicyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyInfoWithDefaults() *PasswordPolicyInfo {
	this := PasswordPolicyInfo{}
	return &this
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *PasswordPolicyInfo) GetPasswordLength() int64 {
	if o == nil || o.PasswordLength == nil {
		var ret int64
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyInfo) GetPasswordLengthOk() (*int64, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *PasswordPolicyInfo) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given int64 and assigns it to the PasswordLength field.
func (o *PasswordPolicyInfo) SetPasswordLength(v int64) {
	o.PasswordLength = &v
}

// GetUseCapitalLetters returns the UseCapitalLetters field value if set, zero value otherwise.
func (o *PasswordPolicyInfo) GetUseCapitalLetters() bool {
	if o == nil || o.UseCapitalLetters == nil {
		var ret bool
		return ret
	}
	return *o.UseCapitalLetters
}

// GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyInfo) GetUseCapitalLettersOk() (*bool, bool) {
	if o == nil || o.UseCapitalLetters == nil {
		return nil, false
	}
	return o.UseCapitalLetters, true
}

// HasUseCapitalLetters returns a boolean if a field has been set.
func (o *PasswordPolicyInfo) HasUseCapitalLetters() bool {
	if o != nil && o.UseCapitalLetters != nil {
		return true
	}

	return false
}

// SetUseCapitalLetters gets a reference to the given bool and assigns it to the UseCapitalLetters field.
func (o *PasswordPolicyInfo) SetUseCapitalLetters(v bool) {
	o.UseCapitalLetters = &v
}

// GetUseLowerLetters returns the UseLowerLetters field value if set, zero value otherwise.
func (o *PasswordPolicyInfo) GetUseLowerLetters() bool {
	if o == nil || o.UseLowerLetters == nil {
		var ret bool
		return ret
	}
	return *o.UseLowerLetters
}

// GetUseLowerLettersOk returns a tuple with the UseLowerLetters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyInfo) GetUseLowerLettersOk() (*bool, bool) {
	if o == nil || o.UseLowerLetters == nil {
		return nil, false
	}
	return o.UseLowerLetters, true
}

// HasUseLowerLetters returns a boolean if a field has been set.
func (o *PasswordPolicyInfo) HasUseLowerLetters() bool {
	if o != nil && o.UseLowerLetters != nil {
		return true
	}

	return false
}

// SetUseLowerLetters gets a reference to the given bool and assigns it to the UseLowerLetters field.
func (o *PasswordPolicyInfo) SetUseLowerLetters(v bool) {
	o.UseLowerLetters = &v
}

// GetUseNumbers returns the UseNumbers field value if set, zero value otherwise.
func (o *PasswordPolicyInfo) GetUseNumbers() bool {
	if o == nil || o.UseNumbers == nil {
		var ret bool
		return ret
	}
	return *o.UseNumbers
}

// GetUseNumbersOk returns a tuple with the UseNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyInfo) GetUseNumbersOk() (*bool, bool) {
	if o == nil || o.UseNumbers == nil {
		return nil, false
	}
	return o.UseNumbers, true
}

// HasUseNumbers returns a boolean if a field has been set.
func (o *PasswordPolicyInfo) HasUseNumbers() bool {
	if o != nil && o.UseNumbers != nil {
		return true
	}

	return false
}

// SetUseNumbers gets a reference to the given bool and assigns it to the UseNumbers field.
func (o *PasswordPolicyInfo) SetUseNumbers(v bool) {
	o.UseNumbers = &v
}

// GetUseSpecialCharacters returns the UseSpecialCharacters field value if set, zero value otherwise.
func (o *PasswordPolicyInfo) GetUseSpecialCharacters() bool {
	if o == nil || o.UseSpecialCharacters == nil {
		var ret bool
		return ret
	}
	return *o.UseSpecialCharacters
}

// GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyInfo) GetUseSpecialCharactersOk() (*bool, bool) {
	if o == nil || o.UseSpecialCharacters == nil {
		return nil, false
	}
	return o.UseSpecialCharacters, true
}

// HasUseSpecialCharacters returns a boolean if a field has been set.
func (o *PasswordPolicyInfo) HasUseSpecialCharacters() bool {
	if o != nil && o.UseSpecialCharacters != nil {
		return true
	}

	return false
}

// SetUseSpecialCharacters gets a reference to the given bool and assigns it to the UseSpecialCharacters field.
func (o *PasswordPolicyInfo) SetUseSpecialCharacters(v bool) {
	o.UseSpecialCharacters = &v
}

func (o PasswordPolicyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordLength != nil {
		toSerialize["password_length"] = o.PasswordLength
	}
	if o.UseCapitalLetters != nil {
		toSerialize["use_capital_letters"] = o.UseCapitalLetters
	}
	if o.UseLowerLetters != nil {
		toSerialize["use_lower_letters"] = o.UseLowerLetters
	}
	if o.UseNumbers != nil {
		toSerialize["use_numbers"] = o.UseNumbers
	}
	if o.UseSpecialCharacters != nil {
		toSerialize["use_special_characters"] = o.UseSpecialCharacters
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordPolicyInfo struct {
	value *PasswordPolicyInfo
	isSet bool
}

func (v NullablePasswordPolicyInfo) Get() *PasswordPolicyInfo {
	return v.value
}

func (v *NullablePasswordPolicyInfo) Set(val *PasswordPolicyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyInfo(val *PasswordPolicyInfo) *NullablePasswordPolicyInfo {
	return &NullablePasswordPolicyInfo{value: val, isSet: true}
}

func (v NullablePasswordPolicyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

