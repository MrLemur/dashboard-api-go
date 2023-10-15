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

// checks if the UpdateDeviceCellularSimsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularSimsRequest{}

// UpdateDeviceCellularSimsRequest struct for UpdateDeviceCellularSimsRequest
type UpdateDeviceCellularSimsRequest struct {
	// List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
	Sims []UpdateDeviceCellularSimsRequestSimsInner `json:"sims,omitempty"`
	SimFailover *UpdateDeviceCellularSimsRequestSimFailover `json:"simFailover,omitempty"`
}

// NewUpdateDeviceCellularSimsRequest instantiates a new UpdateDeviceCellularSimsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularSimsRequest() *UpdateDeviceCellularSimsRequest {
	this := UpdateDeviceCellularSimsRequest{}
	return &this
}

// NewUpdateDeviceCellularSimsRequestWithDefaults instantiates a new UpdateDeviceCellularSimsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularSimsRequestWithDefaults() *UpdateDeviceCellularSimsRequest {
	this := UpdateDeviceCellularSimsRequest{}
	return &this
}

// GetSims returns the Sims field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequest) GetSims() []UpdateDeviceCellularSimsRequestSimsInner {
	if o == nil || IsNil(o.Sims) {
		var ret []UpdateDeviceCellularSimsRequestSimsInner
		return ret
	}
	return o.Sims
}

// GetSimsOk returns a tuple with the Sims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequest) GetSimsOk() ([]UpdateDeviceCellularSimsRequestSimsInner, bool) {
	if o == nil || IsNil(o.Sims) {
		return nil, false
	}
	return o.Sims, true
}

// HasSims returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequest) HasSims() bool {
	if o != nil && !IsNil(o.Sims) {
		return true
	}

	return false
}

// SetSims gets a reference to the given []UpdateDeviceCellularSimsRequestSimsInner and assigns it to the Sims field.
func (o *UpdateDeviceCellularSimsRequest) SetSims(v []UpdateDeviceCellularSimsRequestSimsInner) {
	o.Sims = v
}

// GetSimFailover returns the SimFailover field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequest) GetSimFailover() UpdateDeviceCellularSimsRequestSimFailover {
	if o == nil || IsNil(o.SimFailover) {
		var ret UpdateDeviceCellularSimsRequestSimFailover
		return ret
	}
	return *o.SimFailover
}

// GetSimFailoverOk returns a tuple with the SimFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequest) GetSimFailoverOk() (*UpdateDeviceCellularSimsRequestSimFailover, bool) {
	if o == nil || IsNil(o.SimFailover) {
		return nil, false
	}
	return o.SimFailover, true
}

// HasSimFailover returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequest) HasSimFailover() bool {
	if o != nil && !IsNil(o.SimFailover) {
		return true
	}

	return false
}

// SetSimFailover gets a reference to the given UpdateDeviceCellularSimsRequestSimFailover and assigns it to the SimFailover field.
func (o *UpdateDeviceCellularSimsRequest) SetSimFailover(v UpdateDeviceCellularSimsRequestSimFailover) {
	o.SimFailover = &v
}

func (o UpdateDeviceCellularSimsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularSimsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sims) {
		toSerialize["sims"] = o.Sims
	}
	if !IsNil(o.SimFailover) {
		toSerialize["simFailover"] = o.SimFailover
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCellularSimsRequest struct {
	value *UpdateDeviceCellularSimsRequest
	isSet bool
}

func (v NullableUpdateDeviceCellularSimsRequest) Get() *UpdateDeviceCellularSimsRequest {
	return v.value
}

func (v *NullableUpdateDeviceCellularSimsRequest) Set(val *UpdateDeviceCellularSimsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularSimsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularSimsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularSimsRequest(val *UpdateDeviceCellularSimsRequest) *NullableUpdateDeviceCellularSimsRequest {
	return &NullableUpdateDeviceCellularSimsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularSimsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularSimsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


