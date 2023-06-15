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

// InlineObject205 struct for InlineObject205
type InlineObject205 struct {
	// A set of devices to import (or update)
	Devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices `json:"devices"`
}

// NewInlineObject205 instantiates a new InlineObject205 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject205(devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) *InlineObject205 {
	this := InlineObject205{}
	this.Devices = devices
	return &this
}

// NewInlineObject205WithDefaults instantiates a new InlineObject205 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject205WithDefaults() *InlineObject205 {
	this := InlineObject205{}
	return &this
}

// GetDevices returns the Devices field value
func (o *InlineObject205) GetDevices() []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	if o == nil {
		var ret []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *InlineObject205) GetDevicesOk() ([]OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices, bool) {
	if o == nil {
    return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *InlineObject205) SetDevices(v []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) {
	o.Devices = v
}

func (o InlineObject205) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject205 struct {
	value *InlineObject205
	isSet bool
}

func (v NullableInlineObject205) Get() *InlineObject205 {
	return v.value
}

func (v *NullableInlineObject205) Set(val *InlineObject205) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject205) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject205) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject205(val *InlineObject205) *NullableInlineObject205 {
	return &NullableInlineObject205{value: val, isSet: true}
}

func (v NullableInlineObject205) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject205) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


