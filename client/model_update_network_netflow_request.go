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

// UpdateNetworkNetflowRequest struct for UpdateNetworkNetflowRequest
type UpdateNetworkNetflowRequest struct {
	// Boolean indicating whether NetFlow traffic reporting is enabled (true) or disabled (false).
	ReportingEnabled *bool `json:"reportingEnabled,omitempty"`
	// The IPv4 address of the NetFlow collector.
	CollectorIp *string `json:"collectorIp,omitempty"`
	// The port that the NetFlow collector will be listening on.
	CollectorPort *int32 `json:"collectorPort,omitempty"`
	// Boolean indicating whether Encrypted Traffic Analytics is enabled (true) or disabled (false).
	EtaEnabled *bool `json:"etaEnabled,omitempty"`
	// The port that the Encrypted Traffic Analytics collector will be listening on.
	EtaDstPort *int32 `json:"etaDstPort,omitempty"`
}

// NewUpdateNetworkNetflowRequest instantiates a new UpdateNetworkNetflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkNetflowRequest() *UpdateNetworkNetflowRequest {
	this := UpdateNetworkNetflowRequest{}
	return &this
}

// NewUpdateNetworkNetflowRequestWithDefaults instantiates a new UpdateNetworkNetflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkNetflowRequestWithDefaults() *UpdateNetworkNetflowRequest {
	this := UpdateNetworkNetflowRequest{}
	return &this
}

// GetReportingEnabled returns the ReportingEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkNetflowRequest) GetReportingEnabled() bool {
	if o == nil || isNil(o.ReportingEnabled) {
		var ret bool
		return ret
	}
	return *o.ReportingEnabled
}

// GetReportingEnabledOk returns a tuple with the ReportingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkNetflowRequest) GetReportingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ReportingEnabled) {
    return nil, false
	}
	return o.ReportingEnabled, true
}

// HasReportingEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkNetflowRequest) HasReportingEnabled() bool {
	if o != nil && !isNil(o.ReportingEnabled) {
		return true
	}

	return false
}

// SetReportingEnabled gets a reference to the given bool and assigns it to the ReportingEnabled field.
func (o *UpdateNetworkNetflowRequest) SetReportingEnabled(v bool) {
	o.ReportingEnabled = &v
}

// GetCollectorIp returns the CollectorIp field value if set, zero value otherwise.
func (o *UpdateNetworkNetflowRequest) GetCollectorIp() string {
	if o == nil || isNil(o.CollectorIp) {
		var ret string
		return ret
	}
	return *o.CollectorIp
}

// GetCollectorIpOk returns a tuple with the CollectorIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkNetflowRequest) GetCollectorIpOk() (*string, bool) {
	if o == nil || isNil(o.CollectorIp) {
    return nil, false
	}
	return o.CollectorIp, true
}

// HasCollectorIp returns a boolean if a field has been set.
func (o *UpdateNetworkNetflowRequest) HasCollectorIp() bool {
	if o != nil && !isNil(o.CollectorIp) {
		return true
	}

	return false
}

// SetCollectorIp gets a reference to the given string and assigns it to the CollectorIp field.
func (o *UpdateNetworkNetflowRequest) SetCollectorIp(v string) {
	o.CollectorIp = &v
}

// GetCollectorPort returns the CollectorPort field value if set, zero value otherwise.
func (o *UpdateNetworkNetflowRequest) GetCollectorPort() int32 {
	if o == nil || isNil(o.CollectorPort) {
		var ret int32
		return ret
	}
	return *o.CollectorPort
}

// GetCollectorPortOk returns a tuple with the CollectorPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkNetflowRequest) GetCollectorPortOk() (*int32, bool) {
	if o == nil || isNil(o.CollectorPort) {
    return nil, false
	}
	return o.CollectorPort, true
}

// HasCollectorPort returns a boolean if a field has been set.
func (o *UpdateNetworkNetflowRequest) HasCollectorPort() bool {
	if o != nil && !isNil(o.CollectorPort) {
		return true
	}

	return false
}

// SetCollectorPort gets a reference to the given int32 and assigns it to the CollectorPort field.
func (o *UpdateNetworkNetflowRequest) SetCollectorPort(v int32) {
	o.CollectorPort = &v
}

// GetEtaEnabled returns the EtaEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkNetflowRequest) GetEtaEnabled() bool {
	if o == nil || isNil(o.EtaEnabled) {
		var ret bool
		return ret
	}
	return *o.EtaEnabled
}

// GetEtaEnabledOk returns a tuple with the EtaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkNetflowRequest) GetEtaEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EtaEnabled) {
    return nil, false
	}
	return o.EtaEnabled, true
}

// HasEtaEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkNetflowRequest) HasEtaEnabled() bool {
	if o != nil && !isNil(o.EtaEnabled) {
		return true
	}

	return false
}

// SetEtaEnabled gets a reference to the given bool and assigns it to the EtaEnabled field.
func (o *UpdateNetworkNetflowRequest) SetEtaEnabled(v bool) {
	o.EtaEnabled = &v
}

// GetEtaDstPort returns the EtaDstPort field value if set, zero value otherwise.
func (o *UpdateNetworkNetflowRequest) GetEtaDstPort() int32 {
	if o == nil || isNil(o.EtaDstPort) {
		var ret int32
		return ret
	}
	return *o.EtaDstPort
}

// GetEtaDstPortOk returns a tuple with the EtaDstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkNetflowRequest) GetEtaDstPortOk() (*int32, bool) {
	if o == nil || isNil(o.EtaDstPort) {
    return nil, false
	}
	return o.EtaDstPort, true
}

// HasEtaDstPort returns a boolean if a field has been set.
func (o *UpdateNetworkNetflowRequest) HasEtaDstPort() bool {
	if o != nil && !isNil(o.EtaDstPort) {
		return true
	}

	return false
}

// SetEtaDstPort gets a reference to the given int32 and assigns it to the EtaDstPort field.
func (o *UpdateNetworkNetflowRequest) SetEtaDstPort(v int32) {
	o.EtaDstPort = &v
}

func (o UpdateNetworkNetflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportingEnabled) {
		toSerialize["reportingEnabled"] = o.ReportingEnabled
	}
	if !isNil(o.CollectorIp) {
		toSerialize["collectorIp"] = o.CollectorIp
	}
	if !isNil(o.CollectorPort) {
		toSerialize["collectorPort"] = o.CollectorPort
	}
	if !isNil(o.EtaEnabled) {
		toSerialize["etaEnabled"] = o.EtaEnabled
	}
	if !isNil(o.EtaDstPort) {
		toSerialize["etaDstPort"] = o.EtaDstPort
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkNetflowRequest struct {
	value *UpdateNetworkNetflowRequest
	isSet bool
}

func (v NullableUpdateNetworkNetflowRequest) Get() *UpdateNetworkNetflowRequest {
	return v.value
}

func (v *NullableUpdateNetworkNetflowRequest) Set(val *UpdateNetworkNetflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkNetflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkNetflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkNetflowRequest(val *UpdateNetworkNetflowRequest) *NullableUpdateNetworkNetflowRequest {
	return &NullableUpdateNetworkNetflowRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkNetflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkNetflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


