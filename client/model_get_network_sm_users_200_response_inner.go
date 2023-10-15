/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkSmUsers200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmUsers200ResponseInner{}

// GetNetworkSmUsers200ResponseInner struct for GetNetworkSmUsers200ResponseInner
type GetNetworkSmUsers200ResponseInner struct {
	// The Meraki managed Id of the user record.
	Id *string `json:"id,omitempty"`
	// User email.
	Email *string `json:"email,omitempty"`
	// User full name.
	FullName *string `json:"fullName,omitempty"`
	// The users username.
	Username *string `json:"username,omitempty"`
	// A boolean denoting if the user has a password associated with the record.
	HasPassword *bool `json:"hasPassword,omitempty"`
	// The set of tags the user is scoped to.
	Tags *string `json:"tags,omitempty"`
	// Active Directory Groups the user belongs to.
	AdGroups []string `json:"adGroups,omitempty"`
	// Azure Active Directory Groups the user belongs to.
	AzureAdGroups []string `json:"azureAdGroups,omitempty"`
	// SAML Groups the user belongs to.
	SamlGroups []string `json:"samlGroups,omitempty"`
	// Apple School Manager Groups the user belongs to.
	AsmGroups []string `json:"asmGroups,omitempty"`
	// Whether the user was created using an external integration, or via the Meraki Dashboard.
	IsExternal *bool `json:"isExternal,omitempty"`
	// The user display name.
	DisplayName *string `json:"displayName,omitempty"`
	// A boolean indicating if the user has an associated identity certificate..
	HasIdentityCertificate *bool `json:"hasIdentityCertificate,omitempty"`
	// The url where the users thumbnail is hosted.
	UserThumbnail *string `json:"userThumbnail,omitempty"`
}

// NewGetNetworkSmUsers200ResponseInner instantiates a new GetNetworkSmUsers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmUsers200ResponseInner() *GetNetworkSmUsers200ResponseInner {
	this := GetNetworkSmUsers200ResponseInner{}
	return &this
}

// NewGetNetworkSmUsers200ResponseInnerWithDefaults instantiates a new GetNetworkSmUsers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmUsers200ResponseInnerWithDefaults() *GetNetworkSmUsers200ResponseInner {
	this := GetNetworkSmUsers200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmUsers200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetNetworkSmUsers200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *GetNetworkSmUsers200ResponseInner) SetFullName(v string) {
	o.FullName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetNetworkSmUsers200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *GetNetworkSmUsers200ResponseInner) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *GetNetworkSmUsers200ResponseInner) SetTags(v string) {
	o.Tags = &v
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetAdGroups() []string {
	if o == nil || IsNil(o.AdGroups) {
		var ret []string
		return ret
	}
	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetAdGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given []string and assigns it to the AdGroups field.
func (o *GetNetworkSmUsers200ResponseInner) SetAdGroups(v []string) {
	o.AdGroups = v
}

// GetAzureAdGroups returns the AzureAdGroups field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetAzureAdGroups() []string {
	if o == nil || IsNil(o.AzureAdGroups) {
		var ret []string
		return ret
	}
	return o.AzureAdGroups
}

// GetAzureAdGroupsOk returns a tuple with the AzureAdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetAzureAdGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.AzureAdGroups) {
		return nil, false
	}
	return o.AzureAdGroups, true
}

// HasAzureAdGroups returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasAzureAdGroups() bool {
	if o != nil && !IsNil(o.AzureAdGroups) {
		return true
	}

	return false
}

// SetAzureAdGroups gets a reference to the given []string and assigns it to the AzureAdGroups field.
func (o *GetNetworkSmUsers200ResponseInner) SetAzureAdGroups(v []string) {
	o.AzureAdGroups = v
}

// GetSamlGroups returns the SamlGroups field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetSamlGroups() []string {
	if o == nil || IsNil(o.SamlGroups) {
		var ret []string
		return ret
	}
	return o.SamlGroups
}

// GetSamlGroupsOk returns a tuple with the SamlGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetSamlGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.SamlGroups) {
		return nil, false
	}
	return o.SamlGroups, true
}

// HasSamlGroups returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasSamlGroups() bool {
	if o != nil && !IsNil(o.SamlGroups) {
		return true
	}

	return false
}

// SetSamlGroups gets a reference to the given []string and assigns it to the SamlGroups field.
func (o *GetNetworkSmUsers200ResponseInner) SetSamlGroups(v []string) {
	o.SamlGroups = v
}

