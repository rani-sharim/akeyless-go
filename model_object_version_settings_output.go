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

// ObjectVersionSettingsOutput struct for ObjectVersionSettingsOutput
type ObjectVersionSettingsOutput struct {
	// VersionSettingsObjectType defines object types for account version settings
	ItemType *string `json:"item-type,omitempty"`
	MaxVersions *string `json:"max-versions,omitempty"`
}

// NewObjectVersionSettingsOutput instantiates a new ObjectVersionSettingsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectVersionSettingsOutput() *ObjectVersionSettingsOutput {
	this := ObjectVersionSettingsOutput{}
	return &this
}

// NewObjectVersionSettingsOutputWithDefaults instantiates a new ObjectVersionSettingsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectVersionSettingsOutputWithDefaults() *ObjectVersionSettingsOutput {
	this := ObjectVersionSettingsOutput{}
	return &this
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *ObjectVersionSettingsOutput) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectVersionSettingsOutput) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *ObjectVersionSettingsOutput) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *ObjectVersionSettingsOutput) SetItemType(v string) {
	o.ItemType = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *ObjectVersionSettingsOutput) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectVersionSettingsOutput) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *ObjectVersionSettingsOutput) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *ObjectVersionSettingsOutput) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

func (o ObjectVersionSettingsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemType != nil {
		toSerialize["item-type"] = o.ItemType
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	return json.Marshal(toSerialize)
}

type NullableObjectVersionSettingsOutput struct {
	value *ObjectVersionSettingsOutput
	isSet bool
}

func (v NullableObjectVersionSettingsOutput) Get() *ObjectVersionSettingsOutput {
	return v.value
}

func (v *NullableObjectVersionSettingsOutput) Set(val *ObjectVersionSettingsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectVersionSettingsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectVersionSettingsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectVersionSettingsOutput(val *ObjectVersionSettingsOutput) *NullableObjectVersionSettingsOutput {
	return &NullableObjectVersionSettingsOutput{value: val, isSet: true}
}

func (v NullableObjectVersionSettingsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectVersionSettingsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

