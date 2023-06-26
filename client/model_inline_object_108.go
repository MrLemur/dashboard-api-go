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

// InlineObject108 struct for InlineObject108
type InlineObject108 struct {
	// The wifiMac of the device to be wiped.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The id of the device to be wiped.
	Id *string `json:"id,omitempty"`
	// The serial of the device to be wiped.
	Serial *string `json:"serial,omitempty"`
	// The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices.
	Pin *int32 `json:"pin,omitempty"`
}

// NewInlineObject108 instantiates a new InlineObject108 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject108() *InlineObject108 {
	this := InlineObject108{}
	return &this
}

// NewInlineObject108WithDefaults instantiates a new InlineObject108 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject108WithDefaults() *InlineObject108 {
	this := InlineObject108{}
	return &this
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *InlineObject108) GetWifiMac() string {
	if o == nil || isNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetWifiMacOk() (*string, bool) {
	if o == nil || isNil(o.WifiMac) {
    return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *InlineObject108) HasWifiMac() bool {
	if o != nil && !isNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *InlineObject108) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineObject108) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject108) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject108) SetId(v string) {
	o.Id = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineObject108) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineObject108) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineObject108) SetSerial(v string) {
	o.Serial = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *InlineObject108) GetPin() int32 {
	if o == nil || isNil(o.Pin) {
		var ret int32
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetPinOk() (*int32, bool) {
	if o == nil || isNil(o.Pin) {
    return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *InlineObject108) HasPin() bool {
	if o != nil && !isNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given int32 and assigns it to the Pin field.
func (o *InlineObject108) SetPin(v int32) {
	o.Pin = &v
}

func (o InlineObject108) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject108 struct {
	value *InlineObject108
	isSet bool
}

func (v NullableInlineObject108) Get() *InlineObject108 {
	return v.value
}

func (v *NullableInlineObject108) Set(val *InlineObject108) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject108) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject108) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject108(val *InlineObject108) *NullableInlineObject108 {
	return &NullableInlineObject108{value: val, isSet: true}
}

func (v NullableInlineObject108) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject108) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


