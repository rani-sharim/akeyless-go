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

// GatewayUpdateProducerAws gatewayUpdateProducerAws is a command that Updates aws producer
type GatewayUpdateProducerAws struct {
	AccessMode *string `json:"access-mode,omitempty"`
	// Admin credentials rotation interval (days)
	AdminRotationIntervalDays *int64 `json:"admin-rotation-interval-days,omitempty"`
	// Access Key ID
	AwsAccessKeyId *string `json:"aws-access-key-id,omitempty"`
	// Secret Access Key
	AwsAccessSecretKey *string `json:"aws-access-secret-key,omitempty"`
	// AWS Role ARNs to be used in the Assume Role operation (relevant only for assume_role mode)
	AwsRoleArns *string `json:"aws-role-arns,omitempty"`
	// AWS User console access
	AwsUserConsoleAccess *bool `json:"aws-user-console-access,omitempty"`
	// AWS User groups
	AwsUserGroups *string `json:"aws-user-groups,omitempty"`
	// AWS User policies
	AwsUserPolicies *string `json:"aws-user-policies,omitempty"`
	// AWS User programmatic access
	AwsUserProgrammaticAccess *bool `json:"aws-user-programmatic-access,omitempty"`
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Automatic admin credentials rotation
	EnableAdminRotation *bool `json:"enable-admin-rotation,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Region
	Region *string `json:"region,omitempty"`
	SecureAccessAwsAccountId *string `json:"secure-access-aws-account-id,omitempty"`
	SecureAccessAwsNativeCli *bool `json:"secure-access-aws-native-cli,omitempty"`
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	SecureAccessWebBrowsing *bool `json:"secure-access-web-browsing,omitempty"`
	SecureAccessWebProxy *bool `json:"secure-access-web-proxy,omitempty"`
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

// NewGatewayUpdateProducerAws instantiates a new GatewayUpdateProducerAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerAws(name string) *GatewayUpdateProducerAws {
	this := GatewayUpdateProducerAws{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var awsUserConsoleAccess bool = false
	this.AwsUserConsoleAccess = &awsUserConsoleAccess
	var awsUserProgrammaticAccess bool = true
	this.AwsUserProgrammaticAccess = &awsUserProgrammaticAccess
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	this.Name = name
	var region string = "us-east-2"
	this.Region = &region
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerAwsWithDefaults instantiates a new GatewayUpdateProducerAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerAwsWithDefaults() *GatewayUpdateProducerAws {
	this := GatewayUpdateProducerAws{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var awsUserConsoleAccess bool = false
	this.AwsUserConsoleAccess = &awsUserConsoleAccess
	var awsUserProgrammaticAccess bool = true
	this.AwsUserProgrammaticAccess = &awsUserProgrammaticAccess
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var region string = "us-east-2"
	this.Region = &region
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *GatewayUpdateProducerAws) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *GatewayUpdateProducerAws) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsAccessKeyId() string {
	if o == nil || o.AwsAccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AwsAccessKeyId == nil {
		return nil, false
	}
	return o.AwsAccessKeyId, true
}

// HasAwsAccessKeyId returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsAccessKeyId() bool {
	if o != nil && o.AwsAccessKeyId != nil {
		return true
	}

	return false
}

// SetAwsAccessKeyId gets a reference to the given string and assigns it to the AwsAccessKeyId field.
func (o *GatewayUpdateProducerAws) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = &v
}

// GetAwsAccessSecretKey returns the AwsAccessSecretKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsAccessSecretKey() string {
	if o == nil || o.AwsAccessSecretKey == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessSecretKey
}

// GetAwsAccessSecretKeyOk returns a tuple with the AwsAccessSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsAccessSecretKeyOk() (*string, bool) {
	if o == nil || o.AwsAccessSecretKey == nil {
		return nil, false
	}
	return o.AwsAccessSecretKey, true
}

// HasAwsAccessSecretKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsAccessSecretKey() bool {
	if o != nil && o.AwsAccessSecretKey != nil {
		return true
	}

	return false
}

// SetAwsAccessSecretKey gets a reference to the given string and assigns it to the AwsAccessSecretKey field.
func (o *GatewayUpdateProducerAws) SetAwsAccessSecretKey(v string) {
	o.AwsAccessSecretKey = &v
}

// GetAwsRoleArns returns the AwsRoleArns field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsRoleArns() string {
	if o == nil || o.AwsRoleArns == nil {
		var ret string
		return ret
	}
	return *o.AwsRoleArns
}

// GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsRoleArnsOk() (*string, bool) {
	if o == nil || o.AwsRoleArns == nil {
		return nil, false
	}
	return o.AwsRoleArns, true
}

// HasAwsRoleArns returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsRoleArns() bool {
	if o != nil && o.AwsRoleArns != nil {
		return true
	}

	return false
}

// SetAwsRoleArns gets a reference to the given string and assigns it to the AwsRoleArns field.
func (o *GatewayUpdateProducerAws) SetAwsRoleArns(v string) {
	o.AwsRoleArns = &v
}

// GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsUserConsoleAccess() bool {
	if o == nil || o.AwsUserConsoleAccess == nil {
		var ret bool
		return ret
	}
	return *o.AwsUserConsoleAccess
}

// GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsUserConsoleAccessOk() (*bool, bool) {
	if o == nil || o.AwsUserConsoleAccess == nil {
		return nil, false
	}
	return o.AwsUserConsoleAccess, true
}

// HasAwsUserConsoleAccess returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsUserConsoleAccess() bool {
	if o != nil && o.AwsUserConsoleAccess != nil {
		return true
	}

	return false
}

// SetAwsUserConsoleAccess gets a reference to the given bool and assigns it to the AwsUserConsoleAccess field.
func (o *GatewayUpdateProducerAws) SetAwsUserConsoleAccess(v bool) {
	o.AwsUserConsoleAccess = &v
}

// GetAwsUserGroups returns the AwsUserGroups field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsUserGroups() string {
	if o == nil || o.AwsUserGroups == nil {
		var ret string
		return ret
	}
	return *o.AwsUserGroups
}

// GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsUserGroupsOk() (*string, bool) {
	if o == nil || o.AwsUserGroups == nil {
		return nil, false
	}
	return o.AwsUserGroups, true
}

// HasAwsUserGroups returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsUserGroups() bool {
	if o != nil && o.AwsUserGroups != nil {
		return true
	}

	return false
}

// SetAwsUserGroups gets a reference to the given string and assigns it to the AwsUserGroups field.
func (o *GatewayUpdateProducerAws) SetAwsUserGroups(v string) {
	o.AwsUserGroups = &v
}

// GetAwsUserPolicies returns the AwsUserPolicies field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsUserPolicies() string {
	if o == nil || o.AwsUserPolicies == nil {
		var ret string
		return ret
	}
	return *o.AwsUserPolicies
}

// GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsUserPoliciesOk() (*string, bool) {
	if o == nil || o.AwsUserPolicies == nil {
		return nil, false
	}
	return o.AwsUserPolicies, true
}

// HasAwsUserPolicies returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsUserPolicies() bool {
	if o != nil && o.AwsUserPolicies != nil {
		return true
	}

	return false
}

// SetAwsUserPolicies gets a reference to the given string and assigns it to the AwsUserPolicies field.
func (o *GatewayUpdateProducerAws) SetAwsUserPolicies(v string) {
	o.AwsUserPolicies = &v
}

// GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetAwsUserProgrammaticAccess() bool {
	if o == nil || o.AwsUserProgrammaticAccess == nil {
		var ret bool
		return ret
	}
	return *o.AwsUserProgrammaticAccess
}

// GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetAwsUserProgrammaticAccessOk() (*bool, bool) {
	if o == nil || o.AwsUserProgrammaticAccess == nil {
		return nil, false
	}
	return o.AwsUserProgrammaticAccess, true
}

// HasAwsUserProgrammaticAccess returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasAwsUserProgrammaticAccess() bool {
	if o != nil && o.AwsUserProgrammaticAccess != nil {
		return true
	}

	return false
}

// SetAwsUserProgrammaticAccess gets a reference to the given bool and assigns it to the AwsUserProgrammaticAccess field.
func (o *GatewayUpdateProducerAws) SetAwsUserProgrammaticAccess(v bool) {
	o.AwsUserProgrammaticAccess = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerAws) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *GatewayUpdateProducerAws) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerAws) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerAws) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerAws) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerAws) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GatewayUpdateProducerAws) SetRegion(v string) {
	o.Region = &v
}

// GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessAwsAccountId() string {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessAwsAccountId
}

// GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessAwsAccountIdOk() (*string, bool) {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		return nil, false
	}
	return o.SecureAccessAwsAccountId, true
}

// HasSecureAccessAwsAccountId returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessAwsAccountId() bool {
	if o != nil && o.SecureAccessAwsAccountId != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsAccountId gets a reference to the given string and assigns it to the SecureAccessAwsAccountId field.
func (o *GatewayUpdateProducerAws) SetSecureAccessAwsAccountId(v string) {
	o.SecureAccessAwsAccountId = &v
}

// GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessAwsNativeCli() bool {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAwsNativeCli
}

// GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessAwsNativeCliOk() (*bool, bool) {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		return nil, false
	}
	return o.SecureAccessAwsNativeCli, true
}

// HasSecureAccessAwsNativeCli returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessAwsNativeCli() bool {
	if o != nil && o.SecureAccessAwsNativeCli != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsNativeCli gets a reference to the given bool and assigns it to the SecureAccessAwsNativeCli field.
func (o *GatewayUpdateProducerAws) SetSecureAccessAwsNativeCli(v bool) {
	o.SecureAccessAwsNativeCli = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayUpdateProducerAws) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerAws) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerAws) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *GatewayUpdateProducerAws) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *GatewayUpdateProducerAws) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerAws) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerAws) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerAws) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerAws) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAws) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAws) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAws) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerAws) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerAws) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessMode != nil {
		toSerialize["access-mode"] = o.AccessMode
	}
	if o.AdminRotationIntervalDays != nil {
		toSerialize["admin-rotation-interval-days"] = o.AdminRotationIntervalDays
	}
	if o.AwsAccessKeyId != nil {
		toSerialize["aws-access-key-id"] = o.AwsAccessKeyId
	}
	if o.AwsAccessSecretKey != nil {
		toSerialize["aws-access-secret-key"] = o.AwsAccessSecretKey
	}
	if o.AwsRoleArns != nil {
		toSerialize["aws-role-arns"] = o.AwsRoleArns
	}
	if o.AwsUserConsoleAccess != nil {
		toSerialize["aws-user-console-access"] = o.AwsUserConsoleAccess
	}
	if o.AwsUserGroups != nil {
		toSerialize["aws-user-groups"] = o.AwsUserGroups
	}
	if o.AwsUserPolicies != nil {
		toSerialize["aws-user-policies"] = o.AwsUserPolicies
	}
	if o.AwsUserProgrammaticAccess != nil {
		toSerialize["aws-user-programmatic-access"] = o.AwsUserProgrammaticAccess
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.EnableAdminRotation != nil {
		toSerialize["enable-admin-rotation"] = o.EnableAdminRotation
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
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.SecureAccessAwsAccountId != nil {
		toSerialize["secure-access-aws-account-id"] = o.SecureAccessAwsAccountId
	}
	if o.SecureAccessAwsNativeCli != nil {
		toSerialize["secure-access-aws-native-cli"] = o.SecureAccessAwsNativeCli
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
	}
	if o.SecureAccessWebBrowsing != nil {
		toSerialize["secure-access-web-browsing"] = o.SecureAccessWebBrowsing
	}
	if o.SecureAccessWebProxy != nil {
		toSerialize["secure-access-web-proxy"] = o.SecureAccessWebProxy
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

type NullableGatewayUpdateProducerAws struct {
	value *GatewayUpdateProducerAws
	isSet bool
}

func (v NullableGatewayUpdateProducerAws) Get() *GatewayUpdateProducerAws {
	return v.value
}

func (v *NullableGatewayUpdateProducerAws) Set(val *GatewayUpdateProducerAws) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerAws) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerAws(val *GatewayUpdateProducerAws) *NullableGatewayUpdateProducerAws {
	return &NullableGatewayUpdateProducerAws{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

