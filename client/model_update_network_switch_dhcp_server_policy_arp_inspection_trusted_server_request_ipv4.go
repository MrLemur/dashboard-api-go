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

// UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 The updated IPv4 attributes of the trusted server
type UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 struct {
	// The updated IPv4 address of the trusted server
	Address *string `json:"address,omitempty"`
}

// NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 instantiates a new UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4() *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	this := UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{}
	return &this
}

// NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4WithDefaults instantiates a new UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4WithDefaults() *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	this := UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 struct {
	value *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4
	isSet bool
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) Get() *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) Set(val *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4(val *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) *NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	return &NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


