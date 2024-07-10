# {{classname}}

All URIs are relative to *https://okapi.sma.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductResourceApi.md#CreateProduct) | **Post** /api/kpi/products | 
[**DeleteProduct**](ProductResourceApi.md#DeleteProduct) | **Delete** /api/kpi/products/{id} | 
[**GetAllProducts**](ProductResourceApi.md#GetAllProducts) | **Get** /api/kpi/products | 
[**GetAllProductsWithFlatInformation**](ProductResourceApi.md#GetAllProductsWithFlatInformation) | **Get** /api/kpi/products/all | 
[**GetProduct**](ProductResourceApi.md#GetProduct) | **Get** /api/kpi/products/{id} | 
[**GetProductByName**](ProductResourceApi.md#GetProductByName) | **Get** /api/kpi/products/by-name/{productName} | 
[**PartialUpdateProduct**](ProductResourceApi.md#PartialUpdateProduct) | **Patch** /api/kpi/products | 
[**ProductGroups**](ProductResourceApi.md#ProductGroups) | **Get** /api/kpi/products/groups | 
[**ProductTags**](ProductResourceApi.md#ProductTags) | **Get** /api/kpi/products/tags | 
[**UpdateProduct**](ProductResourceApi.md#UpdateProduct) | **Put** /api/kpi/products | 

# **CreateProduct**
> ProductDto CreateProduct(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductDto**](ProductDto.md)|  | 

### Return type

[**ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProduct**
> DeleteProduct(ctx, id)


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

# **GetAllProducts**
> []ProductDto GetAllProducts(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductResourceApiGetAllProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductResourceApiGetAllProductsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **size** | **optional.Int32**| The size of the page to be returned | [default to 20]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**[]ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProductsWithFlatInformation**
> []ProductDto GetAllProductsWithFlatInformation(ctx, )


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

# **GetProduct**
> ProductDto GetProduct(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductByName**
> ProductDto GetProductByName(ctx, productName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productName** | **string**|  | 

### Return type

[**ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateProduct**
> ProductDto PartialUpdateProduct(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductDto**](ProductDto.md)|  | 

### Return type

[**ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGroups**
> []string ProductGroups(ctx, )


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

# **ProductTags**
> []string ProductTags(ctx, )


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

# **UpdateProduct**
> ProductDto UpdateProduct(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductDto**](ProductDto.md)|  | 

### Return type

[**ProductDto**](ProductDto.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

