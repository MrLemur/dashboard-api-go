/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject201 struct for InlineObject201
type InlineObject201 struct {
	// The name of the VoIP provider
	Name *string `json:"name,omitempty"`
	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address *string `json:"address,omitempty"`
	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
	BestEffortMonitoringEnabled *bool `json:"bestEffortMonitoringEnabled,omitempty"`
}

// NewInlineObject201 instantiates a new InlineObject201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject201() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// NewInlineObject201WithDefaults instantiates a new InlineObject201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject201WithDefaults() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject201) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject201) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject201) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineObject201) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineObject201) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineObject201) SetAddress(v string) {
	o.Address = &v
}

// GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field value if set, zero value otherwise.
func (o *InlineObject201) GetBestEffortMonitoringEnabled() bool {
	if o == nil || isNil(o.BestEffortMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.BestEffortMonitoringEnabled
}

// GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetBestEffortMonitoringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BestEffortMonitoringEnabled) {
    return nil, false
	}
	return o.BestEffortMonitoringEnabled, true
}

// HasBestEffortMonitoringEnabled returns a boolean if a field has been set.
func (o *InlineObject201) HasBestEffortMonitoringEnabled() bool {
	if o != nil && !isNil(o.BestEffortMonitoringEnabled) {
		return true
	}

	return false
}

// SetBestEffortMonitoringEnabled gets a reference to the given bool and assigns it to the BestEffortMonitoringEnabled field.
func (o *InlineObject201) SetBestEffortMonitoringEnabled(v bool) {
	o.BestEffortMonitoringEnabled = &v
}

func (o InlineObject201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.BestEffortMonitoringEnabled) {
		toSerialize["bestEffortMonitoringEnabled"] = o.BestEffortMonitoringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject201 struct {
	value *InlineObject201
	isSet bool
}

func (v NullableInlineObject201) Get() *InlineObject201 {
	return v.value
}

func (v *NullableInlineObject201) Set(val *InlineObject201) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject201) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject201(val *InlineObject201) *NullableInlineObject201 {
	return &NullableInlineObject201{value: val, isSet: true}
}

func (v NullableInlineObject201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


