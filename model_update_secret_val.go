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

// UpdateSecretVal struct for UpdateSecretVal
type UpdateSecretVal struct {
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// The provided value is a multiline value (separated by '\\n')
	Multiline *bool `json:"multiline,omitempty"`
	// Secret name
	Name string `json:"name"`
	// Deprecated
	NewVersion *bool `json:"new-version,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The new secret value
	Value string `json:"value"`
}

// NewUpdateSecretVal instantiates a new UpdateSecretVal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSecretVal(name string, value string) *UpdateSecretVal {
	this := UpdateSecretVal{}
	this.Name = name
	this.Value = value
	return &this
}

// NewUpdateSecretValWithDefaults instantiates a new UpdateSecretVal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSecretValWithDefaults() *UpdateSecretVal {
	this := UpdateSecretVal{}
	return &this
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateSecretVal) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateSecretVal) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateSecretVal) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateSecretVal) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateSecretVal) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateSecretVal) SetKey(v string) {
	o.Key = &v
}

// GetMultiline returns the Multiline field value if set, zero value otherwise.
func (o *UpdateSecretVal) GetMultiline() bool {
	if o == nil || o.Multiline == nil {
		var ret bool
		return ret
	}
	return *o.Multiline
}

// GetMultilineOk returns a tuple with the Multiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetMultilineOk() (*bool, bool) {
	if o == nil || o.Multiline == nil {
		return nil, false
	}
	return o.Multiline, true
}

// HasMultiline returns a boolean if a field has been set.
func (o *UpdateSecretVal) HasMultiline() bool {
	if o != nil && o.Multiline != nil {
		return true
	}

	return false
}

// SetMultiline gets a reference to the given bool and assigns it to the Multiline field.
func (o *UpdateSecretVal) SetMultiline(v bool) {
	o.Multiline = &v
}

// GetName returns the Name field value
func (o *UpdateSecretVal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSecretVal) SetName(v string) {
	o.Name = v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateSecretVal) GetNewVersion() bool {
	if o == nil || o.NewVersion == nil {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetNewVersionOk() (*bool, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateSecretVal) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateSecretVal) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateSecretVal) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateSecretVal) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateSecretVal) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateSecretVal) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateSecretVal) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateSecretVal) SetUidToken(v string) {
	o.UidToken = &v
}

// GetValue returns the Value field value
func (o *UpdateSecretVal) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateSecretVal) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateSecretVal) SetValue(v string) {
	o.Value = v
}

func (o UpdateSecretVal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Multiline != nil {
		toSerialize["multiline"] = o.Multiline
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewVersion != nil {
		toSerialize["new-version"] = o.NewVersion
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSecretVal struct {
	value *UpdateSecretVal
	isSet bool
}

func (v NullableUpdateSecretVal) Get() *UpdateSecretVal {
	return v.value
}

func (v *NullableUpdateSecretVal) Set(val *UpdateSecretVal) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSecretVal) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSecretVal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSecretVal(val *UpdateSecretVal) *NullableUpdateSecretVal {
	return &NullableUpdateSecretVal{value: val, isSet: true}
}

func (v NullableUpdateSecretVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSecretVal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


