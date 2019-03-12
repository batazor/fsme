# Fsme.EventApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sendEventFSM**](EventApi.md#sendEventFSM) | **POST** /{id}/event | 


<a name="sendEventFSM"></a>
# **sendEventFSM**
> Event sendEventFSM(id)



Send event to fsm by ID

### Example
```javascript
import Fsme from 'fsme';

let apiInstance = new Fsme.EventApi();

let id = "id_example"; // String | 

apiInstance.sendEventFSM(id).then((data) => {
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

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/io.goswagger.examples.todo-list.v1+json
 - **Accept**: application/io.goswagger.examples.todo-list.v1+json

