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
)


// InboundFirewallRulesApiService InboundFirewallRulesApi service
type InboundFirewallRulesApiService service

type InboundFirewallRulesApiGetNetworkApplianceFirewallInboundFirewallRulesRequest struct {
	ctx context.Context
	ApiService *InboundFirewallRulesApiService
	networkId string
}

func (r InboundFirewallRulesApiGetNetworkApplianceFirewallInboundFirewallRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallInboundFirewallRulesExecute(r)
}

/*
GetNetworkApplianceFirewallInboundFirewallRules Return the inbound firewall rules for an MX network

Return the inbound firewall rules for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return InboundFirewallRulesApiGetNetworkApplianceFirewallInboundFirewallRulesRequest
*/
func (a *InboundFirewallRulesApiService) GetNetworkApplianceFirewallInboundFirewallRules(ctx context.Context, networkId string) InboundFirewallRulesApiGetNetworkApplianceFirewallInboundFirewallRulesRequest {
	return InboundFirewallRulesApiGetNetworkApplianceFirewallInboundFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InboundFirewallRulesApiService) GetNetworkApplianceFirewallInboundFirewallRulesExecute(r InboundFirewallRulesApiGetNetworkApplianceFirewallInboundFirewallRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InboundFirewallRulesApiService.GetNetworkApplianceFirewallInboundFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/inboundFirewallRules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct {
	ctx context.Context
	ApiService *InboundFirewallRulesApiService
	networkId string
	updateNetworkApplianceFirewallInboundFirewallRulesRequest *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest
}

func (r InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest {
	r.updateNetworkApplianceFirewallInboundFirewallRulesRequest = &updateNetworkApplianceFirewallInboundFirewallRulesRequest
	return r
}

func (r InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceFirewallInboundFirewallRulesExecute(r)
}

/*
UpdateNetworkApplianceFirewallInboundFirewallRules Update the inbound firewall rules of an MX network

Update the inbound firewall rules of an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest
*/
func (a *InboundFirewallRulesApiService) UpdateNetworkApplianceFirewallInboundFirewallRules(ctx context.Context, networkId string) InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest {
	return InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InboundFirewallRulesApiService) UpdateNetworkApplianceFirewallInboundFirewallRulesExecute(r InboundFirewallRulesApiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InboundFirewallRulesApiService.UpdateNetworkApplianceFirewallInboundFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/inboundFirewallRules"
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
	localVarPostBody = r.updateNetworkApplianceFirewallInboundFirewallRulesRequest
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
