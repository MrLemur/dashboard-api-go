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

// UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest struct for UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest struct {
	BandwidthLimits *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) GetBandwidthLimits() UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) GetBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) SetBandwidthLimits(v UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) Get() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) Set(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


