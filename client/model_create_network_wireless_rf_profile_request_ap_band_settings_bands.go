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

// checks if the CreateNetworkWirelessRfProfileRequestApBandSettingsBands type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequestApBandSettingsBands{}

// CreateNetworkWirelessRfProfileRequestApBandSettingsBands Settings related to all bands
type CreateNetworkWirelessRfProfileRequestApBandSettingsBands struct {
	// List of enabled bands. Can include [\"2.4\", \"5\", \"6\"]
	Enabled []string `json:"enabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestApBandSettingsBands instantiates a new CreateNetworkWirelessRfProfileRequestApBandSettingsBands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestApBandSettingsBands() *CreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	this := CreateNetworkWirelessRfProfileRequestApBandSettingsBands{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestApBandSettingsBandsWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestApBandSettingsBands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestApBandSettingsBandsWithDefaults() *CreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	this := CreateNetworkWirelessRfProfileRequestApBandSettingsBands{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestApBandSettingsBands) GetEnabled() []string {
	if o == nil || IsNil(o.Enabled) {
		var ret []string
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestApBandSettingsBands) GetEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestApBandSettingsBands) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given []string and assigns it to the Enabled field.
func (o *CreateNetworkWirelessRfProfileRequestApBandSettingsBands) SetEnabled(v []string) {
	o.Enabled = v
}

func (o CreateNetworkWirelessRfProfileRequestApBandSettingsBands) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequestApBandSettingsBands) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands struct {
	value *CreateNetworkWirelessRfProfileRequestApBandSettingsBands
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands) Get() *CreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands) Set(val *CreateNetworkWirelessRfProfileRequestApBandSettingsBands) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands(val *CreateNetworkWirelessRfProfileRequestApBandSettingsBands) *NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	return &NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestApBandSettingsBands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


