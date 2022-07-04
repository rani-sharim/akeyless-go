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

// CreateAWSTargetOutput struct for CreateAWSTargetOutput
type CreateAWSTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewCreateAWSTargetOutput instantiates a new CreateAWSTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAWSTargetOutput() *CreateAWSTargetOutput {
	this := CreateAWSTargetOutput{}
	return &this
}

// NewCreateAWSTargetOutputWithDefaults instantiates a new CreateAWSTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAWSTargetOutputWithDefaults() *CreateAWSTargetOutput {
	this := CreateAWSTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CreateAWSTargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAWSTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CreateAWSTargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *CreateAWSTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o CreateAWSTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAWSTargetOutput struct {
	value *CreateAWSTargetOutput
	isSet bool
}

func (v NullableCreateAWSTargetOutput) Get() *CreateAWSTargetOutput {
	return v.value
}

func (v *NullableCreateAWSTargetOutput) Set(val *CreateAWSTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAWSTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAWSTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAWSTargetOutput(val *CreateAWSTargetOutput) *NullableCreateAWSTargetOutput {
	return &NullableCreateAWSTargetOutput{value: val, isSet: true}
}

func (v NullableCreateAWSTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAWSTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


