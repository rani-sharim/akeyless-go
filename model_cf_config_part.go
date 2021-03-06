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

// CFConfigPart struct for CFConfigPart
type CFConfigPart struct {
	CustomerFragements *map[string]string `json:"customer_fragements,omitempty"`
}

// NewCFConfigPart instantiates a new CFConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFConfigPart() *CFConfigPart {
	this := CFConfigPart{}
	return &this
}

// NewCFConfigPartWithDefaults instantiates a new CFConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFConfigPartWithDefaults() *CFConfigPart {
	this := CFConfigPart{}
	return &this
}

// GetCustomerFragements returns the CustomerFragements field value if set, zero value otherwise.
func (o *CFConfigPart) GetCustomerFragements() map[string]string {
	if o == nil || o.CustomerFragements == nil {
		var ret map[string]string
		return ret
	}
	return *o.CustomerFragements
}

// GetCustomerFragementsOk returns a tuple with the CustomerFragements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFConfigPart) GetCustomerFragementsOk() (*map[string]string, bool) {
	if o == nil || o.CustomerFragements == nil {
		return nil, false
	}
	return o.CustomerFragements, true
}

// HasCustomerFragements returns a boolean if a field has been set.
func (o *CFConfigPart) HasCustomerFragements() bool {
	if o != nil && o.CustomerFragements != nil {
		return true
	}

	return false
}

// SetCustomerFragements gets a reference to the given map[string]string and assigns it to the CustomerFragements field.
func (o *CFConfigPart) SetCustomerFragements(v map[string]string) {
	o.CustomerFragements = &v
}

func (o CFConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerFragements != nil {
		toSerialize["customer_fragements"] = o.CustomerFragements
	}
	return json.Marshal(toSerialize)
}

type NullableCFConfigPart struct {
	value *CFConfigPart
	isSet bool
}

func (v NullableCFConfigPart) Get() *CFConfigPart {
	return v.value
}

func (v *NullableCFConfigPart) Set(val *CFConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableCFConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableCFConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFConfigPart(val *CFConfigPart) *NullableCFConfigPart {
	return &NullableCFConfigPart{value: val, isSet: true}
}

func (v NullableCFConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


