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

// UpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct for UpdateOrganizationApplianceVpnVpnFirewallRulesRequest
type UpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequest instantiates a new UpdateOrganizationApplianceVpnVpnFirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequest() *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest {
	this := UpdateOrganizationApplianceVpnVpnFirewallRulesRequest{}
	return &this
}

// NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestWithDefaults instantiates a new UpdateOrganizationApplianceVpnVpnFirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestWithDefaults() *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest {
	this := UpdateOrganizationApplianceVpnVpnFirewallRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) GetRules() []UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner {
	if o == nil || isNil(o.Rules) {
		var ret []UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) GetRulesOk() ([]UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) SetRules(v []UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) GetSyslogDefaultRule() bool {
	if o == nil || isNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogDefaultRule) {
    return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) HasSyslogDefaultRule() bool {
	if o != nil && !isNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct {
	value *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest
	isSet bool
}

func (v NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) Get() *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest {
	return v.value
}

func (v *NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) Set(val *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest(val *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) *NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest {
	return &NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


