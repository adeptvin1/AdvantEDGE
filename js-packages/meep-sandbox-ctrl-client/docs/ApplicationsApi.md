# AdvantEdgeSandboxControllerRestApi.ApplicationsApi

All URIs are relative to *https://localhost/sandboxname/sandbox-ctrl/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**applicationsAppInstanceIdDELETE**](ApplicationsApi.md#applicationsAppInstanceIdDELETE) | **DELETE** /applications/{appInstanceId} | 
[**applicationsAppInstanceIdGET**](ApplicationsApi.md#applicationsAppInstanceIdGET) | **GET** /applications/{appInstanceId} | 
[**applicationsAppInstanceIdPUT**](ApplicationsApi.md#applicationsAppInstanceIdPUT) | **PUT** /applications/{appInstanceId} | 
[**applicationsGET**](ApplicationsApi.md#applicationsGET) | **GET** /applications | 
[**applicationsPOST**](ApplicationsApi.md#applicationsPOST) | **POST** /applications | 


<a name="applicationsAppInstanceIdDELETE"></a>
# **applicationsAppInstanceIdDELETE**
> applicationsAppInstanceIdDELETE(appInstanceId)



This method deletes a mec application resource.

### Example
```javascript
var AdvantEdgeSandboxControllerRestApi = require('advant_edge_sandbox_controller_rest_api');

var apiInstance = new AdvantEdgeSandboxControllerRestApi.ApplicationsApi();

var appInstanceId = "appInstanceId_example"; // String | Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC application manager POST method.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.applicationsAppInstanceIdDELETE(appInstanceId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **String**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC application manager POST method. | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="applicationsAppInstanceIdGET"></a>
# **applicationsAppInstanceIdGET**
> ApplicationInfo applicationsAppInstanceIdGET(appInstanceId)



This method retrieves information about a mec application resource.

### Example
```javascript
var AdvantEdgeSandboxControllerRestApi = require('advant_edge_sandbox_controller_rest_api');

var apiInstance = new AdvantEdgeSandboxControllerRestApi.ApplicationsApi();

var appInstanceId = "appInstanceId_example"; // String | Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC application manager POST method.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.applicationsAppInstanceIdGET(appInstanceId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **String**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC application manager POST method. | 

### Return type

[**ApplicationInfo**](ApplicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="applicationsAppInstanceIdPUT"></a>
# **applicationsAppInstanceIdPUT**
> ApplicationInfo applicationsAppInstanceIdPUT(appInstanceIdapplicationInfo)



This method updates the information about a mec application resource.

### Example
```javascript
var AdvantEdgeSandboxControllerRestApi = require('advant_edge_sandbox_controller_rest_api');

var apiInstance = new AdvantEdgeSandboxControllerRestApi.ApplicationsApi();

var appInstanceId = "appInstanceId_example"; // String | Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC application manager POST method.

var applicationInfo = new AdvantEdgeSandboxControllerRestApi.ApplicationInfo(); // ApplicationInfo | Application information


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.applicationsAppInstanceIdPUT(appInstanceIdapplicationInfo, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceId** | **String**| Represents a MEC application instance. Note that the appInstanceId is allocated by the MEC application manager POST method. | 
 **applicationInfo** | [**ApplicationInfo**](ApplicationInfo.md)| Application information | 

### Return type

[**ApplicationInfo**](ApplicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="applicationsGET"></a>
# **applicationsGET**
> [ApplicationInfo] applicationsGET(opts)



This method retrieves information about a list of mec application resources.

### Example
```javascript
var AdvantEdgeSandboxControllerRestApi = require('advant_edge_sandbox_controller_rest_api');

var apiInstance = new AdvantEdgeSandboxControllerRestApi.ApplicationsApi();

var opts = { 
  'app': "app_example", // String | Application name
  'type': "type_example", // String | Application type
  'nodeName': "nodeName_example" // String | Node name
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.applicationsGET(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **String**| Application name | [optional] 
 **type** | **String**| Application type | [optional] 
 **nodeName** | **String**| Node name | [optional] 

### Return type

[**[ApplicationInfo]**](ApplicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="applicationsPOST"></a>
# **applicationsPOST**
> ApplicationInfo applicationsPOST(applicationInfo)



This method is used to create a mec application resource.

### Example
```javascript
var AdvantEdgeSandboxControllerRestApi = require('advant_edge_sandbox_controller_rest_api');

var apiInstance = new AdvantEdgeSandboxControllerRestApi.ApplicationsApi();

var applicationInfo = new AdvantEdgeSandboxControllerRestApi.ApplicationInfo(); // ApplicationInfo | Application information


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.applicationsPOST(applicationInfo, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationInfo** | [**ApplicationInfo**](ApplicationInfo.md)| Application information | 

### Return type

[**ApplicationInfo**](ApplicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

