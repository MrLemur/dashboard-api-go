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

// checks if the UpdateNetworkWirelessSsidVpnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidVpnRequest{}

// UpdateNetworkWirelessSsidVpnRequest struct for UpdateNetworkWirelessSsidVpnRequest
type UpdateNetworkWirelessSsidVpnRequest struct {
	Concentrator *UpdateNetworkWirelessSsidVpnRequestConcentrator `json:"concentrator,omitempty"`
	SplitTunnel *UpdateNetworkWirelessSsidVpnRequestSplitTunnel `json:"splitTunnel,omitempty"`
	Failover *UpdateNetworkWirelessSsidVpnRequestFailover `json:"failover,omitempty"`
}

// NewUpdateNetworkWirelessSsidVpnRequest instantiates a new UpdateNetworkWirelessSsidVpnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidVpnRequest() *UpdateNetworkWirelessSsidVpnRequest {
	this := UpdateNetworkWirelessSsidVpnRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidVpnRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidVpnRequestWithDefaults() *UpdateNetworkWirelessSsidVpnRequest {
	this := UpdateNetworkWirelessSsidVpnRequest{}
	return &this
}

// GetConcentrator returns the Concentrator field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequest) GetConcentrator() UpdateNetworkWirelessSsidVpnRequestConcentrator {
	if o == nil || IsNil(o.Concentrator) {
		var ret UpdateNetworkWirelessSsidVpnRequestConcentrator
		return ret
	}
	return *o.Concentrator
}

// GetConcentratorOk returns a tuple with the Concentrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequest) GetConcentratorOk() (*UpdateNetworkWirelessSsidVpnRequestConcentrator, bool) {
	if o == nil || IsNil(o.Concentrator) {
		return nil, false
	}
	return o.Concentrator, true
}

// HasConcentrator returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequest) HasConcentrator() bool {
	if o != nil && !IsNil(o.Concentrator) {
		return true
	}

	return false
}

// SetConcentrator gets a reference to the given UpdateNetworkWirelessSsidVpnRequestConcentrator and assigns it to the Concentrator field.
func (o *UpdateNetworkWirelessSsidVpnRequest) SetConcentrator(v UpdateNetworkWirelessSsidVpnRequestConcentrator) {
	o.Concentrator = &v
}

// GetSplitTunnel returns the SplitTunnel field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequest) GetSplitTunnel() UpdateNetworkWirelessSsidVpnRequestSplitTunnel {
	if o == nil || IsNil(o.SplitTunnel) {
		var ret UpdateNetworkWirelessSsidVpnRequestSplitTunnel
		return ret
	}
	return *o.SplitTunnel
}

// GetSplitTunnelOk returns a tuple with the SplitTunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequest) GetSplitTunnelOk() (*UpdateNetworkWirelessSsidVpnRequestSplitTunnel, bool) {
	if o == nil || IsNil(o.SplitTunnel) {
		return nil, false
	}
	return o.SplitTunnel, true
}

// HasSplitTunnel returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequest) HasSplitTunnel() bool {
	if o != nil && !IsNil(o.SplitTunnel) {
		return true
	}

	return false
}

// SetSplitTunnel gets a reference to the given UpdateNetworkWirelessSsidVpnRequestSplitTunnel and assigns it to the SplitTunnel field.
func (o *UpdateNetworkWirelessSsidVpnRequest) SetSplitTunnel(v UpdateNetworkWirelessSsidVpnRequestSplitTunnel) {
	o.SplitTunnel = &v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequest) GetFailover() UpdateNetworkWirelessSsidVpnRequestFailover {
	if o == nil || IsNil(o.Failover) {
		var ret UpdateNetworkWirelessSsidVpnRequestFailover
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequest) GetFailoverOk() (*UpdateNetworkWirelessSsidVpnRequestFailover, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequest) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given UpdateNetworkWirelessSsidVpnRequestFailover and assigns it to the Failover field.
func (o *UpdateNetworkWirelessSsidVpnRequest) SetFailover(v UpdateNetworkWirelessSsidVpnRequestFailover) {
	o.Failover = &v
}

func (o UpdateNetworkWirelessSsidVpnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidVpnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Concentrator) {
		toSerialize["concentrator"] = o.Concentrator
	}
	if !IsNil(o.SplitTunnel) {
		toSerialize["splitTunnel"] = o.SplitTunnel
	}
	if !IsNil(o.Failover) {
		toSerialize["failover"] = o.Failover
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidVpnRequest struct {
	value *UpdateNetworkWirelessSsidVpnRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidVpnRequest) Get() *UpdateNetworkWirelessSsidVpnRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequest) Set(val *UpdateNetworkWirelessSsidVpnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidVpnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidVpnRequest(val *UpdateNetworkWirelessSsidVpnRequest) *NullableUpdateNetworkWirelessSsidVpnRequest {
	return &NullableUpdateNetworkWirelessSsidVpnRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidVpnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


