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

// KmipDeleteServer kmipDeleteServer is a command that the kmip server (allowed only if it has no clients nor associated items)
type KmipDeleteServer struct {
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipDeleteServer instantiates a new KmipDeleteServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipDeleteServer() *KmipDeleteServer {
	this := KmipDeleteServer{}
	return &this
}

// NewKmipDeleteServerWithDefaults instantiates a new KmipDeleteServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipDeleteServerWithDefaults() *KmipDeleteServer {
	this := KmipDeleteServer{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipDeleteServer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteServer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipDeleteServer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipDeleteServer) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipDeleteServer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipDeleteServer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipDeleteServer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipDeleteServer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipDeleteServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableKmipDeleteServer struct {
	value *KmipDeleteServer
	isSet bool
}

func (v NullableKmipDeleteServer) Get() *KmipDeleteServer {
	return v.value
}

func (v *NullableKmipDeleteServer) Set(val *KmipDeleteServer) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipDeleteServer) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipDeleteServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipDeleteServer(val *KmipDeleteServer) *NullableKmipDeleteServer {
	return &NullableKmipDeleteServer{value: val, isSet: true}
}

func (v NullableKmipDeleteServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipDeleteServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


