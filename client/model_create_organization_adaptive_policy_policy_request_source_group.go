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

// checks if the CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup{}

// CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup The source adaptive policy group (requires one unique attribute)
type CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup struct {
	// The ID of the source adaptive policy group
	Id *string `json:"id,omitempty"`
	// The name of the source adaptive policy group
	Name *string `json:"name,omitempty"`
	// The SGT of the source adaptive policy group
	Sgt *int32 `json:"sgt,omitempty"`
}

// NewCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup instantiates a new CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup() *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup {
	this := CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup{}
	return &this
}

// NewCreateOrganizationAdaptivePolicyPolicyRequestSourceGroupWithDefaults instantiates a new CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdaptivePolicyPolicyRequestSourceGroupWithDefaults() *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup {
	this := CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) GetSgt() int32 {
	if o == nil || IsNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) GetSgtOk() (*int32, bool) {
	if o == nil || IsNil(o.Sgt) {
		return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) HasSgt() bool {
	if o != nil && !IsNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) SetSgt(v int32) {
	o.Sgt = &v
}

func (o CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Sgt) {
		toSerialize["sgt"] = o.Sgt
	}
	return toSerialize, nil
}

type NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup struct {
	value *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup
	isSet bool
}

func (v NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) Get() *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup {
	return v.value
}

func (v *NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) Set(val *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup(val *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) *NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup {
	return &NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


