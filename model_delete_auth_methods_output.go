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

// DeleteAuthMethodsOutput struct for DeleteAuthMethodsOutput
type DeleteAuthMethodsOutput struct {
	Path *string `json:"path,omitempty"`
}

// NewDeleteAuthMethodsOutput instantiates a new DeleteAuthMethodsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAuthMethodsOutput() *DeleteAuthMethodsOutput {
	this := DeleteAuthMethodsOutput{}
	return &this
}

// NewDeleteAuthMethodsOutputWithDefaults instantiates a new DeleteAuthMethodsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAuthMethodsOutputWithDefaults() *DeleteAuthMethodsOutput {
	this := DeleteAuthMethodsOutput{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DeleteAuthMethodsOutput) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethodsOutput) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DeleteAuthMethodsOutput) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DeleteAuthMethodsOutput) SetPath(v string) {
	o.Path = &v
}

func (o DeleteAuthMethodsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAuthMethodsOutput struct {
	value *DeleteAuthMethodsOutput
	isSet bool
}

func (v NullableDeleteAuthMethodsOutput) Get() *DeleteAuthMethodsOutput {
	return v.value
}

func (v *NullableDeleteAuthMethodsOutput) Set(val *DeleteAuthMethodsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAuthMethodsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAuthMethodsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAuthMethodsOutput(val *DeleteAuthMethodsOutput) *NullableDeleteAuthMethodsOutput {
	return &NullableDeleteAuthMethodsOutput{value: val, isSet: true}
}

func (v NullableDeleteAuthMethodsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAuthMethodsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


