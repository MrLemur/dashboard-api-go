# dashboard-api-go

dashboard-api-go is a Go client library for [MERAKI Center Platform]().

## Usage

```go
import meraki "github.com/meraki/dashboard-api-go/sdk"
``` 

## Introduction

The dashboard-api-go makes it easier to work with the Dashboard Meraki Platform RESTFul APIs from Go.

It supports version 1.33.0

## Getting started

The first think you need to do is to generate an API client. There are two options to do it:

1. Parameters
2. Environment variables

### Parameters

The client could be generated with the following parameters:

- `baseURL`: The base URL, FQDN or IP, of the MERAKI instance.
- `dashboardApiKey`: The meraki_key for access to API.
- `debug`: Boolean to enable debugging
- `sslVerify`: Boolean to enable or disable SSL certificate verification.

```go
client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"MERAKI_KEY",
		"true", "false")
nResponse, _, err := client.Administered.GetAdministeredIDentitiesMe()
	if err != nil {
		fmt.Println(err)
		return
	}
```

### Using environment variables

The client can be configured with the following environment variables:

- `MERAKI_BASE_URL`: The base URL, FQDN or IP, of the MERAKI instance.
- `MERAKI_DASHBOARD_API_KEY`: The meraki_key for access to API.
- `MERAKI_DEBUG`: Boolean to enable debugging
- `MERAKI_SSL_VERIFY`: Boolean to enable or disable SSL certificate verification.

```go
Client, err = dnac.NewClient()
devicesCount, _, err := Client.Devices.GetDeviceCount()
```

## Examples

Here is an example of how we can generate a client, get a device count and then a list of devices filtering them using query params.

```go
client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"Meraki_key",
		"true", "false")
	if err != nil {
		fmt.Println(err)
		return
	}

	nResponse, _, err := client.Organizations.GetOrganizations()
	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println(nResponse)
		return
	}
	fmt.Println("There's no data on response")
```

## Documentation

[dashboard-api-go](https://pkg.go.dev/github.com/meraki/dashboard-api-go)

## Compatibility matrix

| SDK versions | MERAKI Dashboard version supported |
|--------------|------------------------------------|
| 2.y.z        |  1.33.0                            |


## Changelog

All notable changes to this project will be documented in the [CHANGELOG](https://github.com/meraki/dashboard-api-go/blob/main/CHANGELOG.md) file.

The development team may make additional name changes as the library evolves with the MERAKI Dashboard APIs.


## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
