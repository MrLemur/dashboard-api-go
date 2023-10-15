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

// checks if the GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner{}

// GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner struct for GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
type GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner struct {
	// Name of the Device
	Name *string `json:"name,omitempty"`
	// Serial of the Device
	Serial *string `json:"serial,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// The product type
	ProductType *string `json:"productType,omitempty"`
	VlanProfile *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile `json:"vlanProfile,omitempty"`
	Stack *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack `json:"stack,omitempty"`
}

// NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner instantiates a new GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner {
	this := GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner{}
	return &this
}

// NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerWithDefaults instantiates a new GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerWithDefaults() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner {
	this := GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetVlanProfile() GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile {
	if o == nil || IsNil(o.VlanProfile) {
		var ret GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetVlanProfileOk() (*GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile, bool) {
	if o == nil || IsNil(o.VlanProfile) {
		return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) HasVlanProfile() bool {
	if o != nil && !IsNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile and assigns it to the VlanProfile field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) SetVlanProfile(v GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerVlanProfile) {
	o.VlanProfile = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetStack() GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack {
	if o == nil || IsNil(o.Stack) {
		var ret GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) GetStackOk() (*GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) HasStack() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack and assigns it to the Stack field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) SetStack(v GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) {
	o.Stack = &v
}

func (o GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.VlanProfile) {
		toSerialize["vlanProfile"] = o.VlanProfile
	}
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	return toSerialize, nil
}

type NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner struct {
	value *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
	isSet bool
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) Get() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) Set(val *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner(val *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner {
	return &NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


