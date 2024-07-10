# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](LoginResourceApi.md#Login) | **Post** /api/login | 

# **Login**
> string Login(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginResourceApiLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoginResourceApiLoginOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LoginRequestDto**](LoginRequestDto.md)|  | 

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

