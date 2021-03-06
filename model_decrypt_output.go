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

// DecryptOutput struct for DecryptOutput
type DecryptOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewDecryptOutput instantiates a new DecryptOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptOutput() *DecryptOutput {
	this := DecryptOutput{}
	return &this
}

// NewDecryptOutputWithDefaults instantiates a new DecryptOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptOutputWithDefaults() *DecryptOutput {
	this := DecryptOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DecryptOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DecryptOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DecryptOutput) SetResult(v string) {
	o.Result = &v
}

func (o DecryptOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableDecryptOutput struct {
	value *DecryptOutput
	isSet bool
}

func (v NullableDecryptOutput) Get() *DecryptOutput {
	return v.value
}

func (v *NullableDecryptOutput) Set(val *DecryptOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptOutput(val *DecryptOutput) *NullableDecryptOutput {
	return &NullableDecryptOutput{value: val, isSet: true}
}

func (v NullableDecryptOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


