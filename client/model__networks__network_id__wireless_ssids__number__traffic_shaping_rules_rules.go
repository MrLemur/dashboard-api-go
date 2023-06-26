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

// NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules struct for NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules
type NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules struct {
	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. 
	Definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions `json:"definitions"`
	PerClientBandwidthLimits *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`
	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint. 
	DscpTagValue *int32 `json:"dscpTagValue,omitempty"`
	//     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'. 
	PcpTagValue *int32 `json:"pcpTagValue,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules instantiates a new NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules(definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules {
	this := NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules{}
	this.Definitions = definitions
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRulesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRulesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules {
	this := NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules{}
	return &this
}

// GetDefinitions returns the Definitions field value
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetDefinitions() []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions {
	if o == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions
		return ret
	}

	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetDefinitionsOk() ([]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Definitions, true
}

// SetDefinitions sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) SetDefinitions(v []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) {
	o.Definitions = v
}

// GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetPerClientBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	if o == nil || isNil(o.PerClientBandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits
		return ret
	}
	return *o.PerClientBandwidthLimits
}

// GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetPerClientBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits, bool) {
	if o == nil || isNil(o.PerClientBandwidthLimits) {
    return nil, false
	}
	return o.PerClientBandwidthLimits, true
}

// HasPerClientBandwidthLimits returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) HasPerClientBandwidthLimits() bool {
	if o != nil && !isNil(o.PerClientBandwidthLimits) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits and assigns it to the PerClientBandwidthLimits field.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) SetPerClientBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) {
	o.PerClientBandwidthLimits = &v
}

// GetDscpTagValue returns the DscpTagValue field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetDscpTagValue() int32 {
	if o == nil || isNil(o.DscpTagValue) {
		var ret int32
		return ret
	}
	return *o.DscpTagValue
}

// GetDscpTagValueOk returns a tuple with the DscpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetDscpTagValueOk() (*int32, bool) {
	if o == nil || isNil(o.DscpTagValue) {
    return nil, false
	}
	return o.DscpTagValue, true
}

// HasDscpTagValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) HasDscpTagValue() bool {
	if o != nil && !isNil(o.DscpTagValue) {
		return true
	}

	return false
}

// SetDscpTagValue gets a reference to the given int32 and assigns it to the DscpTagValue field.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) SetDscpTagValue(v int32) {
	o.DscpTagValue = &v
}

// GetPcpTagValue returns the PcpTagValue field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetPcpTagValue() int32 {
	if o == nil || isNil(o.PcpTagValue) {
		var ret int32
		return ret
	}
	return *o.PcpTagValue
}

// GetPcpTagValueOk returns a tuple with the PcpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) GetPcpTagValueOk() (*int32, bool) {
	if o == nil || isNil(o.PcpTagValue) {
    return nil, false
	}
	return o.PcpTagValue, true
}

// HasPcpTagValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) HasPcpTagValue() bool {
	if o != nil && !isNil(o.PcpTagValue) {
		return true
	}

	return false
}

// SetPcpTagValue gets a reference to the given int32 and assigns it to the PcpTagValue field.
func (o *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) SetPcpTagValue(v int32) {
	o.PcpTagValue = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.PcpTagValue) {
		toSerialize["pcpTagValue"] = o.PcpTagValue
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules struct {
	value *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) Get() *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) Set(val *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules(val *NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) *NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules {
	return &NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


