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

// checks if the GetNetworkSmDeviceDeviceCommandLogs200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmDeviceDeviceCommandLogs200ResponseInner{}

// GetNetworkSmDeviceDeviceCommandLogs200ResponseInner struct for GetNetworkSmDeviceDeviceCommandLogs200ResponseInner
type GetNetworkSmDeviceDeviceCommandLogs200ResponseInner struct {
	// The type of command sent to the device.
	Action *string `json:"action,omitempty"`
	// The name of the device to which the command is sent.
	Name *string `json:"name,omitempty"`
	// A JSON string object containing command details.
	Details *string `json:"details,omitempty"`
	// The Meraki dashboard user who initiated the command.
	DashboardUser *string `json:"dashboardUser,omitempty"`
	// The time the command was sent to the device.
	Ts *string `json:"ts,omitempty"`
}

// NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInner instantiates a new GetNetworkSmDeviceDeviceCommandLogs200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInner() *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner {
	this := GetNetworkSmDeviceDeviceCommandLogs200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceDeviceCommandLogs200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInnerWithDefaults() *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner {
	this := GetNetworkSmDeviceDeviceCommandLogs200ResponseInner{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetAction(v string) {
	o.Action = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetDetails(v string) {
	o.Details = &v
}

// GetDashboardUser returns the DashboardUser field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDashboardUser() string {
	if o == nil || IsNil(o.DashboardUser) {
		var ret string
		return ret
	}
	return *o.DashboardUser
}

// GetDashboardUserOk returns a tuple with the DashboardUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDashboardUserOk() (*string, bool) {
	if o == nil || IsNil(o.DashboardUser) {
		return nil, false
	}
	return o.DashboardUser, true
}

// HasDashboardUser returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasDashboardUser() bool {
	if o != nil && !IsNil(o.DashboardUser) {
		return true
	}

	return false
}

// SetDashboardUser gets a reference to the given string and assigns it to the DashboardUser field.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetDashboardUser(v string) {
	o.DashboardUser = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.DashboardUser) {
		toSerialize["dashboardUser"] = o.DashboardUser
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner struct {
	value *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner) Get() *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner) Set(val *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner(val *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) *NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner {
	return &NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceDeviceCommandLogs200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


