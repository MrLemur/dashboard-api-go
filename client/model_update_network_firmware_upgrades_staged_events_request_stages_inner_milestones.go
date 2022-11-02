/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones The Staged Upgrade Milestones for the specific stage
type UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones struct {
	// The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)
	ScheduledFor string `json:"scheduledFor"`
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones(scheduledFor string) *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{}
	this.ScheduledFor = scheduledFor
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestonesWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestonesWithDefaults() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) GetScheduledFor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) GetScheduledForOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScheduledFor, true
}

// SetScheduledFor sets field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) SetScheduledFor(v string) {
	o.ScheduledFor = v
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scheduledFor"] = o.ScheduledFor
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones struct {
	value *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) Get() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) Set(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	return &NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


