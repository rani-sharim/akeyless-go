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

// GatewayUpdateProducerGke gatewayUpdateProducerGke is a command that updates gke producer
type GatewayUpdateProducerGke struct {
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// GKE Service Account key file path
	GkeAccountKey *string `json:"gke-account-key,omitempty"`
	// GKE cluster CA certificate
	GkeClusterCert *string `json:"gke-cluster-cert,omitempty"`
	// GKE cluster URL endpoint
	GkeClusterEndpoint *string `json:"gke-cluster-endpoint,omitempty"`
	// GKE cluster name
	GkeClusterName *string `json:"gke-cluster-name,omitempty"`
	// GKE service account email
	GkeServiceAccountEmail *string `json:"gke-service-account-email,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	SecureAccessAllowPortForwading *bool `json:"secure-access-allow-port-forwading,omitempty"`
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	SecureAccessClusterEndpoint *string `json:"secure-access-cluster-endpoint,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// List of the tags attached to this secret
	Tags []string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayUpdateProducerGke instantiates a new GatewayUpdateProducerGke object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerGke(name string) *GatewayUpdateProducerGke {
	this := GatewayUpdateProducerGke{}
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerGkeWithDefaults instantiates a new GatewayUpdateProducerGke object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerGkeWithDefaults() *GatewayUpdateProducerGke {
	this := GatewayUpdateProducerGke{}
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerGke) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetGkeAccountKey returns the GkeAccountKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetGkeAccountKey() string {
	if o == nil || o.GkeAccountKey == nil {
		var ret string
		return ret
	}
	return *o.GkeAccountKey
}

// GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetGkeAccountKeyOk() (*string, bool) {
	if o == nil || o.GkeAccountKey == nil {
		return nil, false
	}
	return o.GkeAccountKey, true
}

// HasGkeAccountKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasGkeAccountKey() bool {
	if o != nil && o.GkeAccountKey != nil {
		return true
	}

	return false
}

// SetGkeAccountKey gets a reference to the given string and assigns it to the GkeAccountKey field.
func (o *GatewayUpdateProducerGke) SetGkeAccountKey(v string) {
	o.GkeAccountKey = &v
}

// GetGkeClusterCert returns the GkeClusterCert field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetGkeClusterCert() string {
	if o == nil || o.GkeClusterCert == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterCert
}

// GetGkeClusterCertOk returns a tuple with the GkeClusterCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetGkeClusterCertOk() (*string, bool) {
	if o == nil || o.GkeClusterCert == nil {
		return nil, false
	}
	return o.GkeClusterCert, true
}

// HasGkeClusterCert returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasGkeClusterCert() bool {
	if o != nil && o.GkeClusterCert != nil {
		return true
	}

	return false
}

// SetGkeClusterCert gets a reference to the given string and assigns it to the GkeClusterCert field.
func (o *GatewayUpdateProducerGke) SetGkeClusterCert(v string) {
	o.GkeClusterCert = &v
}

// GetGkeClusterEndpoint returns the GkeClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetGkeClusterEndpoint() string {
	if o == nil || o.GkeClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterEndpoint
}

// GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetGkeClusterEndpointOk() (*string, bool) {
	if o == nil || o.GkeClusterEndpoint == nil {
		return nil, false
	}
	return o.GkeClusterEndpoint, true
}

// HasGkeClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasGkeClusterEndpoint() bool {
	if o != nil && o.GkeClusterEndpoint != nil {
		return true
	}

	return false
}

// SetGkeClusterEndpoint gets a reference to the given string and assigns it to the GkeClusterEndpoint field.
func (o *GatewayUpdateProducerGke) SetGkeClusterEndpoint(v string) {
	o.GkeClusterEndpoint = &v
}

// GetGkeClusterName returns the GkeClusterName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetGkeClusterName() string {
	if o == nil || o.GkeClusterName == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterName
}

// GetGkeClusterNameOk returns a tuple with the GkeClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetGkeClusterNameOk() (*string, bool) {
	if o == nil || o.GkeClusterName == nil {
		return nil, false
	}
	return o.GkeClusterName, true
}

