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

// GatewayUpdateTmpUsers gatewayUpdateTmpUsers is a command that returns gateway configuration
type GatewayUpdateTmpUsers struct {
	// Producer Name
	Name string `json:"name"`
	// New TTL in Minutes
	NewTtlMin int64 `json:"new-ttl-min"`
	// Tmp Creds ID
	TmpCredsId string `json:"tmp-creds-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateTmpUsers instantiates a new GatewayUpdateTmpUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateTmpUsers(name string, newTtlMin int64, tmpCredsId string) *GatewayUpdateTmpUsers {
	this := GatewayUpdateTmpUsers{}
	this.Name = name
	this.NewTtlMin = newTtlMin
	this.TmpCredsId = tmpCredsId
	return &this
}

// NewGatewayUpdateTmpUsersWithDefaults instantiates a new GatewayUpdateTmpUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateTmpUsersWithDefaults() *GatewayUpdateTmpUsers {
	this := GatewayUpdateTmpUsers{}
	return &this
}

// GetName returns the Name field value
func (o *GatewayUpdateTmpUsers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateTmpUsers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateTmpUsers) SetName(v string) {
	o.Name = v
}

// GetNewTtlMin returns the NewTtlMin field value
func (o *GatewayUpdateTmpUsers) GetNewTtlMin() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewTtlMin
}

// GetNewTtlMinOk returns a tuple with the NewTtlMin field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateTmpUsers) GetNewTtlMinOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewTtlMin, true
}

// SetNewTtlMin sets field value
func (o *GatewayUpdateTmpUsers) SetNewTtlMin(v int64) {
	o.NewTtlMin = v
}

// GetTmpCredsId returns the TmpCredsId field value
func (o *GatewayUpdateTmpUsers) GetTmpCredsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TmpCredsId
}

// GetTmpCredsIdOk returns a tuple with the TmpCredsId field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateTmpUsers) GetTmpCredsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmpCredsId, true
}

// SetTmpCredsId sets field value
func (o *GatewayUpdateTmpUsers) SetTmpCredsId(v string) {
	o.TmpCredsId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateTmpUsers) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateTmpUsers) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateTmpUsers) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateTmpUsers) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateTmpUsers) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateTmpUsers) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateTmpUsers) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateTmpUsers) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateTmpUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["new-ttl-min"] = o.NewTtlMin
	}
	if true {
		toSerialize["tmp-creds-id"] = o.TmpCredsId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateTmpUsers struct {
	value *GatewayUpdateTmpUsers
	isSet bool
}

func (v NullableGatewayUpdateTmpUsers) Get() *GatewayUpdateTmpUsers {
	return v.value
}

func (v *NullableGatewayUpdateTmpUsers) Set(val *GatewayUpdateTmpUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateTmpUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateTmpUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateTmpUsers(val *GatewayUpdateTmpUsers) *NullableGatewayUpdateTmpUsers {
	return &NullableGatewayUpdateTmpUsers{value: val, isSet: true}
}

func (v NullableGatewayUpdateTmpUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateTmpUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

