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

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList AccessList details
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList struct {
	VtyIn *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn `json:"vtyIn,omitempty"`
	VtyOut *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyOut `json:"vtyOut,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList{}
	return &this
}

// GetVtyIn returns the VtyIn field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) GetVtyIn() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn {
	if o == nil || isNil(o.VtyIn) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn
		return ret
	}
	return *o.VtyIn
}

// GetVtyInOk returns a tuple with the VtyIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) GetVtyInOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn, bool) {
	if o == nil || isNil(o.VtyIn) {
    return nil, false
	}
	return o.VtyIn, true
}

// HasVtyIn returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) HasVtyIn() bool {
	if o != nil && !isNil(o.VtyIn) {
		return true
	}

	return false
}

// SetVtyIn gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn and assigns it to the VtyIn field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) SetVtyIn(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) {
	o.VtyIn = &v
}

// GetVtyOut returns the VtyOut field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) GetVtyOut() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyOut {
	if o == nil || isNil(o.VtyOut) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyOut
		return ret
	}
	return *o.VtyOut
}

// GetVtyOutOk returns a tuple with the VtyOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) GetVtyOutOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyOut, bool) {
	if o == nil || isNil(o.VtyOut) {
    return nil, false
	}
	return o.VtyOut, true
}

// HasVtyOut returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) HasVtyOut() bool {
	if o != nil && !isNil(o.VtyOut) {
		return true
	}

	return false
}

// SetVtyOut gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyOut and assigns it to the VtyOut field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) SetVtyOut(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyOut) {
	o.VtyOut = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VtyIn) {
		toSerialize["vtyIn"] = o.VtyIn
	}
	if !isNil(o.VtyOut) {
		toSerialize["vtyOut"] = o.VtyOut
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


