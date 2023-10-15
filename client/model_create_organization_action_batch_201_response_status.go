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

// checks if the CreateOrganizationActionBatch201ResponseStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationActionBatch201ResponseStatus{}

// CreateOrganizationActionBatch201ResponseStatus Status of action batch
type CreateOrganizationActionBatch201ResponseStatus struct {
	// Flag describing whether all actions in the action batch have completed
	Completed *bool `json:"completed,omitempty"`
	// Flag describing whether any actions in the action batch failed
	Failed *bool `json:"failed,omitempty"`
	// List of errors encountered when running actions in the action batch
	Errors []string `json:"errors,omitempty"`
	// Resources created as a result of this action batch
	CreatedResources []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner `json:"createdResources"`
}

// NewCreateOrganizationActionBatch201ResponseStatus instantiates a new CreateOrganizationActionBatch201ResponseStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationActionBatch201ResponseStatus(createdResources []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) *CreateOrganizationActionBatch201ResponseStatus {
	this := CreateOrganizationActionBatch201ResponseStatus{}
	this.CreatedResources = createdResources
	return &this
}

// NewCreateOrganizationActionBatch201ResponseStatusWithDefaults instantiates a new CreateOrganizationActionBatch201ResponseStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationActionBatch201ResponseStatusWithDefaults() *CreateOrganizationActionBatch201ResponseStatus {
	this := CreateOrganizationActionBatch201ResponseStatus{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *CreateOrganizationActionBatch201ResponseStatus) SetCompleted(v bool) {
	o.Completed = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *CreateOrganizationActionBatch201ResponseStatus) SetFailed(v bool) {
	o.Failed = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *CreateOrganizationActionBatch201ResponseStatus) SetErrors(v []string) {
	o.Errors = v
}

// GetCreatedResources returns the CreatedResources field value
func (o *CreateOrganizationActionBatch201ResponseStatus) GetCreatedResources() []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner {
	if o == nil {
		var ret []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner
		return ret
	}

	return o.CreatedResources
}

// GetCreatedResourcesOk returns a tuple with the CreatedResources field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseStatus) GetCreatedResourcesOk() ([]CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedResources, true
}

// SetCreatedResources sets field value
func (o *CreateOrganizationActionBatch201ResponseStatus) SetCreatedResources(v []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) {
	o.CreatedResources = v
}

func (o CreateOrganizationActionBatch201ResponseStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationActionBatch201ResponseStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["createdResources"] = o.CreatedResources
	return toSerialize, nil
}

type NullableCreateOrganizationActionBatch201ResponseStatus struct {
	value *CreateOrganizationActionBatch201ResponseStatus
	isSet bool
}

func (v NullableCreateOrganizationActionBatch201ResponseStatus) Get() *CreateOrganizationActionBatch201ResponseStatus {
	return v.value
}

func (v *NullableCreateOrganizationActionBatch201ResponseStatus) Set(val *CreateOrganizationActionBatch201ResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationActionBatch201ResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationActionBatch201ResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationActionBatch201ResponseStatus(val *CreateOrganizationActionBatch201ResponseStatus) *NullableCreateOrganizationActionBatch201ResponseStatus {
	return &NullableCreateOrganizationActionBatch201ResponseStatus{value: val, isSet: true}
}

func (v NullableCreateOrganizationActionBatch201ResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationActionBatch201ResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


