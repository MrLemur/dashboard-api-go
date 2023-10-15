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

// checks if the CreateOrganizationBrandingPolicyRequestCustomLogoImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationBrandingPolicyRequestCustomLogoImage{}

// CreateOrganizationBrandingPolicyRequestCustomLogoImage Properties for setting the image.
type CreateOrganizationBrandingPolicyRequestCustomLogoImage struct {
	// The file contents (a base 64 encoded string) of your new logo.
	Contents *string `json:"contents,omitempty"`
	// The format of the encoded contents.  Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
}

// NewCreateOrganizationBrandingPolicyRequestCustomLogoImage instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogoImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationBrandingPolicyRequestCustomLogoImage() *CreateOrganizationBrandingPolicyRequestCustomLogoImage {
	this := CreateOrganizationBrandingPolicyRequestCustomLogoImage{}
	return &this
}

// NewCreateOrganizationBrandingPolicyRequestCustomLogoImageWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogoImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationBrandingPolicyRequestCustomLogoImageWithDefaults() *CreateOrganizationBrandingPolicyRequestCustomLogoImage {
	this := CreateOrganizationBrandingPolicyRequestCustomLogoImage{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) SetContents(v string) {
	o.Contents = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogoImage) SetFormat(v string) {
	o.Format = &v
}

func (o CreateOrganizationBrandingPolicyRequestCustomLogoImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationBrandingPolicyRequestCustomLogoImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	return toSerialize, nil
}

type NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage struct {
	value *CreateOrganizationBrandingPolicyRequestCustomLogoImage
	isSet bool
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage) Get() *CreateOrganizationBrandingPolicyRequestCustomLogoImage {
	return v.value
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage) Set(val *CreateOrganizationBrandingPolicyRequestCustomLogoImage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationBrandingPolicyRequestCustomLogoImage(val *CreateOrganizationBrandingPolicyRequestCustomLogoImage) *NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage {
	return &NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage{value: val, isSet: true}
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogoImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


