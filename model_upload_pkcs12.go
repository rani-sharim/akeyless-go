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

// UploadPKCS12 struct for UploadPKCS12
type UploadPKCS12 struct {
	// The customer fragment ID that will be used to split the key (if empty, the key will be created independently of a customer fragment)
	CustomerFrgId *string `json:"customer-frg-id,omitempty"`
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// PKCS#12 input file (private key and certificate only)
	In string `json:"in"`
	// A metadata about the key
	Metadata *string `json:"metadata,omitempty"`
	// Name of key to be created
	Name string `json:"name"`
	// Passphrase to unlock the pkcs#12 bundle
	Passphrase string `json:"passphrase"`
	// The number of fragments that the item will be split into
	SplitLevel *int64 `json:"split-level,omitempty"`
	// List of the tags attached to this key
	Tag []string `json:"tag,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUploadPKCS12 instantiates a new UploadPKCS12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadPKCS12(in string, name string, passphrase string) *UploadPKCS12 {
	this := UploadPKCS12{}
	this.In = in
	this.Name = name
	this.Passphrase = passphrase
	var splitLevel int64 = 2
	this.SplitLevel = &splitLevel
	return &this
}

// NewUploadPKCS12WithDefaults instantiates a new UploadPKCS12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadPKCS12WithDefaults() *UploadPKCS12 {
	this := UploadPKCS12{}
	var splitLevel int64 = 2
	this.SplitLevel = &splitLevel
	return &this
}

// GetCustomerFrgId returns the CustomerFrgId field value if set, zero value otherwise.
func (o *UploadPKCS12) GetCustomerFrgId() string {
	if o == nil || o.CustomerFrgId == nil {
		var ret string
		return ret
	}
	return *o.CustomerFrgId
}

// GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetCustomerFrgIdOk() (*string, bool) {
	if o == nil || o.CustomerFrgId == nil {
		return nil, false
	}
	return o.CustomerFrgId, true
}

// HasCustomerFrgId returns a boolean if a field has been set.
func (o *UploadPKCS12) HasCustomerFrgId() bool {
	if o != nil && o.CustomerFrgId != nil {
		return true
	}

	return false
}

// SetCustomerFrgId gets a reference to the given string and assigns it to the CustomerFrgId field.
func (o *UploadPKCS12) SetCustomerFrgId(v string) {
	o.CustomerFrgId = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *UploadPKCS12) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *UploadPKCS12) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *UploadPKCS12) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetIn returns the In field value
func (o *UploadPKCS12) GetIn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.In
}

// GetInOk returns a tuple with the In field value
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetInOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.In, true
}

// SetIn sets field value
func (o *UploadPKCS12) SetIn(v string) {
	o.In = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UploadPKCS12) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UploadPKCS12) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *UploadPKCS12) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *UploadPKCS12) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UploadPKCS12) SetName(v string) {
	o.Name = v
}

// GetPassphrase returns the Passphrase field value
func (o *UploadPKCS12) GetPassphrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetPassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *UploadPKCS12) SetPassphrase(v string) {
	o.Passphrase = v
}

// GetSplitLevel returns the SplitLevel field value if set, zero value otherwise.
func (o *UploadPKCS12) GetSplitLevel() int64 {
	if o == nil || o.SplitLevel == nil {
		var ret int64
		return ret
	}
	return *o.SplitLevel
}

// GetSplitLevelOk returns a tuple with the SplitLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetSplitLevelOk() (*int64, bool) {
	if o == nil || o.SplitLevel == nil {
		return nil, false
	}
	return o.SplitLevel, true
}

// HasSplitLevel returns a boolean if a field has been set.
func (o *UploadPKCS12) HasSplitLevel() bool {
	if o != nil && o.SplitLevel != nil {
		return true
	}

	return false
}

// SetSplitLevel gets a reference to the given int64 and assigns it to the SplitLevel field.
func (o *UploadPKCS12) SetSplitLevel(v int64) {
	o.SplitLevel = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *UploadPKCS12) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetTagOk() ([]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *UploadPKCS12) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *UploadPKCS12) SetTag(v []string) {
	o.Tag = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UploadPKCS12) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UploadPKCS12) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UploadPKCS12) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UploadPKCS12) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPKCS12) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UploadPKCS12) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UploadPKCS12) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UploadPKCS12) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerFrgId != nil {
		toSerialize["customer-frg-id"] = o.CustomerFrgId
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if true {
		toSerialize["in"] = o.In
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.SplitLevel != nil {
		toSerialize["split-level"] = o.SplitLevel
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUploadPKCS12 struct {
	value *UploadPKCS12
	isSet bool
}

func (v NullableUploadPKCS12) Get() *UploadPKCS12 {
	return v.value
}

func (v *NullableUploadPKCS12) Set(val *UploadPKCS12) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadPKCS12) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadPKCS12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadPKCS12(val *UploadPKCS12) *NullableUploadPKCS12 {
	return &NullableUploadPKCS12{value: val, isSet: true}
}

func (v NullableUploadPKCS12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadPKCS12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

