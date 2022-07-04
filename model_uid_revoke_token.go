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

// UidRevokeToken struct for UidRevokeToken
type UidRevokeToken struct {
	// The universal identity auth method name
	AuthMethodName *string `json:"auth-method-name,omitempty"`
	// the universal identity token/token-id to revoke
	RevokeToken string `json:"revoke-token"`
	// revokeSelf/revokeAll (delete only this token/this token and his children)
	RevokeType string `json:"revoke-type"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUidRevokeToken instantiates a new UidRevokeToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUidRevokeToken(revokeToken string, revokeType string) *UidRevokeToken {
	this := UidRevokeToken{}
	this.RevokeToken = revokeToken
	this.RevokeType = revokeType
	return &this
}

// NewUidRevokeTokenWithDefaults instantiates a new UidRevokeToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUidRevokeTokenWithDefaults() *UidRevokeToken {
	this := UidRevokeToken{}
	return &this
}

// GetAuthMethodName returns the AuthMethodName field value if set, zero value otherwise.
func (o *UidRevokeToken) GetAuthMethodName() string {
	if o == nil || o.AuthMethodName == nil {
		var ret string
		return ret
	}
	return *o.AuthMethodName
}

// GetAuthMethodNameOk returns a tuple with the AuthMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidRevokeToken) GetAuthMethodNameOk() (*string, bool) {
	if o == nil || o.AuthMethodName == nil {
		return nil, false
	}
	return o.AuthMethodName, true
}

// HasAuthMethodName returns a boolean if a field has been set.
func (o *UidRevokeToken) HasAuthMethodName() bool {
	if o != nil && o.AuthMethodName != nil {
		return true
	}

	return false
}

// SetAuthMethodName gets a reference to the given string and assigns it to the AuthMethodName field.
func (o *UidRevokeToken) SetAuthMethodName(v string) {
	o.AuthMethodName = &v
}

// GetRevokeToken returns the RevokeToken field value
func (o *UidRevokeToken) GetRevokeToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevokeToken
}

// GetRevokeTokenOk returns a tuple with the RevokeToken field value
// and a boolean to check if the value has been set.
func (o *UidRevokeToken) GetRevokeTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevokeToken, true
}

// SetRevokeToken sets field value
func (o *UidRevokeToken) SetRevokeToken(v string) {
	o.RevokeToken = v
}

// GetRevokeType returns the RevokeType field value
func (o *UidRevokeToken) GetRevokeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevokeType
}

// GetRevokeTypeOk returns a tuple with the RevokeType field value
// and a boolean to check if the value has been set.
func (o *UidRevokeToken) GetRevokeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevokeType, true
}

// SetRevokeType sets field value
func (o *UidRevokeToken) SetRevokeType(v string) {
	o.RevokeType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UidRevokeToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidRevokeToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UidRevokeToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UidRevokeToken) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UidRevokeToken) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidRevokeToken) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UidRevokeToken) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UidRevokeToken) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UidRevokeToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthMethodName != nil {
		toSerialize["auth-method-name"] = o.AuthMethodName
	}
	if true {
		toSerialize["revoke-token"] = o.RevokeToken
	}
	if true {
		toSerialize["revoke-type"] = o.RevokeType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUidRevokeToken struct {
	value *UidRevokeToken
	isSet bool
}

func (v NullableUidRevokeToken) Get() *UidRevokeToken {
	return v.value
}

func (v *NullableUidRevokeToken) Set(val *UidRevokeToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUidRevokeToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUidRevokeToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUidRevokeToken(val *UidRevokeToken) *NullableUidRevokeToken {
	return &NullableUidRevokeToken{value: val, isSet: true}
}

func (v NullableUidRevokeToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUidRevokeToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


