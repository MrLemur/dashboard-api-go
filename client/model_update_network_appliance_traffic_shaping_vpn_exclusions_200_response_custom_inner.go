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

// checks if the UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner{}

// UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner struct for UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner
type UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner struct {
	// Protocol.
	Protocol string `json:"protocol"`
	// Destination address; hostname required for DNS, IPv4 otherwise.
	Destination string `json:"destination"`
	// Destination port.
	Port string `json:"port"`
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner(protocol string, destination string, port string) *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner{}
	this.Protocol = protocol
	this.Destination = destination
	this.Port = port
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestination returns the Destination field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) SetDestination(v string) {
	o.Destination = v
}

// GetPort returns the Port field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) SetPort(v string) {
	o.Port = v
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["destination"] = o.Destination
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner struct {
	value *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) Get() *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) Set(val *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner(val *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner {
	return &NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


