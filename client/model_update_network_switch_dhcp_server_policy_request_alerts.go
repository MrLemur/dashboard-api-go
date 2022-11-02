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

// UpdateNetworkSwitchDhcpServerPolicyRequestAlerts Alert settings for DHCP servers
type UpdateNetworkSwitchDhcpServerPolicyRequestAlerts struct {
	Email *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail `json:"email,omitempty"`
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestAlerts instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequestAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDhcpServerPolicyRequestAlerts() *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts {
	this := UpdateNetworkSwitchDhcpServerPolicyRequestAlerts{}
	return &this
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestAlertsWithDefaults instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequestAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDhcpServerPolicyRequestAlertsWithDefaults() *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts {
	this := UpdateNetworkSwitchDhcpServerPolicyRequestAlerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) GetEmail() UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) GetEmailOk() (*UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail and assigns it to the Email field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) SetEmail(v UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) {
	o.Email = &v
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts struct {
	value *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts
	isSet bool
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts) Get() *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts) Set(val *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts(val *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts {
	return &NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


