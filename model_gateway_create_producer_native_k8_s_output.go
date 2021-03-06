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

// GatewayCreateProducerNativeK8SOutput struct for GatewayCreateProducerNativeK8SOutput
type GatewayCreateProducerNativeK8SOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayCreateProducerNativeK8SOutput instantiates a new GatewayCreateProducerNativeK8SOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerNativeK8SOutput() *GatewayCreateProducerNativeK8SOutput {
	this := GatewayCreateProducerNativeK8SOutput{}
	return &this
}

// NewGatewayCreateProducerNativeK8SOutputWithDefaults instantiates a new GatewayCreateProducerNativeK8SOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerNativeK8SOutputWithDefaults() *GatewayCreateProducerNativeK8SOutput {
	this := GatewayCreateProducerNativeK8SOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8SOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8SOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8SOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayCreateProducerNativeK8SOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayCreateProducerNativeK8SOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerNativeK8SOutput struct {
	value *GatewayCreateProducerNativeK8SOutput
	isSet bool
}

func (v NullableGatewayCreateProducerNativeK8SOutput) Get() *GatewayCreateProducerNativeK8SOutput {
	return v.value
}

func (v *NullableGatewayCreateProducerNativeK8SOutput) Set(val *GatewayCreateProducerNativeK8SOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerNativeK8SOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerNativeK8SOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerNativeK8SOutput(val *GatewayCreateProducerNativeK8SOutput) *NullableGatewayCreateProducerNativeK8SOutput {
	return &NullableGatewayCreateProducerNativeK8SOutput{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerNativeK8SOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerNativeK8SOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


