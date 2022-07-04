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

// AllowedAccess struct for AllowedAccess
type AllowedAccess struct {
	AccId *string `json:"acc_id,omitempty"`
	AccessRulesType *string `json:"access_rules_type,omitempty"`
	AllowedApi *bool `json:"allowed_api,omitempty"`
	AllowedsLogin *bool `json:"alloweds_login,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty"`
	Hash *string `json:"hash,omitempty"`
	IsValid *bool `json:"is_valid,omitempty"`
	Name *string `json:"name,omitempty"`
	SubClaims *map[string][]string `json:"sub_claims,omitempty"`
}

// NewAllowedAccess instantiates a new AllowedAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedAccess() *AllowedAccess {
	this := AllowedAccess{}
	return &this
}

// NewAllowedAccessWithDefaults instantiates a new AllowedAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedAccessWithDefaults() *AllowedAccess {
	this := AllowedAccess{}
	return &this
}

// GetAccId returns the AccId field value if set, zero value otherwise.
func (o *AllowedAccess) GetAccId() string {
	if o == nil || o.AccId == nil {
		var ret string
		return ret
	}
	return *o.AccId
}

// GetAccIdOk returns a tuple with the AccId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetAccIdOk() (*string, bool) {
	if o == nil || o.AccId == nil {
		return nil, false
	}
	return o.AccId, true
}

// HasAccId returns a boolean if a field has been set.
func (o *AllowedAccess) HasAccId() bool {
	if o != nil && o.AccId != nil {
		return true
	}

	return false
}

// SetAccId gets a reference to the given string and assigns it to the AccId field.
func (o *AllowedAccess) SetAccId(v string) {
	o.AccId = &v
}

// GetAccessRulesType returns the AccessRulesType field value if set, zero value otherwise.
func (o *AllowedAccess) GetAccessRulesType() string {
	if o == nil || o.AccessRulesType == nil {
		var ret string
		return ret
	}
	return *o.AccessRulesType
}

// GetAccessRulesTypeOk returns a tuple with the AccessRulesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetAccessRulesTypeOk() (*string, bool) {
	if o == nil || o.AccessRulesType == nil {
		return nil, false
	}
	return o.AccessRulesType, true
}

// HasAccessRulesType returns a boolean if a field has been set.
func (o *AllowedAccess) HasAccessRulesType() bool {
	if o != nil && o.AccessRulesType != nil {
		return true
	}

	return false
}

// SetAccessRulesType gets a reference to the given string and assigns it to the AccessRulesType field.
func (o *AllowedAccess) SetAccessRulesType(v string) {
	o.AccessRulesType = &v
}

// GetAllowedApi returns the AllowedApi field value if set, zero value otherwise.
func (o *AllowedAccess) GetAllowedApi() bool {
	if o == nil || o.AllowedApi == nil {
		var ret bool
		return ret
	}
	return *o.AllowedApi
}

// GetAllowedApiOk returns a tuple with the AllowedApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetAllowedApiOk() (*bool, bool) {
	if o == nil || o.AllowedApi == nil {
		return nil, false
	}
	return o.AllowedApi, true
}

// HasAllowedApi returns a boolean if a field has been set.
func (o *AllowedAccess) HasAllowedApi() bool {
	if o != nil && o.AllowedApi != nil {
		return true
	}

	return false
}

// SetAllowedApi gets a reference to the given bool and assigns it to the AllowedApi field.
func (o *AllowedAccess) SetAllowedApi(v bool) {
	o.AllowedApi = &v
}

// GetAllowedsLogin returns the AllowedsLogin field value if set, zero value otherwise.
func (o *AllowedAccess) GetAllowedsLogin() bool {
	if o == nil || o.AllowedsLogin == nil {
		var ret bool
		return ret
	}
	return *o.AllowedsLogin
}

// GetAllowedsLoginOk returns a tuple with the AllowedsLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetAllowedsLoginOk() (*bool, bool) {
	if o == nil || o.AllowedsLogin == nil {
		return nil, false
	}
	return o.AllowedsLogin, true
}

// HasAllowedsLogin returns a boolean if a field has been set.
func (o *AllowedAccess) HasAllowedsLogin() bool {
	if o != nil && o.AllowedsLogin != nil {
		return true
	}

	return false
}

// SetAllowedsLogin gets a reference to the given bool and assigns it to the AllowedsLogin field.
func (o *AllowedAccess) SetAllowedsLogin(v bool) {
	o.AllowedsLogin = &v
}

// GetErrMsg returns the ErrMsg field value if set, zero value otherwise.
func (o *AllowedAccess) GetErrMsg() string {
	if o == nil || o.ErrMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrMsg
}

// GetErrMsgOk returns a tuple with the ErrMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetErrMsgOk() (*string, bool) {
	if o == nil || o.ErrMsg == nil {
		return nil, false
	}
	return o.ErrMsg, true
}

// HasErrMsg returns a boolean if a field has been set.
func (o *AllowedAccess) HasErrMsg() bool {
	if o != nil && o.ErrMsg != nil {
		return true
	}

	return false
}

// SetErrMsg gets a reference to the given string and assigns it to the ErrMsg field.
func (o *AllowedAccess) SetErrMsg(v string) {
	o.ErrMsg = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *AllowedAccess) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *AllowedAccess) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *AllowedAccess) SetHash(v string) {
	o.Hash = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *AllowedAccess) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *AllowedAccess) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *AllowedAccess) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllowedAccess) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllowedAccess) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllowedAccess) SetName(v string) {
	o.Name = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *AllowedAccess) GetSubClaims() map[string][]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string][]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedAccess) GetSubClaimsOk() (*map[string][]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *AllowedAccess) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string][]string and assigns it to the SubClaims field.
func (o *AllowedAccess) SetSubClaims(v map[string][]string) {
	o.SubClaims = &v
}

func (o AllowedAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccId != nil {
		toSerialize["acc_id"] = o.AccId
	}
	if o.AccessRulesType != nil {
		toSerialize["access_rules_type"] = o.AccessRulesType
	}
	if o.AllowedApi != nil {
		toSerialize["allowed_api"] = o.AllowedApi
	}
	if o.AllowedsLogin != nil {
		toSerialize["alloweds_login"] = o.AllowedsLogin
	}
	if o.ErrMsg != nil {
		toSerialize["err_msg"] = o.ErrMsg
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.IsValid != nil {
		toSerialize["is_valid"] = o.IsValid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SubClaims != nil {
		toSerialize["sub_claims"] = o.SubClaims
	}
	return json.Marshal(toSerialize)
}

type NullableAllowedAccess struct {
	value *AllowedAccess
	isSet bool
}

func (v NullableAllowedAccess) Get() *AllowedAccess {
	return v.value
}

func (v *NullableAllowedAccess) Set(val *AllowedAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedAccess(val *AllowedAccess) *NullableAllowedAccess {
	return &NullableAllowedAccess{value: val, isSet: true}
}

func (v NullableAllowedAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


