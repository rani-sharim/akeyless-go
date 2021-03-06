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

// KmipMoveServer kmipMoveServer is a command that Moves the root location of the kmip server and all associated items to a new root location
type KmipMoveServer struct {
	NewRoot *string `json:"new-root,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipMoveServer instantiates a new KmipMoveServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipMoveServer() *KmipMoveServer {
	this := KmipMoveServer{}
	return &this
}

// NewKmipMoveServerWithDefaults instantiates a new KmipMoveServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipMoveServerWithDefaults() *KmipMoveServer {
	this := KmipMoveServer{}
	return &this
}

// GetNewRoot returns the NewRoot field value if set, zero value otherwise.
func (o *KmipMoveServer) GetNewRoot() string {
	if o == nil || o.NewRoot == nil {
		var ret string
		return ret
	}
	return *o.NewRoot
}

// GetNewRootOk returns a tuple with the NewRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipMoveServer) GetNewRootOk() (*string, bool) {
	if o == nil || o.NewRoot == nil {
		return nil, false
	}
	return o.NewRoot, true
}

// HasNewRoot returns a boolean if a field has been set.
func (o *KmipMoveServer) HasNewRoot() bool {
	if o != nil && o.NewRoot != nil {
		return true
	}

	return false
}

// SetNewRoot gets a reference to the given string and assigns it to the NewRoot field.
func (o *KmipMoveServer) SetNewRoot(v string) {
	o.NewRoot = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipMoveServer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipMoveServer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipMoveServer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipMoveServer) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipMoveServer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipMoveServer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipMoveServer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipMoveServer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipMoveServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewRoot != nil {
		toSerialize["new-root"] = o.NewRoot
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableKmipMoveServer struct {
	value *KmipMoveServer
	isSet bool
}

func (v NullableKmipMoveServer) Get() *KmipMoveServer {
	return v.value
}

func (v *NullableKmipMoveServer) Set(val *KmipMoveServer) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipMoveServer) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipMoveServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipMoveServer(val *KmipMoveServer) *NullableKmipMoveServer {
	return &NullableKmipMoveServer{value: val, isSet: true}
}

func (v NullableKmipMoveServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipMoveServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


