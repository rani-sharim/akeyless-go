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

// Connect Connect is a command that performs secure remote access
type Connect struct {
	Helper map[string]interface{} `json:"Helper,omitempty"`
	// used to override .akeyless-connect.rc in tests
	RcFileOverride *string `json:"RcFileOverride,omitempty"`
	// The Bastion API path
	BastionCtrlPath *string `json:"bastion-ctrl-path,omitempty"`
	// The Bastion API Port
	BastionCtrlPort *string `json:"bastion-ctrl-port,omitempty"`
	// The Bastion API protocol
	BastionCtrlProto *string `json:"bastion-ctrl-proto,omitempty"`
	// The Bastion API prefix
	BastionCtrlSubdomain *string `json:"bastion-ctrl-subdomain,omitempty"`
	// The Akeyless certificate issuer name
	CertIssuerName *string `json:"cert-issuer-name,omitempty"`
	// The file from which the identity (private key) for public key authentication is read
	IdentityFile *string `json:"identity-file,omitempty"`
	// The Secret name (for database and AWS producers - producer name)
	Name *string `json:"name,omitempty"`
	// The Use to add offical SSH arguments (except -i)
	SshExtraArgs *string `json:"ssh-extra-args,omitempty"`
	// Set this option to output legacy ('ssh-rsa-cert-v01@openssh.com') signing algorithm name in the ssh certificate.
	SshLegacySigningAlg *bool `json:"ssh-legacy-signing-alg,omitempty"`
	// The target
	Target *string `json:"target,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The jump box server
	ViaBastion *string `json:"via-bastion,omitempty"`
}

// NewConnect instantiates a new Connect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnect() *Connect {
	this := Connect{}
	var bastionCtrlPort string = "9900"
	this.BastionCtrlPort = &bastionCtrlPort
	var bastionCtrlProto string = "http"
	this.BastionCtrlProto = &bastionCtrlProto
	return &this
}

// NewConnectWithDefaults instantiates a new Connect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectWithDefaults() *Connect {
	this := Connect{}
	var bastionCtrlPort string = "9900"
	this.BastionCtrlPort = &bastionCtrlPort
	var bastionCtrlProto string = "http"
	this.BastionCtrlProto = &bastionCtrlProto
	return &this
}

// GetHelper returns the Helper field value if set, zero value otherwise.
func (o *Connect) GetHelper() map[string]interface{} {
	if o == nil || o.Helper == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Helper
}

// GetHelperOk returns a tuple with the Helper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetHelperOk() (map[string]interface{}, bool) {
	if o == nil || o.Helper == nil {
		return nil, false
	}
	return o.Helper, true
}

// HasHelper returns a boolean if a field has been set.
func (o *Connect) HasHelper() bool {
	if o != nil && o.Helper != nil {
		return true
	}

	return false
}

// SetHelper gets a reference to the given map[string]interface{} and assigns it to the Helper field.
func (o *Connect) SetHelper(v map[string]interface{}) {
	o.Helper = v
}

// GetRcFileOverride returns the RcFileOverride field value if set, zero value otherwise.
func (o *Connect) GetRcFileOverride() string {
	if o == nil || o.RcFileOverride == nil {
		var ret string
		return ret
	}
	return *o.RcFileOverride
}

// GetRcFileOverrideOk returns a tuple with the RcFileOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetRcFileOverrideOk() (*string, bool) {
	if o == nil || o.RcFileOverride == nil {
		return nil, false
	}
	return o.RcFileOverride, true
}

// HasRcFileOverride returns a boolean if a field has been set.
func (o *Connect) HasRcFileOverride() bool {
	if o != nil && o.RcFileOverride != nil {
		return true
	}

	return false
}

// SetRcFileOverride gets a reference to the given string and assigns it to the RcFileOverride field.
func (o *Connect) SetRcFileOverride(v string) {
	o.RcFileOverride = &v
}

// GetBastionCtrlPath returns the BastionCtrlPath field value if set, zero value otherwise.
func (o *Connect) GetBastionCtrlPath() string {
	if o == nil || o.BastionCtrlPath == nil {
		var ret string
		return ret
	}
	return *o.BastionCtrlPath
}

// GetBastionCtrlPathOk returns a tuple with the BastionCtrlPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetBastionCtrlPathOk() (*string, bool) {
	if o == nil || o.BastionCtrlPath == nil {
		return nil, false
	}
	return o.BastionCtrlPath, true
}

// HasBastionCtrlPath returns a boolean if a field has been set.
func (o *Connect) HasBastionCtrlPath() bool {
	if o != nil && o.BastionCtrlPath != nil {
		return true
	}

	return false
}

// SetBastionCtrlPath gets a reference to the given string and assigns it to the BastionCtrlPath field.
func (o *Connect) SetBastionCtrlPath(v string) {
	o.BastionCtrlPath = &v
}

// GetBastionCtrlPort returns the BastionCtrlPort field value if set, zero value otherwise.
func (o *Connect) GetBastionCtrlPort() string {
	if o == nil || o.BastionCtrlPort == nil {
		var ret string
		return ret
	}
	return *o.BastionCtrlPort
}

// GetBastionCtrlPortOk returns a tuple with the BastionCtrlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetBastionCtrlPortOk() (*string, bool) {
	if o == nil || o.BastionCtrlPort == nil {
		return nil, false
	}
	return o.BastionCtrlPort, true
}

// HasBastionCtrlPort returns a boolean if a field has been set.
func (o *Connect) HasBastionCtrlPort() bool {
	if o != nil && o.BastionCtrlPort != nil {
		return true
	}

	return false
}

// SetBastionCtrlPort gets a reference to the given string and assigns it to the BastionCtrlPort field.
func (o *Connect) SetBastionCtrlPort(v string) {
	o.BastionCtrlPort = &v
}

// GetBastionCtrlProto returns the BastionCtrlProto field value if set, zero value otherwise.
func (o *Connect) GetBastionCtrlProto() string {
	if o == nil || o.BastionCtrlProto == nil {
		var ret string
		return ret
	}
	return *o.BastionCtrlProto
}

// GetBastionCtrlProtoOk returns a tuple with the BastionCtrlProto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetBastionCtrlProtoOk() (*string, bool) {
	if o == nil || o.BastionCtrlProto == nil {
		return nil, false
	}
	return o.BastionCtrlProto, true
}

// HasBastionCtrlProto returns a boolean if a field has been set.
func (o *Connect) HasBastionCtrlProto() bool {
	if o != nil && o.BastionCtrlProto != nil {
		return true
	}

	return false
}

// SetBastionCtrlProto gets a reference to the given string and assigns it to the BastionCtrlProto field.
func (o *Connect) SetBastionCtrlProto(v string) {
	o.BastionCtrlProto = &v
}

// GetBastionCtrlSubdomain returns the BastionCtrlSubdomain field value if set, zero value otherwise.
func (o *Connect) GetBastionCtrlSubdomain() string {
	if o == nil || o.BastionCtrlSubdomain == nil {
		var ret string
		return ret
	}
	return *o.BastionCtrlSubdomain
}

// GetBastionCtrlSubdomainOk returns a tuple with the BastionCtrlSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetBastionCtrlSubdomainOk() (*string, bool) {
	if o == nil || o.BastionCtrlSubdomain == nil {
		return nil, false
	}
	return o.BastionCtrlSubdomain, true
}

// HasBastionCtrlSubdomain returns a boolean if a field has been set.
func (o *Connect) HasBastionCtrlSubdomain() bool {
	if o != nil && o.BastionCtrlSubdomain != nil {
		return true
	}

	return false
}

// SetBastionCtrlSubdomain gets a reference to the given string and assigns it to the BastionCtrlSubdomain field.
func (o *Connect) SetBastionCtrlSubdomain(v string) {
	o.BastionCtrlSubdomain = &v
}

// GetCertIssuerName returns the CertIssuerName field value if set, zero value otherwise.
func (o *Connect) GetCertIssuerName() string {
	if o == nil || o.CertIssuerName == nil {
		var ret string
		return ret
	}
	return *o.CertIssuerName
}

// GetCertIssuerNameOk returns a tuple with the CertIssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetCertIssuerNameOk() (*string, bool) {
	if o == nil || o.CertIssuerName == nil {
		return nil, false
	}
	return o.CertIssuerName, true
}

// HasCertIssuerName returns a boolean if a field has been set.
func (o *Connect) HasCertIssuerName() bool {
	if o != nil && o.CertIssuerName != nil {
		return true
	}

	return false
}

// SetCertIssuerName gets a reference to the given string and assigns it to the CertIssuerName field.
func (o *Connect) SetCertIssuerName(v string) {
	o.CertIssuerName = &v
}

// GetIdentityFile returns the IdentityFile field value if set, zero value otherwise.
func (o *Connect) GetIdentityFile() string {
	if o == nil || o.IdentityFile == nil {
		var ret string
		return ret
	}
	return *o.IdentityFile
}

// GetIdentityFileOk returns a tuple with the IdentityFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetIdentityFileOk() (*string, bool) {
	if o == nil || o.IdentityFile == nil {
		return nil, false
	}
	return o.IdentityFile, true
}

// HasIdentityFile returns a boolean if a field has been set.
func (o *Connect) HasIdentityFile() bool {
	if o != nil && o.IdentityFile != nil {
		return true
	}

	return false
}

// SetIdentityFile gets a reference to the given string and assigns it to the IdentityFile field.
func (o *Connect) SetIdentityFile(v string) {
	o.IdentityFile = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Connect) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Connect) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Connect) SetName(v string) {
	o.Name = &v
}

// GetSshExtraArgs returns the SshExtraArgs field value if set, zero value otherwise.
func (o *Connect) GetSshExtraArgs() string {
	if o == nil || o.SshExtraArgs == nil {
		var ret string
		return ret
	}
	return *o.SshExtraArgs
}

// GetSshExtraArgsOk returns a tuple with the SshExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetSshExtraArgsOk() (*string, bool) {
	if o == nil || o.SshExtraArgs == nil {
		return nil, false
	}
	return o.SshExtraArgs, true
}

// HasSshExtraArgs returns a boolean if a field has been set.
func (o *Connect) HasSshExtraArgs() bool {
	if o != nil && o.SshExtraArgs != nil {
		return true
	}

	return false
}

// SetSshExtraArgs gets a reference to the given string and assigns it to the SshExtraArgs field.
func (o *Connect) SetSshExtraArgs(v string) {
	o.SshExtraArgs = &v
}

// GetSshLegacySigningAlg returns the SshLegacySigningAlg field value if set, zero value otherwise.
func (o *Connect) GetSshLegacySigningAlg() bool {
	if o == nil || o.SshLegacySigningAlg == nil {
		var ret bool
		return ret
	}
	return *o.SshLegacySigningAlg
}

// GetSshLegacySigningAlgOk returns a tuple with the SshLegacySigningAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetSshLegacySigningAlgOk() (*bool, bool) {
	if o == nil || o.SshLegacySigningAlg == nil {
		return nil, false
	}
	return o.SshLegacySigningAlg, true
}

// HasSshLegacySigningAlg returns a boolean if a field has been set.
func (o *Connect) HasSshLegacySigningAlg() bool {
	if o != nil && o.SshLegacySigningAlg != nil {
		return true
	}

	return false
}

// SetSshLegacySigningAlg gets a reference to the given bool and assigns it to the SshLegacySigningAlg field.
func (o *Connect) SetSshLegacySigningAlg(v bool) {
	o.SshLegacySigningAlg = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Connect) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Connect) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *Connect) SetTarget(v string) {
	o.Target = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Connect) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Connect) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Connect) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Connect) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Connect) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Connect) SetUidToken(v string) {
	o.UidToken = &v
}

// GetViaBastion returns the ViaBastion field value if set, zero value otherwise.
func (o *Connect) GetViaBastion() string {
	if o == nil || o.ViaBastion == nil {
		var ret string
		return ret
	}
	return *o.ViaBastion
}

// GetViaBastionOk returns a tuple with the ViaBastion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connect) GetViaBastionOk() (*string, bool) {
	if o == nil || o.ViaBastion == nil {
		return nil, false
	}
	return o.ViaBastion, true
}

// HasViaBastion returns a boolean if a field has been set.
func (o *Connect) HasViaBastion() bool {
	if o != nil && o.ViaBastion != nil {
		return true
	}

	return false
}

// SetViaBastion gets a reference to the given string and assigns it to the ViaBastion field.
func (o *Connect) SetViaBastion(v string) {
	o.ViaBastion = &v
}

func (o Connect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Helper != nil {
		toSerialize["Helper"] = o.Helper
	}
	if o.RcFileOverride != nil {
		toSerialize["RcFileOverride"] = o.RcFileOverride
	}
	if o.BastionCtrlPath != nil {
		toSerialize["bastion-ctrl-path"] = o.BastionCtrlPath
	}
	if o.BastionCtrlPort != nil {
		toSerialize["bastion-ctrl-port"] = o.BastionCtrlPort
	}
	if o.BastionCtrlProto != nil {
		toSerialize["bastion-ctrl-proto"] = o.BastionCtrlProto
	}
	if o.BastionCtrlSubdomain != nil {
		toSerialize["bastion-ctrl-subdomain"] = o.BastionCtrlSubdomain
	}
	if o.CertIssuerName != nil {
		toSerialize["cert-issuer-name"] = o.CertIssuerName
	}
	if o.IdentityFile != nil {
		toSerialize["identity-file"] = o.IdentityFile
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SshExtraArgs != nil {
		toSerialize["ssh-extra-args"] = o.SshExtraArgs
	}
	if o.SshLegacySigningAlg != nil {
		toSerialize["ssh-legacy-signing-alg"] = o.SshLegacySigningAlg
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.ViaBastion != nil {
		toSerialize["via-bastion"] = o.ViaBastion
	}
	return json.Marshal(toSerialize)
}

type NullableConnect struct {
	value *Connect
	isSet bool
}

func (v NullableConnect) Get() *Connect {
	return v.value
}

func (v *NullableConnect) Set(val *Connect) {
	v.value = val
	v.isSet = true
}

func (v NullableConnect) IsSet() bool {
	return v.isSet
}

func (v *NullableConnect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnect(val *Connect) *NullableConnect {
	return &NullableConnect{value: val, isSet: true}
}

func (v NullableConnect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


