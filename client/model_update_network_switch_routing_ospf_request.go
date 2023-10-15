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

// checks if the UpdateNetworkSwitchRoutingOspfRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchRoutingOspfRequest{}

// UpdateNetworkSwitchRoutingOspfRequest struct for UpdateNetworkSwitchRoutingOspfRequest
type UpdateNetworkSwitchRoutingOspfRequest struct {
	// Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	HelloTimerInSeconds *int32 `json:"helloTimerInSeconds,omitempty"`
	// Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	DeadTimerInSeconds *int32 `json:"deadTimerInSeconds,omitempty"`
	// OSPF areas
	Areas []UpdateNetworkSwitchRoutingOspfRequestAreasInner `json:"areas,omitempty"`
	V3 *UpdateNetworkSwitchRoutingOspfRequestV3 `json:"v3,omitempty"`
	// Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
	Md5AuthenticationEnabled *bool `json:"md5AuthenticationEnabled,omitempty"`
	Md5AuthenticationKey *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`
}

// NewUpdateNetworkSwitchRoutingOspfRequest instantiates a new UpdateNetworkSwitchRoutingOspfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingOspfRequest() *UpdateNetworkSwitchRoutingOspfRequest {
	this := UpdateNetworkSwitchRoutingOspfRequest{}
	return &this
}

// NewUpdateNetworkSwitchRoutingOspfRequestWithDefaults instantiates a new UpdateNetworkSwitchRoutingOspfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingOspfRequestWithDefaults() *UpdateNetworkSwitchRoutingOspfRequest {
	this := UpdateNetworkSwitchRoutingOspfRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHelloTimerInSeconds returns the HelloTimerInSeconds field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetHelloTimerInSeconds() int32 {
	if o == nil || IsNil(o.HelloTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.HelloTimerInSeconds
}

// GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetHelloTimerInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.HelloTimerInSeconds) {
		return nil, false
	}
	return o.HelloTimerInSeconds, true
}

// HasHelloTimerInSeconds returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasHelloTimerInSeconds() bool {
	if o != nil && !IsNil(o.HelloTimerInSeconds) {
		return true
	}

	return false
}

// SetHelloTimerInSeconds gets a reference to the given int32 and assigns it to the HelloTimerInSeconds field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetHelloTimerInSeconds(v int32) {
	o.HelloTimerInSeconds = &v
}

// GetDeadTimerInSeconds returns the DeadTimerInSeconds field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetDeadTimerInSeconds() int32 {
	if o == nil || IsNil(o.DeadTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.DeadTimerInSeconds
}

// GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetDeadTimerInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.DeadTimerInSeconds) {
		return nil, false
	}
	return o.DeadTimerInSeconds, true
}

// HasDeadTimerInSeconds returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasDeadTimerInSeconds() bool {
	if o != nil && !IsNil(o.DeadTimerInSeconds) {
		return true
	}

	return false
}

// SetDeadTimerInSeconds gets a reference to the given int32 and assigns it to the DeadTimerInSeconds field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetDeadTimerInSeconds(v int32) {
	o.DeadTimerInSeconds = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetAreas() []UpdateNetworkSwitchRoutingOspfRequestAreasInner {
	if o == nil || IsNil(o.Areas) {
		var ret []UpdateNetworkSwitchRoutingOspfRequestAreasInner
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetAreasOk() ([]UpdateNetworkSwitchRoutingOspfRequestAreasInner, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []UpdateNetworkSwitchRoutingOspfRequestAreasInner and assigns it to the Areas field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetAreas(v []UpdateNetworkSwitchRoutingOspfRequestAreasInner) {
	o.Areas = v
}

// GetV3 returns the V3 field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetV3() UpdateNetworkSwitchRoutingOspfRequestV3 {
	if o == nil || IsNil(o.V3) {
		var ret UpdateNetworkSwitchRoutingOspfRequestV3
		return ret
	}
	return *o.V3
}

// GetV3Ok returns a tuple with the V3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetV3Ok() (*UpdateNetworkSwitchRoutingOspfRequestV3, bool) {
	if o == nil || IsNil(o.V3) {
		return nil, false
	}
	return o.V3, true
}

// HasV3 returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasV3() bool {
	if o != nil && !IsNil(o.V3) {
		return true
	}

	return false
}

// SetV3 gets a reference to the given UpdateNetworkSwitchRoutingOspfRequestV3 and assigns it to the V3 field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetV3(v UpdateNetworkSwitchRoutingOspfRequestV3) {
	o.V3 = &v
}

// GetMd5AuthenticationEnabled returns the Md5AuthenticationEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetMd5AuthenticationEnabled() bool {
	if o == nil || IsNil(o.Md5AuthenticationEnabled) {
		var ret bool
		return ret
	}
	return *o.Md5AuthenticationEnabled
}

// GetMd5AuthenticationEnabledOk returns a tuple with the Md5AuthenticationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetMd5AuthenticationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Md5AuthenticationEnabled) {
		return nil, false
	}
	return o.Md5AuthenticationEnabled, true
}

// HasMd5AuthenticationEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasMd5AuthenticationEnabled() bool {
	if o != nil && !IsNil(o.Md5AuthenticationEnabled) {
		return true
	}

	return false
}

// SetMd5AuthenticationEnabled gets a reference to the given bool and assigns it to the Md5AuthenticationEnabled field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetMd5AuthenticationEnabled(v bool) {
	o.Md5AuthenticationEnabled = &v
}

// GetMd5AuthenticationKey returns the Md5AuthenticationKey field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetMd5AuthenticationKey() UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey {
	if o == nil || IsNil(o.Md5AuthenticationKey) {
		var ret UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey
		return ret
	}
	return *o.Md5AuthenticationKey
}

// GetMd5AuthenticationKeyOk returns a tuple with the Md5AuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) GetMd5AuthenticationKeyOk() (*UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey, bool) {
	if o == nil || IsNil(o.Md5AuthenticationKey) {
		return nil, false
	}
	return o.Md5AuthenticationKey, true
}

// HasMd5AuthenticationKey returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequest) HasMd5AuthenticationKey() bool {
	if o != nil && !IsNil(o.Md5AuthenticationKey) {
		return true
	}

	return false
}

// SetMd5AuthenticationKey gets a reference to the given UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey and assigns it to the Md5AuthenticationKey field.
func (o *UpdateNetworkSwitchRoutingOspfRequest) SetMd5AuthenticationKey(v UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) {
	o.Md5AuthenticationKey = &v
}

func (o UpdateNetworkSwitchRoutingOspfRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchRoutingOspfRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HelloTimerInSeconds) {
		toSerialize["helloTimerInSeconds"] = o.HelloTimerInSeconds
	}
	if !IsNil(o.DeadTimerInSeconds) {
		toSerialize["deadTimerInSeconds"] = o.DeadTimerInSeconds
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	if !IsNil(o.V3) {
		toSerialize["v3"] = o.V3
	}
	if !IsNil(o.Md5AuthenticationEnabled) {
		toSerialize["md5AuthenticationEnabled"] = o.Md5AuthenticationEnabled
	}
	if !IsNil(o.Md5AuthenticationKey) {
		toSerialize["md5AuthenticationKey"] = o.Md5AuthenticationKey
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchRoutingOspfRequest struct {
	value *UpdateNetworkSwitchRoutingOspfRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequest) Get() *UpdateNetworkSwitchRoutingOspfRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequest) Set(val *UpdateNetworkSwitchRoutingOspfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingOspfRequest(val *UpdateNetworkSwitchRoutingOspfRequest) *NullableUpdateNetworkSwitchRoutingOspfRequest {
	return &NullableUpdateNetworkSwitchRoutingOspfRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


