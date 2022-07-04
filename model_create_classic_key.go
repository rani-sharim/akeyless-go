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

// CreateClassicKey CreateClassicKey is a command that creates classic key
type CreateClassicKey struct {
	// Classic Key type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, RSA1024, RSA2048, RSA3072, RSA4096, EC256, EC384]
	Alg string `json:"alg"`
	// Certificate in a PEM format.
	CertFileData *string `json:"cert-file-data,omitempty"`
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Base64-encoded classic key value
	KeyData *string `json:"key-data,omitempty"`
	// Metadata about the classic key
	Metadata *string `json:"metadata,omitempty"`
	// ClassicKey name
	Name string `json:"name"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	ProtectionKeyName *string `json:"protection-key-name,omitempty"`
	// List of the tags attached to this classic key
	Tags []string `json:"tags,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateClassicKey instantiates a new CreateClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClassicKey(alg string, name string) *CreateClassicKey {
	this := CreateClassicKey{}
	this.Alg = alg
	this.Name = name
	return &this
}

// NewCreateClassicKeyWithDefaults instantiates a new CreateClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClassicKeyWithDefaults() *CreateClassicKey {
	this := CreateClassicKey{}
	return &this
}

// GetAlg returns the Alg field value
func (o *CreateClassicKey) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *CreateClassicKey) SetAlg(v string) {
	o.Alg = v
}

// GetCertFileData returns the CertFileData field value if set, zero value otherwise.
func (o *CreateClassicKey) GetCertFileData() string {
	if o == nil || o.CertFileData == nil {
		var ret string
		return ret
	}
	return *o.CertFileData
}

// GetCertFileDataOk returns a tuple with the CertFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetCertFileDataOk() (*string, bool) {
	if o == nil || o.CertFileData == nil {
		return nil, false
	}
	return o.CertFileData, true
}

// HasCertFileData returns a boolean if a field has been set.
func (o *CreateClassicKey) HasCertFileData() bool {
	if o != nil && o.CertFileData != nil {
		return true
	}

	return false
}

// SetCertFileData gets a reference to the given string and assigns it to the CertFileData field.
func (o *CreateClassicKey) SetCertFileData(v string) {
	o.CertFileData = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateClassicKey) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateClassicKey) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateClassicKey) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *CreateClassicKey) GetKeyData() string {
	if o == nil || o.KeyData == nil {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetKeyDataOk() (*string, bool) {
	if o == nil || o.KeyData == nil {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *CreateClassicKey) HasKeyData() bool {
	if o != nil && o.KeyData != nil {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *CreateClassicKey) SetKeyData(v string) {
	o.KeyData = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateClassicKey) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateClassicKey) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateClassicKey) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateClassicKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateClassicKey) SetName(v string) {
	o.Name = v
}

// GetProtectionKeyName returns the ProtectionKeyName field value if set, zero value otherwise.
func (o *CreateClassicKey) GetProtectionKeyName() string {
	if o == nil || o.ProtectionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKeyName
}

// GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetProtectionKeyNameOk() (*string, bool) {
	if o == nil || o.ProtectionKeyName == nil {
		return nil, false
	}
	return o.ProtectionKeyName, true
}

// HasProtectionKeyName returns a boolean if a field has been set.
func (o *CreateClassicKey) HasProtectionKeyName() bool {
	if o != nil && o.ProtectionKeyName != nil {
		return true
	}

	return false
}

// SetProtectionKeyName gets a reference to the given string and assigns it to the ProtectionKeyName field.
func (o *CreateClassicKey) SetProtectionKeyName(v string) {
	o.ProtectionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateClassicKey) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateClassicKey) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateClassicKey) SetTags(v []string) {
	o.Tags = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.CertFileData != nil {
		toSerialize["cert-file-data"] = o.CertFileData
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.KeyData != nil {
		toSerialize["key-data"] = o.KeyData
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProtectionKeyName != nil {
		toSerialize["protection-key-name"] = o.ProtectionKeyName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClassicKey struct {
	value *CreateClassicKey
	isSet bool
}

func (v NullableCreateClassicKey) Get() *CreateClassicKey {
	return v.value
}

func (v *NullableCreateClassicKey) Set(val *CreateClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClassicKey(val *CreateClassicKey) *NullableCreateClassicKey {
	return &NullableCreateClassicKey{value: val, isSet: true}
}

func (v NullableCreateClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


