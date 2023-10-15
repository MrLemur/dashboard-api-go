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

// checks if the CreateNetworkVlanProfileRequestVlanGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkVlanProfileRequestVlanGroupsInner{}

// CreateNetworkVlanProfileRequestVlanGroupsInner struct for CreateNetworkVlanProfileRequestVlanGroupsInner
type CreateNetworkVlanProfileRequestVlanGroupsInner struct {
	// Name of the VLAN, string length must be from 1 to 32 characters
	Name string `json:"name"`
	// Comma-separated VLAN IDs or ID ranges
	VlanIds string `json:"vlanIds"`
}

// NewCreateNetworkVlanProfileRequestVlanGroupsInner instantiates a new CreateNetworkVlanProfileRequestVlanGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkVlanProfileRequestVlanGroupsInner(name string, vlanIds string) *CreateNetworkVlanProfileRequestVlanGroupsInner {
	this := CreateNetworkVlanProfileRequestVlanGroupsInner{}
	this.Name = name
	this.VlanIds = vlanIds
	return &this
}

// NewCreateNetworkVlanProfileRequestVlanGroupsInnerWithDefaults instantiates a new CreateNetworkVlanProfileRequestVlanGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkVlanProfileRequestVlanGroupsInnerWithDefaults() *CreateNetworkVlanProfileRequestVlanGroupsInner {
	this := CreateNetworkVlanProfileRequestVlanGroupsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) SetName(v string) {
	o.Name = v
}

// GetVlanIds returns the VlanIds field value
func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetVlanIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VlanIds
}

// GetVlanIdsOk returns a tuple with the VlanIds field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetVlanIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanIds, true
}

// SetVlanIds sets field value
func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) SetVlanIds(v string) {
	o.VlanIds = v
}

func (o CreateNetworkVlanProfileRequestVlanGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkVlanProfileRequestVlanGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["vlanIds"] = o.VlanIds
	return toSerialize, nil
}

type NullableCreateNetworkVlanProfileRequestVlanGroupsInner struct {
	value *CreateNetworkVlanProfileRequestVlanGroupsInner
	isSet bool
}

func (v NullableCreateNetworkVlanProfileRequestVlanGroupsInner) Get() *CreateNetworkVlanProfileRequestVlanGroupsInner {
	return v.value
}

func (v *NullableCreateNetworkVlanProfileRequestVlanGroupsInner) Set(val *CreateNetworkVlanProfileRequestVlanGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkVlanProfileRequestVlanGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkVlanProfileRequestVlanGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkVlanProfileRequestVlanGroupsInner(val *CreateNetworkVlanProfileRequestVlanGroupsInner) *NullableCreateNetworkVlanProfileRequestVlanGroupsInner {
	return &NullableCreateNetworkVlanProfileRequestVlanGroupsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkVlanProfileRequestVlanGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkVlanProfileRequestVlanGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


