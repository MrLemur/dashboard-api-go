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

// checks if the GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile{}

// GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile The VLAN Profile
type GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile struct {
	// IName of the VLAN Profile
	Iname *string `json:"iname,omitempty"`
	// Name of the VLAN Profile
	Name *string `json:"name,omitempty"`
	// Is this VLAN profile the default for the network?
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile instantiates a new GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile {
	this := GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile{}
	return &this
}

// NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfileWithDefaults instantiates a new GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfileWithDefaults() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile {
	this := GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile{}
	return &this
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) GetIname() string {
	if o == nil || IsNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) GetInameOk() (*string, bool) {
	if o == nil || IsNil(o.Iname) {
		return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) HasIname() bool {
	if o != nil && !IsNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) SetIname(v string) {
	o.Iname = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) SetName(v string) {
	o.Name = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile struct {
	value *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile
	isSet bool
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) Get() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile {
	return v.value
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) Set(val *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile(val *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile {
	return &NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile{value: val, isSet: true}
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


