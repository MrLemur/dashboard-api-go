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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp UDP attributes of the packet.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp struct {
	// UDP length of the packet.
	Length *int32 `json:"length,omitempty"`
	// UDP checksum of the packet.
	Checksum *string `json:"checksum,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdpWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdpWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) SetLength(v int32) {
	o.Length = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) SetChecksum(v string) {
	o.Checksum = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


