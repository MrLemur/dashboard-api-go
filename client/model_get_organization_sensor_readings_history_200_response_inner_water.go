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

// GetOrganizationSensorReadingsHistory200ResponseInnerWater Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.
type GetOrganizationSensorReadingsHistory200ResponseInnerWater struct {
	// True if water is detected.
	Present *bool `json:"present,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerWater instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerWater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerWater() *GetOrganizationSensorReadingsHistory200ResponseInnerWater {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerWater{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerWaterWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerWater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerWaterWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerWater {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerWater{}
	return &this
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerWater) GetPresent() bool {
	if o == nil || isNil(o.Present) {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerWater) GetPresentOk() (*bool, bool) {
	if o == nil || isNil(o.Present) {
    return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerWater) HasPresent() bool {
	if o != nil && !isNil(o.Present) {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerWater) SetPresent(v bool) {
	o.Present = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerWater) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Present) {
		toSerialize["present"] = o.Present
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerWater
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerWater {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerWater) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerWater(val *GetOrganizationSensorReadingsHistory200ResponseInnerWater) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerWater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


