# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllPwnedNoteKeys**](ProductReleaseResourceApi.md#AllPwnedNoteKeys) | **Get** /api/kpi/product-releases/pwned-note-keys | 
[**CreateProductRelease**](ProductReleaseResourceApi.md#CreateProductRelease) | **Post** /api/kpi/product-releases | 
[**DeleteProductRelease**](ProductReleaseResourceApi.md#DeleteProductRelease) | **Delete** /api/kpi/product-releases/{id} | 
[**GeProductReleasesByName**](ProductReleaseResourceApi.md#GeProductReleasesByName) | **Get** /api/kpi/product-releases/by-name/{productName}/{revision} | 
[**GetAllLatestProductReleases**](ProductReleaseResourceApi.md#GetAllLatestProductReleases) | **Get** /api/kpi/product-releases/latest | 
[**GetAllProductReleases**](ProductReleaseResourceApi.md#GetAllProductReleases) | **Get** /api/kpi/product-releases | 
[**GetAllProductReleasesForProduct**](ProductReleaseResourceApi.md#GetAllProductReleasesForProduct) | **Get** /api/kpi/product-releases/by-product/{productId} | 
[**GetAllProductReleasesForProductName**](ProductReleaseResourceApi.md#GetAllProductReleasesForProductName) | **Get** /api/kpi/product-releases/by-name/{productName} | 
[**GetLatestProductReleasesForProduct**](ProductReleaseResourceApi.md#GetLatestProductReleasesForProduct) | **Get** /api/kpi/product-releases/latest/by-product/{productId} | 
[**GetPreviousModuleVersion**](ProductReleaseResourceApi.md#GetPreviousModuleVersion) | **Get** /api/kpi/product-releases/previous/{id} | 
[**GetProductRelease**](ProductReleaseResourceApi.md#GetProductRelease) | **Get** /api/kpi/product-releases/{id} | 
[**PartialUpdateProductRelease**](ProductReleaseResourceApi.md#PartialUpdateProductRelease) | **Patch** /api/kpi/product-releases | 
[**UpdateProductRelease**](ProductReleaseResourceApi.md#UpdateProductRelease) | **Put** /api/kpi/product-releases | 

# **AllPwnedNoteKeys**
> []string AllPwnedNoteKeys(ctx, )


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

# **CreateProductRelease**
> ProductReleaseDto CreateProductRelease(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductReleaseDto**](ProductReleaseDto.md)|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductRelease**
> DeleteProductRelease(ctx, id)


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

# **GeProductReleasesByName**
> ProductReleaseDto GeProductReleasesByName(ctx, productName, revision)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productName** | **string**|  | 
  **revision** | **string**|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllLatestProductReleases**
> []ProductReleaseDto GetAllLatestProductReleases(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductReleaseResourceApiGetAllLatestProductReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductReleaseResourceApiGetAllLatestProductReleasesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versionKind** | **optional.String**|  | [default to ALL]
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProductReleases**
> []ProductReleaseDto GetAllProductReleases(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductReleaseResourceApiGetAllProductReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductReleaseResourceApiGetAllProductReleasesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProductReleasesForProduct**
> []ProductReleaseDto GetAllProductReleasesForProduct(ctx, productId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
 **optional** | ***ProductReleaseResourceApiGetAllProductReleasesForProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductReleaseResourceApiGetAllProductReleasesForProductOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProductReleasesForProductName**
> []ProductReleaseDto GetAllProductReleasesForProductName(ctx, productName, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productName** | **string**|  | 
 **optional** | ***ProductReleaseResourceApiGetAllProductReleasesForProductNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductReleaseResourceApiGetAllProductReleasesForProductNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestProductReleasesForProduct**
> ProductReleaseDto GetLatestProductReleasesForProduct(ctx, productId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
 **optional** | ***ProductReleaseResourceApiGetLatestProductReleasesForProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductReleaseResourceApiGetLatestProductReleasesForProductOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionKind** | **optional.String**|  | [default to ALL]

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreviousModuleVersion**
> ProductReleaseDto GetPreviousModuleVersion(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProductReleaseResourceApiGetPreviousModuleVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductReleaseResourceApiGetPreviousModuleVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionKind** | **optional.String**|  | [default to ALL]

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductRelease**
> ProductReleaseDto GetProductRelease(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateProductRelease**
> ProductReleaseDto PartialUpdateProductRelease(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductReleaseDto**](ProductReleaseDto.md)|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductRelease**
> ProductReleaseDto UpdateProductRelease(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductReleaseDto**](ProductReleaseDto.md)|  | 

### Return type

[**ProductReleaseDto**](ProductReleaseDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

