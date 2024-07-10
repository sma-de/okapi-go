# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModuleVersion**](ModuleVersionResourceApi.md#CreateModuleVersion) | **Post** /api/kpi/module-versions | 
[**DeleteModuleVersion**](ModuleVersionResourceApi.md#DeleteModuleVersion) | **Delete** /api/kpi/module-versions/{id} | 
[**GetAllModuleVersions**](ModuleVersionResourceApi.md#GetAllModuleVersions) | **Get** /api/kpi/module-versions | 
[**GetAllModuleVersionsForModule**](ModuleVersionResourceApi.md#GetAllModuleVersionsForModule) | **Get** /api/kpi/module-versions/by-module/{moduleId} | 
[**GetAllModuleVersionsForModuleName**](ModuleVersionResourceApi.md#GetAllModuleVersionsForModuleName) | **Get** /api/kpi/module-versions/by-name/{moduleName} | 
[**GetAllModuleVersionsForProductRelease**](ModuleVersionResourceApi.md#GetAllModuleVersionsForProductRelease) | **Get** /api/kpi/module-versions/by-product-release/{productReleaseId} | 
[**GetLatestModuleVersionsForModule**](ModuleVersionResourceApi.md#GetLatestModuleVersionsForModule) | **Get** /api/kpi/module-versions/latest | 
[**GetLatestModuleVersionsForModule1**](ModuleVersionResourceApi.md#GetLatestModuleVersionsForModule1) | **Get** /api/kpi/module-versions/latest/by-module/{moduleId} | 
[**GetModuleVersionAndCalculatedKpi**](ModuleVersionResourceApi.md#GetModuleVersionAndCalculatedKpi) | **Get** /api/kpi/module-versions/{id} | 
[**GetModuleVersionForModuleNameWithRevision**](ModuleVersionResourceApi.md#GetModuleVersionForModuleNameWithRevision) | **Get** /api/kpi/module-versions/by-module/{moduleName}/{revision} | 
[**GetPreviousModuleVersion1**](ModuleVersionResourceApi.md#GetPreviousModuleVersion1) | **Get** /api/kpi/module-versions/previous/{id} | 
[**PartialUpdateModuleVersion**](ModuleVersionResourceApi.md#PartialUpdateModuleVersion) | **Patch** /api/kpi/module-versions | 
[**UpdateModuleVersion**](ModuleVersionResourceApi.md#UpdateModuleVersion) | **Put** /api/kpi/module-versions | 

# **CreateModuleVersion**
> ModuleVersionDto CreateModuleVersion(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleVersionDto**](ModuleVersionDto.md)|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteModuleVersion**
> DeleteModuleVersion(ctx, id)


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

# **GetAllModuleVersions**
> []ModuleVersionDto GetAllModuleVersions(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModuleVersionResourceApiGetAllModuleVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleVersionResourceApiGetAllModuleVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllModuleVersionsForModule**
> []ModuleVersionDto GetAllModuleVersionsForModule(ctx, moduleId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleId** | **string**|  | 
 **optional** | ***ModuleVersionResourceApiGetAllModuleVersionsForModuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleVersionResourceApiGetAllModuleVersionsForModuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllModuleVersionsForModuleName**
> []ModuleVersionDto GetAllModuleVersionsForModuleName(ctx, moduleName, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**|  | 
 **optional** | ***ModuleVersionResourceApiGetAllModuleVersionsForModuleNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleVersionResourceApiGetAllModuleVersionsForModuleNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllModuleVersionsForProductRelease**
> []ModuleVersionDto GetAllModuleVersionsForProductRelease(ctx, productReleaseId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productReleaseId** | **string**|  | 
 **optional** | ***ModuleVersionResourceApiGetAllModuleVersionsForProductReleaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleVersionResourceApiGetAllModuleVersionsForProductReleaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestModuleVersionsForModule**
> []ModuleVersionDto GetLatestModuleVersionsForModule(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ModuleVersionResourceApiGetLatestModuleVersionsForModuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleVersionResourceApiGetLatestModuleVersionsForModuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestModuleVersionsForModule1**
> ModuleVersionDto GetLatestModuleVersionsForModule1(ctx, moduleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleId** | **string**|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModuleVersionAndCalculatedKpi**
> ModuleVersionDto GetModuleVersionAndCalculatedKpi(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModuleVersionForModuleNameWithRevision**
> ModuleVersionDto GetModuleVersionForModuleNameWithRevision(ctx, moduleName, revision, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**|  | 
  **revision** | **string**|  | 
 **optional** | ***ModuleVersionResourceApiGetModuleVersionForModuleNameWithRevisionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ModuleVersionResourceApiGetModuleVersionForModuleNameWithRevisionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreviousModuleVersion1**
> ModuleVersionDto GetPreviousModuleVersion1(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateModuleVersion**
> ModuleVersionDto PartialUpdateModuleVersion(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleVersionDto**](ModuleVersionDto.md)|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateModuleVersion**
> ModuleVersionDto UpdateModuleVersion(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleVersionDto**](ModuleVersionDto.md)|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

