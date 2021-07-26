# \AppServicesApi

All URIs are relative to *https://localhost/sandboxname/mec_service_mgmt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppServicesGET**](AppServicesApi.md#AppServicesGET) | **Get** /applications/{appInstanceId}/services | 
[**AppServicesPOST**](AppServicesApi.md#AppServicesPOST) | **Post** /applications/{appInstanceId}/services | 
[**AppServicesServiceIdDELETE**](AppServicesApi.md#AppServicesServiceIdDELETE) | **Delete** /applications/{appInstanceId}/services/{serviceId} | 
[**AppServicesServiceIdGET**](AppServicesApi.md#AppServicesServiceIdGET) | **Get** /applications/{appInstanceId}/services/{serviceId} | 
[**AppServicesServiceIdPUT**](AppServicesApi.md#AppServicesServiceIdPUT) | **Put** /applications/{appInstanceId}/services/{serviceId} | 


# **AppServicesGET**
> []ServiceInfo AppServicesGET(ctx, appInstanceId, optional)


This method retrieves information about a list of mecService resources. This method is typically used in \"service availability query\" procedure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appInstanceId** | **string**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC platform manager. | 
 **optional** | ***AppServicesGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppServicesGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serInstanceId** | [**optional.Interface of []string**](string.md)| A MEC application instance may use multiple ser_instance_ids as an input parameter to query the availability of a list of MEC service instances. Either \&quot;ser_instance_id\&quot; or \&quot;ser_name\&quot; or \&quot;ser_category_id\&quot; or none of them shall be present. | 
 **serName** | [**optional.Interface of []string**](string.md)| A MEC application instance may use multiple ser_names as an input parameter to query the availability of a list of MEC service instances. Either \&quot;ser_instance_id\&quot; or \&quot;ser_name\&quot; or \&quot;ser_category_id\&quot; or none of them shall be present. | 
 **serCategoryId** | **optional.String**| A MEC application instance may use ser_category_id as an input parameter to query the availability of a list of MEC service instances in a serCategory. Either \&quot;ser_instance_id\&quot; or \&quot;ser_name\&quot; or \&quot;ser_category_id\&quot; or none of them shall be present. | 
 **consumedLocalOnly** | **optional.Bool**| Indicate whether the service can only be consumed by the MEC  applications located in the same locality (as defined by  scopeOfLocality) as this service instance. | 
 **isLocal** | **optional.Bool**| Indicate whether the service is located in the same locality (as  defined by scopeOfLocality) as the consuming MEC application. | 
 **scopeOfLocality** | **optional.String**| A MEC application instance may use scope_of_locality as an input  parameter to query the availability of a list of MEC service instances  with a certain scope of locality. | 

### Return type

[**[]ServiceInfo**](ServiceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppServicesPOST**
> ServiceInfo AppServicesPOST(ctx, body, appInstanceId)


This method is used to create a mecService resource. This method is typically used in \"service availability update and new service registration\" procedure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInfoPost**](ServiceInfoPost.md)| New ServiceInfo with updated &quot;state&quot; is included as entity body of the request | 
  **appInstanceId** | **string**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC platform manager. | 

### Return type

[**ServiceInfo**](ServiceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppServicesServiceIdDELETE**
> AppServicesServiceIdDELETE(ctx, appInstanceId, serviceId)


This method deletes a mecService resource. This method is typically used in the service deregistration procedure. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appInstanceId** | **string**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC platform manager. | 
  **serviceId** | **string**| Represents a MEC service instance. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppServicesServiceIdGET**
> ServiceInfo AppServicesServiceIdGET(ctx, appInstanceId, serviceId)


This method retrieves information about a mecService resource. This method is typically used in \"service availability query\" procedure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appInstanceId** | **string**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC platform manager. | 
  **serviceId** | **string**| Represents a MEC service instance. | 

### Return type

[**ServiceInfo**](ServiceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppServicesServiceIdPUT**
> ServiceInfo AppServicesServiceIdPUT(ctx, body, appInstanceId, serviceId)


This method updates the information about a mecService resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInfo**](ServiceInfo.md)| New ServiceInfo with updated &quot;state&quot; is included as entity body of the request | 
  **appInstanceId** | **string**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC platform manager. | 
  **serviceId** | **string**| Represents a MEC service instance. | 

### Return type

[**ServiceInfo**](ServiceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
