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

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser User credentials used to connect to the device
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser struct {
	// The public key for the registered user
	PublicKey *string `json:"publicKey,omitempty"`
	// The username added to Catalyst device
	Username *string `json:"username,omitempty"`
	Secret *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserSecret `json:"secret,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
    return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) SetUsername(v string) {
	o.Username = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) GetSecret() CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserSecret {
	if o == nil || isNil(o.Secret) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserSecret
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) GetSecretOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserSecret, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserSecret and assigns it to the Secret field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) SetSecret(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUserSecret) {
	o.Secret = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


