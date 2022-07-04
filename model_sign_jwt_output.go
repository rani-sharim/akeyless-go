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

// SignJWTOutput struct for SignJWTOutput
type SignJWTOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewSignJWTOutput instantiates a new SignJWTOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignJWTOutput() *SignJWTOutput {
	this := SignJWTOutput{}
	return &this
}

// NewSignJWTOutputWithDefaults instantiates a new SignJWTOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignJWTOutputWithDefaults() *SignJWTOutput {
	this := SignJWTOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SignJWTOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignJWTOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SignJWTOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *SignJWTOutput) SetResult(v string) {
	o.Result = &v
}

func (o SignJWTOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableSignJWTOutput struct {
	value *SignJWTOutput
	isSet bool
}

func (v NullableSignJWTOutput) Get() *SignJWTOutput {
	return v.value
}

func (v *NullableSignJWTOutput) Set(val *SignJWTOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableSignJWTOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableSignJWTOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignJWTOutput(val *SignJWTOutput) *NullableSignJWTOutput {
	return &NullableSignJWTOutput{value: val, isSet: true}
}

func (v NullableSignJWTOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignJWTOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


