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

// UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner struct for UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner
type UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner struct {
	// Serial number of access point to be configured with alternate management IP
	Serial string `json:"serial"`
	// Wireless alternate management interface device IP. Provide an empty string to remove alternate management IP assignment
	AlternateManagementIp string `json:"alternateManagementIp"`
	// Subnet mask must be in IP format
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Gateway must be in IP format
	Gateway *string `json:"gateway,omitempty"`
	// Primary DNS must be in IP format
	Dns1 *string `json:"dns1,omitempty"`
	// Optional secondary DNS must be in IP format
	Dns2 *string `json:"dns2,omitempty"`
}

// NewUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner instantiates a new UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner(serial string, alternateManagementIp string) *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner {
	this := UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner{}
	this.Serial = serial
	this.AlternateManagementIp = alternateManagementIp
	return &this
}

// NewUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInnerWithDefaults instantiates a new UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInnerWithDefaults() *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner {
	this := UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner{}
	return &this
}

// GetSerial returns the Serial field value
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) SetSerial(v string) {
	o.Serial = v
}

// GetAlternateManagementIp returns the AlternateManagementIp field value
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetAlternateManagementIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateManagementIp
}

// GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetAlternateManagementIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlternateManagementIp, true
}

// SetAlternateManagementIp sets field value
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) SetAlternateManagementIp(v string) {
	o.AlternateManagementIp = v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetSubnetMask() string {
	if o == nil || isNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetSubnetMaskOk() (*string, bool) {
	if o == nil || isNil(o.SubnetMask) {
    return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) HasSubnetMask() bool {
	if o != nil && !isNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetDns1 returns the Dns1 field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetDns1() string {
	if o == nil || isNil(o.Dns1) {
		var ret string
		return ret
	}
	return *o.Dns1
}

// GetDns1Ok returns a tuple with the Dns1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetDns1Ok() (*string, bool) {
	if o == nil || isNil(o.Dns1) {
    return nil, false
	}
	return o.Dns1, true
}

// HasDns1 returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) HasDns1() bool {
	if o != nil && !isNil(o.Dns1) {
		return true
	}

	return false
}

// SetDns1 gets a reference to the given string and assigns it to the Dns1 field.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) SetDns1(v string) {
	o.Dns1 = &v
}

// GetDns2 returns the Dns2 field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetDns2() string {
	if o == nil || isNil(o.Dns2) {
		var ret string
		return ret
	}
	return *o.Dns2
}

// GetDns2Ok returns a tuple with the Dns2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) GetDns2Ok() (*string, bool) {
	if o == nil || isNil(o.Dns2) {
    return nil, false
	}
	return o.Dns2, true
}

// HasDns2 returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) HasDns2() bool {
	if o != nil && !isNil(o.Dns2) {
		return true
	}

	return false
}

// SetDns2 gets a reference to the given string and assigns it to the Dns2 field.
func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) SetDns2(v string) {
	o.Dns2 = &v
}

func (o UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["alternateManagementIp"] = o.AlternateManagementIp
	}
	if !isNil(o.SubnetMask) {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.Dns1) {
		toSerialize["dns1"] = o.Dns1
	}
	if !isNil(o.Dns2) {
		toSerialize["dns2"] = o.Dns2
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner struct {
	value *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) Get() *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) Set(val *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner(val *UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) *NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner {
	return &NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


