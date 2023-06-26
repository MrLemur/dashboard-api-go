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

// NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality struct {
	// Alerting threshold as indoor air quality score.
	Score *int32 `json:"score,omitempty"`
	// Alerting threshold as a qualitative indoor air quality level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality() *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQualityWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQualityWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) GetScore() int32 {
	if o == nil || isNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) GetScoreOk() (*int32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) SetScore(v int32) {
	o.Score = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality(val *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


