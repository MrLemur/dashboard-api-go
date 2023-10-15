/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationAdmins200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdmins200ResponseInner{}

// GetOrganizationAdmins200ResponseInner struct for GetOrganizationAdmins200ResponseInner
type GetOrganizationAdmins200ResponseInner struct {
	// Admin's ID
	Id *string `json:"id,omitempty"`
	// Admin's username
	Name *string `json:"name,omitempty"`
	// Admin's email address
	Email *string `json:"email,omitempty"`
	// Admin's level of access to the organization
	OrgAccess *string `json:"orgAccess,omitempty"`
	// Status of the admin's account
	AccountStatus *string `json:"accountStatus,omitempty"`
	// Indicates whether two-factor authentication is enabled
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
	// Indicates whether the admin has an API key
	HasApiKey *bool `json:"hasApiKey,omitempty"`
	// Time when the admin was last active
	LastActive *time.Time `json:"lastActive,omitempty"`
	// Admin tag information
	Tags []GetOrganizationAdmins200ResponseInnerTagsInner `json:"tags,omitempty"`
	// Admin network access information
	Networks []GetOrganizationAdmins200ResponseInnerNetworksInner `json:"networks,omitempty"`
	// Admin's authentication method
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
}

// NewGetOrganizationAdmins200ResponseInner instantiates a new GetOrganizationAdmins200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdmins200ResponseInner() *GetOrganizationAdmins200ResponseInner {
	this := GetOrganizationAdmins200ResponseInner{}
	return &this
}

// NewGetOrganizationAdmins200ResponseInnerWithDefaults instantiates a new GetOrganizationAdmins200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdmins200ResponseInnerWithDefaults() *GetOrganizationAdmins200ResponseInner {
	this := GetOrganizationAdmins200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationAdmins200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationAdmins200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetOrganizationAdmins200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetOrgAccess() string {
	if o == nil || IsNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetOrgAccessOk() (*string, bool) {
	if o == nil || IsNil(o.OrgAccess) {
		return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasOrgAccess() bool {
	if o != nil && !IsNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *GetOrganizationAdmins200ResponseInner) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetAccountStatus() string {
	if o == nil || IsNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetAccountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasAccountStatus() bool {
	if o != nil && !IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *GetOrganizationAdmins200ResponseInner) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetTwoFactorAuthEnabled() bool {
	if o == nil || IsNil(o.TwoFactorAuthEnabled) {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthEnabled
}

// GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetTwoFactorAuthEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TwoFactorAuthEnabled) {
		return nil, false
	}
	return o.TwoFactorAuthEnabled, true
}

// HasTwoFactorAuthEnabled returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasTwoFactorAuthEnabled() bool {
	if o != nil && !IsNil(o.TwoFactorAuthEnabled) {
		return true
	}

	return false
}

// SetTwoFactorAuthEnabled gets a reference to the given bool and assigns it to the TwoFactorAuthEnabled field.
func (o *GetOrganizationAdmins200ResponseInner) SetTwoFactorAuthEnabled(v bool) {
	o.TwoFactorAuthEnabled = &v
}

// GetHasApiKey returns the HasApiKey field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetHasApiKey() bool {
	if o == nil || IsNil(o.HasApiKey) {
		var ret bool
		return ret
	}
	return *o.HasApiKey
}

// GetHasApiKeyOk returns a tuple with the HasApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetHasApiKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.HasApiKey) {
		return nil, false
	}
	return o.HasApiKey, true
}

// HasHasApiKey returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasHasApiKey() bool {
	if o != nil && !IsNil(o.HasApiKey) {
		return true
	}

	return false
}

// SetHasApiKey gets a reference to the given bool and assigns it to the HasApiKey field.
func (o *GetOrganizationAdmins200ResponseInner) SetHasApiKey(v bool) {
	o.HasApiKey = &v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetLastActive() time.Time {
	if o == nil || IsNil(o.LastActive) {
		var ret time.Time
		return ret
	}
	return *o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetLastActiveOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastActive) {
		return nil, false
	}
	return o.LastActive, true
}

// HasLastActive returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasLastActive() bool {
	if o != nil && !IsNil(o.LastActive) {
		return true
	}

	return false
}

// SetLastActive gets a reference to the given time.Time and assigns it to the LastActive field.
func (o *GetOrganizationAdmins200ResponseInner) SetLastActive(v time.Time) {
	o.LastActive = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetTags() []GetOrganizationAdmins200ResponseInnerTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []GetOrganizationAdmins200ResponseInnerTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetTagsOk() ([]GetOrganizationAdmins200ResponseInnerTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []GetOrganizationAdmins200ResponseInnerTagsInner and assigns it to the Tags field.
func (o *GetOrganizationAdmins200ResponseInner) SetTags(v []GetOrganizationAdmins200ResponseInnerTagsInner) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetNetworks() []GetOrganizationAdmins200ResponseInnerNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []GetOrganizationAdmins200ResponseInnerNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetNetworksOk() ([]GetOrganizationAdmins200ResponseInnerNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []GetOrganizationAdmins200ResponseInnerNetworksInner and assigns it to the Networks field.
func (o *GetOrganizationAdmins200ResponseInner) SetNetworks(v []GetOrganizationAdmins200ResponseInnerNetworksInner) {
	o.Networks = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInner) GetAuthenticationMethod() string {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInner) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInner) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *GetOrganizationAdmins200ResponseInner) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o GetOrganizationAdmins200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdmins200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.OrgAccess) {
		toSerialize["orgAccess"] = o.OrgAccess
	}
	if !IsNil(o.AccountStatus) {
		toSerialize["accountStatus"] = o.AccountStatus
	}
	if !IsNil(o.TwoFactorAuthEnabled) {
		toSerialize["twoFactorAuthEnabled"] = o.TwoFactorAuthEnabled
	}
	if !IsNil(o.HasApiKey) {
		toSerialize["hasApiKey"] = o.HasApiKey
	}
	if !IsNil(o.LastActive) {
		toSerialize["lastActive"] = o.LastActive
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdmins200ResponseInner struct {
	value *GetOrganizationAdmins200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationAdmins200ResponseInner) Get() *GetOrganizationAdmins200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationAdmins200ResponseInner) Set(val *GetOrganizationAdmins200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdmins200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdmins200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdmins200ResponseInner(val *GetOrganizationAdmins200ResponseInner) *NullableGetOrganizationAdmins200ResponseInner {
	return &NullableGetOrganizationAdmins200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdmins200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdmins200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


