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

// ReverseRBACOutput struct for ReverseRBACOutput
type ReverseRBACOutput struct {
	Clients *map[string]ReverseRBACClient `json:"clients,omitempty"`
}

// NewReverseRBACOutput instantiates a new ReverseRBACOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseRBACOutput() *ReverseRBACOutput {
	this := ReverseRBACOutput{}
	return &this
}

// NewReverseRBACOutputWithDefaults instantiates a new ReverseRBACOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseRBACOutputWithDefaults() *ReverseRBACOutput {
	this := ReverseRBACOutput{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ReverseRBACOutput) GetClients() map[string]ReverseRBACClient {
	if o == nil || o.Clients == nil {
		var ret map[string]ReverseRBACClient
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseRBACOutput) GetClientsOk() (*map[string]ReverseRBACClient, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ReverseRBACOutput) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given map[string]ReverseRBACClient and assigns it to the Clients field.
func (o *ReverseRBACOutput) SetClients(v map[string]ReverseRBACClient) {
	o.Clients = &v
}

func (o ReverseRBACOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableReverseRBACOutput struct {
	value *ReverseRBACOutput
	isSet bool
}

func (v NullableReverseRBACOutput) Get() *ReverseRBACOutput {
	return v.value
}

func (v *NullableReverseRBACOutput) Set(val *ReverseRBACOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseRBACOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseRBACOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseRBACOutput(val *ReverseRBACOutput) *NullableReverseRBACOutput {
	return &NullableReverseRBACOutput{value: val, isSet: true}
}

func (v NullableReverseRBACOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseRBACOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


