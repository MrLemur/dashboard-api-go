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

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	// Id to check the status of your ping request.
	PingId *string `json:"pingId,omitempty"`
	// GET this url to check the status of your ping request.
	Url *string `json:"url,omitempty"`
	Request *InlineResponse2011Request `json:"request,omitempty"`
	// Status of the ping request.
	Status *string `json:"status,omitempty"`
	Results *InlineResponse2003Results `json:"results,omitempty"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetPingId returns the PingId field value if set, zero value otherwise.
func (o *InlineResponse2003) GetPingId() string {
	if o == nil || isNil(o.PingId) {
		var ret string
		return ret
	}
	return *o.PingId
}

// GetPingIdOk returns a tuple with the PingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetPingIdOk() (*string, bool) {
	if o == nil || isNil(o.PingId) {
    return nil, false
	}
	return o.PingId, true
}

// HasPingId returns a boolean if a field has been set.
func (o *InlineResponse2003) HasPingId() bool {
	if o != nil && !isNil(o.PingId) {
		return true
	}

	return false
}

// SetPingId gets a reference to the given string and assigns it to the PingId field.
func (o *InlineResponse2003) SetPingId(v string) {
	o.PingId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2003) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2003) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2003) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2003) GetRequest() InlineResponse2011Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2011Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetRequestOk() (*InlineResponse2011Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2003) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2011Request and assigns it to the Request field.
func (o *InlineResponse2003) SetRequest(v InlineResponse2011Request) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2003) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2003) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2003) SetStatus(v string) {
	o.Status = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse2003) GetResults() InlineResponse2003Results {
	if o == nil || isNil(o.Results) {
		var ret InlineResponse2003Results
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetResultsOk() (*InlineResponse2003Results, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse2003) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given InlineResponse2003Results and assigns it to the Results field.
func (o *InlineResponse2003) SetResults(v InlineResponse2003Results) {
	o.Results = &v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PingId) {
		toSerialize["pingId"] = o.PingId
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


