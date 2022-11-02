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

// UpdateNetworkSettingsRequest struct for UpdateNetworkSettingsRequest
type UpdateNetworkSettingsRequest struct {
	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled *bool `json:"localStatusPageEnabled,omitempty"`
	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled *bool `json:"remoteStatusPageEnabled,omitempty"`
	SecurePort *UpdateNetworkSettingsRequestSecurePort `json:"securePort,omitempty"`
	LocalStatusPage *UpdateNetworkSettingsRequestLocalStatusPage `json:"localStatusPage,omitempty"`
}

// NewUpdateNetworkSettingsRequest instantiates a new UpdateNetworkSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSettingsRequest() *UpdateNetworkSettingsRequest {
	this := UpdateNetworkSettingsRequest{}
	return &this
}

// NewUpdateNetworkSettingsRequestWithDefaults instantiates a new UpdateNetworkSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSettingsRequestWithDefaults() *UpdateNetworkSettingsRequest {
	this := UpdateNetworkSettingsRequest{}
	return &this
}

// GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequest) GetLocalStatusPageEnabled() bool {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.LocalStatusPageEnabled
}

// GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequest) GetLocalStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
    return nil, false
	}
	return o.LocalStatusPageEnabled, true
}

// HasLocalStatusPageEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequest) HasLocalStatusPageEnabled() bool {
	if o != nil && !isNil(o.LocalStatusPageEnabled) {
		return true
	}

	return false
}

// SetLocalStatusPageEnabled gets a reference to the given bool and assigns it to the LocalStatusPageEnabled field.
func (o *UpdateNetworkSettingsRequest) SetLocalStatusPageEnabled(v bool) {
	o.LocalStatusPageEnabled = &v
}

// GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequest) GetRemoteStatusPageEnabled() bool {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.RemoteStatusPageEnabled
}

// GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequest) GetRemoteStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
    return nil, false
	}
	return o.RemoteStatusPageEnabled, true
}

// HasRemoteStatusPageEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequest) HasRemoteStatusPageEnabled() bool {
	if o != nil && !isNil(o.RemoteStatusPageEnabled) {
		return true
	}

	return false
}

// SetRemoteStatusPageEnabled gets a reference to the given bool and assigns it to the RemoteStatusPageEnabled field.
func (o *UpdateNetworkSettingsRequest) SetRemoteStatusPageEnabled(v bool) {
	o.RemoteStatusPageEnabled = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequest) GetSecurePort() UpdateNetworkSettingsRequestSecurePort {
	if o == nil || isNil(o.SecurePort) {
		var ret UpdateNetworkSettingsRequestSecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequest) GetSecurePortOk() (*UpdateNetworkSettingsRequestSecurePort, bool) {
	if o == nil || isNil(o.SecurePort) {
    return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequest) HasSecurePort() bool {
	if o != nil && !isNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given UpdateNetworkSettingsRequestSecurePort and assigns it to the SecurePort field.
func (o *UpdateNetworkSettingsRequest) SetSecurePort(v UpdateNetworkSettingsRequestSecurePort) {
	o.SecurePort = &v
}

// GetLocalStatusPage returns the LocalStatusPage field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequest) GetLocalStatusPage() UpdateNetworkSettingsRequestLocalStatusPage {
	if o == nil || isNil(o.LocalStatusPage) {
		var ret UpdateNetworkSettingsRequestLocalStatusPage
		return ret
	}
	return *o.LocalStatusPage
}

// GetLocalStatusPageOk returns a tuple with the LocalStatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequest) GetLocalStatusPageOk() (*UpdateNetworkSettingsRequestLocalStatusPage, bool) {
	if o == nil || isNil(o.LocalStatusPage) {
    return nil, false
	}
	return o.LocalStatusPage, true
}

// HasLocalStatusPage returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequest) HasLocalStatusPage() bool {
	if o != nil && !isNil(o.LocalStatusPage) {
		return true
	}

	return false
}

// SetLocalStatusPage gets a reference to the given UpdateNetworkSettingsRequestLocalStatusPage and assigns it to the LocalStatusPage field.
func (o *UpdateNetworkSettingsRequest) SetLocalStatusPage(v UpdateNetworkSettingsRequestLocalStatusPage) {
	o.LocalStatusPage = &v
}

func (o UpdateNetworkSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocalStatusPageEnabled) {
		toSerialize["localStatusPageEnabled"] = o.LocalStatusPageEnabled
	}
	if !isNil(o.RemoteStatusPageEnabled) {
		toSerialize["remoteStatusPageEnabled"] = o.RemoteStatusPageEnabled
	}
	if !isNil(o.SecurePort) {
		toSerialize["securePort"] = o.SecurePort
	}
	if !isNil(o.LocalStatusPage) {
		toSerialize["localStatusPage"] = o.LocalStatusPage
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSettingsRequest struct {
	value *UpdateNetworkSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkSettingsRequest) Get() *UpdateNetworkSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkSettingsRequest) Set(val *UpdateNetworkSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSettingsRequest(val *UpdateNetworkSettingsRequest) *NullableUpdateNetworkSettingsRequest {
	return &NullableUpdateNetworkSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


