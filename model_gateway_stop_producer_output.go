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

// GatewayStopProducerOutput struct for GatewayStopProducerOutput
type GatewayStopProducerOutput struct {
	ProducerName *string `json:"producer_name,omitempty"`
}

// NewGatewayStopProducerOutput instantiates a new GatewayStopProducerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayStopProducerOutput() *GatewayStopProducerOutput {
	this := GatewayStopProducerOutput{}
	return &this
}

// NewGatewayStopProducerOutputWithDefaults instantiates a new GatewayStopProducerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayStopProducerOutputWithDefaults() *GatewayStopProducerOutput {
	this := GatewayStopProducerOutput{}
	return &this
}

// GetProducerName returns the ProducerName field value if set, zero value otherwise.
func (o *GatewayStopProducerOutput) GetProducerName() string {
	if o == nil || o.ProducerName == nil {
		var ret string
		return ret
	}
	return *o.ProducerName
}

// GetProducerNameOk returns a tuple with the ProducerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStopProducerOutput) GetProducerNameOk() (*string, bool) {
	if o == nil || o.ProducerName == nil {
		return nil, false
	}
	return o.ProducerName, true
}

// HasProducerName returns a boolean if a field has been set.
func (o *GatewayStopProducerOutput) HasProducerName() bool {
	if o != nil && o.ProducerName != nil {
		return true
	}

	return false
}

// SetProducerName gets a reference to the given string and assigns it to the ProducerName field.
func (o *GatewayStopProducerOutput) SetProducerName(v string) {
	o.ProducerName = &v
}

func (o GatewayStopProducerOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerName != nil {
		toSerialize["producer_name"] = o.ProducerName
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayStopProducerOutput struct {
	value *GatewayStopProducerOutput
	isSet bool
}

func (v NullableGatewayStopProducerOutput) Get() *GatewayStopProducerOutput {
	return v.value
}

func (v *NullableGatewayStopProducerOutput) Set(val *GatewayStopProducerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayStopProducerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayStopProducerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayStopProducerOutput(val *GatewayStopProducerOutput) *NullableGatewayStopProducerOutput {
	return &NullableGatewayStopProducerOutput{value: val, isSet: true}
}

func (v NullableGatewayStopProducerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayStopProducerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

