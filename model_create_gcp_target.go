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

// CreateGcpTarget struct for CreateGcpTarget
type CreateGcpTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// Base64-encoded service account private key text
	GcpKey *string `json:"gcp-key,omitempty"`
	// GCP service account email
	GcpSaEmail *string `json:"gcp-sa-email,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	UseGwCloudIdentity *bool `json:"use-gw-cloud-identity,omitempty"`
}

// NewCreateGcpTarget instantiates a new CreateGcpTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGcpTarget(name string) *CreateGcpTarget {
	this := CreateGcpTarget{}
	this.Name = name
	return &this
}

// NewCreateGcpTargetWithDefaults instantiates a new CreateGcpTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGcpTargetWithDefaults() *CreateGcpTarget {
	this := CreateGcpTarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateGcpTarget) SetComment(v string) {
	o.Comment = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *CreateGcpTarget) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetGcpSaEmail returns the GcpSaEmail field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetGcpSaEmail() string {
	if o == nil || o.GcpSaEmail == nil {
		var ret string
		return ret
	}
	return *o.GcpSaEmail
}

// GetGcpSaEmailOk returns a tuple with the GcpSaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetGcpSaEmailOk() (*string, bool) {
	if o == nil || o.GcpSaEmail == nil {
		return nil, false
	}
	return o.GcpSaEmail, true
}

// HasGcpSaEmail returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasGcpSaEmail() bool {
	if o != nil && o.GcpSaEmail != nil {
		return true
	}

	return false
}

// SetGcpSaEmail gets a reference to the given string and assigns it to the GcpSaEmail field.
func (o *CreateGcpTarget) SetGcpSaEmail(v string) {
	o.GcpSaEmail = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateGcpTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *CreateGcpTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGcpTarget) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateGcpTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateGcpTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUseGwCloudIdentity returns the UseGwCloudIdentity field value if set, zero value otherwise.
func (o *CreateGcpTarget) GetUseGwCloudIdentity() bool {
	if o == nil || o.UseGwCloudIdentity == nil {
		var ret bool
		return ret
	}
	return *o.UseGwCloudIdentity
}

// GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGcpTarget) GetUseGwCloudIdentityOk() (*bool, bool) {
	if o == nil || o.UseGwCloudIdentity == nil {
		return nil, false
	}
	return o.UseGwCloudIdentity, true
}

// HasUseGwCloudIdentity returns a boolean if a field has been set.
func (o *CreateGcpTarget) HasUseGwCloudIdentity() bool {
	if o != nil && o.UseGwCloudIdentity != nil {
		return true
	}

	return false
}

// SetUseGwCloudIdentity gets a reference to the given bool and assigns it to the UseGwCloudIdentity field.
func (o *CreateGcpTarget) SetUseGwCloudIdentity(v bool) {
	o.UseGwCloudIdentity = &v
}

func (o CreateGcpTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.GcpKey != nil {
		toSerialize["gcp-key"] = o.GcpKey
	}
	if o.GcpSaEmail != nil {
		toSerialize["gcp-sa-email"] = o.GcpSaEmail
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UseGwCloudIdentity != nil {
		toSerialize["use-gw-cloud-identity"] = o.UseGwCloudIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGcpTarget struct {
	value *CreateGcpTarget
	isSet bool
}

func (v NullableCreateGcpTarget) Get() *CreateGcpTarget {
	return v.value
}

func (v *NullableCreateGcpTarget) Set(val *CreateGcpTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGcpTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGcpTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGcpTarget(val *CreateGcpTarget) *NullableCreateGcpTarget {
	return &NullableCreateGcpTarget{value: val, isSet: true}
}

func (v NullableCreateGcpTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGcpTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


