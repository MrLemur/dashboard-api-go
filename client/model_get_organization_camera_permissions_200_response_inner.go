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

// checks if the GetOrganizationCameraPermissions200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationCameraPermissions200ResponseInner{}

// GetOrganizationCameraPermissions200ResponseInner struct for GetOrganizationCameraPermissions200ResponseInner
type GetOrganizationCameraPermissions200ResponseInner struct {
	// Permission scope id
	Id *string `json:"id,omitempty"`
	// Name of permission scope
	Name *string `json:"name,omitempty"`
	// Permission scope level
	Level *string `json:"level,omitempty"`
}

// NewGetOrganizationCameraPermissions200ResponseInner instantiates a new GetOrganizationCameraPermissions200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCameraPermissions200ResponseInner() *GetOrganizationCameraPermissions200ResponseInner {
	this := GetOrganizationCameraPermissions200ResponseInner{}
	return &this
}

// NewGetOrganizationCameraPermissions200ResponseInnerWithDefaults instantiates a new GetOrganizationCameraPermissions200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCameraPermissions200ResponseInnerWithDefaults() *GetOrganizationCameraPermissions200ResponseInner {
	this := GetOrganizationCameraPermissions200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationCameraPermissions200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraPermissions200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationCameraPermissions200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationCameraPermissions200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationCameraPermissions200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraPermissions200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationCameraPermissions200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationCameraPermissions200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetOrganizationCameraPermissions200ResponseInner) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraPermissions200ResponseInner) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetOrganizationCameraPermissions200ResponseInner) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *GetOrganizationCameraPermissions200ResponseInner) SetLevel(v string) {
	o.Level = &v
}

func (o GetOrganizationCameraPermissions200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationCameraPermissions200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableGetOrganizationCameraPermissions200ResponseInner struct {
	value *GetOrganizationCameraPermissions200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationCameraPermissions200ResponseInner) Get() *GetOrganizationCameraPermissions200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationCameraPermissions200ResponseInner) Set(val *GetOrganizationCameraPermissions200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCameraPermissions200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCameraPermissions200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCameraPermissions200ResponseInner(val *GetOrganizationCameraPermissions200ResponseInner) *NullableGetOrganizationCameraPermissions200ResponseInner {
	return &NullableGetOrganizationCameraPermissions200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationCameraPermissions200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCameraPermissions200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


