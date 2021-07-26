# \AppConfirmTerminationApi

All URIs are relative to *https://localhost/sandboxname/mec_app_support/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsConfirmTerminationPOST**](AppConfirmTerminationApi.md#ApplicationsConfirmTerminationPOST) | **Post** /applications/{appInstanceId}/confirm_termination | 


# **ApplicationsConfirmTerminationPOST**
> ApplicationsConfirmTerminationPOST(ctx, body, appInstanceId)


This method is used to confirm the application level termination  of an application instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppTerminationConfirmation**](AppTerminationConfirmation.md)|  | 
  **appInstanceId** | **string**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC platform manager. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
