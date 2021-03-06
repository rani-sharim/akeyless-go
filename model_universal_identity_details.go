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

// UniversalIdentityDetails struct for UniversalIdentityDetails
type UniversalIdentityDetails struct {
	MaxDepth *int32 `json:"max_depth,omitempty"`
	NumberOfTokens *int64 `json:"number_of_tokens,omitempty"`
	Root *UIDTokenDetails `json:"root,omitempty"`
}

// NewUniversalIdentityDetails instantiates a new UniversalIdentityDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniversalIdentityDetails() *UniversalIdentityDetails {
	this := UniversalIdentityDetails{}
	return &this
}

// NewUniversalIdentityDetailsWithDefaults instantiates a new UniversalIdentityDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniversalIdentityDetailsWithDefaults() *UniversalIdentityDetails {
	this := UniversalIdentityDetails{}
	return &this
}

// GetMaxDepth returns the MaxDepth field value if set, zero value otherwise.
func (o *UniversalIdentityDetails) GetMaxDepth() int32 {
	if o == nil || o.MaxDepth == nil {
		var ret int32
		return ret
	}
	return *o.MaxDepth
}

// GetMaxDepthOk returns a tuple with the MaxDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalIdentityDetails) GetMaxDepthOk() (*int32, bool) {
	if o == nil || o.MaxDepth == nil {
		return nil, false
	}
	return o.MaxDepth, true
}

// HasMaxDepth returns a boolean if a field has been set.
func (o *UniversalIdentityDetails) HasMaxDepth() bool {
	if o != nil && o.MaxDepth != nil {
		return true
	}

	return false
}

// SetMaxDepth gets a reference to the given int32 and assigns it to the MaxDepth field.
func (o *UniversalIdentityDetails) SetMaxDepth(v int32) {
	o.MaxDepth = &v
}

// GetNumberOfTokens returns the NumberOfTokens field value if set, zero value otherwise.
func (o *UniversalIdentityDetails) GetNumberOfTokens() int64 {
	if o == nil || o.NumberOfTokens == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfTokens
}

// GetNumberOfTokensOk returns a tuple with the NumberOfTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalIdentityDetails) GetNumberOfTokensOk() (*int64, bool) {
	if o == nil || o.NumberOfTokens == nil {
		return nil, false
	}
	return o.NumberOfTokens, true
}

// HasNumberOfTokens returns a boolean if a field has been set.
func (o *UniversalIdentityDetails) HasNumberOfTokens() bool {
	if o != nil && o.NumberOfTokens != nil {
		return true
	}

	return false
}

// SetNumberOfTokens gets a reference to the given int64 and assigns it to the NumberOfTokens field.
func (o *UniversalIdentityDetails) SetNumberOfTokens(v int64) {
	o.NumberOfTokens = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *UniversalIdentityDetails) GetRoot() UIDTokenDetails {
	if o == nil || o.Root == nil {
		var ret UIDTokenDetails
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalIdentityDetails) GetRootOk() (*UIDTokenDetails, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *UniversalIdentityDetails) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given UIDTokenDetails and assigns it to the Root field.
func (o *UniversalIdentityDetails) SetRoot(v UIDTokenDetails) {
	o.Root = &v
}

func (o UniversalIdentityDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxDepth != nil {
		toSerialize["max_depth"] = o.MaxDepth
	}
	if o.NumberOfTokens != nil {
		toSerialize["number_of_tokens"] = o.NumberOfTokens
	}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	return json.Marshal(toSerialize)
}

type NullableUniversalIdentityDetails struct {
	value *UniversalIdentityDetails
	isSet bool
}

func (v NullableUniversalIdentityDetails) Get() *UniversalIdentityDetails {
	return v.value
}

func (v *NullableUniversalIdentityDetails) Set(val *UniversalIdentityDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUniversalIdentityDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUniversalIdentityDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniversalIdentityDetails(val *UniversalIdentityDetails) *NullableUniversalIdentityDetails {
	return &NullableUniversalIdentityDetails{value: val, isSet: true}
}

func (v NullableUniversalIdentityDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniversalIdentityDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


