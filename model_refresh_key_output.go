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

// RefreshKeyOutput struct for RefreshKeyOutput
type RefreshKeyOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewRefreshKeyOutput instantiates a new RefreshKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshKeyOutput() *RefreshKeyOutput {
	this := RefreshKeyOutput{}
	return &this
}

// NewRefreshKeyOutputWithDefaults instantiates a new RefreshKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshKeyOutputWithDefaults() *RefreshKeyOutput {
	this := RefreshKeyOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RefreshKeyOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshKeyOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RefreshKeyOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *RefreshKeyOutput) SetResult(v string) {
	o.Result = &v
}

func (o RefreshKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshKeyOutput struct {
	value *RefreshKeyOutput
	isSet bool
}

func (v NullableRefreshKeyOutput) Get() *RefreshKeyOutput {
	return v.value
}

func (v *NullableRefreshKeyOutput) Set(val *RefreshKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshKeyOutput(val *RefreshKeyOutput) *NullableRefreshKeyOutput {
	return &NullableRefreshKeyOutput{value: val, isSet: true}
}

func (v NullableRefreshKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


