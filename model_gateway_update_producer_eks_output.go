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

// GatewayUpdateProducerEksOutput struct for GatewayUpdateProducerEksOutput
type GatewayUpdateProducerEksOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayUpdateProducerEksOutput instantiates a new GatewayUpdateProducerEksOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerEksOutput() *GatewayUpdateProducerEksOutput {
	this := GatewayUpdateProducerEksOutput{}
	return &this
}

// NewGatewayUpdateProducerEksOutputWithDefaults instantiates a new GatewayUpdateProducerEksOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerEksOutputWithDefaults() *GatewayUpdateProducerEksOutput {
	this := GatewayUpdateProducerEksOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEksOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEksOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEksOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayUpdateProducerEksOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayUpdateProducerEksOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerEksOutput struct {
	value *GatewayUpdateProducerEksOutput
	isSet bool
}

func (v NullableGatewayUpdateProducerEksOutput) Get() *GatewayUpdateProducerEksOutput {
	return v.value
}

func (v *NullableGatewayUpdateProducerEksOutput) Set(val *GatewayUpdateProducerEksOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerEksOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerEksOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerEksOutput(val *GatewayUpdateProducerEksOutput) *NullableGatewayUpdateProducerEksOutput {
	return &NullableGatewayUpdateProducerEksOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerEksOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerEksOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


