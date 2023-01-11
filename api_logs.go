/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// LogsApiService LogsApi service
type LogsApiService service

type ApiLogsDestroyRequest struct {
	ctx context.Context
	ApiService *LogsApiService
	id string
}

func (r ApiLogsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.LogsDestroyExecute(r)
}

/*
LogsDestroy Method for LogsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiLogsDestroyRequest
*/
func (a *LogsApiService) LogsDestroy(ctx context.Context, id string) ApiLogsDestroyRequest {
	return ApiLogsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *LogsApiService) LogsDestroyExecute(r ApiLogsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.LogsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/logs/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiLogsListRequest struct {
	ctx context.Context
	ApiService *LogsApiService
	causedBy *int32
	dateAfter *string
	dateBefore *string
	datetime *time.Time
	details *string
	device *int32
	deviceSnOrName *string
	deviceGroup *int32
	deviceGroupName *string
	deviceGroupLabel *int32
	event *string
	owner *string
	page *int32
	search *string
}

func (r ApiLogsListRequest) CausedBy(causedBy int32) ApiLogsListRequest {
	r.causedBy = &causedBy
	return r
}

func (r ApiLogsListRequest) DateAfter(dateAfter string) ApiLogsListRequest {
	r.dateAfter = &dateAfter
	return r
}

func (r ApiLogsListRequest) DateBefore(dateBefore string) ApiLogsListRequest {
	r.dateBefore = &dateBefore
	return r
}

func (r ApiLogsListRequest) Datetime(datetime time.Time) ApiLogsListRequest {
	r.datetime = &datetime
	return r
}

func (r ApiLogsListRequest) Details(details string) ApiLogsListRequest {
	r.details = &details
	return r
}

func (r ApiLogsListRequest) Device(device int32) ApiLogsListRequest {
	r.device = &device
	return r
}

func (r ApiLogsListRequest) DeviceSnOrName(deviceSnOrName string) ApiLogsListRequest {
	r.deviceSnOrName = &deviceSnOrName
	return r
}

func (r ApiLogsListRequest) DeviceGroup(deviceGroup int32) ApiLogsListRequest {
	r.deviceGroup = &deviceGroup
	return r
}

func (r ApiLogsListRequest) DeviceGroupName(deviceGroupName string) ApiLogsListRequest {
	r.deviceGroupName = &deviceGroupName
	return r
}

func (r ApiLogsListRequest) DeviceGroupLabel(deviceGroupLabel int32) ApiLogsListRequest {
	r.deviceGroupLabel = &deviceGroupLabel
	return r
}

func (r ApiLogsListRequest) Event(event string) ApiLogsListRequest {
	r.event = &event
	return r
}

// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
func (r ApiLogsListRequest) Owner(owner string) ApiLogsListRequest {
	r.owner = &owner
	return r
}

// A page number within the paginated result set.
func (r ApiLogsListRequest) Page(page int32) ApiLogsListRequest {
	r.page = &page
	return r
}

// A search term.
func (r ApiLogsListRequest) Search(search string) ApiLogsListRequest {
	r.search = &search
	return r
}

func (r ApiLogsListRequest) Execute() (*PaginatedLogList, *http.Response, error) {
	return r.ApiService.LogsListExecute(r)
}

/*
LogsList Method for LogsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLogsListRequest
*/
func (a *LogsApiService) LogsList(ctx context.Context) ApiLogsListRequest {
	return ApiLogsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedLogList
func (a *LogsApiService) LogsListExecute(r ApiLogsListRequest) (*PaginatedLogList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedLogList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.LogsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/logs/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.causedBy != nil {
		localVarQueryParams.Add("caused_by", parameterToString(*r.causedBy, ""))
	}
	if r.dateAfter != nil {
		localVarQueryParams.Add("date_after", parameterToString(*r.dateAfter, ""))
	}
	if r.dateBefore != nil {
		localVarQueryParams.Add("date_before", parameterToString(*r.dateBefore, ""))
	}
	if r.datetime != nil {
		localVarQueryParams.Add("datetime", parameterToString(*r.datetime, ""))
	}
	if r.details != nil {
		localVarQueryParams.Add("details", parameterToString(*r.details, ""))
	}
	if r.device != nil {
		localVarQueryParams.Add("device", parameterToString(*r.device, ""))
	}
	if r.deviceSnOrName != nil {
		localVarQueryParams.Add("device__sn_or_name", parameterToString(*r.deviceSnOrName, ""))
	}
	if r.deviceGroup != nil {
		localVarQueryParams.Add("device_group", parameterToString(*r.deviceGroup, ""))
	}
	if r.deviceGroupName != nil {
		localVarQueryParams.Add("device_group__name", parameterToString(*r.deviceGroupName, ""))
	}
	if r.deviceGroupLabel != nil {
		localVarQueryParams.Add("device_group_label", parameterToString(*r.deviceGroupLabel, ""))
	}
	if r.event != nil {
		localVarQueryParams.Add("event", parameterToString(*r.event, ""))
	}
	if r.owner != nil {
		localVarQueryParams.Add("owner", parameterToString(*r.owner, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiLogsRetrieveRequest struct {
	ctx context.Context
	ApiService *LogsApiService
	id int32
}

func (r ApiLogsRetrieveRequest) Execute() (*Log, *http.Response, error) {
	return r.ApiService.LogsRetrieveExecute(r)
}

/*
LogsRetrieve Method for LogsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this log.
 @return ApiLogsRetrieveRequest
*/
func (a *LogsApiService) LogsRetrieve(ctx context.Context, id int32) ApiLogsRetrieveRequest {
	return ApiLogsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Log
func (a *LogsApiService) LogsRetrieveExecute(r ApiLogsRetrieveRequest) (*Log, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Log
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.LogsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/logs/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
