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

// GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork Network info.
type GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork struct {
	// ID for the network that the device is associated with.
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork instantiates a new GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork {
	this := GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetworkWithDefaults() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork {
	this := GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork struct {
	value *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) Get() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) Set(val *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork(val *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork {
	return &NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


