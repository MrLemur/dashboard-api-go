/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{}

// UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner struct for UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner
type UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner struct {
	Group *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup `json:"group,omitempty"`
	Milestones *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones `json:"milestones,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerWithDefaults() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) GetGroup() UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup {
	if o == nil || IsNil(o.Group) {
		var ret UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) GetGroupOk() (*UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup and assigns it to the Group field.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) SetGroup(v UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) {
	o.Group = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) GetMilestones() UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	if o == nil || IsNil(o.Milestones) {
		var ret UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) GetMilestonesOk() (*UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones, bool) {
	if o == nil || IsNil(o.Milestones) {
		return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) HasMilestones() bool {
	if o != nil && !IsNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones and assigns it to the Milestones field.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) SetMilestones(v UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) {
	o.Milestones = &v
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Milestones) {
		toSerialize["milestones"] = o.Milestones
	}
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner struct {
	value *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) Get() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) Set(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	return &NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


