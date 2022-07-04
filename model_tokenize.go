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

// Tokenize tokenize is a command that encrypts text with a tokenizer
type Tokenize struct {
	// Data to be encrypted
	Plaintext string `json:"plaintext"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The name of the tokenizer to use in the encryption process
	TokenizerName string `json:"tokenizer-name"`
	// Base64 encoded tweak for vaultless encryption
	Tweak *string `json:"tweak,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewTokenize instantiates a new Tokenize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenize(plaintext string, tokenizerName string) *Tokenize {
	this := Tokenize{}
	this.Plaintext = plaintext
	this.TokenizerName = tokenizerName
	return &this
}

// NewTokenizeWithDefaults instantiates a new Tokenize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizeWithDefaults() *Tokenize {
	this := Tokenize{}
	return &this
}

// GetPlaintext returns the Plaintext field value
func (o *Tokenize) GetPlaintext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value
// and a boolean to check if the value has been set.
func (o *Tokenize) GetPlaintextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plaintext, true
}

// SetPlaintext sets field value
func (o *Tokenize) SetPlaintext(v string) {
	o.Plaintext = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Tokenize) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tokenize) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Tokenize) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Tokenize) SetToken(v string) {
	o.Token = &v
}

// GetTokenizerName returns the TokenizerName field value
func (o *Tokenize) GetTokenizerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenizerName
}

// GetTokenizerNameOk returns a tuple with the TokenizerName field value
// and a boolean to check if the value has been set.
func (o *Tokenize) GetTokenizerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenizerName, true
}

// SetTokenizerName sets field value
func (o *Tokenize) SetTokenizerName(v string) {
	o.TokenizerName = v
}

// GetTweak returns the Tweak field value if set, zero value otherwise.
func (o *Tokenize) GetTweak() string {
	if o == nil || o.Tweak == nil {
		var ret string
		return ret
	}
	return *o.Tweak
}

// GetTweakOk returns a tuple with the Tweak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tokenize) GetTweakOk() (*string, bool) {
	if o == nil || o.Tweak == nil {
		return nil, false
	}
	return o.Tweak, true
}

// HasTweak returns a boolean if a field has been set.
func (o *Tokenize) HasTweak() bool {
	if o != nil && o.Tweak != nil {
		return true
	}

	return false
}

// SetTweak gets a reference to the given string and assigns it to the Tweak field.
func (o *Tokenize) SetTweak(v string) {
	o.Tweak = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Tokenize) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tokenize) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Tokenize) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Tokenize) SetUidToken(v string) {
	o.UidToken = &v
}

func (o Tokenize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plaintext"] = o.Plaintext
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["tokenizer-name"] = o.TokenizerName
	}
	if o.Tweak != nil {
		toSerialize["tweak"] = o.Tweak
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenize struct {
	value *Tokenize
	isSet bool
}

func (v NullableTokenize) Get() *Tokenize {
	return v.value
}

func (v *NullableTokenize) Set(val *Tokenize) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenize) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenize(val *Tokenize) *NullableTokenize {
	return &NullableTokenize{value: val, isSet: true}
}

func (v NullableTokenize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


