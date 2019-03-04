# \FacadeApi

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorDetail**](FacadeApi.md#AuthorDetail) | **Get** /facade/author/{id} | Author Detail
[**ListPosts**](FacadeApi.md#ListPosts) | **Get** /facade/post | List Posts
[**PostDetail**](FacadeApi.md#PostDetail) | **Get** /facade/post/{id} | Post Detail


# **AuthorDetail**
> FacadeModel AuthorDetail(ctx, id)
Author Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| For what author detail is requested | 

### Return type

[**FacadeModel**](Facade.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPosts**
> PostListModel ListPosts(ctx, )
List Posts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PostListModel**](PostList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDetail**
> FacadeModel PostDetail(ctx, id)
Post Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| For what post detail is requested | 

### Return type

[**FacadeModel**](Facade.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

