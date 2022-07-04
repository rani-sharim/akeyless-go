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

// KmipRenewServerCertificate struct for KmipRenewServerCertificate
type KmipRenewServerCertificate struct {
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipRenewServerCertificate instantiates a new KmipRenewServerCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipRenewServerCertificate() *KmipRenewServerCertificate {
	this := KmipRenewServerCertificate{}
	return &this
}

// NewKmipRenewServerCertificateWithDefaults instantiates a new KmipRenewServerCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipRenewServerCertificateWithDefaults() *KmipRenewServerCertificate {
	this := KmipRenewServerCertificate{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipRenewServerCertificate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipRenewServerCertificate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipRenewServerCertificate) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipRenewServerCertificate) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipRenewServerCertificate) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipRenewServerCertificate) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipRenewServerCertificate) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipRenewServerCertificate) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipRenewServerCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableKmipRenewServerCertificate struct {
	value *KmipRenewServerCertificate
	isSet bool
}

func (v NullableKmipRenewServerCertificate) Get() *KmipRenewServerCertificate {
	return v.value
}

func (v *NullableKmipRenewServerCertificate) Set(val *KmipRenewServerCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipRenewServerCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipRenewServerCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipRenewServerCertificate(val *KmipRenewServerCertificate) *NullableKmipRenewServerCertificate {
	return &NullableKmipRenewServerCertificate{value: val, isSet: true}
}

func (v NullableKmipRenewServerCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipRenewServerCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


