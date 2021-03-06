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

// UpdateNativeK8STargetOutput struct for UpdateNativeK8STargetOutput
type UpdateNativeK8STargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewUpdateNativeK8STargetOutput instantiates a new UpdateNativeK8STargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNativeK8STargetOutput() *UpdateNativeK8STargetOutput {
	this := UpdateNativeK8STargetOutput{}
	return &this
}

// NewUpdateNativeK8STargetOutputWithDefaults instantiates a new UpdateNativeK8STargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNativeK8STargetOutputWithDefaults() *UpdateNativeK8STargetOutput {
	this := UpdateNativeK8STargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateNativeK8STargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNativeK8STargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateNativeK8STargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *UpdateNativeK8STargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o UpdateNativeK8STargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNativeK8STargetOutput struct {
	value *UpdateNativeK8STargetOutput
	isSet bool
}

func (v NullableUpdateNativeK8STargetOutput) Get() *UpdateNativeK8STargetOutput {
	return v.value
}

func (v *NullableUpdateNativeK8STargetOutput) Set(val *UpdateNativeK8STargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNativeK8STargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNativeK8STargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNativeK8STargetOutput(val *UpdateNativeK8STargetOutput) *NullableUpdateNativeK8STargetOutput {
	return &NullableUpdateNativeK8STargetOutput{value: val, isSet: true}
}

func (v NullableUpdateNativeK8STargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNativeK8STargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


