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

// GatewayCreateProducerOracleDbOutput struct for GatewayCreateProducerOracleDbOutput
type GatewayCreateProducerOracleDbOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayCreateProducerOracleDbOutput instantiates a new GatewayCreateProducerOracleDbOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerOracleDbOutput() *GatewayCreateProducerOracleDbOutput {
	this := GatewayCreateProducerOracleDbOutput{}
	return &this
}

// NewGatewayCreateProducerOracleDbOutputWithDefaults instantiates a new GatewayCreateProducerOracleDbOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerOracleDbOutputWithDefaults() *GatewayCreateProducerOracleDbOutput {
	this := GatewayCreateProducerOracleDbOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDbOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || o.ProducerDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDbOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.ProducerDetails == nil {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDbOutput) HasProducerDetails() bool {
	if o != nil && o.ProducerDetails != nil {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayCreateProducerOracleDbOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayCreateProducerOracleDbOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProducerDetails != nil {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerOracleDbOutput struct {
	value *GatewayCreateProducerOracleDbOutput
	isSet bool
}

func (v NullableGatewayCreateProducerOracleDbOutput) Get() *GatewayCreateProducerOracleDbOutput {
	return v.value
}

func (v *NullableGatewayCreateProducerOracleDbOutput) Set(val *GatewayCreateProducerOracleDbOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerOracleDbOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerOracleDbOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerOracleDbOutput(val *GatewayCreateProducerOracleDbOutput) *NullableGatewayCreateProducerOracleDbOutput {
	return &NullableGatewayCreateProducerOracleDbOutput{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerOracleDbOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerOracleDbOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

