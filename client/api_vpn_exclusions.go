/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// VpnExclusionsApiService VpnExclusionsApi service
type VpnExclusionsApiService service

type VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest struct {
	ctx context.Context
	ApiService *VpnExclusionsApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest) PerPage(perPage int32) VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest) StartingAfter(startingAfter string) VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest) EndingBefore(endingBefore string) VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter the results by network IDs
func (r VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest) NetworkIds(networkIds []string) VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

func (r VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest) Execute() ([]UpdateNetworkApplianceTrafficShapingVpnExclusions200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkExecute(r)
}

/*
GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork Display VPN exclusion rules for MX networks.

Display VPN exclusion rules for MX networks.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest
*/
func (a *VpnExclusionsApiService) GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(ctx context.Context, organizationId string) VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest {
	return VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
func (a *VpnExclusionsApiService) GetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkExecute(r VpnExclusionsApiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest) ([]UpdateNetworkApplianceTrafficShapingVpnExclusions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VpnExclusionsApiService.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/trafficShaping/vpnExclusions/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct {
	ctx context.Context
	ApiService *VpnExclusionsApiService
	networkId string
	updateNetworkApplianceTrafficShapingVpnExclusionsRequest *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest
}

func (r VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest(updateNetworkApplianceTrafficShapingVpnExclusionsRequest UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest {
	r.updateNetworkApplianceTrafficShapingVpnExclusionsRequest = &updateNetworkApplianceTrafficShapingVpnExclusionsRequest
	return r
}

func (r VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) Execute() (*UpdateNetworkApplianceTrafficShapingVpnExclusions200Response, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceTrafficShapingVpnExclusionsExecute(r)
}

/*
UpdateNetworkApplianceTrafficShapingVpnExclusions Update VPN exclusion rules for an MX network.

Update VPN exclusion rules for an MX network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest
*/
func (a *VpnExclusionsApiService) UpdateNetworkApplianceTrafficShapingVpnExclusions(ctx context.Context, networkId string) VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest {
	return VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
func (a *VpnExclusionsApiService) UpdateNetworkApplianceTrafficShapingVpnExclusionsExecute(r VpnExclusionsApiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) (*UpdateNetworkApplianceTrafficShapingVpnExclusions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VpnExclusionsApiService.UpdateNetworkApplianceTrafficShapingVpnExclusions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/trafficShaping/vpnExclusions"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateNetworkApplianceTrafficShapingVpnExclusionsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
