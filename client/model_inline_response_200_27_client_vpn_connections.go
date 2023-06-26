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

// InlineResponse20027ClientVpnConnections struct for InlineResponse20027ClientVpnConnections
type InlineResponse20027ClientVpnConnections struct {
	// The IP address of the VPN the client last connected to
	RemoteIp *string `json:"remoteIp,omitempty"`
	// The time the client last connected to the VPN
	ConnectedAt *int32 `json:"connectedAt,omitempty"`
	// The time the client last disconnectd from the VPN
	DisconnectedAt *int32 `json:"disconnectedAt,omitempty"`
}

// NewInlineResponse20027ClientVpnConnections instantiates a new InlineResponse20027ClientVpnConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027ClientVpnConnections() *InlineResponse20027ClientVpnConnections {
	this := InlineResponse20027ClientVpnConnections{}
	return &this
}

// NewInlineResponse20027ClientVpnConnectionsWithDefaults instantiates a new InlineResponse20027ClientVpnConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027ClientVpnConnectionsWithDefaults() *InlineResponse20027ClientVpnConnections {
	this := InlineResponse20027ClientVpnConnections{}
	return &this
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *InlineResponse20027ClientVpnConnections) GetRemoteIp() string {
	if o == nil || isNil(o.RemoteIp) {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027ClientVpnConnections) GetRemoteIpOk() (*string, bool) {
	if o == nil || isNil(o.RemoteIp) {
    return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *InlineResponse20027ClientVpnConnections) HasRemoteIp() bool {
	if o != nil && !isNil(o.RemoteIp) {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *InlineResponse20027ClientVpnConnections) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetConnectedAt returns the ConnectedAt field value if set, zero value otherwise.
func (o *InlineResponse20027ClientVpnConnections) GetConnectedAt() int32 {
	if o == nil || isNil(o.ConnectedAt) {
		var ret int32
		return ret
	}
	return *o.ConnectedAt
}

// GetConnectedAtOk returns a tuple with the ConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027ClientVpnConnections) GetConnectedAtOk() (*int32, bool) {
	if o == nil || isNil(o.ConnectedAt) {
    return nil, false
	}
	return o.ConnectedAt, true
}

// HasConnectedAt returns a boolean if a field has been set.
func (o *InlineResponse20027ClientVpnConnections) HasConnectedAt() bool {
	if o != nil && !isNil(o.ConnectedAt) {
		return true
	}

	return false
}

// SetConnectedAt gets a reference to the given int32 and assigns it to the ConnectedAt field.
func (o *InlineResponse20027ClientVpnConnections) SetConnectedAt(v int32) {
	o.ConnectedAt = &v
}

// GetDisconnectedAt returns the DisconnectedAt field value if set, zero value otherwise.
func (o *InlineResponse20027ClientVpnConnections) GetDisconnectedAt() int32 {
	if o == nil || isNil(o.DisconnectedAt) {
		var ret int32
		return ret
	}
	return *o.DisconnectedAt
}

// GetDisconnectedAtOk returns a tuple with the DisconnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027ClientVpnConnections) GetDisconnectedAtOk() (*int32, bool) {
	if o == nil || isNil(o.DisconnectedAt) {
    return nil, false
	}
	return o.DisconnectedAt, true
}

// HasDisconnectedAt returns a boolean if a field has been set.
func (o *InlineResponse20027ClientVpnConnections) HasDisconnectedAt() bool {
	if o != nil && !isNil(o.DisconnectedAt) {
		return true
	}

	return false
}

// SetDisconnectedAt gets a reference to the given int32 and assigns it to the DisconnectedAt field.
func (o *InlineResponse20027ClientVpnConnections) SetDisconnectedAt(v int32) {
	o.DisconnectedAt = &v
}

func (o InlineResponse20027ClientVpnConnections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemoteIp) {
		toSerialize["remoteIp"] = o.RemoteIp
	}
	if !isNil(o.ConnectedAt) {
		toSerialize["connectedAt"] = o.ConnectedAt
	}
	if !isNil(o.DisconnectedAt) {
		toSerialize["disconnectedAt"] = o.DisconnectedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027ClientVpnConnections struct {
	value *InlineResponse20027ClientVpnConnections
	isSet bool
}

func (v NullableInlineResponse20027ClientVpnConnections) Get() *InlineResponse20027ClientVpnConnections {
	return v.value
}

func (v *NullableInlineResponse20027ClientVpnConnections) Set(val *InlineResponse20027ClientVpnConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027ClientVpnConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027ClientVpnConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027ClientVpnConnections(val *InlineResponse20027ClientVpnConnections) *NullableInlineResponse20027ClientVpnConnections {
	return &NullableInlineResponse20027ClientVpnConnections{value: val, isSet: true}
}

func (v NullableInlineResponse20027ClientVpnConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027ClientVpnConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


