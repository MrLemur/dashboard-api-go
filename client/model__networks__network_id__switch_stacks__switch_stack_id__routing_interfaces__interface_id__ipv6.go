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

// NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 The IPv6 settings of the interface.
type NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 struct {
	// The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// The IPv6 prefix of the interface. Required if IPv6 object is included and interface does not already have ipv6.prefix configured
	Prefix *string `json:"prefix,omitempty"`
	// The IPv6 address of the interface. Required if assignmentMode is included and set as 'static'. Must not be included if assignmentMode is 'eui-64'.
	Address *string `json:"address,omitempty"`
	// The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.
	Gateway *string `json:"gateway,omitempty"`
}

// NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 instantiates a new NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 {
	this := NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6{}
	return &this
}

// NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6WithDefaults instantiates a new NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6WithDefaults() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 {
	this := NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6{}
	return &this
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetAssignmentMode() string {
	if o == nil || isNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.AssignmentMode) {
    return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) HasAssignmentMode() bool {
	if o != nil && !isNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) SetPrefix(v string) {
	o.Prefix = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) SetGateway(v string) {
	o.Gateway = &v
}

func (o NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 struct {
	value *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) Get() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) Set(val *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6(val *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6 {
	return &NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


