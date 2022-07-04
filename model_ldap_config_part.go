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

// LdapConfigPart struct for LdapConfigPart
type LdapConfigPart struct {
	LdapAccessId *string `json:"ldap_access_id,omitempty"`
	LdapAnonymousSearch *bool `json:"ldap_anonymous_search,omitempty"`
	LdapBindDn *string `json:"ldap_bind_dn,omitempty"`
	LdapBindPassword *string `json:"ldap_bind_password,omitempty"`
	LdapCert *string `json:"ldap_cert,omitempty"`
	LdapEnable *bool `json:"ldap_enable,omitempty"`
	LdapGroupAttr *string `json:"ldap_group_attr,omitempty"`
	LdapGroupDn *string `json:"ldap_group_dn,omitempty"`
	LdapGroupFilter *string `json:"ldap_group_filter,omitempty"`
	LdapPrivateKey *string `json:"ldap_private_key,omitempty"`
	LdapTokenExpiration *string `json:"ldap_token_expiration,omitempty"`
	LdapUrl *string `json:"ldap_url,omitempty"`
	LdapUserAttr *string `json:"ldap_user_attr,omitempty"`
	LdapUserDn *string `json:"ldap_user_dn,omitempty"`
}

// NewLdapConfigPart instantiates a new LdapConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapConfigPart() *LdapConfigPart {
	this := LdapConfigPart{}
	return &this
}

// NewLdapConfigPartWithDefaults instantiates a new LdapConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapConfigPartWithDefaults() *LdapConfigPart {
	this := LdapConfigPart{}
	return &this
}

// GetLdapAccessId returns the LdapAccessId field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapAccessId() string {
	if o == nil || o.LdapAccessId == nil {
		var ret string
		return ret
	}
	return *o.LdapAccessId
}

// GetLdapAccessIdOk returns a tuple with the LdapAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapAccessIdOk() (*string, bool) {
	if o == nil || o.LdapAccessId == nil {
		return nil, false
	}
	return o.LdapAccessId, true
}

// HasLdapAccessId returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapAccessId() bool {
	if o != nil && o.LdapAccessId != nil {
		return true
	}

	return false
}

// SetLdapAccessId gets a reference to the given string and assigns it to the LdapAccessId field.
func (o *LdapConfigPart) SetLdapAccessId(v string) {
	o.LdapAccessId = &v
}

// GetLdapAnonymousSearch returns the LdapAnonymousSearch field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapAnonymousSearch() bool {
	if o == nil || o.LdapAnonymousSearch == nil {
		var ret bool
		return ret
	}
	return *o.LdapAnonymousSearch
}

// GetLdapAnonymousSearchOk returns a tuple with the LdapAnonymousSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapAnonymousSearchOk() (*bool, bool) {
	if o == nil || o.LdapAnonymousSearch == nil {
		return nil, false
	}
	return o.LdapAnonymousSearch, true
}

// HasLdapAnonymousSearch returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapAnonymousSearch() bool {
	if o != nil && o.LdapAnonymousSearch != nil {
		return true
	}

	return false
}

// SetLdapAnonymousSearch gets a reference to the given bool and assigns it to the LdapAnonymousSearch field.
func (o *LdapConfigPart) SetLdapAnonymousSearch(v bool) {
	o.LdapAnonymousSearch = &v
}

// GetLdapBindDn returns the LdapBindDn field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapBindDn() string {
	if o == nil || o.LdapBindDn == nil {
		var ret string
		return ret
	}
	return *o.LdapBindDn
}

// GetLdapBindDnOk returns a tuple with the LdapBindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapBindDnOk() (*string, bool) {
	if o == nil || o.LdapBindDn == nil {
		return nil, false
	}
	return o.LdapBindDn, true
}

// HasLdapBindDn returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapBindDn() bool {
	if o != nil && o.LdapBindDn != nil {
		return true
	}

	return false
}

// SetLdapBindDn gets a reference to the given string and assigns it to the LdapBindDn field.
func (o *LdapConfigPart) SetLdapBindDn(v string) {
	o.LdapBindDn = &v
}

// GetLdapBindPassword returns the LdapBindPassword field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapBindPassword() string {
	if o == nil || o.LdapBindPassword == nil {
		var ret string
		return ret
	}
	return *o.LdapBindPassword
}

// GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapBindPasswordOk() (*string, bool) {
	if o == nil || o.LdapBindPassword == nil {
		return nil, false
	}
	return o.LdapBindPassword, true
}

// HasLdapBindPassword returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapBindPassword() bool {
	if o != nil && o.LdapBindPassword != nil {
		return true
	}

	return false
}

// SetLdapBindPassword gets a reference to the given string and assigns it to the LdapBindPassword field.
func (o *LdapConfigPart) SetLdapBindPassword(v string) {
	o.LdapBindPassword = &v
}

