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

// GatewayUpdateProducerGcpOutput struct for GatewayUpdateProducerGcpOutput
type GatewayUpdateProducerGcpOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayUpdateProducerGcpOutput instantiates a new GatewayUpdateProducerGcpOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerGcpOutput() *GatewayUpdateProducerGcpOutput {
	this := GatewayUpdateProducerGcpOutput{}
	return &this
}

// NewGatewayUpdateProducerGcpOutputWithDefaults instantiates a new GatewayUpdateProducerGcpOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerGcpOutputWithDefaults() *GatewayUpdateProducerGcpOutput {
	this := GatewayUpdateProducerGcpOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcpOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcpOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcpOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayUpdateProducerGcpOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayUpdateProducerGcpOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerGcpOutput struct {
	value *GatewayUpdateProducerGcpOutput
	isSet bool
}

func (v NullableGatewayUpdateProducerGcpOutput) Get() *GatewayUpdateProducerGcpOutput {
	return v.value
}

func (v *NullableGatewayUpdateProducerGcpOutput) Set(val *GatewayUpdateProducerGcpOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerGcpOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerGcpOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerGcpOutput(val *GatewayUpdateProducerGcpOutput) *NullableGatewayUpdateProducerGcpOutput {
	return &NullableGatewayUpdateProducerGcpOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerGcpOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerGcpOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


