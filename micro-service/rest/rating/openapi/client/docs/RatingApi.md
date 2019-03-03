# \RatingApi

All URIs are relative to *http://localhost:4003*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRating**](RatingApi.md#CreateRating) | **Post** /rating | Create rating
[**DeleteRating**](RatingApi.md#DeleteRating) | **Delete** /rating/{id} | Delete rating
[**GetRatingById**](RatingApi.md#GetRatingById) | **Get** /rating/{id} | Get rating by id
[**ListRatings**](RatingApi.md#ListRatings) | **Get** /rating | List ratings


# **CreateRating**
> RatingModel CreateRating(ctx, createRatingModel)
Create rating

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createRatingModel** | [**CreateRatingModel**](CreateRatingModel.md)| Created rating object | 

### Return type

[**RatingModel**](Rating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRating**
> DeleteRating(ctx, id)
Delete rating

Delete a rating by id

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

# **GetRatingById**
> RatingModel GetRatingById(ctx, id)
Get rating by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The id that needs to be fetched. Use 0 for testing. | 

### Return type

[**RatingModel**](Rating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRatings**
> RatingListModel ListRatings(ctx, postId)
List ratings

a list of ratings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postId** | **int64**| Only get ratings of the supplied post | 

### Return type

[**RatingListModel**](RatingList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

