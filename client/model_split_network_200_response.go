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

// SplitNetwork200Response struct for SplitNetwork200Response
type SplitNetwork200Response struct {
	// Networks after the split
	ResultingNetworks []GetNetwork200Response `json:"resultingNetworks,omitempty"`
}

// NewSplitNetwork200Response instantiates a new SplitNetwork200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitNetwork200Response() *SplitNetwork200Response {
	this := SplitNetwork200Response{}
	return &this
}

// NewSplitNetwork200ResponseWithDefaults instantiates a new SplitNetwork200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitNetwork200ResponseWithDefaults() *SplitNetwork200Response {
	this := SplitNetwork200Response{}
	return &this
}

// GetResultingNetworks returns the ResultingNetworks field value if set, zero value otherwise.
func (o *SplitNetwork200Response) GetResultingNetworks() []GetNetwork200Response {
	if o == nil || isNil(o.ResultingNetworks) {
		var ret []GetNetwork200Response
		return ret
	}
	return o.ResultingNetworks
}

// GetResultingNetworksOk returns a tuple with the ResultingNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitNetwork200Response) GetResultingNetworksOk() ([]GetNetwork200Response, bool) {
	if o == nil || isNil(o.ResultingNetworks) {
    return nil, false
	}
	return o.ResultingNetworks, true
}

// HasResultingNetworks returns a boolean if a field has been set.
func (o *SplitNetwork200Response) HasResultingNetworks() bool {
	if o != nil && !isNil(o.ResultingNetworks) {
		return true
	}

	return false
}

// SetResultingNetworks gets a reference to the given []GetNetwork200Response and assigns it to the ResultingNetworks field.
func (o *SplitNetwork200Response) SetResultingNetworks(v []GetNetwork200Response) {
	o.ResultingNetworks = v
}

func (o SplitNetwork200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetworks) {
		toSerialize["resultingNetworks"] = o.ResultingNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableSplitNetwork200Response struct {
	value *SplitNetwork200Response
	isSet bool
}

func (v NullableSplitNetwork200Response) Get() *SplitNetwork200Response {
	return v.value
}

func (v *NullableSplitNetwork200Response) Set(val *SplitNetwork200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitNetwork200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitNetwork200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitNetwork200Response(val *SplitNetwork200Response) *NullableSplitNetwork200Response {
	return &NullableSplitNetwork200Response{value: val, isSet: true}
}

func (v NullableSplitNetwork200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitNetwork200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


