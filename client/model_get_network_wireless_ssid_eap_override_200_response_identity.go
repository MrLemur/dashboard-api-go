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

// checks if the GetNetworkWirelessSsidEapOverride200ResponseIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidEapOverride200ResponseIdentity{}

// GetNetworkWirelessSsidEapOverride200ResponseIdentity EAP settings for identity requests.
type GetNetworkWirelessSsidEapOverride200ResponseIdentity struct {
	// Maximum number of EAP retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewGetNetworkWirelessSsidEapOverride200ResponseIdentity instantiates a new GetNetworkWirelessSsidEapOverride200ResponseIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidEapOverride200ResponseIdentity() *GetNetworkWirelessSsidEapOverride200ResponseIdentity {
	this := GetNetworkWirelessSsidEapOverride200ResponseIdentity{}
	return &this
}

// NewGetNetworkWirelessSsidEapOverride200ResponseIdentityWithDefaults instantiates a new GetNetworkWirelessSsidEapOverride200ResponseIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidEapOverride200ResponseIdentityWithDefaults() *GetNetworkWirelessSsidEapOverride200ResponseIdentity {
	this := GetNetworkWirelessSsidEapOverride200ResponseIdentity{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) GetRetries() int32 {
	if o == nil || IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) GetRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *GetNetworkWirelessSsidEapOverride200ResponseIdentity) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o GetNetworkWirelessSsidEapOverride200ResponseIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidEapOverride200ResponseIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity struct {
	value *GetNetworkWirelessSsidEapOverride200ResponseIdentity
	isSet bool
}

func (v NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity) Get() *GetNetworkWirelessSsidEapOverride200ResponseIdentity {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity) Set(val *GetNetworkWirelessSsidEapOverride200ResponseIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidEapOverride200ResponseIdentity(val *GetNetworkWirelessSsidEapOverride200ResponseIdentity) *NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity {
	return &NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidEapOverride200ResponseIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


