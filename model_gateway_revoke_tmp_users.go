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

// GatewayRevokeTmpUsers gatewayRevokeTmpUsers is a command that revoke producer tmp user
type GatewayRevokeTmpUsers struct {
	// Host
	Host *string `json:"host,omitempty"`
	// Producer Name
	Name string `json:"name"`
	// Soft Delete
	SoftDelete *bool `json:"soft-delete,omitempty"`
	// Tmp Creds ID
	TmpCredsId string `json:"tmp-creds-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayRevokeTmpUsers instantiates a new GatewayRevokeTmpUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayRevokeTmpUsers(name string, tmpCredsId string) *GatewayRevokeTmpUsers {
	this := GatewayRevokeTmpUsers{}
	this.Name = name
	this.TmpCredsId = tmpCredsId
	return &this
}

// NewGatewayRevokeTmpUsersWithDefaults instantiates a new GatewayRevokeTmpUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayRevokeTmpUsersWithDefaults() *GatewayRevokeTmpUsers {
	this := GatewayRevokeTmpUsers{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GatewayRevokeTmpUsers) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayRevokeTmpUsers) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GatewayRevokeTmpUsers) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GatewayRevokeTmpUsers) SetHost(v string) {
	o.Host = &v
}

// GetName returns the Name field value
func (o *GatewayRevokeTmpUsers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayRevokeTmpUsers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayRevokeTmpUsers) SetName(v string) {
	o.Name = v
}

// GetSoftDelete returns the SoftDelete field value if set, zero value otherwise.
func (o *GatewayRevokeTmpUsers) GetSoftDelete() bool {
	if o == nil || o.SoftDelete == nil {
		var ret bool
		return ret
	}
	return *o.SoftDelete
}

// GetSoftDeleteOk returns a tuple with the SoftDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayRevokeTmpUsers) GetSoftDeleteOk() (*bool, bool) {
	if o == nil || o.SoftDelete == nil {
		return nil, false
	}
	return o.SoftDelete, true
}

// HasSoftDelete returns a boolean if a field has been set.
func (o *GatewayRevokeTmpUsers) HasSoftDelete() bool {
	if o != nil && o.SoftDelete != nil {
		return true
	}

	return false
}

// SetSoftDelete gets a reference to the given bool and assigns it to the SoftDelete field.
func (o *GatewayRevokeTmpUsers) SetSoftDelete(v bool) {
	o.SoftDelete = &v
}

// GetTmpCredsId returns the TmpCredsId field value
func (o *GatewayRevokeTmpUsers) GetTmpCredsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TmpCredsId
}

// GetTmpCredsIdOk returns a tuple with the TmpCredsId field value
// and a boolean to check if the value has been set.
func (o *GatewayRevokeTmpUsers) GetTmpCredsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmpCredsId, true
}

// SetTmpCredsId sets field value
func (o *GatewayRevokeTmpUsers) SetTmpCredsId(v string) {
	o.TmpCredsId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayRevokeTmpUsers) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayRevokeTmpUsers) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayRevokeTmpUsers) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayRevokeTmpUsers) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayRevokeTmpUsers) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayRevokeTmpUsers) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayRevokeTmpUsers) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayRevokeTmpUsers) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayRevokeTmpUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SoftDelete != nil {
		toSerialize["soft-delete"] = o.SoftDelete
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

type NullableGatewayRevokeTmpUsers struct {
	value *GatewayRevokeTmpUsers
	isSet bool
}

func (v NullableGatewayRevokeTmpUsers) Get() *GatewayRevokeTmpUsers {
	return v.value
}

func (v *NullableGatewayRevokeTmpUsers) Set(val *GatewayRevokeTmpUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayRevokeTmpUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayRevokeTmpUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayRevokeTmpUsers(val *GatewayRevokeTmpUsers) *NullableGatewayRevokeTmpUsers {
	return &NullableGatewayRevokeTmpUsers{value: val, isSet: true}
}

func (v NullableGatewayRevokeTmpUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayRevokeTmpUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