// HasGkeClusterName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasGkeClusterName() bool {
	if o != nil && o.GkeClusterName != nil {
		return true
	}

	return false
}

// SetGkeClusterName gets a reference to the given string and assigns it to the GkeClusterName field.
func (o *GatewayUpdateProducerGke) SetGkeClusterName(v string) {
	o.GkeClusterName = &v
}

// GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetGkeServiceAccountEmail() string {
	if o == nil || o.GkeServiceAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.GkeServiceAccountEmail
}

// GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetGkeServiceAccountEmailOk() (*string, bool) {
	if o == nil || o.GkeServiceAccountEmail == nil {
		return nil, false
	}
	return o.GkeServiceAccountEmail, true
}

// HasGkeServiceAccountEmail returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasGkeServiceAccountEmail() bool {
	if o != nil && o.GkeServiceAccountEmail != nil {
		return true
	}

	return false
}

// SetGkeServiceAccountEmail gets a reference to the given string and assigns it to the GkeServiceAccountEmail field.
func (o *GatewayUpdateProducerGke) SetGkeServiceAccountEmail(v string) {
	o.GkeServiceAccountEmail = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerGke) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerGke) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerGke) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerGke) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetSecureAccessAllowPortForwading() bool {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowPortForwading
}

// GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetSecureAccessAllowPortForwadingOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		return nil, false
	}
	return o.SecureAccessAllowPortForwading, true
}

// HasSecureAccessAllowPortForwading returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasSecureAccessAllowPortForwading() bool {
	if o != nil && o.SecureAccessAllowPortForwading != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowPortForwading gets a reference to the given bool and assigns it to the SecureAccessAllowPortForwading field.
func (o *GatewayUpdateProducerGke) SetSecureAccessAllowPortForwading(v bool) {
	o.SecureAccessAllowPortForwading = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayUpdateProducerGke) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetSecureAccessClusterEndpoint() string {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessClusterEndpoint
}

// GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetSecureAccessClusterEndpointOk() (*string, bool) {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		return nil, false
	}
	return o.SecureAccessClusterEndpoint, true
}

// HasSecureAccessClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasSecureAccessClusterEndpoint() bool {
	if o != nil && o.SecureAccessClusterEndpoint != nil {
		return true
	}

	return false
}

// SetSecureAccessClusterEndpoint gets a reference to the given string and assigns it to the SecureAccessClusterEndpoint field.
func (o *GatewayUpdateProducerGke) SetSecureAccessClusterEndpoint(v string) {
	o.SecureAccessClusterEndpoint = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerGke) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerGke) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerGke) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerGke) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerGke) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerGke) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGke) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGke) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGke) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerGke) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerGke) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.GkeAccountKey != nil {
		toSerialize["gke-account-key"] = o.GkeAccountKey
	}
	if o.GkeClusterCert != nil {
		toSerialize["gke-cluster-cert"] = o.GkeClusterCert
	}
	if o.GkeClusterEndpoint != nil {
		toSerialize["gke-cluster-endpoint"] = o.GkeClusterEndpoint
	}
	if o.GkeClusterName != nil {
		toSerialize["gke-cluster-name"] = o.GkeClusterName
	}
	if o.GkeServiceAccountEmail != nil {
		toSerialize["gke-service-account-email"] = o.GkeServiceAccountEmail
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.SecureAccessAllowPortForwading != nil {
		toSerialize["secure-access-allow-port-forwading"] = o.SecureAccessAllowPortForwading
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessClusterEndpoint != nil {
		toSerialize["secure-access-cluster-endpoint"] = o.SecureAccessClusterEndpoint
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerGke struct {
	value *GatewayUpdateProducerGke
	isSet bool
}

func (v NullableGatewayUpdateProducerGke) Get() *GatewayUpdateProducerGke {
	return v.value
}

func (v *NullableGatewayUpdateProducerGke) Set(val *GatewayUpdateProducerGke) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerGke) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerGke) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerGke(val *GatewayUpdateProducerGke) *NullableGatewayUpdateProducerGke {
	return &NullableGatewayUpdateProducerGke{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerGke) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerGke) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


