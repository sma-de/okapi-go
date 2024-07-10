# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashPing**](UtilResourceApi.md#DashPing) | **Get** /api/kpi/dash/ping | 
[**DashSonar**](UtilResourceApi.md#DashSonar) | **Get** /api/kpi/dash/sonar | 
[**GetBadge**](UtilResourceApi.md#GetBadge) | **Get** /api/kpi/badge/{badge} | 
[**GetBadgeForModuleVersion**](UtilResourceApi.md#GetBadgeForModuleVersion) | **Get** /api/kpi/badge/module/{moduleName}/{moduleVersion} | 
[**GetBadgeForProductVersion**](UtilResourceApi.md#GetBadgeForProductVersion) | **Get** /api/kpi/badge/product/{productName}/{productVersion} | 
[**GetBadgeForProductVersion1**](UtilResourceApi.md#GetBadgeForProductVersion1) | **Get** /api/kpi/badge/product/by-susy/{susyId}/{productVersion} | 
[**GetIp**](UtilResourceApi.md#GetIp) | **Get** /api/kpi/ip | 
[**ImportDashDate**](UtilResourceApi.md#ImportDashDate) | **Get** /api/kpi/force-import | 
[**MapProductGroups**](UtilResourceApi.md#MapProductGroups) | **Get** /api/kpi/map-product-groups | 
[**ReloadModuleVersionMetrics**](UtilResourceApi.md#ReloadModuleVersionMetrics) | **Post** /api/kpi/reload/metrics/module-version/{moduleVersionId} | 
[**ReloadProductReloadMetrics**](UtilResourceApi.md#ReloadProductReloadMetrics) | **Post** /api/kpi/reload/metrics/product-release/{productReleaseId} | 
[**ReloadReleaseManagerStatus**](UtilResourceApi.md#ReloadReleaseManagerStatus) | **Post** /api/kpi/reload/release-manager-status | 
[**ResetModuleKpi**](UtilResourceApi.md#ResetModuleKpi) | **Post** /api/kpi/reset/kpi/module | 
[**ResetProductKpi**](UtilResourceApi.md#ResetProductKpi) | **Post** /api/kpi/reset/kpi/product | 

# **DashPing**
> PingDto DashPing(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingDto**](PingDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashSonar**
> SonarDto DashSonar(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SonarDto**](SonarDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBadge**
> *os.File GetBadge(ctx, badge)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **badge** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBadgeForModuleVersion**
> *os.File GetBadgeForModuleVersion(ctx, moduleName, moduleVersion)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**|  | 
  **moduleVersion** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBadgeForProductVersion**
> *os.File GetBadgeForProductVersion(ctx, productName, productVersion)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productName** | **string**|  | 
  **productVersion** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBadgeForProductVersion1**
> *os.File GetBadgeForProductVersion1(ctx, susyId, productVersion)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **susyId** | **int32**|  | 
  **productVersion** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIp**
> string GetIp(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportDashDate**
> ImportDashDate(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MapProductGroups**
> []ProductDto MapProductGroups(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadModuleVersionMetrics**
> ModuleVersionDto ReloadModuleVersionMetrics(ctx, moduleVersionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleVersionId** | **string**|  | 

### Return type

[**ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadProductReloadMetrics**
> ProductReleaseDto ReloadProductReloadMetrics(ctx, productReleaseId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productReleaseId** | **string**|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadReleaseManagerStatus**
> []ProductReleaseDto ReloadReleaseManagerStatus(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetModuleKpi**
> []ModuleVersionDto ResetModuleKpi(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModuleVersionDto**](ModuleVersionDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetProductKpi**
> []ProductDto ResetProductKpi(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

