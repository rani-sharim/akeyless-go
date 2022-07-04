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

// UpdateAccountSettingsOutput struct for UpdateAccountSettingsOutput
type UpdateAccountSettingsOutput struct {
	Updated *bool `json:"updated,omitempty"`
}

// NewUpdateAccountSettingsOutput instantiates a new UpdateAccountSettingsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountSettingsOutput() *UpdateAccountSettingsOutput {
	this := UpdateAccountSettingsOutput{}
	return &this
}

// NewUpdateAccountSettingsOutputWithDefaults instantiates a new UpdateAccountSettingsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountSettingsOutputWithDefaults() *UpdateAccountSettingsOutput {
	this := UpdateAccountSettingsOutput{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UpdateAccountSettingsOutput) GetUpdated() bool {
	if o == nil || o.Updated == nil {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountSettingsOutput) GetUpdatedOk() (*bool, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UpdateAccountSettingsOutput) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *UpdateAccountSettingsOutput) SetUpdated(v bool) {
	o.Updated = &v
}

func (o UpdateAccountSettingsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountSettingsOutput struct {
	value *UpdateAccountSettingsOutput
	isSet bool
}

func (v NullableUpdateAccountSettingsOutput) Get() *UpdateAccountSettingsOutput {
	return v.value
}

func (v *NullableUpdateAccountSettingsOutput) Set(val *UpdateAccountSettingsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountSettingsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountSettingsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountSettingsOutput(val *UpdateAccountSettingsOutput) *NullableUpdateAccountSettingsOutput {
	return &NullableUpdateAccountSettingsOutput{value: val, isSet: true}
}

func (v NullableUpdateAccountSettingsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountSettingsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

