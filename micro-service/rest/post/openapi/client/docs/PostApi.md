# \PostApi

All URIs are relative to *http://localhost:4002*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePost**](PostApi.md#CreatePost) | **Post** /post | Create post
[**DeletePost**](PostApi.md#DeletePost) | **Delete** /post/{id} | Delete post
[**GetPostById**](PostApi.md#GetPostById) | **Get** /post/{id} | Get post by id
[**ListPosts**](PostApi.md#ListPosts) | **Get** /post | List posts


# **CreatePost**
> PostModel CreatePost(ctx, createPostModel)
Create post

This can only be done by the logged in post.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPostModel** | [**CreatePostModel**](CreatePostModel.md)| Created post object | 

### Return type

[**PostModel**](Post.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePost**
> DeletePost(ctx, id)
Delete post

Delete a post by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The id that needs to be deleted. Use 0 for testing. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPostById**
> PostModel GetPostById(ctx, id)
Get post by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The id that needs to be fetched. Use 0 for testing. | 

### Return type

[**PostModel**](Post.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPosts**
> PostListModel ListPosts(ctx, optional)
List posts

a list of posts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPostsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **optional.Int64**| Only get post of the supplied author | 

### Return type

[**PostListModel**](PostList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

