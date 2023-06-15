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

// InlineResponse20049 struct for InlineResponse20049
type InlineResponse20049 struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
	// The network to which the devices was moved.
	NewNetwork *string `json:"newNetwork,omitempty"`
}

// NewInlineResponse20049 instantiates a new InlineResponse20049 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049WithDefaults() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineResponse20049) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineResponse20049) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineResponse20049) SetIds(v []string) {
	o.Ids = v
}

// GetNewNetwork returns the NewNetwork field value if set, zero value otherwise.
func (o *InlineResponse20049) GetNewNetwork() string {
	if o == nil || isNil(o.NewNetwork) {
		var ret string
		return ret
	}
	return *o.NewNetwork
}

// GetNewNetworkOk returns a tuple with the NewNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetNewNetworkOk() (*string, bool) {
	if o == nil || isNil(o.NewNetwork) {
    return nil, false
	}
	return o.NewNetwork, true
}

// HasNewNetwork returns a boolean if a field has been set.
func (o *InlineResponse20049) HasNewNetwork() bool {
	if o != nil && !isNil(o.NewNetwork) {
		return true
	}

	return false
}

// SetNewNetwork gets a reference to the given string and assigns it to the NewNetwork field.
func (o *InlineResponse20049) SetNewNetwork(v string) {
	o.NewNetwork = &v
}

func (o InlineResponse20049) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.NewNetwork) {
		toSerialize["newNetwork"] = o.NewNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20049 struct {
	value *InlineResponse20049
	isSet bool
}

func (v NullableInlineResponse20049) Get() *InlineResponse20049 {
	return v.value
}

func (v *NullableInlineResponse20049) Set(val *InlineResponse20049) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049(val *InlineResponse20049) *NullableInlineResponse20049 {
	return &NullableInlineResponse20049{value: val, isSet: true}
}

func (v NullableInlineResponse20049) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


