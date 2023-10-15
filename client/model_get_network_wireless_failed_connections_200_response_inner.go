/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessFailedConnections200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessFailedConnections200ResponseInner{}

// GetNetworkWirelessFailedConnections200ResponseInner struct for GetNetworkWirelessFailedConnections200ResponseInner
type GetNetworkWirelessFailedConnections200ResponseInner struct {
	// SSID Number
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// LAN
	Vlan *int32 `json:"vlan,omitempty"`
	// Client Mac
	ClientMac *string `json:"clientMac,omitempty"`
	// Serial Number
	Serial *string `json:"serial,omitempty"`
	// The failed onboarding step. One of: assoc, auth, dhcp, dns.
	FailureStep *string `json:"failureStep,omitempty"`
	// The failure type in the onboarding step
	Type *string `json:"type,omitempty"`
	// The timestamp when the client mac failed
	Ts *time.Time `json:"ts,omitempty"`
}

// NewGetNetworkWirelessFailedConnections200ResponseInner instantiates a new GetNetworkWirelessFailedConnections200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessFailedConnections200ResponseInner() *GetNetworkWirelessFailedConnections200ResponseInner {
	this := GetNetworkWirelessFailedConnections200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessFailedConnections200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessFailedConnections200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessFailedConnections200ResponseInnerWithDefaults() *GetNetworkWirelessFailedConnections200ResponseInner {
	this := GetNetworkWirelessFailedConnections200ResponseInner{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSsidNumber() int32 {
	if o == nil || IsNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSsidNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SsidNumber) {
		return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasSsidNumber() bool {
	if o != nil && !IsNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetVlan(v int32) {
	o.Vlan = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetClientMac() string {
	if o == nil || IsNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetClientMacOk() (*string, bool) {
	if o == nil || IsNil(o.ClientMac) {
		return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasClientMac() bool {
	if o != nil && !IsNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetFailureStep returns the FailureStep field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetFailureStep() string {
	if o == nil || IsNil(o.FailureStep) {
		var ret string
		return ret
	}
	return *o.FailureStep
}

// GetFailureStepOk returns a tuple with the FailureStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetFailureStepOk() (*string, bool) {
	if o == nil || IsNil(o.FailureStep) {
		return nil, false
	}
	return o.FailureStep, true
}

// HasFailureStep returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasFailureStep() bool {
	if o != nil && !IsNil(o.FailureStep) {
		return true
	}

	return false
}

// SetFailureStep gets a reference to the given string and assigns it to the FailureStep field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetFailureStep(v string) {
	o.FailureStep = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetTs(v time.Time) {
	o.Ts = &v
}

func (o GetNetworkWirelessFailedConnections200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessFailedConnections200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.ClientMac) {
		toSerialize["clientMac"] = o.ClientMac
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.FailureStep) {
		toSerialize["failureStep"] = o.FailureStep
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessFailedConnections200ResponseInner struct {
	value *GetNetworkWirelessFailedConnections200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessFailedConnections200ResponseInner) Get() *GetNetworkWirelessFailedConnections200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessFailedConnections200ResponseInner) Set(val *GetNetworkWirelessFailedConnections200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessFailedConnections200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessFailedConnections200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessFailedConnections200ResponseInner(val *GetNetworkWirelessFailedConnections200ResponseInner) *NullableGetNetworkWirelessFailedConnections200ResponseInner {
	return &NullableGetNetworkWirelessFailedConnections200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessFailedConnections200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessFailedConnections200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


