/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject84 struct for InlineObject84
type InlineObject84 struct {
	// All completed or in-progress stages in the network with their new start times. All pending stages will be canceled
	Stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages `json:"stages"`
	// The reason for rolling back the staged upgrade
	Reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons `json:"reasons,omitempty"`
}

// NewInlineObject84 instantiates a new InlineObject84 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject84(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) *InlineObject84 {
	this := InlineObject84{}
	this.Stages = stages
	return &this
}

// NewInlineObject84WithDefaults instantiates a new InlineObject84 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject84WithDefaults() *InlineObject84 {
	this := InlineObject84{}
	return &this
}

// GetStages returns the Stages field value
func (o *InlineObject84) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedEventsStages
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *InlineObject84) GetStagesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool) {
	if o == nil {
    return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *InlineObject84) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) {
	o.Stages = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *InlineObject84) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	if o == nil || isNil(o.Reasons) {
		var ret []NetworksNetworkIdFirmwareUpgradesRollbacksReasons
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject84) GetReasonsOk() ([]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool) {
	if o == nil || isNil(o.Reasons) {
    return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *InlineObject84) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []NetworksNetworkIdFirmwareUpgradesRollbacksReasons and assigns it to the Reasons field.
func (o *InlineObject84) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons) {
	o.Reasons = v
}

func (o InlineObject84) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stages"] = o.Stages
	}
	if !isNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject84 struct {
	value *InlineObject84
	isSet bool
}

func (v NullableInlineObject84) Get() *InlineObject84 {
	return v.value
}

func (v *NullableInlineObject84) Set(val *InlineObject84) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject84) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject84) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject84(val *InlineObject84) *NullableInlineObject84 {
	return &NullableInlineObject84{value: val, isSet: true}
}

func (v NullableInlineObject84) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject84) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


