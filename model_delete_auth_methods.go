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

// DeleteAuthMethods deleteAuthMethods is a command that deletes multiple auth methods from a given path.
type DeleteAuthMethods struct {
	// Path to delete the auth methods from
	Path string `json:"path"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDeleteAuthMethods instantiates a new DeleteAuthMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAuthMethods(path string) *DeleteAuthMethods {
	this := DeleteAuthMethods{}
	this.Path = path
	return &this
}

// NewDeleteAuthMethodsWithDefaults instantiates a new DeleteAuthMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAuthMethodsWithDefaults() *DeleteAuthMethods {
	this := DeleteAuthMethods{}
	return &this
}

// GetPath returns the Path field value
func (o *DeleteAuthMethods) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethods) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DeleteAuthMethods) SetPath(v string) {
	o.Path = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteAuthMethods) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethods) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteAuthMethods) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteAuthMethods) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteAuthMethods) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethods) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteAuthMethods) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteAuthMethods) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DeleteAuthMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAuthMethods struct {
	value *DeleteAuthMethods
	isSet bool
}

func (v NullableDeleteAuthMethods) Get() *DeleteAuthMethods {
	return v.value
}

func (v *NullableDeleteAuthMethods) Set(val *DeleteAuthMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAuthMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAuthMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAuthMethods(val *DeleteAuthMethods) *NullableDeleteAuthMethods {
	return &NullableDeleteAuthMethods{value: val, isSet: true}
}

func (v NullableDeleteAuthMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAuthMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


