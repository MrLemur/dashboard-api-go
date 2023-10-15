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

// checks if the GetDeviceSwitchPortsStatuses200ResponseInnerLldp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchPortsStatuses200ResponseInnerLldp{}

// GetDeviceSwitchPortsStatuses200ResponseInnerLldp The Link Layer Discovery Protocol (LLDP) information of the connected device.
type GetDeviceSwitchPortsStatuses200ResponseInnerLldp struct {
	// The device's system name.
	SystemName *string `json:"systemName,omitempty"`
	// The device's system description.
	SystemDescription *string `json:"systemDescription,omitempty"`
	// The device's chassis ID.
	ChassisId *string `json:"chassisId,omitempty"`
	// Identifies the port from which the LLDP packet was sent
	PortId *string `json:"portId,omitempty"`
	// The device's management VLAN.
	ManagementVlan *int32 `json:"managementVlan,omitempty"`
	// The port's VLAN.
	PortVlan *int32 `json:"portVlan,omitempty"`
	// The device's management IP.
	ManagementAddress *string `json:"managementAddress,omitempty"`
	// Description of the port from which the LLDP packet was sent.
	PortDescription *string `json:"portDescription,omitempty"`
	// Identifies the device type, which indicates the functional capabilities of the device.
	SystemCapabilities *string `json:"systemCapabilities,omitempty"`
}

// NewGetDeviceSwitchPortsStatuses200ResponseInnerLldp instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerLldp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchPortsStatuses200ResponseInnerLldp() *GetDeviceSwitchPortsStatuses200ResponseInnerLldp {
	this := GetDeviceSwitchPortsStatuses200ResponseInnerLldp{}
	return &this
}

// NewGetDeviceSwitchPortsStatuses200ResponseInnerLldpWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerLldp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchPortsStatuses200ResponseInnerLldpWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInnerLldp {
	this := GetDeviceSwitchPortsStatuses200ResponseInnerLldp{}
	return &this
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetSystemName(v string) {
	o.SystemName = &v
}

// GetSystemDescription returns the SystemDescription field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemDescription() string {
	if o == nil || IsNil(o.SystemDescription) {
		var ret string
		return ret
	}
	return *o.SystemDescription
}

// GetSystemDescriptionOk returns a tuple with the SystemDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SystemDescription) {
		return nil, false
	}
	return o.SystemDescription, true
}

// HasSystemDescription returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasSystemDescription() bool {
	if o != nil && !IsNil(o.SystemDescription) {
		return true
	}

	return false
}

// SetSystemDescription gets a reference to the given string and assigns it to the SystemDescription field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetSystemDescription(v string) {
	o.SystemDescription = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetChassisId() string {
	if o == nil || IsNil(o.ChassisId) {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetChassisIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetPortId(v string) {
	o.PortId = &v
}

// GetManagementVlan returns the ManagementVlan field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementVlan() int32 {
	if o == nil || IsNil(o.ManagementVlan) {
		var ret int32
		return ret
	}
	return *o.ManagementVlan
}

// GetManagementVlanOk returns a tuple with the ManagementVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.ManagementVlan) {
		return nil, false
	}
	return o.ManagementVlan, true
}

// HasManagementVlan returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasManagementVlan() bool {
	if o != nil && !IsNil(o.ManagementVlan) {
		return true
	}

	return false
}

// SetManagementVlan gets a reference to the given int32 and assigns it to the ManagementVlan field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetManagementVlan(v int32) {
	o.ManagementVlan = &v
}

// GetPortVlan returns the PortVlan field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortVlan() int32 {
	if o == nil || IsNil(o.PortVlan) {
		var ret int32
		return ret
	}
	return *o.PortVlan
}

// GetPortVlanOk returns a tuple with the PortVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.PortVlan) {
		return nil, false
	}
	return o.PortVlan, true
}

// HasPortVlan returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasPortVlan() bool {
	if o != nil && !IsNil(o.PortVlan) {
		return true
	}

	return false
}

// SetPortVlan gets a reference to the given int32 and assigns it to the PortVlan field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetPortVlan(v int32) {
	o.PortVlan = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementAddress() string {
	if o == nil || IsNil(o.ManagementAddress) {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementAddress) {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasManagementAddress() bool {
	if o != nil && !IsNil(o.ManagementAddress) {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetPortDescription returns the PortDescription field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortDescription() string {
	if o == nil || IsNil(o.PortDescription) {
		var ret string
		return ret
	}
	return *o.PortDescription
}

// GetPortDescriptionOk returns a tuple with the PortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PortDescription) {
		return nil, false
	}
	return o.PortDescription, true
}

// HasPortDescription returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasPortDescription() bool {
	if o != nil && !IsNil(o.PortDescription) {
		return true
	}

	return false
}

// SetPortDescription gets a reference to the given string and assigns it to the PortDescription field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetPortDescription(v string) {
	o.PortDescription = &v
}

// GetSystemCapabilities returns the SystemCapabilities field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemCapabilities() string {
	if o == nil || IsNil(o.SystemCapabilities) {
		var ret string
		return ret
	}
	return *o.SystemCapabilities
}

// GetSystemCapabilitiesOk returns a tuple with the SystemCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemCapabilitiesOk() (*string, bool) {
	if o == nil || IsNil(o.SystemCapabilities) {
		return nil, false
	}
	return o.SystemCapabilities, true
}

// HasSystemCapabilities returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasSystemCapabilities() bool {
	if o != nil && !IsNil(o.SystemCapabilities) {
		return true
	}

	return false
}

// SetSystemCapabilities gets a reference to the given string and assigns it to the SystemCapabilities field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetSystemCapabilities(v string) {
	o.SystemCapabilities = &v
}

func (o GetDeviceSwitchPortsStatuses200ResponseInnerLldp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchPortsStatuses200ResponseInnerLldp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !IsNil(o.SystemDescription) {
		toSerialize["systemDescription"] = o.SystemDescription
	}
	if !IsNil(o.ChassisId) {
		toSerialize["chassisId"] = o.ChassisId
	}
	if !IsNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !IsNil(o.ManagementVlan) {
		toSerialize["managementVlan"] = o.ManagementVlan
	}
	if !IsNil(o.PortVlan) {
		toSerialize["portVlan"] = o.PortVlan
	}
	if !IsNil(o.ManagementAddress) {
		toSerialize["managementAddress"] = o.ManagementAddress
	}
	if !IsNil(o.PortDescription) {
		toSerialize["portDescription"] = o.PortDescription
	}
	if !IsNil(o.SystemCapabilities) {
		toSerialize["systemCapabilities"] = o.SystemCapabilities
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp struct {
	value *GetDeviceSwitchPortsStatuses200ResponseInnerLldp
	isSet bool
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp) Get() *GetDeviceSwitchPortsStatuses200ResponseInnerLldp {
	return v.value
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp) Set(val *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp(val *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) *NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp {
	return &NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerLldp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


