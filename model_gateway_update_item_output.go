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

// GatewayUpdateItemOutput struct for GatewayUpdateItemOutput
type GatewayUpdateItemOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewGatewayUpdateItemOutput instantiates a new GatewayUpdateItemOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateItemOutput() *GatewayUpdateItemOutput {
	this := GatewayUpdateItemOutput{}
	return &this
}

// NewGatewayUpdateItemOutputWithDefaults instantiates a new GatewayUpdateItemOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateItemOutputWithDefaults() *GatewayUpdateItemOutput {
	this := GatewayUpdateItemOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GatewayUpdateItemOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItemOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GatewayUpdateItemOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GatewayUpdateItemOutput) SetName(v string) {
	o.Name = &v
}

func (o GatewayUpdateItemOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateItemOutput struct {
	value *GatewayUpdateItemOutput
	isSet bool
}

func (v NullableGatewayUpdateItemOutput) Get() *GatewayUpdateItemOutput {
	return v.value
}

func (v *NullableGatewayUpdateItemOutput) Set(val *GatewayUpdateItemOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateItemOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateItemOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateItemOutput(val *GatewayUpdateItemOutput) *NullableGatewayUpdateItemOutput {
	return &NullableGatewayUpdateItemOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateItemOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateItemOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


