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

// NetworksNetworkIdApplianceTrafficShapingRulesRules struct for NetworksNetworkIdApplianceTrafficShapingRulesRules
type NetworksNetworkIdApplianceTrafficShapingRulesRules struct {
	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. 
	Definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions `json:"definitions"`
	PerClientBandwidthLimits *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`
	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint. 
	DscpTagValue *int32 `json:"dscpTagValue,omitempty"`
	//     A string, indicating the priority level for packets bound to your rule.     Can be 'low', 'normal' or 'high'. 
	Priority *string `json:"priority,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingRulesRules instantiates a new NetworksNetworkIdApplianceTrafficShapingRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingRulesRules(definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) *NetworksNetworkIdApplianceTrafficShapingRulesRules {
	this := NetworksNetworkIdApplianceTrafficShapingRulesRules{}
	this.Definitions = definitions
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingRulesRulesWithDefaults() *NetworksNetworkIdApplianceTrafficShapingRulesRules {
	this := NetworksNetworkIdApplianceTrafficShapingRulesRules{}
	return &this
}

// GetDefinitions returns the Definitions field value
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDefinitions() []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions {
	if o == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions
		return ret
	}

	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDefinitionsOk() ([]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Definitions, true
}

// SetDefinitions sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetDefinitions(v []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) {
	o.Definitions = v
}

// GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPerClientBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	if o == nil || isNil(o.PerClientBandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits
		return ret
	}
	return *o.PerClientBandwidthLimits
}

// GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPerClientBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits, bool) {
	if o == nil || isNil(o.PerClientBandwidthLimits) {
    return nil, false
	}
	return o.PerClientBandwidthLimits, true
}

// HasPerClientBandwidthLimits returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) HasPerClientBandwidthLimits() bool {
	if o != nil && !isNil(o.PerClientBandwidthLimits) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits and assigns it to the PerClientBandwidthLimits field.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetPerClientBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) {
	o.PerClientBandwidthLimits = &v
}

// GetDscpTagValue returns the DscpTagValue field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDscpTagValue() int32 {
	if o == nil || isNil(o.DscpTagValue) {
		var ret int32
		return ret
	}
	return *o.DscpTagValue
}

// GetDscpTagValueOk returns a tuple with the DscpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDscpTagValueOk() (*int32, bool) {
	if o == nil || isNil(o.DscpTagValue) {
    return nil, false
	}
	return o.DscpTagValue, true
}

// HasDscpTagValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) HasDscpTagValue() bool {
	if o != nil && !isNil(o.DscpTagValue) {
		return true
	}

	return false
}

// SetDscpTagValue gets a reference to the given int32 and assigns it to the DscpTagValue field.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetDscpTagValue(v int32) {
	o.DscpTagValue = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPriority() string {
	if o == nil || isNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPriorityOk() (*string, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetPriority(v string) {
	o.Priority = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingRulesRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["definitions"] = o.Definitions
	}
	if !isNil(o.PerClientBandwidthLimits) {
		toSerialize["perClientBandwidthLimits"] = o.PerClientBandwidthLimits
	}
	if !isNil(o.DscpTagValue) {
		toSerialize["dscpTagValue"] = o.DscpTagValue
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingRulesRules struct {
	value *NetworksNetworkIdApplianceTrafficShapingRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingRulesRules) Get() *NetworksNetworkIdApplianceTrafficShapingRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingRulesRules) Set(val *NetworksNetworkIdApplianceTrafficShapingRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingRulesRules(val *NetworksNetworkIdApplianceTrafficShapingRulesRules) *NullableNetworksNetworkIdApplianceTrafficShapingRulesRules {
	return &NullableNetworksNetworkIdApplianceTrafficShapingRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


