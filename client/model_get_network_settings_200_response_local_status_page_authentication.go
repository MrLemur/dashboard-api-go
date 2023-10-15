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

// checks if the GetNetworkSettings200ResponseLocalStatusPageAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200ResponseLocalStatusPageAuthentication{}

// GetNetworkSettings200ResponseLocalStatusPageAuthentication A hash of Local Status page(s)' authentication options applied to the Network.
type GetNetworkSettings200ResponseLocalStatusPageAuthentication struct {
	// Enables / disables the authentication on Local Status page(s).
	Enabled *bool `json:"enabled,omitempty"`
	// The username used for Local Status Page(s).
	Username *string `json:"username,omitempty"`
}

// NewGetNetworkSettings200ResponseLocalStatusPageAuthentication instantiates a new GetNetworkSettings200ResponseLocalStatusPageAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200ResponseLocalStatusPageAuthentication() *GetNetworkSettings200ResponseLocalStatusPageAuthentication {
	this := GetNetworkSettings200ResponseLocalStatusPageAuthentication{}
	return &this
}

// NewGetNetworkSettings200ResponseLocalStatusPageAuthenticationWithDefaults instantiates a new GetNetworkSettings200ResponseLocalStatusPageAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseLocalStatusPageAuthenticationWithDefaults() *GetNetworkSettings200ResponseLocalStatusPageAuthentication {
	this := GetNetworkSettings200ResponseLocalStatusPageAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetNetworkSettings200ResponseLocalStatusPageAuthentication) SetUsername(v string) {
	o.Username = &v
}

func (o GetNetworkSettings200ResponseLocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200ResponseLocalStatusPageAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication struct {
	value *GetNetworkSettings200ResponseLocalStatusPageAuthentication
	isSet bool
}

func (v NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication) Get() *GetNetworkSettings200ResponseLocalStatusPageAuthentication {
	return v.value
}

func (v *NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication) Set(val *GetNetworkSettings200ResponseLocalStatusPageAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200ResponseLocalStatusPageAuthentication(val *GetNetworkSettings200ResponseLocalStatusPageAuthentication) *NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication {
	return &NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200ResponseLocalStatusPageAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


