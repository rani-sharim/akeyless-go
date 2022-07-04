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

// MigrationStatusReplyObj struct for MigrationStatusReplyObj
type MigrationStatusReplyObj struct {
	DurationTime *string `json:"duration_time,omitempty"`
	LastStatusMessage *string `json:"last_status_message,omitempty"`
	MaxNameLength *int64 `json:"max_name_length,omitempty"`
	MaxValueLength *int64 `json:"max_value_length,omitempty"`
	MigrationId *string `json:"migration_id,omitempty"`
	MigrationItems *MigrationItems `json:"migration_items,omitempty"`
	MigrationName *string `json:"migration_name,omitempty"`
	MigrationState *string `json:"migration_state,omitempty"`
	MigrationType *string `json:"migration_type,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
}

// NewMigrationStatusReplyObj instantiates a new MigrationStatusReplyObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationStatusReplyObj() *MigrationStatusReplyObj {
	this := MigrationStatusReplyObj{}
	return &this
}

// NewMigrationStatusReplyObjWithDefaults instantiates a new MigrationStatusReplyObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationStatusReplyObjWithDefaults() *MigrationStatusReplyObj {
	this := MigrationStatusReplyObj{}
	return &this
}

// GetDurationTime returns the DurationTime field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetDurationTime() string {
	if o == nil || o.DurationTime == nil {
		var ret string
		return ret
	}
	return *o.DurationTime
}

// GetDurationTimeOk returns a tuple with the DurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetDurationTimeOk() (*string, bool) {
	if o == nil || o.DurationTime == nil {
		return nil, false
	}
	return o.DurationTime, true
}

// HasDurationTime returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasDurationTime() bool {
	if o != nil && o.DurationTime != nil {
		return true
	}

	return false
}

// SetDurationTime gets a reference to the given string and assigns it to the DurationTime field.
func (o *MigrationStatusReplyObj) SetDurationTime(v string) {
	o.DurationTime = &v
}

// GetLastStatusMessage returns the LastStatusMessage field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetLastStatusMessage() string {
	if o == nil || o.LastStatusMessage == nil {
		var ret string
		return ret
	}
	return *o.LastStatusMessage
}

// GetLastStatusMessageOk returns a tuple with the LastStatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetLastStatusMessageOk() (*string, bool) {
	if o == nil || o.LastStatusMessage == nil {
		return nil, false
	}
	return o.LastStatusMessage, true
}

// HasLastStatusMessage returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasLastStatusMessage() bool {
	if o != nil && o.LastStatusMessage != nil {
		return true
	}

	return false
}

// SetLastStatusMessage gets a reference to the given string and assigns it to the LastStatusMessage field.
func (o *MigrationStatusReplyObj) SetLastStatusMessage(v string) {
	o.LastStatusMessage = &v
}

// GetMaxNameLength returns the MaxNameLength field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMaxNameLength() int64 {
	if o == nil || o.MaxNameLength == nil {
		var ret int64
		return ret
	}
	return *o.MaxNameLength
}

// GetMaxNameLengthOk returns a tuple with the MaxNameLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMaxNameLengthOk() (*int64, bool) {
	if o == nil || o.MaxNameLength == nil {
		return nil, false
	}
	return o.MaxNameLength, true
}

// HasMaxNameLength returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMaxNameLength() bool {
	if o != nil && o.MaxNameLength != nil {
		return true
	}

	return false
}

// SetMaxNameLength gets a reference to the given int64 and assigns it to the MaxNameLength field.
func (o *MigrationStatusReplyObj) SetMaxNameLength(v int64) {
	o.MaxNameLength = &v
}

// GetMaxValueLength returns the MaxValueLength field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMaxValueLength() int64 {
	if o == nil || o.MaxValueLength == nil {
		var ret int64
		return ret
	}
	return *o.MaxValueLength
}

// GetMaxValueLengthOk returns a tuple with the MaxValueLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMaxValueLengthOk() (*int64, bool) {
	if o == nil || o.MaxValueLength == nil {
		return nil, false
	}
	return o.MaxValueLength, true
}

// HasMaxValueLength returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMaxValueLength() bool {
	if o != nil && o.MaxValueLength != nil {
		return true
	}

	return false
}

// SetMaxValueLength gets a reference to the given int64 and assigns it to the MaxValueLength field.
func (o *MigrationStatusReplyObj) SetMaxValueLength(v int64) {
	o.MaxValueLength = &v
}

// GetMigrationId returns the MigrationId field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMigrationId() string {
	if o == nil || o.MigrationId == nil {
		var ret string
		return ret
	}
	return *o.MigrationId
}

// GetMigrationIdOk returns a tuple with the MigrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMigrationIdOk() (*string, bool) {
	if o == nil || o.MigrationId == nil {
		return nil, false
	}
	return o.MigrationId, true
}

// HasMigrationId returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMigrationId() bool {
	if o != nil && o.MigrationId != nil {
		return true
	}

	return false
}

// SetMigrationId gets a reference to the given string and assigns it to the MigrationId field.
func (o *MigrationStatusReplyObj) SetMigrationId(v string) {
	o.MigrationId = &v
}

// GetMigrationItems returns the MigrationItems field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMigrationItems() MigrationItems {
	if o == nil || o.MigrationItems == nil {
		var ret MigrationItems
		return ret
	}
	return *o.MigrationItems
}

// GetMigrationItemsOk returns a tuple with the MigrationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMigrationItemsOk() (*MigrationItems, bool) {
	if o == nil || o.MigrationItems == nil {
		return nil, false
	}
	return o.MigrationItems, true
}

// HasMigrationItems returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMigrationItems() bool {
	if o != nil && o.MigrationItems != nil {
		return true
	}

	return false
}

// SetMigrationItems gets a reference to the given MigrationItems and assigns it to the MigrationItems field.
func (o *MigrationStatusReplyObj) SetMigrationItems(v MigrationItems) {
	o.MigrationItems = &v
}

// GetMigrationName returns the MigrationName field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMigrationName() string {
	if o == nil || o.MigrationName == nil {
		var ret string
		return ret
	}
	return *o.MigrationName
}

// GetMigrationNameOk returns a tuple with the MigrationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMigrationNameOk() (*string, bool) {
	if o == nil || o.MigrationName == nil {
		return nil, false
	}
	return o.MigrationName, true
}

// HasMigrationName returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMigrationName() bool {
	if o != nil && o.MigrationName != nil {
		return true
	}

	return false
}

// SetMigrationName gets a reference to the given string and assigns it to the MigrationName field.
func (o *MigrationStatusReplyObj) SetMigrationName(v string) {
	o.MigrationName = &v
}

// GetMigrationState returns the MigrationState field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMigrationState() string {
	if o == nil || o.MigrationState == nil {
		var ret string
		return ret
	}
	return *o.MigrationState
}

// GetMigrationStateOk returns a tuple with the MigrationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMigrationStateOk() (*string, bool) {
	if o == nil || o.MigrationState == nil {
		return nil, false
	}
	return o.MigrationState, true
}

// HasMigrationState returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMigrationState() bool {
	if o != nil && o.MigrationState != nil {
		return true
	}

	return false
}

// SetMigrationState gets a reference to the given string and assigns it to the MigrationState field.
func (o *MigrationStatusReplyObj) SetMigrationState(v string) {
	o.MigrationState = &v
}

// GetMigrationType returns the MigrationType field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetMigrationType() string {
	if o == nil || o.MigrationType == nil {
		var ret string
		return ret
	}
	return *o.MigrationType
}

// GetMigrationTypeOk returns a tuple with the MigrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetMigrationTypeOk() (*string, bool) {
	if o == nil || o.MigrationType == nil {
		return nil, false
	}
	return o.MigrationType, true
}

// HasMigrationType returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasMigrationType() bool {
	if o != nil && o.MigrationType != nil {
		return true
	}

	return false
}

// SetMigrationType gets a reference to the given string and assigns it to the MigrationType field.
func (o *MigrationStatusReplyObj) SetMigrationType(v string) {
	o.MigrationType = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MigrationStatusReplyObj) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatusReplyObj) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MigrationStatusReplyObj) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *MigrationStatusReplyObj) SetStartTime(v string) {
	o.StartTime = &v
}

func (o MigrationStatusReplyObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DurationTime != nil {
		toSerialize["duration_time"] = o.DurationTime
	}
	if o.LastStatusMessage != nil {
		toSerialize["last_status_message"] = o.LastStatusMessage
	}
	if o.MaxNameLength != nil {
		toSerialize["max_name_length"] = o.MaxNameLength
	}
	if o.MaxValueLength != nil {
		toSerialize["max_value_length"] = o.MaxValueLength
	}
	if o.MigrationId != nil {
		toSerialize["migration_id"] = o.MigrationId
	}
	if o.MigrationItems != nil {
		toSerialize["migration_items"] = o.MigrationItems
	}
	if o.MigrationName != nil {
		toSerialize["migration_name"] = o.MigrationName
	}
	if o.MigrationState != nil {
		toSerialize["migration_state"] = o.MigrationState
	}
	if o.MigrationType != nil {
		toSerialize["migration_type"] = o.MigrationType
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableMigrationStatusReplyObj struct {
	value *MigrationStatusReplyObj
	isSet bool
}

func (v NullableMigrationStatusReplyObj) Get() *MigrationStatusReplyObj {
	return v.value
}

func (v *NullableMigrationStatusReplyObj) Set(val *MigrationStatusReplyObj) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationStatusReplyObj) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationStatusReplyObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationStatusReplyObj(val *MigrationStatusReplyObj) *NullableMigrationStatusReplyObj {
	return &NullableMigrationStatusReplyObj{value: val, isSet: true}
}

func (v NullableMigrationStatusReplyObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationStatusReplyObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


