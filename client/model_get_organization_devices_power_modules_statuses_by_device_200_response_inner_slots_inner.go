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

// GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner struct for GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner
type GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner struct {
	// Which slot the AC power supply occupies. Possible values are: 0, 1, 2.
	Number *int32 `json:"number,omitempty"`
	// The power supply unit serial number.
	Serial *string `json:"serial,omitempty"`
	// The power supply unit model.
	Model *string `json:"model,omitempty"`
	// Status of the power supply unit. Possible values are: connected, not connected, powering.
	Status *string `json:"status,omitempty"`
}

// NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner instantiates a new GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner {
	this := GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner{}
	return &this
}

// NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInnerWithDefaults instantiates a new GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInnerWithDefaults() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner {
	this := GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) SetNumber(v int32) {
	o.Number = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) SetModel(v string) {
	o.Model = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner struct {
	value *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner
	isSet bool
}

func (v NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) Get() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) Set(val *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner(val *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner {
	return &NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


