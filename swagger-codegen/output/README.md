# Go API client for swagger

Servers management

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ServersApi* | [**CreateServer**](docs/ServersApi.md#createserver) | **Post** /servers/ | Crete a server with its name
*ServersApi* | [**DeleteServer**](docs/ServersApi.md#deleteserver) | **Delete** /servers/{id} | Delete a server given its identifier
*ServersApi* | [**GetServer**](docs/ServersApi.md#getserver) | **Get** /servers/{id} | Fetch a server given its identifier
*ServersApi* | [**ListServers**](docs/ServersApi.md#listservers) | **Get** /servers/ | List all servers
*ServersApi* | [**UpdateServer**](docs/ServersApi.md#updateserver) | **Put** /servers/{id} | Update a server given its identifier


## Documentation For Models

 - [ServerRequest](docs/ServerRequest.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