// GetLdapCert returns the LdapCert field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapCert() string {
	if o == nil || o.LdapCert == nil {
		var ret string
		return ret
	}
	return *o.LdapCert
}

// GetLdapCertOk returns a tuple with the LdapCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapCertOk() (*string, bool) {
	if o == nil || o.LdapCert == nil {
		return nil, false
	}
	return o.LdapCert, true
}

// HasLdapCert returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapCert() bool {
	if o != nil && o.LdapCert != nil {
		return true
	}

	return false
}

// SetLdapCert gets a reference to the given string and assigns it to the LdapCert field.
func (o *LdapConfigPart) SetLdapCert(v string) {
	o.LdapCert = &v
}

// GetLdapEnable returns the LdapEnable field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapEnable() bool {
	if o == nil || o.LdapEnable == nil {
		var ret bool
		return ret
	}
	return *o.LdapEnable
}

// GetLdapEnableOk returns a tuple with the LdapEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapEnableOk() (*bool, bool) {
	if o == nil || o.LdapEnable == nil {
		return nil, false
	}
	return o.LdapEnable, true
}

// HasLdapEnable returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapEnable() bool {
	if o != nil && o.LdapEnable != nil {
		return true
	}

	return false
}

// SetLdapEnable gets a reference to the given bool and assigns it to the LdapEnable field.
func (o *LdapConfigPart) SetLdapEnable(v bool) {
	o.LdapEnable = &v
}

// GetLdapGroupAttr returns the LdapGroupAttr field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapGroupAttr() string {
	if o == nil || o.LdapGroupAttr == nil {
		var ret string
		return ret
	}
	return *o.LdapGroupAttr
}

// GetLdapGroupAttrOk returns a tuple with the LdapGroupAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapGroupAttrOk() (*string, bool) {
	if o == nil || o.LdapGroupAttr == nil {
		return nil, false
	}
	return o.LdapGroupAttr, true
}

// HasLdapGroupAttr returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapGroupAttr() bool {
	if o != nil && o.LdapGroupAttr != nil {
		return true
	}

	return false
}

// SetLdapGroupAttr gets a reference to the given string and assigns it to the LdapGroupAttr field.
func (o *LdapConfigPart) SetLdapGroupAttr(v string) {
	o.LdapGroupAttr = &v
}

// GetLdapGroupDn returns the LdapGroupDn field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapGroupDn() string {
	if o == nil || o.LdapGroupDn == nil {
		var ret string
		return ret
	}
	return *o.LdapGroupDn
}

// GetLdapGroupDnOk returns a tuple with the LdapGroupDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapGroupDnOk() (*string, bool) {
	if o == nil || o.LdapGroupDn == nil {
		return nil, false
	}
	return o.LdapGroupDn, true
}

// HasLdapGroupDn returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapGroupDn() bool {
	if o != nil && o.LdapGroupDn != nil {
		return true
	}

	return false
}

// SetLdapGroupDn gets a reference to the given string and assigns it to the LdapGroupDn field.
func (o *LdapConfigPart) SetLdapGroupDn(v string) {
	o.LdapGroupDn = &v
}

// GetLdapGroupFilter returns the LdapGroupFilter field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapGroupFilter() string {
	if o == nil || o.LdapGroupFilter == nil {
		var ret string
		return ret
	}
	return *o.LdapGroupFilter
}

// GetLdapGroupFilterOk returns a tuple with the LdapGroupFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapGroupFilterOk() (*string, bool) {
	if o == nil || o.LdapGroupFilter == nil {
		return nil, false
	}
	return o.LdapGroupFilter, true
}

// HasLdapGroupFilter returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapGroupFilter() bool {
	if o != nil && o.LdapGroupFilter != nil {
		return true
	}

	return false
}

// SetLdapGroupFilter gets a reference to the given string and assigns it to the LdapGroupFilter field.
func (o *LdapConfigPart) SetLdapGroupFilter(v string) {
	o.LdapGroupFilter = &v
}

// GetLdapPrivateKey returns the LdapPrivateKey field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapPrivateKey() string {
	if o == nil || o.LdapPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.LdapPrivateKey
}

// GetLdapPrivateKeyOk returns a tuple with the LdapPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapPrivateKeyOk() (*string, bool) {
	if o == nil || o.LdapPrivateKey == nil {
		return nil, false
	}
	return o.LdapPrivateKey, true
}

// HasLdapPrivateKey returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapPrivateKey() bool {
	if o != nil && o.LdapPrivateKey != nil {
		return true
	}

	return false
}

// SetLdapPrivateKey gets a reference to the given string and assigns it to the LdapPrivateKey field.
func (o *LdapConfigPart) SetLdapPrivateKey(v string) {
	o.LdapPrivateKey = &v
}

// GetLdapTokenExpiration returns the LdapTokenExpiration field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapTokenExpiration() string {
	if o == nil || o.LdapTokenExpiration == nil {
		var ret string
		return ret
	}
	return *o.LdapTokenExpiration
}

