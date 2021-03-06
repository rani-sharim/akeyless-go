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

// GatewayCreateProducerMongoOutput struct for GatewayCreateProducerMongoOutput
type GatewayCreateProducerMongoOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayCreateProducerMongoOutput instantiates a new GatewayCreateProducerMongoOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerMongoOutput() *GatewayCreateProducerMongoOutput {
	this := GatewayCreateProducerMongoOutput{}
	return &this
}

// NewGatewayCreateProducerMongoOutputWithDefaults instantiates a new GatewayCreateProducerMongoOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerMongoOutputWithDefaults() *GatewayCreateProducerMongoOutput {
	this := GatewayCreateProducerMongoOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongoOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongoOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongoOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayCreateProducerMongoOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayCreateProducerMongoOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerMongoOutput struct {
	value *GatewayCreateProducerMongoOutput
	isSet bool
}

func (v NullableGatewayCreateProducerMongoOutput) Get() *GatewayCreateProducerMongoOutput {
	return v.value
}

func (v *NullableGatewayCreateProducerMongoOutput) Set(val *GatewayCreateProducerMongoOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerMongoOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerMongoOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerMongoOutput(val *GatewayCreateProducerMongoOutput) *NullableGatewayCreateProducerMongoOutput {
	return &NullableGatewayCreateProducerMongoOutput{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerMongoOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerMongoOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


