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
	"time"
)

// TmpUserData struct for TmpUserData
type TmpUserData struct {
	AccessId *string `json:"access_id,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	CustomTtl *int64 `json:"custom_ttl,omitempty"`
	DynamicSecretType *string `json:"dynamic_secret_type,omitempty"`
	Host *string `json:"host,omitempty"`
	Id *string `json:"id,omitempty"`
	SubClaims *map[string][]string `json:"sub_claims,omitempty"`
}

// NewTmpUserData instantiates a new TmpUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmpUserData() *TmpUserData {
	this := TmpUserData{}
	return &this
}

// NewTmpUserDataWithDefaults instantiates a new TmpUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmpUserDataWithDefaults() *TmpUserData {
	this := TmpUserData{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *TmpUserData) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *TmpUserData) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *TmpUserData) SetAccessId(v string) {
	o.AccessId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *TmpUserData) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *TmpUserData) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *TmpUserData) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetCustomTtl returns the CustomTtl field value if set, zero value otherwise.
func (o *TmpUserData) GetCustomTtl() int64 {
	if o == nil || o.CustomTtl == nil {
		var ret int64
		return ret
	}
	return *o.CustomTtl
}

// GetCustomTtlOk returns a tuple with the CustomTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetCustomTtlOk() (*int64, bool) {
	if o == nil || o.CustomTtl == nil {
		return nil, false
	}
	return o.CustomTtl, true
}

// HasCustomTtl returns a boolean if a field has been set.
func (o *TmpUserData) HasCustomTtl() bool {
	if o != nil && o.CustomTtl != nil {
		return true
	}

	return false
}

// SetCustomTtl gets a reference to the given int64 and assigns it to the CustomTtl field.
func (o *TmpUserData) SetCustomTtl(v int64) {
	o.CustomTtl = &v
}

// GetDynamicSecretType returns the DynamicSecretType field value if set, zero value otherwise.
func (o *TmpUserData) GetDynamicSecretType() string {
	if o == nil || o.DynamicSecretType == nil {
		var ret string
		return ret
	}
	return *o.DynamicSecretType
}

// GetDynamicSecretTypeOk returns a tuple with the DynamicSecretType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetDynamicSecretTypeOk() (*string, bool) {
	if o == nil || o.DynamicSecretType == nil {
		return nil, false
	}
	return o.DynamicSecretType, true
}

// HasDynamicSecretType returns a boolean if a field has been set.
func (o *TmpUserData) HasDynamicSecretType() bool {
	if o != nil && o.DynamicSecretType != nil {
		return true
	}

	return false
}

// SetDynamicSecretType gets a reference to the given string and assigns it to the DynamicSecretType field.
func (o *TmpUserData) SetDynamicSecretType(v string) {
	o.DynamicSecretType = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *TmpUserData) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *TmpUserData) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *TmpUserData) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TmpUserData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TmpUserData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TmpUserData) SetId(v string) {
	o.Id = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *TmpUserData) GetSubClaims() map[string][]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string][]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmpUserData) GetSubClaimsOk() (*map[string][]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *TmpUserData) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string][]string and assigns it to the SubClaims field.
func (o *TmpUserData) SetSubClaims(v map[string][]string) {
	o.SubClaims = &v
}

func (o TmpUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.CustomTtl != nil {
		toSerialize["custom_ttl"] = o.CustomTtl
	}
	if o.DynamicSecretType != nil {
		toSerialize["dynamic_secret_type"] = o.DynamicSecretType
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SubClaims != nil {
		toSerialize["sub_claims"] = o.SubClaims
	}
	return json.Marshal(toSerialize)
}

type NullableTmpUserData struct {
	value *TmpUserData
	isSet bool
}

func (v NullableTmpUserData) Get() *TmpUserData {
	return v.value
}

func (v *NullableTmpUserData) Set(val *TmpUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableTmpUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableTmpUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmpUserData(val *TmpUserData) *NullableTmpUserData {
	return &NullableTmpUserData{value: val, isSet: true}
}

func (v NullableTmpUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmpUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


