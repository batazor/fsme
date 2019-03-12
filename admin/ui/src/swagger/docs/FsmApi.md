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



Create a new fsm

### Example
```javascript
import Fsme from 'fsme';

let apiInstance = new Fsme.FsmApi();

let opts = { 
  'body': new Fsme.Fsm() // Fsm | 
};
apiInstance.addFSM(opts).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

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



Delete a fsm

### Example
```javascript
import Fsme from 'fsme';

let apiInstance = new Fsme.FsmApi();

let id = "id_example"; // String | 

apiInstance.destroyFSM(id).then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

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



Get a fsm by ID

### Example
```javascript
import Fsme from 'fsme';

let apiInstance = new Fsme.FsmApi();

let id = "id_example"; // String | 

apiInstance.getFSM(id).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

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



Get list fsm

### Example
```javascript
import Fsme from 'fsme';

let apiInstance = new Fsme.FsmApi();

let opts = { 
  'enable': true, // Boolean | 
  'limit': 10 // Number | 
};
apiInstance.getFSMList(opts).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

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



Update a fsm

### Example
```javascript
import Fsme from 'fsme';

let apiInstance = new Fsme.FsmApi();

let id = "id_example"; // String | 

let opts = { 
  'body': new Fsme.Fsm() // Fsm | 
};
apiInstance.updateFSM(id, opts).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

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

