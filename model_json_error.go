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

// JSONError struct for JSONError
type JSONError struct {
	Error *string `json:"error,omitempty"`
}

// NewJSONError instantiates a new JSONError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONError() *JSONError {
	this := JSONError{}
	return &this
}

// NewJSONErrorWithDefaults instantiates a new JSONError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONErrorWithDefaults() *JSONError {
	this := JSONError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *JSONError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *JSONError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *JSONError) SetError(v string) {
	o.Error = &v
}

func (o JSONError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableJSONError struct {
	value *JSONError
	isSet bool
}

func (v NullableJSONError) Get() *JSONError {
	return v.value
}

func (v *NullableJSONError) Set(val *JSONError) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONError) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONError(val *JSONError) *NullableJSONError {
	return &NullableJSONError{value: val, isSet: true}
}

func (v NullableJSONError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


