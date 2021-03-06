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

// UpdateTargetDetailsOutput struct for UpdateTargetDetailsOutput
type UpdateTargetDetailsOutput struct {
	Updated *bool `json:"updated,omitempty"`
}

// NewUpdateTargetDetailsOutput instantiates a new UpdateTargetDetailsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTargetDetailsOutput() *UpdateTargetDetailsOutput {
	this := UpdateTargetDetailsOutput{}
	return &this
}

// NewUpdateTargetDetailsOutputWithDefaults instantiates a new UpdateTargetDetailsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTargetDetailsOutputWithDefaults() *UpdateTargetDetailsOutput {
	this := UpdateTargetDetailsOutput{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UpdateTargetDetailsOutput) GetUpdated() bool {
	if o == nil || o.Updated == nil {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetDetailsOutput) GetUpdatedOk() (*bool, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UpdateTargetDetailsOutput) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *UpdateTargetDetailsOutput) SetUpdated(v bool) {
	o.Updated = &v
}

func (o UpdateTargetDetailsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTargetDetailsOutput struct {
	value *UpdateTargetDetailsOutput
	isSet bool
}

func (v NullableUpdateTargetDetailsOutput) Get() *UpdateTargetDetailsOutput {
	return v.value
}

func (v *NullableUpdateTargetDetailsOutput) Set(val *UpdateTargetDetailsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTargetDetailsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTargetDetailsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTargetDetailsOutput(val *UpdateTargetDetailsOutput) *NullableUpdateTargetDetailsOutput {
	return &NullableUpdateTargetDetailsOutput{value: val, isSet: true}
}

func (v NullableUpdateTargetDetailsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTargetDetailsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


