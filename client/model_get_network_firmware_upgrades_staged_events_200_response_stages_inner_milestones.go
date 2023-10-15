/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones{}

// GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones The Staged Upgrade Milestones for the stage
type GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones struct {
	// Scheduled start time for the group
	ScheduledFor *time.Time `json:"scheduledFor,omitempty"`
	// Start time for the group
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Finish time for the group
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Time that the group was canceled
	CanceledAt *time.Time `json:"canceledAt,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestonesWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestonesWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetScheduledFor() time.Time {
	if o == nil || IsNil(o.ScheduledFor) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetScheduledForOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledFor) {
		return nil, false
	}
	return o.ScheduledFor, true
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) HasScheduledFor() bool {
	if o != nil && !IsNil(o.ScheduledFor) {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given time.Time and assigns it to the ScheduledFor field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) SetScheduledFor(v time.Time) {
	o.ScheduledFor = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CanceledAt) {
		return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) HasCanceledAt() bool {
	if o != nil && !IsNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScheduledFor) {
		toSerialize["scheduledFor"] = o.ScheduledFor
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !IsNil(o.CanceledAt) {
		toSerialize["canceledAt"] = o.CanceledAt
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


