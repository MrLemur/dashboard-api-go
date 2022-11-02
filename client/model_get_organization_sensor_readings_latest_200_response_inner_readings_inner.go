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

// GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner struct for GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner
type GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner struct {
	// Time at which the reading occurred, in ISO8601 format.
	Ts *string `json:"ts,omitempty"`
	// Type of sensor reading.
	Metric *string `json:"metric,omitempty"`
	Battery *GetOrganizationSensorReadingsHistory200ResponseInnerBattery `json:"battery,omitempty"`
	Button *GetOrganizationSensorReadingsHistory200ResponseInnerButton `json:"button,omitempty"`
	Door *GetOrganizationSensorReadingsHistory200ResponseInnerDoor `json:"door,omitempty"`
	Humidity *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity `json:"humidity,omitempty"`
	IndoorAirQuality *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality `json:"indoorAirQuality,omitempty"`
	Noise *GetOrganizationSensorReadingsHistory200ResponseInnerNoise `json:"noise,omitempty"`
	Pm25 *GetOrganizationSensorReadingsHistory200ResponseInnerPm25 `json:"pm25,omitempty"`
	Temperature *GetOrganizationSensorReadingsHistory200ResponseInnerTemperature `json:"temperature,omitempty"`
	Tvoc *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc `json:"tvoc,omitempty"`
	Water *GetOrganizationSensorReadingsHistory200ResponseInnerWater `json:"water,omitempty"`
}

// NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner instantiates a new GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner() *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner {
	this := GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner{}
	return &this
}

// NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInnerWithDefaults instantiates a new GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInnerWithDefaults() *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner {
	this := GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetTs(v string) {
	o.Ts = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetMetric() string {
	if o == nil || isNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetMetricOk() (*string, bool) {
	if o == nil || isNil(o.Metric) {
    return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasMetric() bool {
	if o != nil && !isNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetMetric(v string) {
	o.Metric = &v
}

// GetBattery returns the Battery field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetBattery() GetOrganizationSensorReadingsHistory200ResponseInnerBattery {
	if o == nil || isNil(o.Battery) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerBattery
		return ret
	}
	return *o.Battery
}

// GetBatteryOk returns a tuple with the Battery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetBatteryOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerBattery, bool) {
	if o == nil || isNil(o.Battery) {
    return nil, false
	}
	return o.Battery, true
}

// HasBattery returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasBattery() bool {
	if o != nil && !isNil(o.Battery) {
		return true
	}

	return false
}

// SetBattery gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerBattery and assigns it to the Battery field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetBattery(v GetOrganizationSensorReadingsHistory200ResponseInnerBattery) {
	o.Battery = &v
}

// GetButton returns the Button field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetButton() GetOrganizationSensorReadingsHistory200ResponseInnerButton {
	if o == nil || isNil(o.Button) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerButton
		return ret
	}
	return *o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetButtonOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerButton, bool) {
	if o == nil || isNil(o.Button) {
    return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasButton() bool {
	if o != nil && !isNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerButton and assigns it to the Button field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetButton(v GetOrganizationSensorReadingsHistory200ResponseInnerButton) {
	o.Button = &v
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetDoor() GetOrganizationSensorReadingsHistory200ResponseInnerDoor {
	if o == nil || isNil(o.Door) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerDoor
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetDoorOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerDoor, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerDoor and assigns it to the Door field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetDoor(v GetOrganizationSensorReadingsHistory200ResponseInnerDoor) {
	o.Door = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetHumidity() GetOrganizationSensorReadingsHistory200ResponseInnerHumidity {
	if o == nil || isNil(o.Humidity) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerHumidity
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetHumidityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerHumidity, bool) {
	if o == nil || isNil(o.Humidity) {
    return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasHumidity() bool {
	if o != nil && !isNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerHumidity and assigns it to the Humidity field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetHumidity(v GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetIndoorAirQuality() GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality {
	if o == nil || isNil(o.IndoorAirQuality) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetIndoorAirQualityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality, bool) {
	if o == nil || isNil(o.IndoorAirQuality) {
    return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasIndoorAirQuality() bool {
	if o != nil && !isNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality and assigns it to the IndoorAirQuality field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetIndoorAirQuality(v GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetNoise() GetOrganizationSensorReadingsHistory200ResponseInnerNoise {
	if o == nil || isNil(o.Noise) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetNoiseOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNoise, bool) {
	if o == nil || isNil(o.Noise) {
    return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasNoise() bool {
	if o != nil && !isNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerNoise and assigns it to the Noise field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetNoise(v GetOrganizationSensorReadingsHistory200ResponseInnerNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetPm25() GetOrganizationSensorReadingsHistory200ResponseInnerPm25 {
	if o == nil || isNil(o.Pm25) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerPm25
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetPm25Ok() (*GetOrganizationSensorReadingsHistory200ResponseInnerPm25, bool) {
	if o == nil || isNil(o.Pm25) {
    return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasPm25() bool {
	if o != nil && !isNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerPm25 and assigns it to the Pm25 field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetPm25(v GetOrganizationSensorReadingsHistory200ResponseInnerPm25) {
	o.Pm25 = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTemperature() GetOrganizationSensorReadingsHistory200ResponseInnerTemperature {
	if o == nil || isNil(o.Temperature) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerTemperature
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTemperatureOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTemperature, bool) {
	if o == nil || isNil(o.Temperature) {
    return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasTemperature() bool {
	if o != nil && !isNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerTemperature and assigns it to the Temperature field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetTemperature(v GetOrganizationSensorReadingsHistory200ResponseInnerTemperature) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTvoc() GetOrganizationSensorReadingsHistory200ResponseInnerTvoc {
	if o == nil || isNil(o.Tvoc) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerTvoc
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTvocOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTvoc, bool) {
	if o == nil || isNil(o.Tvoc) {
    return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasTvoc() bool {
	if o != nil && !isNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerTvoc and assigns it to the Tvoc field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetTvoc(v GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) {
	o.Tvoc = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetWater() GetOrganizationSensorReadingsHistory200ResponseInnerWater {
	if o == nil || isNil(o.Water) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerWater
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetWaterOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerWater, bool) {
	if o == nil || isNil(o.Water) {
    return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasWater() bool {
	if o != nil && !isNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerWater and assigns it to the Water field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetWater(v GetOrganizationSensorReadingsHistory200ResponseInnerWater) {
	o.Water = &v
}

func (o GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !isNil(o.Battery) {
		toSerialize["battery"] = o.Battery
	}
	if !isNil(o.Button) {
		toSerialize["button"] = o.Button
	}
	if !isNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !isNil(o.Humidity) {
		toSerialize["humidity"] = o.Humidity
	}
	if !isNil(o.IndoorAirQuality) {
		toSerialize["indoorAirQuality"] = o.IndoorAirQuality
	}
	if !isNil(o.Noise) {
		toSerialize["noise"] = o.Noise
	}
	if !isNil(o.Pm25) {
		toSerialize["pm25"] = o.Pm25
	}
	if !isNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !isNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !isNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner struct {
	value *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) Get() *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) Set(val *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner(val *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) *NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner {
	return &NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


