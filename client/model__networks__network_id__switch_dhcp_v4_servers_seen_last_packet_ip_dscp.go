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

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp DSCP attributes of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp struct {
	// DSCP tag number of the packet.
	Tag *int32 `json:"tag,omitempty"`
	// ECN number of the packet.
	Ecn *int32 `json:"ecn,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscpWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscpWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) GetTag() int32 {
	if o == nil || isNil(o.Tag) {
		var ret int32
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) GetTagOk() (*int32, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given int32 and assigns it to the Tag field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) SetTag(v int32) {
	o.Tag = &v
}

// GetEcn returns the Ecn field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) GetEcn() int32 {
	if o == nil || isNil(o.Ecn) {
		var ret int32
		return ret
	}
	return *o.Ecn
}

// GetEcnOk returns a tuple with the Ecn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) GetEcnOk() (*int32, bool) {
	if o == nil || isNil(o.Ecn) {
    return nil, false
	}
	return o.Ecn, true
}

// HasEcn returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) HasEcn() bool {
	if o != nil && !isNil(o.Ecn) {
		return true
	}

	return false
}

// SetEcn gets a reference to the given int32 and assigns it to the Ecn field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) SetEcn(v int32) {
	o.Ecn = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Ecn) {
		toSerialize["ecn"] = o.Ecn
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


