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

// checks if the GetNetworkSyslogServers200ResponseServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSyslogServers200ResponseServersInner{}

// GetNetworkSyslogServers200ResponseServersInner struct for GetNetworkSyslogServers200ResponseServersInner
type GetNetworkSyslogServers200ResponseServersInner struct {
	// The IP address of the syslog server
	Host *string `json:"host,omitempty"`
	// The port of the syslog server
	Port *int32 `json:"port,omitempty"`
	// A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
	Roles []string `json:"roles,omitempty"`
}

// NewGetNetworkSyslogServers200ResponseServersInner instantiates a new GetNetworkSyslogServers200ResponseServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSyslogServers200ResponseServersInner() *GetNetworkSyslogServers200ResponseServersInner {
	this := GetNetworkSyslogServers200ResponseServersInner{}
	return &this
}

// NewGetNetworkSyslogServers200ResponseServersInnerWithDefaults instantiates a new GetNetworkSyslogServers200ResponseServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSyslogServers200ResponseServersInnerWithDefaults() *GetNetworkSyslogServers200ResponseServersInner {
	this := GetNetworkSyslogServers200ResponseServersInner{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetNetworkSyslogServers200ResponseServersInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSyslogServers200ResponseServersInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetNetworkSyslogServers200ResponseServersInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetNetworkSyslogServers200ResponseServersInner) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkSyslogServers200ResponseServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSyslogServers200ResponseServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkSyslogServers200ResponseServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNetworkSyslogServers200ResponseServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GetNetworkSyslogServers200ResponseServersInner) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSyslogServers200ResponseServersInner) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GetNetworkSyslogServers200ResponseServersInner) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GetNetworkSyslogServers200ResponseServersInner) SetRoles(v []string) {
	o.Roles = v
}

func (o GetNetworkSyslogServers200ResponseServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSyslogServers200ResponseServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableGetNetworkSyslogServers200ResponseServersInner struct {
	value *GetNetworkSyslogServers200ResponseServersInner
	isSet bool
}

func (v NullableGetNetworkSyslogServers200ResponseServersInner) Get() *GetNetworkSyslogServers200ResponseServersInner {
	return v.value
}

func (v *NullableGetNetworkSyslogServers200ResponseServersInner) Set(val *GetNetworkSyslogServers200ResponseServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSyslogServers200ResponseServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSyslogServers200ResponseServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSyslogServers200ResponseServersInner(val *GetNetworkSyslogServers200ResponseServersInner) *NullableGetNetworkSyslogServers200ResponseServersInner {
	return &NullableGetNetworkSyslogServers200ResponseServersInner{value: val, isSet: true}
}

func (v NullableGetNetworkSyslogServers200ResponseServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSyslogServers200ResponseServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


