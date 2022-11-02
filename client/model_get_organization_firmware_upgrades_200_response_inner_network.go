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

// GetOrganizationFirmwareUpgrades200ResponseInnerNetwork Network of the upgrade
type GetOrganizationFirmwareUpgrades200ResponseInnerNetwork struct {
	// ID of network
	Id *string `json:"id,omitempty"`
	// The network
	Name *string `json:"name,omitempty"`
}

// NewGetOrganizationFirmwareUpgrades200ResponseInnerNetwork instantiates a new GetOrganizationFirmwareUpgrades200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgrades200ResponseInnerNetwork() *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork {
	this := GetOrganizationFirmwareUpgrades200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationFirmwareUpgrades200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationFirmwareUpgrades200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgrades200ResponseInnerNetworkWithDefaults() *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork {
	this := GetOrganizationFirmwareUpgrades200ResponseInnerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) SetName(v string) {
	o.Name = &v
}

func (o GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork struct {
	value *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork) Get() *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork) Set(val *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork(val *GetOrganizationFirmwareUpgrades200ResponseInnerNetwork) *NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork {
	return &NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgrades200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


