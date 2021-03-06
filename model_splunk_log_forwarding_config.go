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

// SplunkLogForwardingConfig struct for SplunkLogForwardingConfig
type SplunkLogForwardingConfig struct {
	SplunkIndex *string `json:"splunk_index,omitempty"`
	SplunkSource *string `json:"splunk_source,omitempty"`
	SplunkSourcetype *string `json:"splunk_sourcetype,omitempty"`
	SplunkToken *string `json:"splunk_token,omitempty"`
	SplunkUrl *string `json:"splunk_url,omitempty"`
}

// NewSplunkLogForwardingConfig instantiates a new SplunkLogForwardingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplunkLogForwardingConfig() *SplunkLogForwardingConfig {
	this := SplunkLogForwardingConfig{}
	return &this
}

// NewSplunkLogForwardingConfigWithDefaults instantiates a new SplunkLogForwardingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplunkLogForwardingConfigWithDefaults() *SplunkLogForwardingConfig {
	this := SplunkLogForwardingConfig{}
	return &this
}

// GetSplunkIndex returns the SplunkIndex field value if set, zero value otherwise.
func (o *SplunkLogForwardingConfig) GetSplunkIndex() string {
	if o == nil || o.SplunkIndex == nil {
		var ret string
		return ret
	}
	return *o.SplunkIndex
}

// GetSplunkIndexOk returns a tuple with the SplunkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkLogForwardingConfig) GetSplunkIndexOk() (*string, bool) {
	if o == nil || o.SplunkIndex == nil {
		return nil, false
	}
	return o.SplunkIndex, true
}

// HasSplunkIndex returns a boolean if a field has been set.
func (o *SplunkLogForwardingConfig) HasSplunkIndex() bool {
	if o != nil && o.SplunkIndex != nil {
		return true
	}

	return false
}

// SetSplunkIndex gets a reference to the given string and assigns it to the SplunkIndex field.
func (o *SplunkLogForwardingConfig) SetSplunkIndex(v string) {
	o.SplunkIndex = &v
}

// GetSplunkSource returns the SplunkSource field value if set, zero value otherwise.
func (o *SplunkLogForwardingConfig) GetSplunkSource() string {
	if o == nil || o.SplunkSource == nil {
		var ret string
		return ret
	}
	return *o.SplunkSource
}

// GetSplunkSourceOk returns a tuple with the SplunkSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkLogForwardingConfig) GetSplunkSourceOk() (*string, bool) {
	if o == nil || o.SplunkSource == nil {
		return nil, false
	}
	return o.SplunkSource, true
}

// HasSplunkSource returns a boolean if a field has been set.
func (o *SplunkLogForwardingConfig) HasSplunkSource() bool {
	if o != nil && o.SplunkSource != nil {
		return true
	}

	return false
}

// SetSplunkSource gets a reference to the given string and assigns it to the SplunkSource field.
func (o *SplunkLogForwardingConfig) SetSplunkSource(v string) {
	o.SplunkSource = &v
}

// GetSplunkSourcetype returns the SplunkSourcetype field value if set, zero value otherwise.
func (o *SplunkLogForwardingConfig) GetSplunkSourcetype() string {
	if o == nil || o.SplunkSourcetype == nil {
		var ret string
		return ret
	}
	return *o.SplunkSourcetype
}

// GetSplunkSourcetypeOk returns a tuple with the SplunkSourcetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkLogForwardingConfig) GetSplunkSourcetypeOk() (*string, bool) {
	if o == nil || o.SplunkSourcetype == nil {
		return nil, false
	}
	return o.SplunkSourcetype, true
}

// HasSplunkSourcetype returns a boolean if a field has been set.
func (o *SplunkLogForwardingConfig) HasSplunkSourcetype() bool {
	if o != nil && o.SplunkSourcetype != nil {
		return true
	}

	return false
}

// SetSplunkSourcetype gets a reference to the given string and assigns it to the SplunkSourcetype field.
func (o *SplunkLogForwardingConfig) SetSplunkSourcetype(v string) {
	o.SplunkSourcetype = &v
}

// GetSplunkToken returns the SplunkToken field value if set, zero value otherwise.
func (o *SplunkLogForwardingConfig) GetSplunkToken() string {
	if o == nil || o.SplunkToken == nil {
		var ret string
		return ret
	}
	return *o.SplunkToken
}

// GetSplunkTokenOk returns a tuple with the SplunkToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkLogForwardingConfig) GetSplunkTokenOk() (*string, bool) {
	if o == nil || o.SplunkToken == nil {
		return nil, false
	}
	return o.SplunkToken, true
}

// HasSplunkToken returns a boolean if a field has been set.
func (o *SplunkLogForwardingConfig) HasSplunkToken() bool {
	if o != nil && o.SplunkToken != nil {
		return true
	}

	return false
}

// SetSplunkToken gets a reference to the given string and assigns it to the SplunkToken field.
func (o *SplunkLogForwardingConfig) SetSplunkToken(v string) {
	o.SplunkToken = &v
}

// GetSplunkUrl returns the SplunkUrl field value if set, zero value otherwise.
func (o *SplunkLogForwardingConfig) GetSplunkUrl() string {
	if o == nil || o.SplunkUrl == nil {
		var ret string
		return ret
	}
	return *o.SplunkUrl
}

// GetSplunkUrlOk returns a tuple with the SplunkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplunkLogForwardingConfig) GetSplunkUrlOk() (*string, bool) {
	if o == nil || o.SplunkUrl == nil {
		return nil, false
	}
	return o.SplunkUrl, true
}

// HasSplunkUrl returns a boolean if a field has been set.
func (o *SplunkLogForwardingConfig) HasSplunkUrl() bool {
	if o != nil && o.SplunkUrl != nil {
		return true
	}

	return false
}

// SetSplunkUrl gets a reference to the given string and assigns it to the SplunkUrl field.
func (o *SplunkLogForwardingConfig) SetSplunkUrl(v string) {
	o.SplunkUrl = &v
}

func (o SplunkLogForwardingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SplunkIndex != nil {
		toSerialize["splunk_index"] = o.SplunkIndex
	}
	if o.SplunkSource != nil {
		toSerialize["splunk_source"] = o.SplunkSource
	}
	if o.SplunkSourcetype != nil {
		toSerialize["splunk_sourcetype"] = o.SplunkSourcetype
	}
	if o.SplunkToken != nil {
		toSerialize["splunk_token"] = o.SplunkToken
	}
	if o.SplunkUrl != nil {
		toSerialize["splunk_url"] = o.SplunkUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSplunkLogForwardingConfig struct {
	value *SplunkLogForwardingConfig
	isSet bool
}

func (v NullableSplunkLogForwardingConfig) Get() *SplunkLogForwardingConfig {
	return v.value
}

func (v *NullableSplunkLogForwardingConfig) Set(val *SplunkLogForwardingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSplunkLogForwardingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSplunkLogForwardingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplunkLogForwardingConfig(val *SplunkLogForwardingConfig) *NullableSplunkLogForwardingConfig {
	return &NullableSplunkLogForwardingConfig{value: val, isSet: true}
}

func (v NullableSplunkLogForwardingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplunkLogForwardingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


