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

// K8SPayload struct for K8SPayload
type K8SPayload struct {
	Ca []int32 `json:"ca,omitempty"`
	ClientCert []int32 `json:"client_cert,omitempty"`
	ClientKey []int32 `json:"client_key,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Password *string `json:"password,omitempty"`
	Server *string `json:"server,omitempty"`
	SkipSystem *bool `json:"skip_system,omitempty"`
	Token *string `json:"token,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewK8SPayload instantiates a new K8SPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8SPayload() *K8SPayload {
	this := K8SPayload{}
	return &this
}

// NewK8SPayloadWithDefaults instantiates a new K8SPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8SPayloadWithDefaults() *K8SPayload {
	this := K8SPayload{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *K8SPayload) GetCa() []int32 {
	if o == nil || o.Ca == nil {
		var ret []int32
		return ret
	}
	return o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetCaOk() ([]int32, bool) {
	if o == nil || o.Ca == nil {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *K8SPayload) HasCa() bool {
	if o != nil && o.Ca != nil {
		return true
	}

	return false
}

// SetCa gets a reference to the given []int32 and assigns it to the Ca field.
func (o *K8SPayload) SetCa(v []int32) {
	o.Ca = v
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise.
func (o *K8SPayload) GetClientCert() []int32 {
	if o == nil || o.ClientCert == nil {
		var ret []int32
		return ret
	}
	return o.ClientCert
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetClientCertOk() ([]int32, bool) {
	if o == nil || o.ClientCert == nil {
		return nil, false
	}
	return o.ClientCert, true
}

// HasClientCert returns a boolean if a field has been set.
func (o *K8SPayload) HasClientCert() bool {
	if o != nil && o.ClientCert != nil {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given []int32 and assigns it to the ClientCert field.
func (o *K8SPayload) SetClientCert(v []int32) {
	o.ClientCert = v
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise.
func (o *K8SPayload) GetClientKey() []int32 {
	if o == nil || o.ClientKey == nil {
		var ret []int32
		return ret
	}
	return o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetClientKeyOk() ([]int32, bool) {
	if o == nil || o.ClientKey == nil {
		return nil, false
	}
	return o.ClientKey, true
}

// HasClientKey returns a boolean if a field has been set.
func (o *K8SPayload) HasClientKey() bool {
	if o != nil && o.ClientKey != nil {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given []int32 and assigns it to the ClientKey field.
func (o *K8SPayload) SetClientKey(v []int32) {
	o.ClientKey = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *K8SPayload) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *K8SPayload) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *K8SPayload) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *K8SPayload) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *K8SPayload) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *K8SPayload) SetPassword(v string) {
	o.Password = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *K8SPayload) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *K8SPayload) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *K8SPayload) SetServer(v string) {
	o.Server = &v
}

// GetSkipSystem returns the SkipSystem field value if set, zero value otherwise.
func (o *K8SPayload) GetSkipSystem() bool {
	if o == nil || o.SkipSystem == nil {
		var ret bool
		return ret
	}
	return *o.SkipSystem
}

// GetSkipSystemOk returns a tuple with the SkipSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetSkipSystemOk() (*bool, bool) {
	if o == nil || o.SkipSystem == nil {
		return nil, false
	}
	return o.SkipSystem, true
}

// HasSkipSystem returns a boolean if a field has been set.
func (o *K8SPayload) HasSkipSystem() bool {
	if o != nil && o.SkipSystem != nil {
		return true
	}

	return false
}

// SetSkipSystem gets a reference to the given bool and assigns it to the SkipSystem field.
func (o *K8SPayload) SetSkipSystem(v bool) {
	o.SkipSystem = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *K8SPayload) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *K8SPayload) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *K8SPayload) SetToken(v string) {
	o.Token = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *K8SPayload) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8SPayload) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *K8SPayload) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *K8SPayload) SetUsername(v string) {
	o.Username = &v
}

func (o K8SPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ca != nil {
		toSerialize["ca"] = o.Ca
	}
	if o.ClientCert != nil {
		toSerialize["client_cert"] = o.ClientCert
	}
	if o.ClientKey != nil {
		toSerialize["client_key"] = o.ClientKey
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.SkipSystem != nil {
		toSerialize["skip_system"] = o.SkipSystem
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableK8SPayload struct {
	value *K8SPayload
	isSet bool
}

func (v NullableK8SPayload) Get() *K8SPayload {
	return v.value
}

func (v *NullableK8SPayload) Set(val *K8SPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableK8SPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableK8SPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8SPayload(val *K8SPayload) *NullableK8SPayload {
	return &NullableK8SPayload{value: val, isSet: true}
}

func (v NullableK8SPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8SPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


