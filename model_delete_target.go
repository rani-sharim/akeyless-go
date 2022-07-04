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

// DeleteTarget struct for DeleteTarget
type DeleteTarget struct {
	// Enforce deletion
	ForceDeletion *bool `json:"force-deletion,omitempty"`
	// Target name
	Name string `json:"name"`
	// Target version
	TargetVersion *int32 `json:"target-version,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDeleteTarget instantiates a new DeleteTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTarget(name string) *DeleteTarget {
	this := DeleteTarget{}
	var forceDeletion bool = false
	this.ForceDeletion = &forceDeletion
	this.Name = name
	return &this
}

// NewDeleteTargetWithDefaults instantiates a new DeleteTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTargetWithDefaults() *DeleteTarget {
	this := DeleteTarget{}
	var forceDeletion bool = false
	this.ForceDeletion = &forceDeletion
	return &this
}

// GetForceDeletion returns the ForceDeletion field value if set, zero value otherwise.
func (o *DeleteTarget) GetForceDeletion() bool {
	if o == nil || o.ForceDeletion == nil {
		var ret bool
		return ret
	}
	return *o.ForceDeletion
}

// GetForceDeletionOk returns a tuple with the ForceDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTarget) GetForceDeletionOk() (*bool, bool) {
	if o == nil || o.ForceDeletion == nil {
		return nil, false
	}
	return o.ForceDeletion, true
}

// HasForceDeletion returns a boolean if a field has been set.
func (o *DeleteTarget) HasForceDeletion() bool {
	if o != nil && o.ForceDeletion != nil {
		return true
	}

	return false
}

// SetForceDeletion gets a reference to the given bool and assigns it to the ForceDeletion field.
func (o *DeleteTarget) SetForceDeletion(v bool) {
	o.ForceDeletion = &v
}

// GetName returns the Name field value
func (o *DeleteTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeleteTarget) SetName(v string) {
	o.Name = v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *DeleteTarget) GetTargetVersion() int32 {
	if o == nil || o.TargetVersion == nil {
		var ret int32
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTarget) GetTargetVersionOk() (*int32, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *DeleteTarget) HasTargetVersion() bool {
	if o != nil && o.TargetVersion != nil {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given int32 and assigns it to the TargetVersion field.
func (o *DeleteTarget) SetTargetVersion(v int32) {
	o.TargetVersion = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DeleteTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceDeletion != nil {
		toSerialize["force-deletion"] = o.ForceDeletion
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.TargetVersion != nil {
		toSerialize["target-version"] = o.TargetVersion
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteTarget struct {
	value *DeleteTarget
	isSet bool
}

func (v NullableDeleteTarget) Get() *DeleteTarget {
	return v.value
}

func (v *NullableDeleteTarget) Set(val *DeleteTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTarget(val *DeleteTarget) *NullableDeleteTarget {
	return &NullableDeleteTarget{value: val, isSet: true}
}

func (v NullableDeleteTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


