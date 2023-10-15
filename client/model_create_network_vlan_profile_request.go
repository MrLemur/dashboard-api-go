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

// checks if the CreateNetworkVlanProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkVlanProfileRequest{}

// CreateNetworkVlanProfileRequest struct for CreateNetworkVlanProfileRequest
type CreateNetworkVlanProfileRequest struct {
	// Name of the profile, string length must be from 1 to 255 characters
	Name string `json:"name"`
	// An array of named VLANs
	VlanNames []CreateNetworkVlanProfileRequestVlanNamesInner `json:"vlanNames"`
	// An array of VLAN groups
	VlanGroups []CreateNetworkVlanProfileRequestVlanGroupsInner `json:"vlanGroups"`
	// IName of the profile
	Iname string `json:"iname"`
}

// NewCreateNetworkVlanProfileRequest instantiates a new CreateNetworkVlanProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkVlanProfileRequest(name string, vlanNames []CreateNetworkVlanProfileRequestVlanNamesInner, vlanGroups []CreateNetworkVlanProfileRequestVlanGroupsInner, iname string) *CreateNetworkVlanProfileRequest {
	this := CreateNetworkVlanProfileRequest{}
	this.Name = name
	this.VlanNames = vlanNames
	this.VlanGroups = vlanGroups
	this.Iname = iname
	return &this
}

// NewCreateNetworkVlanProfileRequestWithDefaults instantiates a new CreateNetworkVlanProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkVlanProfileRequestWithDefaults() *CreateNetworkVlanProfileRequest {
	this := CreateNetworkVlanProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkVlanProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkVlanProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkVlanProfileRequest) SetName(v string) {
	o.Name = v
}

// GetVlanNames returns the VlanNames field value
func (o *CreateNetworkVlanProfileRequest) GetVlanNames() []CreateNetworkVlanProfileRequestVlanNamesInner {
	if o == nil {
		var ret []CreateNetworkVlanProfileRequestVlanNamesInner
		return ret
	}

	return o.VlanNames
}

// GetVlanNamesOk returns a tuple with the VlanNames field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkVlanProfileRequest) GetVlanNamesOk() ([]CreateNetworkVlanProfileRequestVlanNamesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanNames, true
}

// SetVlanNames sets field value
func (o *CreateNetworkVlanProfileRequest) SetVlanNames(v []CreateNetworkVlanProfileRequestVlanNamesInner) {
	o.VlanNames = v
}

// GetVlanGroups returns the VlanGroups field value
func (o *CreateNetworkVlanProfileRequest) GetVlanGroups() []CreateNetworkVlanProfileRequestVlanGroupsInner {
	if o == nil {
		var ret []CreateNetworkVlanProfileRequestVlanGroupsInner
		return ret
	}

	return o.VlanGroups
}

// GetVlanGroupsOk returns a tuple with the VlanGroups field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkVlanProfileRequest) GetVlanGroupsOk() ([]CreateNetworkVlanProfileRequestVlanGroupsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanGroups, true
}

// SetVlanGroups sets field value
func (o *CreateNetworkVlanProfileRequest) SetVlanGroups(v []CreateNetworkVlanProfileRequestVlanGroupsInner) {
	o.VlanGroups = v
}

// GetIname returns the Iname field value
func (o *CreateNetworkVlanProfileRequest) GetIname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iname
}

// GetInameOk returns a tuple with the Iname field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkVlanProfileRequest) GetInameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iname, true
}

// SetIname sets field value
func (o *CreateNetworkVlanProfileRequest) SetIname(v string) {
	o.Iname = v
}

func (o CreateNetworkVlanProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkVlanProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["vlanNames"] = o.VlanNames
	toSerialize["vlanGroups"] = o.VlanGroups
	toSerialize["iname"] = o.Iname
	return toSerialize, nil
}

type NullableCreateNetworkVlanProfileRequest struct {
	value *CreateNetworkVlanProfileRequest
	isSet bool
}

func (v NullableCreateNetworkVlanProfileRequest) Get() *CreateNetworkVlanProfileRequest {
	return v.value
}

func (v *NullableCreateNetworkVlanProfileRequest) Set(val *CreateNetworkVlanProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkVlanProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkVlanProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkVlanProfileRequest(val *CreateNetworkVlanProfileRequest) *NullableCreateNetworkVlanProfileRequest {
	return &NullableCreateNetworkVlanProfileRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkVlanProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkVlanProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


