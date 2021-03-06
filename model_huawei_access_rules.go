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

// HuaweiAccessRules struct for HuaweiAccessRules
type HuaweiAccessRules struct {
	// The auth URL.
	AuthEndpoint *string `json:"auth_endpoint,omitempty"`
	// The list of domain ids that the login is restricted to.
	DomainId []string `json:"domain_id,omitempty"`
	// The list of domainNames that the login is restricted to.
	DomainName []string `json:"domain_name,omitempty"`
	// The list of tenantIDs  that the login is restricted to.
	TenantId []string `json:"tenant_id,omitempty"`
	// The list of tenantNames  that the login is restricted to.
	TenantName []string `json:"tenant_name,omitempty"`
	// The list of user ids that the login is restricted to.
	UserId []string `json:"user_id,omitempty"`
	// The list of user names that the login is restricted to.
	UserName []string `json:"user_name,omitempty"`
}

// NewHuaweiAccessRules instantiates a new HuaweiAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHuaweiAccessRules() *HuaweiAccessRules {
	this := HuaweiAccessRules{}
	return &this
}

// NewHuaweiAccessRulesWithDefaults instantiates a new HuaweiAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHuaweiAccessRulesWithDefaults() *HuaweiAccessRules {
	this := HuaweiAccessRules{}
	return &this
}

// GetAuthEndpoint returns the AuthEndpoint field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetAuthEndpoint() string {
	if o == nil || o.AuthEndpoint == nil {
		var ret string
		return ret
	}
	return *o.AuthEndpoint
}

// GetAuthEndpointOk returns a tuple with the AuthEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetAuthEndpointOk() (*string, bool) {
	if o == nil || o.AuthEndpoint == nil {
		return nil, false
	}
	return o.AuthEndpoint, true
}

// HasAuthEndpoint returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasAuthEndpoint() bool {
	if o != nil && o.AuthEndpoint != nil {
		return true
	}

	return false
}

// SetAuthEndpoint gets a reference to the given string and assigns it to the AuthEndpoint field.
func (o *HuaweiAccessRules) SetAuthEndpoint(v string) {
	o.AuthEndpoint = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetDomainId() []string {
	if o == nil || o.DomainId == nil {
		var ret []string
		return ret
	}
	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetDomainIdOk() ([]string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given []string and assigns it to the DomainId field.
func (o *HuaweiAccessRules) SetDomainId(v []string) {
	o.DomainId = v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetDomainName() []string {
	if o == nil || o.DomainName == nil {
		var ret []string
		return ret
	}
	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetDomainNameOk() ([]string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given []string and assigns it to the DomainName field.
func (o *HuaweiAccessRules) SetDomainName(v []string) {
	o.DomainName = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetTenantId() []string {
	if o == nil || o.TenantId == nil {
		var ret []string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetTenantIdOk() ([]string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given []string and assigns it to the TenantId field.
func (o *HuaweiAccessRules) SetTenantId(v []string) {
	o.TenantId = v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetTenantName() []string {
	if o == nil || o.TenantName == nil {
		var ret []string
		return ret
	}
	return o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetTenantNameOk() ([]string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasTenantName() bool {
	if o != nil && o.TenantName != nil {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given []string and assigns it to the TenantName field.
func (o *HuaweiAccessRules) SetTenantName(v []string) {
	o.TenantName = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetUserId() []string {
	if o == nil || o.UserId == nil {
		var ret []string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetUserIdOk() ([]string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given []string and assigns it to the UserId field.
func (o *HuaweiAccessRules) SetUserId(v []string) {
	o.UserId = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *HuaweiAccessRules) GetUserName() []string {
	if o == nil || o.UserName == nil {
		var ret []string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HuaweiAccessRules) GetUserNameOk() ([]string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *HuaweiAccessRules) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given []string and assigns it to the UserName field.
func (o *HuaweiAccessRules) SetUserName(v []string) {
	o.UserName = v
}

func (o HuaweiAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthEndpoint != nil {
		toSerialize["auth_endpoint"] = o.AuthEndpoint
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.DomainName != nil {
		toSerialize["domain_name"] = o.DomainName
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.TenantName != nil {
		toSerialize["tenant_name"] = o.TenantName
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableHuaweiAccessRules struct {
	value *HuaweiAccessRules
	isSet bool
}

func (v NullableHuaweiAccessRules) Get() *HuaweiAccessRules {
	return v.value
}

func (v *NullableHuaweiAccessRules) Set(val *HuaweiAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableHuaweiAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableHuaweiAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHuaweiAccessRules(val *HuaweiAccessRules) *NullableHuaweiAccessRules {
	return &NullableHuaweiAccessRules{value: val, isSet: true}
}

func (v NullableHuaweiAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHuaweiAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


