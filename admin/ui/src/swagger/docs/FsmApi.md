# Fsme.FsmApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addFSM**](FsmApi.md#addFSM) | **POST** / | 
[**destroyFSM**](FsmApi.md#destroyFSM) | **DELETE** /{id} | 
[**getFSM**](FsmApi.md#getFSM) | **GET** /{id} | 
[**getFSMList**](FsmApi.md#getFSMList) | **GET** / | 
[**updateFSM**](FsmApi.md#updateFSM) | **PATCH** /{id} | 


<a name="addFSM"></a>
# **addFSM**
> Fsm addFSM(opts)



### Example
```javascript
var Fsme = require('fsme');

var apiInstance = new Fsme.FsmApi();

var opts = { 
  'body': new Fsme.Fsm() // Fsm | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.addFSM(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Fsm**](Fsm.md)|  | [optional] 

### Return type

[**Fsm**](Fsm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/io.goswagger.examples.todo-list.v1+json
 - **Accept**: application/io.goswagger.examples.todo-list.v1+json

<a name="destroyFSM"></a>
# **destroyFSM**
> destroyFSM(id)



### Example
```javascript
var Fsme = require('fsme');

var apiInstance = new Fsme.FsmApi();

var id = "id_example"; // String | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.destroyFSM(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/io.goswagger.examples.todo-list.v1+json
 - **Accept**: application/io.goswagger.examples.todo-list.v1+json

<a name="getFSM"></a>
# **getFSM**
> Fsm getFSM(id)



### Example
```javascript
var Fsme = require('fsme');

var apiInstance = new Fsme.FsmApi();

var id = "id_example"; // String | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getFSM(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**Fsm**](Fsm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/io.goswagger.examples.todo-list.v1+json
 - **Accept**: application/io.goswagger.examples.todo-list.v1+json

<a name="getFSMList"></a>
# **getFSMList**
> [Fsm] getFSMList(opts)



### Example
```javascript
var Fsme = require('fsme');

var apiInstance = new Fsme.FsmApi();

var opts = { 
  'enable': true, // Boolean | 
  'limit': 10 // Number | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getFSMList(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enable** | **Boolean**|  | [optional] 
 **limit** | **Number**|  | [optional] [default to 10]

### Return type

[**[Fsm]**](Fsm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/io.goswagger.examples.todo-list.v1+json
 - **Accept**: application/io.goswagger.examples.todo-list.v1+json

<a name="updateFSM"></a>
# **updateFSM**
> Fsm updateFSM(id, opts)



### Example
```javascript
var Fsme = require('fsme');

var apiInstance = new Fsme.FsmApi();

var id = "id_example"; // String | 

var opts = { 
  'body': new Fsme.Fsm() // Fsm | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.updateFSM(id, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 
 **body** | [**Fsm**](Fsm.md)|  | [optional] 

### Return type

[**Fsm**](Fsm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/io.goswagger.examples.todo-list.v1+json
 - **Accept**: application/io.goswagger.examples.todo-list.v1+json