// GetAsmGroups returns the AsmGroups field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetAsmGroups() []string {
	if o == nil || IsNil(o.AsmGroups) {
		var ret []string
		return ret
	}
	return o.AsmGroups
}

// GetAsmGroupsOk returns a tuple with the AsmGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetAsmGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.AsmGroups) {
		return nil, false
	}
	return o.AsmGroups, true
}

// HasAsmGroups returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasAsmGroups() bool {
	if o != nil && !IsNil(o.AsmGroups) {
		return true
	}

	return false
}

// SetAsmGroups gets a reference to the given []string and assigns it to the AsmGroups field.
func (o *GetNetworkSmUsers200ResponseInner) SetAsmGroups(v []string) {
	o.AsmGroups = v
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetIsExternal() bool {
	if o == nil || IsNil(o.IsExternal) {
		var ret bool
		return ret
	}
	return *o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetIsExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExternal) {
		return nil, false
	}
	return o.IsExternal, true
}

// HasIsExternal returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasIsExternal() bool {
	if o != nil && !IsNil(o.IsExternal) {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given bool and assigns it to the IsExternal field.
func (o *GetNetworkSmUsers200ResponseInner) SetIsExternal(v bool) {
	o.IsExternal = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GetNetworkSmUsers200ResponseInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHasIdentityCertificate returns the HasIdentityCertificate field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetHasIdentityCertificate() bool {
	if o == nil || IsNil(o.HasIdentityCertificate) {
		var ret bool
		return ret
	}
	return *o.HasIdentityCertificate
}

// GetHasIdentityCertificateOk returns a tuple with the HasIdentityCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetHasIdentityCertificateOk() (*bool, bool) {
	if o == nil || IsNil(o.HasIdentityCertificate) {
		return nil, false
	}
	return o.HasIdentityCertificate, true
}

// HasHasIdentityCertificate returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasHasIdentityCertificate() bool {
	if o != nil && !IsNil(o.HasIdentityCertificate) {
		return true
	}

	return false
}

// SetHasIdentityCertificate gets a reference to the given bool and assigns it to the HasIdentityCertificate field.
func (o *GetNetworkSmUsers200ResponseInner) SetHasIdentityCertificate(v bool) {
	o.HasIdentityCertificate = &v
}

// GetUserThumbnail returns the UserThumbnail field value if set, zero value otherwise.
func (o *GetNetworkSmUsers200ResponseInner) GetUserThumbnail() string {
	if o == nil || IsNil(o.UserThumbnail) {
		var ret string
		return ret
	}
	return *o.UserThumbnail
}

// GetUserThumbnailOk returns a tuple with the UserThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUsers200ResponseInner) GetUserThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.UserThumbnail) {
		return nil, false
	}
	return o.UserThumbnail, true
}

// HasUserThumbnail returns a boolean if a field has been set.
func (o *GetNetworkSmUsers200ResponseInner) HasUserThumbnail() bool {
	if o != nil && !IsNil(o.UserThumbnail) {
		return true
	}

	return false
}

// SetUserThumbnail gets a reference to the given string and assigns it to the UserThumbnail field.
func (o *GetNetworkSmUsers200ResponseInner) SetUserThumbnail(v string) {
	o.UserThumbnail = &v
}

func (o GetNetworkSmUsers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmUsers200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	if !IsNil(o.AzureAdGroups) {
		toSerialize["azureAdGroups"] = o.AzureAdGroups
	}
	if !IsNil(o.SamlGroups) {
		toSerialize["samlGroups"] = o.SamlGroups
	}
	if !IsNil(o.AsmGroups) {
		toSerialize["asmGroups"] = o.AsmGroups
	}
	if !IsNil(o.IsExternal) {
		toSerialize["isExternal"] = o.IsExternal
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.HasIdentityCertificate) {
		toSerialize["hasIdentityCertificate"] = o.HasIdentityCertificate
	}
	if !IsNil(o.UserThumbnail) {
		toSerialize["userThumbnail"] = o.UserThumbnail
	}
	return toSerialize, nil
}

type NullableGetNetworkSmUsers200ResponseInner struct {
	value *GetNetworkSmUsers200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmUsers200ResponseInner) Get() *GetNetworkSmUsers200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmUsers200ResponseInner) Set(val *GetNetworkSmUsers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmUsers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmUsers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmUsers200ResponseInner(val *GetNetworkSmUsers200ResponseInner) *NullableGetNetworkSmUsers200ResponseInner {
	return &NullableGetNetworkSmUsers200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmUsers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmUsers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


