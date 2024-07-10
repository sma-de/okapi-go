# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportModuleVersion**](ImportResourceApi.md#ImportModuleVersion) | **Post** /api/kpi/import/module-version | 
[**ImportProductRelease**](ImportResourceApi.md#ImportProductRelease) | **Post** /api/kpi/import/product-release | 

# **ImportModuleVersion**
> []ModuleVersionDto ImportModuleVersion(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportResourceApiImportModuleVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportResourceApiImportModuleVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ModuleVersionImportDto**](ModuleVersionImportDto.md)|  | 

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportProductRelease**
> ProductReleaseDto ImportProductRelease(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductReleaseImportDto**](ProductReleaseImportDto.md)|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

