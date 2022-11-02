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

// CreateOrganizationEarlyAccessFeaturesOptInRequest struct for CreateOrganizationEarlyAccessFeaturesOptInRequest
type CreateOrganizationEarlyAccessFeaturesOptInRequest struct {
	// Short name of the early access feature
	ShortName string `json:"shortName"`
	// A list of network IDs to apply the opt-in to
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"`
}

// NewCreateOrganizationEarlyAccessFeaturesOptInRequest instantiates a new CreateOrganizationEarlyAccessFeaturesOptInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationEarlyAccessFeaturesOptInRequest(shortName string) *CreateOrganizationEarlyAccessFeaturesOptInRequest {
	this := CreateOrganizationEarlyAccessFeaturesOptInRequest{}
	this.ShortName = shortName
	return &this
}

// NewCreateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults instantiates a new CreateOrganizationEarlyAccessFeaturesOptInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults() *CreateOrganizationEarlyAccessFeaturesOptInRequest {
	this := CreateOrganizationEarlyAccessFeaturesOptInRequest{}
	return &this
}

// GetShortName returns the ShortName field value
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetShortNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) SetShortName(v string) {
	o.ShortName = v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetLimitScopeToNetworks() []string {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []string
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetLimitScopeToNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []string and assigns it to the LimitScopeToNetworks field.
func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) SetLimitScopeToNetworks(v []string) {
	o.LimitScopeToNetworks = v
}

func (o CreateOrganizationEarlyAccessFeaturesOptInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.LimitScopeToNetworks) {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationEarlyAccessFeaturesOptInRequest struct {
	value *CreateOrganizationEarlyAccessFeaturesOptInRequest
	isSet bool
}

func (v NullableCreateOrganizationEarlyAccessFeaturesOptInRequest) Get() *CreateOrganizationEarlyAccessFeaturesOptInRequest {
	return v.value
}

func (v *NullableCreateOrganizationEarlyAccessFeaturesOptInRequest) Set(val *CreateOrganizationEarlyAccessFeaturesOptInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationEarlyAccessFeaturesOptInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationEarlyAccessFeaturesOptInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationEarlyAccessFeaturesOptInRequest(val *CreateOrganizationEarlyAccessFeaturesOptInRequest) *NullableCreateOrganizationEarlyAccessFeaturesOptInRequest {
	return &NullableCreateOrganizationEarlyAccessFeaturesOptInRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationEarlyAccessFeaturesOptInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationEarlyAccessFeaturesOptInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


