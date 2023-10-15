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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp DSCP attributes of the packet.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp struct {
	// DSCP tag number of the packet.
	Tag *int32 `json:"tag,omitempty"`
	// ECN number of the packet.
	Ecn *int32 `json:"ecn,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscpWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscpWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) GetTag() int32 {
	if o == nil || IsNil(o.Tag) {
		var ret int32
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) GetTagOk() (*int32, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given int32 and assigns it to the Tag field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) SetTag(v int32) {
	o.Tag = &v
}

// GetEcn returns the Ecn field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) GetEcn() int32 {
	if o == nil || IsNil(o.Ecn) {
		var ret int32
		return ret
	}
	return *o.Ecn
}

// GetEcnOk returns a tuple with the Ecn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) GetEcnOk() (*int32, bool) {
	if o == nil || IsNil(o.Ecn) {
		return nil, false
	}
	return o.Ecn, true
}

// HasEcn returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) HasEcn() bool {
	if o != nil && !IsNil(o.Ecn) {
		return true
	}

	return false
}

// SetEcn gets a reference to the given int32 and assigns it to the Ecn field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) SetEcn(v int32) {
	o.Ecn = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Ecn) {
		toSerialize["ecn"] = o.Ecn
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIpDscp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


