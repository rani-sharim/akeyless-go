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

// Item struct for Item
type Item struct {
	AutoRotate *bool `json:"auto_rotate,omitempty"`
	CertIssuerSignerKeyName *string `json:"cert_issuer_signer_key_name,omitempty"`
	CertificateIssueDetails *CertificateIssueInfo `json:"certificate_issue_details,omitempty"`
	Certificates *string `json:"certificates,omitempty"`
	ClientPermissions []string `json:"client_permissions,omitempty"`
	CustomerFragmentId *string `json:"customer_fragment_id,omitempty"`
	DeleteProtection *bool `json:"delete_protection,omitempty"`
	DeletionDate *time.Time `json:"deletion_date,omitempty"`
	DisplayId *string `json:"display_id,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	ItemAccessibility *int64 `json:"item_accessibility,omitempty"`
	ItemGeneralInfo *ItemGeneralInfo `json:"item_general_info,omitempty"`
	ItemId *int64 `json:"item_id,omitempty"`
	ItemMetadata *string `json:"item_metadata,omitempty"`
	ItemName *string `json:"item_name,omitempty"`
	ItemSize *int64 `json:"item_size,omitempty"`
	// ItemState defines the different states an Item can be in
	ItemState *string `json:"item_state,omitempty"`
	ItemSubType *string `json:"item_sub_type,omitempty"`
	ItemTags []string `json:"item_tags,omitempty"`
	ItemTargetsAssoc []ItemTargetAssociation `json:"item_targets_assoc,omitempty"`
	ItemType *string `json:"item_type,omitempty"`
	ItemVersions []ItemVersion `json:"item_versions,omitempty"`
	LastVersion *int32 `json:"last_version,omitempty"`
	NextRotationDate *time.Time `json:"next_rotation_date,omitempty"`
	ProtectionKeyName *string `json:"protection_key_name,omitempty"`
	ProtectionKeyType *string `json:"protection_key_type,omitempty"`
	PublicValue *string `json:"public_value,omitempty"`
	RotationInterval *int64 `json:"rotation_interval,omitempty"`
	TargetVersions []TargetItemVersion `json:"target_versions,omitempty"`
	WithCustomerFragment *bool `json:"with_customer_fragment,omitempty"`
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem() *Item {
	this := Item{}
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *Item) GetAutoRotate() bool {
	if o == nil || o.AutoRotate == nil {
		var ret bool
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetAutoRotateOk() (*bool, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *Item) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given bool and assigns it to the AutoRotate field.
func (o *Item) SetAutoRotate(v bool) {
	o.AutoRotate = &v
}

// GetCertIssuerSignerKeyName returns the CertIssuerSignerKeyName field value if set, zero value otherwise.
func (o *Item) GetCertIssuerSignerKeyName() string {
	if o == nil || o.CertIssuerSignerKeyName == nil {
		var ret string
		return ret
	}
	return *o.CertIssuerSignerKeyName
}

// GetCertIssuerSignerKeyNameOk returns a tuple with the CertIssuerSignerKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetCertIssuerSignerKeyNameOk() (*string, bool) {
	if o == nil || o.CertIssuerSignerKeyName == nil {
		return nil, false
	}
	return o.CertIssuerSignerKeyName, true
}

// HasCertIssuerSignerKeyName returns a boolean if a field has been set.
func (o *Item) HasCertIssuerSignerKeyName() bool {
	if o != nil && o.CertIssuerSignerKeyName != nil {
		return true
	}

	return false
}

// SetCertIssuerSignerKeyName gets a reference to the given string and assigns it to the CertIssuerSignerKeyName field.
func (o *Item) SetCertIssuerSignerKeyName(v string) {
	o.CertIssuerSignerKeyName = &v
}

// GetCertificateIssueDetails returns the CertificateIssueDetails field value if set, zero value otherwise.
func (o *Item) GetCertificateIssueDetails() CertificateIssueInfo {
	if o == nil || o.CertificateIssueDetails == nil {
		var ret CertificateIssueInfo
		return ret
	}
	return *o.CertificateIssueDetails
}

// GetCertificateIssueDetailsOk returns a tuple with the CertificateIssueDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetCertificateIssueDetailsOk() (*CertificateIssueInfo, bool) {
	if o == nil || o.CertificateIssueDetails == nil {
		return nil, false
	}
	return o.CertificateIssueDetails, true
}

// HasCertificateIssueDetails returns a boolean if a field has been set.
func (o *Item) HasCertificateIssueDetails() bool {
	if o != nil && o.CertificateIssueDetails != nil {
		return true
	}

	return false
}

// SetCertificateIssueDetails gets a reference to the given CertificateIssueInfo and assigns it to the CertificateIssueDetails field.
func (o *Item) SetCertificateIssueDetails(v CertificateIssueInfo) {
	o.CertificateIssueDetails = &v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *Item) GetCertificates() string {
	if o == nil || o.Certificates == nil {
		var ret string
		return ret
	}
	return *o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetCertificatesOk() (*string, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *Item) HasCertificates() bool {
	if o != nil && o.Certificates != nil {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given string and assigns it to the Certificates field.
func (o *Item) SetCertificates(v string) {
	o.Certificates = &v
}

// GetClientPermissions returns the ClientPermissions field value if set, zero value otherwise.
func (o *Item) GetClientPermissions() []string {
	if o == nil || o.ClientPermissions == nil {
		var ret []string
		return ret
	}
	return o.ClientPermissions
}

// GetClientPermissionsOk returns a tuple with the ClientPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetClientPermissionsOk() ([]string, bool) {
	if o == nil || o.ClientPermissions == nil {
		return nil, false
	}
	return o.ClientPermissions, true
}

// HasClientPermissions returns a boolean if a field has been set.
func (o *Item) HasClientPermissions() bool {
	if o != nil && o.ClientPermissions != nil {
		return true
	}

	return false
}

// SetClientPermissions gets a reference to the given []string and assigns it to the ClientPermissions field.
func (o *Item) SetClientPermissions(v []string) {
	o.ClientPermissions = v
}

// GetCustomerFragmentId returns the CustomerFragmentId field value if set, zero value otherwise.
func (o *Item) GetCustomerFragmentId() string {
	if o == nil || o.CustomerFragmentId == nil {
		var ret string
		return ret
	}
	return *o.CustomerFragmentId
}

// GetCustomerFragmentIdOk returns a tuple with the CustomerFragmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetCustomerFragmentIdOk() (*string, bool) {
	if o == nil || o.CustomerFragmentId == nil {
		return nil, false
	}
	return o.CustomerFragmentId, true
}

// HasCustomerFragmentId returns a boolean if a field has been set.
func (o *Item) HasCustomerFragmentId() bool {
	if o != nil && o.CustomerFragmentId != nil {
		return true
	}

	return false
}

// SetCustomerFragmentId gets a reference to the given string and assigns it to the CustomerFragmentId field.
func (o *Item) SetCustomerFragmentId(v string) {
	o.CustomerFragmentId = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *Item) GetDeleteProtection() bool {
	if o == nil || o.DeleteProtection == nil {
		var ret bool
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDeleteProtectionOk() (*bool, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *Item) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given bool and assigns it to the DeleteProtection field.
func (o *Item) SetDeleteProtection(v bool) {
	o.DeleteProtection = &v
}

// GetDeletionDate returns the DeletionDate field value if set, zero value otherwise.
func (o *Item) GetDeletionDate() time.Time {
	if o == nil || o.DeletionDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionDate
}

// GetDeletionDateOk returns a tuple with the DeletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDeletionDateOk() (*time.Time, bool) {
	if o == nil || o.DeletionDate == nil {
		return nil, false
	}
	return o.DeletionDate, true
}

// HasDeletionDate returns a boolean if a field has been set.
func (o *Item) HasDeletionDate() bool {
	if o != nil && o.DeletionDate != nil {
		return true
	}

	return false
}

// SetDeletionDate gets a reference to the given time.Time and assigns it to the DeletionDate field.
func (o *Item) SetDeletionDate(v time.Time) {
	o.DeletionDate = &v
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *Item) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *Item) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *Item) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *Item) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *Item) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *Item) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetItemAccessibility returns the ItemAccessibility field value if set, zero value otherwise.
func (o *Item) GetItemAccessibility() int64 {
	if o == nil || o.ItemAccessibility == nil {
		var ret int64
		return ret
	}
	return *o.ItemAccessibility
}

// GetItemAccessibilityOk returns a tuple with the ItemAccessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemAccessibilityOk() (*int64, bool) {
	if o == nil || o.ItemAccessibility == nil {
		return nil, false
	}
	return o.ItemAccessibility, true
}

// HasItemAccessibility returns a boolean if a field has been set.
func (o *Item) HasItemAccessibility() bool {
	if o != nil && o.ItemAccessibility != nil {
		return true
	}

	return false
}

// SetItemAccessibility gets a reference to the given int64 and assigns it to the ItemAccessibility field.
func (o *Item) SetItemAccessibility(v int64) {
	o.ItemAccessibility = &v
}

// GetItemGeneralInfo returns the ItemGeneralInfo field value if set, zero value otherwise.
func (o *Item) GetItemGeneralInfo() ItemGeneralInfo {
	if o == nil || o.ItemGeneralInfo == nil {
		var ret ItemGeneralInfo
		return ret
	}
	return *o.ItemGeneralInfo
}

// GetItemGeneralInfoOk returns a tuple with the ItemGeneralInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemGeneralInfoOk() (*ItemGeneralInfo, bool) {
	if o == nil || o.ItemGeneralInfo == nil {
		return nil, false
	}
	return o.ItemGeneralInfo, true
}

// HasItemGeneralInfo returns a boolean if a field has been set.
func (o *Item) HasItemGeneralInfo() bool {
	if o != nil && o.ItemGeneralInfo != nil {
		return true
	}

	return false
}

// SetItemGeneralInfo gets a reference to the given ItemGeneralInfo and assigns it to the ItemGeneralInfo field.
func (o *Item) SetItemGeneralInfo(v ItemGeneralInfo) {
	o.ItemGeneralInfo = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *Item) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *Item) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *Item) SetItemId(v int64) {
	o.ItemId = &v
}

// GetItemMetadata returns the ItemMetadata field value if set, zero value otherwise.
func (o *Item) GetItemMetadata() string {
	if o == nil || o.ItemMetadata == nil {
		var ret string
		return ret
	}
	return *o.ItemMetadata
}

// GetItemMetadataOk returns a tuple with the ItemMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemMetadataOk() (*string, bool) {
	if o == nil || o.ItemMetadata == nil {
		return nil, false
	}
	return o.ItemMetadata, true
}

// HasItemMetadata returns a boolean if a field has been set.
func (o *Item) HasItemMetadata() bool {
	if o != nil && o.ItemMetadata != nil {
		return true
	}

	return false
}

// SetItemMetadata gets a reference to the given string and assigns it to the ItemMetadata field.
func (o *Item) SetItemMetadata(v string) {
	o.ItemMetadata = &v
}

// GetItemName returns the ItemName field value if set, zero value otherwise.
func (o *Item) GetItemName() string {
	if o == nil || o.ItemName == nil {
		var ret string
		return ret
	}
	return *o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemNameOk() (*string, bool) {
	if o == nil || o.ItemName == nil {
		return nil, false
	}
	return o.ItemName, true
}

// HasItemName returns a boolean if a field has been set.
func (o *Item) HasItemName() bool {
	if o != nil && o.ItemName != nil {
		return true
	}

	return false
}

// SetItemName gets a reference to the given string and assigns it to the ItemName field.
func (o *Item) SetItemName(v string) {
	o.ItemName = &v
}

// GetItemSize returns the ItemSize field value if set, zero value otherwise.
func (o *Item) GetItemSize() int64 {
	if o == nil || o.ItemSize == nil {
		var ret int64
		return ret
	}
	return *o.ItemSize
}

// GetItemSizeOk returns a tuple with the ItemSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemSizeOk() (*int64, bool) {
	if o == nil || o.ItemSize == nil {
		return nil, false
	}
	return o.ItemSize, true
}

// HasItemSize returns a boolean if a field has been set.
func (o *Item) HasItemSize() bool {
	if o != nil && o.ItemSize != nil {
		return true
	}

	return false
}

// SetItemSize gets a reference to the given int64 and assigns it to the ItemSize field.
func (o *Item) SetItemSize(v int64) {
	o.ItemSize = &v
}

// GetItemState returns the ItemState field value if set, zero value otherwise.
func (o *Item) GetItemState() string {
	if o == nil || o.ItemState == nil {
		var ret string
		return ret
	}
	return *o.ItemState
}

// GetItemStateOk returns a tuple with the ItemState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemStateOk() (*string, bool) {
	if o == nil || o.ItemState == nil {
		return nil, false
	}
	return o.ItemState, true
}

// HasItemState returns a boolean if a field has been set.
func (o *Item) HasItemState() bool {
	if o != nil && o.ItemState != nil {
		return true
	}

	return false
}

// SetItemState gets a reference to the given string and assigns it to the ItemState field.
func (o *Item) SetItemState(v string) {
	o.ItemState = &v
}

// GetItemSubType returns the ItemSubType field value if set, zero value otherwise.
func (o *Item) GetItemSubType() string {
	if o == nil || o.ItemSubType == nil {
		var ret string
		return ret
	}
	return *o.ItemSubType
}

// GetItemSubTypeOk returns a tuple with the ItemSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemSubTypeOk() (*string, bool) {
	if o == nil || o.ItemSubType == nil {
		return nil, false
	}
	return o.ItemSubType, true
}

// HasItemSubType returns a boolean if a field has been set.
func (o *Item) HasItemSubType() bool {
	if o != nil && o.ItemSubType != nil {
		return true
	}

	return false
}

// SetItemSubType gets a reference to the given string and assigns it to the ItemSubType field.
func (o *Item) SetItemSubType(v string) {
	o.ItemSubType = &v
}

// GetItemTags returns the ItemTags field value if set, zero value otherwise.
func (o *Item) GetItemTags() []string {
	if o == nil || o.ItemTags == nil {
		var ret []string
		return ret
	}
	return o.ItemTags
}

// GetItemTagsOk returns a tuple with the ItemTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemTagsOk() ([]string, bool) {
	if o == nil || o.ItemTags == nil {
		return nil, false
	}
	return o.ItemTags, true
}

// HasItemTags returns a boolean if a field has been set.
func (o *Item) HasItemTags() bool {
	if o != nil && o.ItemTags != nil {
		return true
	}

	return false
}

// SetItemTags gets a reference to the given []string and assigns it to the ItemTags field.
func (o *Item) SetItemTags(v []string) {
	o.ItemTags = v
}

// GetItemTargetsAssoc returns the ItemTargetsAssoc field value if set, zero value otherwise.
func (o *Item) GetItemTargetsAssoc() []ItemTargetAssociation {
	if o == nil || o.ItemTargetsAssoc == nil {
		var ret []ItemTargetAssociation
		return ret
	}
	return o.ItemTargetsAssoc
}

// GetItemTargetsAssocOk returns a tuple with the ItemTargetsAssoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemTargetsAssocOk() ([]ItemTargetAssociation, bool) {
	if o == nil || o.ItemTargetsAssoc == nil {
		return nil, false
	}
	return o.ItemTargetsAssoc, true
}

// HasItemTargetsAssoc returns a boolean if a field has been set.
func (o *Item) HasItemTargetsAssoc() bool {
	if o != nil && o.ItemTargetsAssoc != nil {
		return true
	}

	return false
}

// SetItemTargetsAssoc gets a reference to the given []ItemTargetAssociation and assigns it to the ItemTargetsAssoc field.
func (o *Item) SetItemTargetsAssoc(v []ItemTargetAssociation) {
	o.ItemTargetsAssoc = v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *Item) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *Item) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *Item) SetItemType(v string) {
	o.ItemType = &v
}

// GetItemVersions returns the ItemVersions field value if set, zero value otherwise.
func (o *Item) GetItemVersions() []ItemVersion {
	if o == nil || o.ItemVersions == nil {
		var ret []ItemVersion
		return ret
	}
	return o.ItemVersions
}

// GetItemVersionsOk returns a tuple with the ItemVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemVersionsOk() ([]ItemVersion, bool) {
	if o == nil || o.ItemVersions == nil {
		return nil, false
	}
	return o.ItemVersions, true
}

// HasItemVersions returns a boolean if a field has been set.
func (o *Item) HasItemVersions() bool {
	if o != nil && o.ItemVersions != nil {
		return true
	}

	return false
}

// SetItemVersions gets a reference to the given []ItemVersion and assigns it to the ItemVersions field.
func (o *Item) SetItemVersions(v []ItemVersion) {
	o.ItemVersions = v
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *Item) GetLastVersion() int32 {
	if o == nil || o.LastVersion == nil {
		var ret int32
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetLastVersionOk() (*int32, bool) {
	if o == nil || o.LastVersion == nil {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *Item) HasLastVersion() bool {
	if o != nil && o.LastVersion != nil {
		return true
	}

	return false
}

// SetLastVersion gets a reference to the given int32 and assigns it to the LastVersion field.
func (o *Item) SetLastVersion(v int32) {
	o.LastVersion = &v
}

// GetNextRotationDate returns the NextRotationDate field value if set, zero value otherwise.
func (o *Item) GetNextRotationDate() time.Time {
	if o == nil || o.NextRotationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRotationDate
}

// GetNextRotationDateOk returns a tuple with the NextRotationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetNextRotationDateOk() (*time.Time, bool) {
	if o == nil || o.NextRotationDate == nil {
		return nil, false
	}
	return o.NextRotationDate, true
}

// HasNextRotationDate returns a boolean if a field has been set.
func (o *Item) HasNextRotationDate() bool {
	if o != nil && o.NextRotationDate != nil {
		return true
	}

	return false
}

// SetNextRotationDate gets a reference to the given time.Time and assigns it to the NextRotationDate field.
func (o *Item) SetNextRotationDate(v time.Time) {
	o.NextRotationDate = &v
}

// GetProtectionKeyName returns the ProtectionKeyName field value if set, zero value otherwise.
func (o *Item) GetProtectionKeyName() string {
	if o == nil || o.ProtectionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKeyName
}

// GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetProtectionKeyNameOk() (*string, bool) {
	if o == nil || o.ProtectionKeyName == nil {
		return nil, false
	}
	return o.ProtectionKeyName, true
}

// HasProtectionKeyName returns a boolean if a field has been set.
func (o *Item) HasProtectionKeyName() bool {
	if o != nil && o.ProtectionKeyName != nil {
		return true
	}

	return false
}

// SetProtectionKeyName gets a reference to the given string and assigns it to the ProtectionKeyName field.
func (o *Item) SetProtectionKeyName(v string) {
	o.ProtectionKeyName = &v
}

// GetProtectionKeyType returns the ProtectionKeyType field value if set, zero value otherwise.
func (o *Item) GetProtectionKeyType() string {
	if o == nil || o.ProtectionKeyType == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKeyType
}

// GetProtectionKeyTypeOk returns a tuple with the ProtectionKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetProtectionKeyTypeOk() (*string, bool) {
	if o == nil || o.ProtectionKeyType == nil {
		return nil, false
	}
	return o.ProtectionKeyType, true
}

// HasProtectionKeyType returns a boolean if a field has been set.
func (o *Item) HasProtectionKeyType() bool {
	if o != nil && o.ProtectionKeyType != nil {
		return true
	}

	return false
}

// SetProtectionKeyType gets a reference to the given string and assigns it to the ProtectionKeyType field.
func (o *Item) SetProtectionKeyType(v string) {
	o.ProtectionKeyType = &v
}

// GetPublicValue returns the PublicValue field value if set, zero value otherwise.
func (o *Item) GetPublicValue() string {
	if o == nil || o.PublicValue == nil {
		var ret string
		return ret
	}
	return *o.PublicValue
}

// GetPublicValueOk returns a tuple with the PublicValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetPublicValueOk() (*string, bool) {
	if o == nil || o.PublicValue == nil {
		return nil, false
	}
	return o.PublicValue, true
}

// HasPublicValue returns a boolean if a field has been set.
func (o *Item) HasPublicValue() bool {
	if o != nil && o.PublicValue != nil {
		return true
	}

	return false
}

// SetPublicValue gets a reference to the given string and assigns it to the PublicValue field.
func (o *Item) SetPublicValue(v string) {
	o.PublicValue = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *Item) GetRotationInterval() int64 {
	if o == nil || o.RotationInterval == nil {
		var ret int64
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetRotationIntervalOk() (*int64, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *Item) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given int64 and assigns it to the RotationInterval field.
func (o *Item) SetRotationInterval(v int64) {
	o.RotationInterval = &v
}

// GetTargetVersions returns the TargetVersions field value if set, zero value otherwise.
func (o *Item) GetTargetVersions() []TargetItemVersion {
	if o == nil || o.TargetVersions == nil {
		var ret []TargetItemVersion
		return ret
	}
	return o.TargetVersions
}

// GetTargetVersionsOk returns a tuple with the TargetVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetTargetVersionsOk() ([]TargetItemVersion, bool) {
	if o == nil || o.TargetVersions == nil {
		return nil, false
	}
	return o.TargetVersions, true
}

// HasTargetVersions returns a boolean if a field has been set.
func (o *Item) HasTargetVersions() bool {
	if o != nil && o.TargetVersions != nil {
		return true
	}

	return false
}

// SetTargetVersions gets a reference to the given []TargetItemVersion and assigns it to the TargetVersions field.
func (o *Item) SetTargetVersions(v []TargetItemVersion) {
	o.TargetVersions = v
}

// GetWithCustomerFragment returns the WithCustomerFragment field value if set, zero value otherwise.
func (o *Item) GetWithCustomerFragment() bool {
	if o == nil || o.WithCustomerFragment == nil {
		var ret bool
		return ret
	}
	return *o.WithCustomerFragment
}

// GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetWithCustomerFragmentOk() (*bool, bool) {
	if o == nil || o.WithCustomerFragment == nil {
		return nil, false
	}
	return o.WithCustomerFragment, true
}

// HasWithCustomerFragment returns a boolean if a field has been set.
func (o *Item) HasWithCustomerFragment() bool {
	if o != nil && o.WithCustomerFragment != nil {
		return true
	}

	return false
}

// SetWithCustomerFragment gets a reference to the given bool and assigns it to the WithCustomerFragment field.
func (o *Item) SetWithCustomerFragment(v bool) {
	o.WithCustomerFragment = &v
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoRotate != nil {
		toSerialize["auto_rotate"] = o.AutoRotate
	}
	if o.CertIssuerSignerKeyName != nil {
		toSerialize["cert_issuer_signer_key_name"] = o.CertIssuerSignerKeyName
	}
	if o.CertificateIssueDetails != nil {
		toSerialize["certificate_issue_details"] = o.CertificateIssueDetails
	}
	if o.Certificates != nil {
		toSerialize["certificates"] = o.Certificates
	}
	if o.ClientPermissions != nil {
		toSerialize["client_permissions"] = o.ClientPermissions
	}
	if o.CustomerFragmentId != nil {
		toSerialize["customer_fragment_id"] = o.CustomerFragmentId
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.DeletionDate != nil {
		toSerialize["deletion_date"] = o.DeletionDate
	}
	if o.DisplayId != nil {
		toSerialize["display_id"] = o.DisplayId
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.ItemAccessibility != nil {
		toSerialize["item_accessibility"] = o.ItemAccessibility
	}
	if o.ItemGeneralInfo != nil {
		toSerialize["item_general_info"] = o.ItemGeneralInfo
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	if o.ItemMetadata != nil {
		toSerialize["item_metadata"] = o.ItemMetadata
	}
	if o.ItemName != nil {
		toSerialize["item_name"] = o.ItemName
	}
	if o.ItemSize != nil {
		toSerialize["item_size"] = o.ItemSize
	}
	if o.ItemState != nil {
		toSerialize["item_state"] = o.ItemState
	}
	if o.ItemSubType != nil {
		toSerialize["item_sub_type"] = o.ItemSubType
	}
	if o.ItemTags != nil {
		toSerialize["item_tags"] = o.ItemTags
	}
	if o.ItemTargetsAssoc != nil {
		toSerialize["item_targets_assoc"] = o.ItemTargetsAssoc
	}
	if o.ItemType != nil {
		toSerialize["item_type"] = o.ItemType
	}
	if o.ItemVersions != nil {
		toSerialize["item_versions"] = o.ItemVersions
	}
	if o.LastVersion != nil {
		toSerialize["last_version"] = o.LastVersion
	}
	if o.NextRotationDate != nil {
		toSerialize["next_rotation_date"] = o.NextRotationDate
	}
	if o.ProtectionKeyName != nil {
		toSerialize["protection_key_name"] = o.ProtectionKeyName
	}
	if o.ProtectionKeyType != nil {
		toSerialize["protection_key_type"] = o.ProtectionKeyType
	}
	if o.PublicValue != nil {
		toSerialize["public_value"] = o.PublicValue
	}
	if o.RotationInterval != nil {
		toSerialize["rotation_interval"] = o.RotationInterval
	}
	if o.TargetVersions != nil {
		toSerialize["target_versions"] = o.TargetVersions
	}
	if o.WithCustomerFragment != nil {
		toSerialize["with_customer_fragment"] = o.WithCustomerFragment
	}
	return json.Marshal(toSerialize)
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

