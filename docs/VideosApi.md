# \VideosApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVideoRecorder**](VideosApi.md#GetVideoRecorder) | **Get** /videos/live/getRecorder | Get Live Video Recorder HTML
[**MarkLiveRecordingComplete**](VideosApi.md#MarkLiveRecordingComplete) | **Post** /videos/live/markComplete | Completes a live recording
[**SignUpload**](VideosApi.md#SignUpload) | **Post** /video/signedUpload | Generate Signed Url


# **GetVideoRecorder**
> VideoRecorderMethodResponse GetVideoRecorder($width, $videoId)

Get Live Video Recorder HTML

Returns an object with a number of properties to help you put a video recorder on your site.         This is to be used in conjunction with the VideoRecordedLive call one the user has confirmed in your UI that         the video is how they want it.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **int32**| The width of the recorder to present. | [optional] 
 **videoId** | **string**| The id of the video to record | [optional] 

### Return type

[**VideoRecorderMethodResponse**](VideoRecorderMethodResponse.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkLiveRecordingComplete**
> VideoPublicRepresentation MarkLiveRecordingComplete($videoId, $filename, $title)

Completes a live recording

Used in conjunction with the live recorder method to mark a video recording as complete.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string**| The id of the video to mark as done. | 
 **filename** | **string**| The filename that was chosen as the final video. | 
 **title** | **string**| The title to give the video | 

### Return type

[**VideoPublicRepresentation**](VideoPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignUpload**
> string SignUpload($policy, $v4)

Generate Signed Url

Generates a signed url to be used for video uploads.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**SignUploadRequest**](SignUploadRequest.md)| The policy to sign | 
 **v4** | **bool**| Whether to do v4 signing | [optional] 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

