# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModule**](ModuleResourceApi.md#CreateModule) | **Post** /api/kpi/modules | 
[**DeleteModule**](ModuleResourceApi.md#DeleteModule) | **Delete** /api/kpi/modules/{id} | 
[**GetAllModules**](ModuleResourceApi.md#GetAllModules) | **Get** /api/kpi/modules | 
[**GetAllModulesByTag**](ModuleResourceApi.md#GetAllModulesByTag) | **Get** /api/kpi/modules/tag/{tag} | 
[**GetAllModulesWithFlatInformation**](ModuleResourceApi.md#GetAllModulesWithFlatInformation) | **Get** /api/kpi/modules/all | 
[**GetModule**](ModuleResourceApi.md#GetModule) | **Get** /api/kpi/modules/{id} | 
[**GetModuleByName**](ModuleResourceApi.md#GetModuleByName) | **Get** /api/kpi/modules/by-name/{moduleName} | 
[**MergeModules**](ModuleResourceApi.md#MergeModules) | **Post** /api/kpi/modules/merge/{fromId}/{intoId} | 
[**ModuleTags**](ModuleResourceApi.md#ModuleTags) | **Get** /api/kpi/modules/tags | 
[**PartialUpdateModule**](ModuleResourceApi.md#PartialUpdateModule) | **Patch** /api/kpi/modules | 
[**UpdateModule**](ModuleResourceApi.md#UpdateModule) | **Put** /api/kpi/modules | 

# **CreateModule**
> ModuleDto CreateModule(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleDto**](ModuleDto.md)|  | 

### Return type

[**ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteModule**
> DeleteModule(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllModules**
> []ModuleDto GetAllModules(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModuleResourceApiGetAllModulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleResourceApiGetAllModulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllModulesByTag**
> []ModuleDto GetAllModulesByTag(ctx, tag)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**|  | 

### Return type

[**[]ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllModulesWithFlatInformation**
> []ModuleDto GetAllModulesWithFlatInformation(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModule**
> ModuleDto GetModule(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModuleByName**
> ModuleDto GetModuleByName(ctx, moduleName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**|  | 

### Return type

[**ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeModules**
> ModuleDto MergeModules(ctx, fromId, intoId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromId** | **string**|  | 
  **intoId** | **string**|  | 

### Return type

[**ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModuleTags**
> []string ModuleTags(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateModule**
> ModuleDto PartialUpdateModule(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleDto**](ModuleDto.md)|  | 

### Return type

[**ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateModule**
> ModuleDto UpdateModule(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleDto**](ModuleDto.md)|  | 

### Return type

[**ModuleDto**](ModuleDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

