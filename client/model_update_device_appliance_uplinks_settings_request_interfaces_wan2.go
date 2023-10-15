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

// checks if the UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2{}

// UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 WAN 2 settings.
type UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 struct {
	// Enable or disable the interface.
	Enabled *bool `json:"enabled,omitempty"`
	VlanTagging *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging `json:"vlanTagging,omitempty"`
	Svis *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis `json:"svis,omitempty"`
	Pppoe *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe `json:"pppoe,omitempty"`
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2{}
	return &this
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2WithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2WithDefaults() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetVlanTagging() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging {
	if o == nil || IsNil(o.VlanTagging) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetVlanTaggingOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging, bool) {
	if o == nil || IsNil(o.VlanTagging) {
		return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasVlanTagging() bool {
	if o != nil && !IsNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging and assigns it to the VlanTagging field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetVlanTagging(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) {
	o.VlanTagging = &v
}

// GetSvis returns the Svis field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetSvis() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis {
	if o == nil || IsNil(o.Svis) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis
		return ret
	}
	return *o.Svis
}

// GetSvisOk returns a tuple with the Svis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetSvisOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis, bool) {
	if o == nil || IsNil(o.Svis) {
		return nil, false
	}
	return o.Svis, true
}

// HasSvis returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasSvis() bool {
	if o != nil && !IsNil(o.Svis) {
		return true
	}

	return false
}

// SetSvis gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis and assigns it to the Svis field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetSvis(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis) {
	o.Svis = &v
}

// GetPppoe returns the Pppoe field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetPppoe() UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe {
	if o == nil || IsNil(o.Pppoe) {
		var ret UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe
		return ret
	}
	return *o.Pppoe
}

// GetPppoeOk returns a tuple with the Pppoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetPppoeOk() (*UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe, bool) {
	if o == nil || IsNil(o.Pppoe) {
		return nil, false
	}
	return o.Pppoe, true
}

// HasPppoe returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasPppoe() bool {
	if o != nil && !IsNil(o.Pppoe) {
		return true
	}

	return false
}

// SetPppoe gets a reference to the given UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe and assigns it to the Pppoe field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetPppoe(v UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) {
	o.Pppoe = &v
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.VlanTagging) {
		toSerialize["vlanTagging"] = o.VlanTagging
	}
	if !IsNil(o.Svis) {
		toSerialize["svis"] = o.Svis
	}
	if !IsNil(o.Pppoe) {
		toSerialize["pppoe"] = o.Pppoe
	}
	return toSerialize, nil
}

type NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 struct {
	value *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2
	isSet bool
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) Get() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 {
	return v.value
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) Set(val *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2(val *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 {
	return &NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


