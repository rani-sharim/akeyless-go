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

// GatewayDeleteProducer gatewayDeleteProducer is a command that deletes producer
type GatewayDeleteProducer struct {
	// Producer name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayDeleteProducer instantiates a new GatewayDeleteProducer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeleteProducer(name string) *GatewayDeleteProducer {
	this := GatewayDeleteProducer{}
	this.Name = name
	return &this
}

// NewGatewayDeleteProducerWithDefaults instantiates a new GatewayDeleteProducer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeleteProducerWithDefaults() *GatewayDeleteProducer {
	this := GatewayDeleteProducer{}
	return &this
}

// GetName returns the Name field value
func (o *GatewayDeleteProducer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayDeleteProducer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayDeleteProducer) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayDeleteProducer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteProducer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayDeleteProducer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayDeleteProducer) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayDeleteProducer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteProducer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayDeleteProducer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayDeleteProducer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayDeleteProducer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayDeleteProducer struct {
	value *GatewayDeleteProducer
	isSet bool
}

func (v NullableGatewayDeleteProducer) Get() *GatewayDeleteProducer {
	return v.value
}

func (v *NullableGatewayDeleteProducer) Set(val *GatewayDeleteProducer) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeleteProducer) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeleteProducer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeleteProducer(val *GatewayDeleteProducer) *NullableGatewayDeleteProducer {
	return &NullableGatewayDeleteProducer{value: val, isSet: true}
}

func (v NullableGatewayDeleteProducer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeleteProducer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