// GetLdapTokenExpirationOk returns a tuple with the LdapTokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapTokenExpirationOk() (*string, bool) {
	if o == nil || o.LdapTokenExpiration == nil {
		return nil, false
	}
	return o.LdapTokenExpiration, true
}

// HasLdapTokenExpiration returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapTokenExpiration() bool {
	if o != nil && o.LdapTokenExpiration != nil {
		return true
	}

	return false
}

// SetLdapTokenExpiration gets a reference to the given string and assigns it to the LdapTokenExpiration field.
func (o *LdapConfigPart) SetLdapTokenExpiration(v string) {
	o.LdapTokenExpiration = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapUrl() string {
	if o == nil || o.LdapUrl == nil {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapUrlOk() (*string, bool) {
	if o == nil || o.LdapUrl == nil {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapUrl() bool {
	if o != nil && o.LdapUrl != nil {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *LdapConfigPart) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetLdapUserAttr returns the LdapUserAttr field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapUserAttr() string {
	if o == nil || o.LdapUserAttr == nil {
		var ret string
		return ret
	}
	return *o.LdapUserAttr
}

// GetLdapUserAttrOk returns a tuple with the LdapUserAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapUserAttrOk() (*string, bool) {
	if o == nil || o.LdapUserAttr == nil {
		return nil, false
	}
	return o.LdapUserAttr, true
}

// HasLdapUserAttr returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapUserAttr() bool {
	if o != nil && o.LdapUserAttr != nil {
		return true
	}

	return false
}

// SetLdapUserAttr gets a reference to the given string and assigns it to the LdapUserAttr field.
func (o *LdapConfigPart) SetLdapUserAttr(v string) {
	o.LdapUserAttr = &v
}

// GetLdapUserDn returns the LdapUserDn field value if set, zero value otherwise.
func (o *LdapConfigPart) GetLdapUserDn() string {
	if o == nil || o.LdapUserDn == nil {
		var ret string
		return ret
	}
	return *o.LdapUserDn
}

// GetLdapUserDnOk returns a tuple with the LdapUserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigPart) GetLdapUserDnOk() (*string, bool) {
	if o == nil || o.LdapUserDn == nil {
		return nil, false
	}
	return o.LdapUserDn, true
}

// HasLdapUserDn returns a boolean if a field has been set.
func (o *LdapConfigPart) HasLdapUserDn() bool {
	if o != nil && o.LdapUserDn != nil {
		return true
	}

	return false
}

// SetLdapUserDn gets a reference to the given string and assigns it to the LdapUserDn field.
func (o *LdapConfigPart) SetLdapUserDn(v string) {
	o.LdapUserDn = &v
}

func (o LdapConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LdapAccessId != nil {
		toSerialize["ldap_access_id"] = o.LdapAccessId
	}
	if o.LdapAnonymousSearch != nil {
		toSerialize["ldap_anonymous_search"] = o.LdapAnonymousSearch
	}
	if o.LdapBindDn != nil {
		toSerialize["ldap_bind_dn"] = o.LdapBindDn
	}
	if o.LdapBindPassword != nil {
		toSerialize["ldap_bind_password"] = o.LdapBindPassword
	}
	if o.LdapCert != nil {
		toSerialize["ldap_cert"] = o.LdapCert
	}
	if o.LdapEnable != nil {
		toSerialize["ldap_enable"] = o.LdapEnable
	}
	if o.LdapGroupAttr != nil {
		toSerialize["ldap_group_attr"] = o.LdapGroupAttr
	}
	if o.LdapGroupDn != nil {
		toSerialize["ldap_group_dn"] = o.LdapGroupDn
	}
	if o.LdapGroupFilter != nil {
		toSerialize["ldap_group_filter"] = o.LdapGroupFilter
	}
	if o.LdapPrivateKey != nil {
		toSerialize["ldap_private_key"] = o.LdapPrivateKey
	}
	if o.LdapTokenExpiration != nil {
		toSerialize["ldap_token_expiration"] = o.LdapTokenExpiration
	}
	if o.LdapUrl != nil {
		toSerialize["ldap_url"] = o.LdapUrl
	}
	if o.LdapUserAttr != nil {
		toSerialize["ldap_user_attr"] = o.LdapUserAttr
	}
	if o.LdapUserDn != nil {
		toSerialize["ldap_user_dn"] = o.LdapUserDn
	}
	return json.Marshal(toSerialize)
}

type NullableLdapConfigPart struct {
	value *LdapConfigPart
	isSet bool
}

func (v NullableLdapConfigPart) Get() *LdapConfigPart {
	return v.value
}

func (v *NullableLdapConfigPart) Set(val *LdapConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapConfigPart(val *LdapConfigPart) *NullableLdapConfigPart {
	return &NullableLdapConfigPart{value: val, isSet: true}
}

func (v NullableLdapConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


