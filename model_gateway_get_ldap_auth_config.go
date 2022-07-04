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

// GatewayGetLdapAuthConfig gatewayGetLdapAuth is a command that gets ldap auth config
type GatewayGetLdapAuthConfig struct {
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayGetLdapAuthConfig instantiates a new GatewayGetLdapAuthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayGetLdapAuthConfig() *GatewayGetLdapAuthConfig {
	this := GatewayGetLdapAuthConfig{}
	return &this
}

// NewGatewayGetLdapAuthConfigWithDefaults instantiates a new GatewayGetLdapAuthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayGetLdapAuthConfigWithDefaults() *GatewayGetLdapAuthConfig {
	this := GatewayGetLdapAuthConfig{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayGetLdapAuthConfig) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetLdapAuthConfig) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayGetLdapAuthConfig) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayGetLdapAuthConfig) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayGetLdapAuthConfig) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetLdapAuthConfig) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayGetLdapAuthConfig) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayGetLdapAuthConfig) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayGetLdapAuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayGetLdapAuthConfig struct {
	value *GatewayGetLdapAuthConfig
	isSet bool
}

func (v NullableGatewayGetLdapAuthConfig) Get() *GatewayGetLdapAuthConfig {
	return v.value
}

func (v *NullableGatewayGetLdapAuthConfig) Set(val *GatewayGetLdapAuthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayGetLdapAuthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayGetLdapAuthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayGetLdapAuthConfig(val *GatewayGetLdapAuthConfig) *NullableGatewayGetLdapAuthConfig {
	return &NullableGatewayGetLdapAuthConfig{value: val, isSet: true}
}

func (v NullableGatewayGetLdapAuthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayGetLdapAuthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


