/*
PromptQL API

PromptQL API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiExecuteProgramExecuteProgramPostRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	executeRequest *ExecuteRequest
}

func (r ApiExecuteProgramExecuteProgramPostRequest) ExecuteRequest(executeRequest ExecuteRequest) ApiExecuteProgramExecuteProgramPostRequest {
	r.executeRequest = &executeRequest
	return r
}

func (r ApiExecuteProgramExecuteProgramPostRequest) Execute() (*PromptQlExecutionResult, *http.Response, error) {
	return r.ApiService.ExecuteProgramExecuteProgramPostExecute(r)
}

/*
ExecuteProgramExecuteProgramPost Execute Program

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExecuteProgramExecuteProgramPostRequest
*/
func (a *DefaultAPIService) ExecuteProgramExecuteProgramPost(ctx context.Context) ApiExecuteProgramExecuteProgramPostRequest {
	return ApiExecuteProgramExecuteProgramPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}


// ExecuteProgramExecuteProgramPostRequest creates an http request.
// @return PromptQlExecutionResult
func (a *DefaultAPIService) ExecuteProgramExecuteProgramPostRequest(r ApiExecuteProgramExecuteProgramPostRequest) (*http.Request, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ExecuteProgramExecuteProgramPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/execute_program"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.executeRequest == nil {
		return nil, reportError("executeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.executeRequest

	return a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
}

// ExecuteProgramExecuteProgramPostExecuteRequest executes the http request. Returns the http response.
func (a *DefaultAPIService) ExecuteProgramExecuteProgramPostExecuteRequest(req *http.Request) (*http.Response, error) {
	localVarHTTPResponse, err := a.client.cfg.HTTPClient.Do(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		
		if err != nil {
			return localVarHTTPResponse, err
		}
		
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// Execute executes the request
//  @return PromptQlExecutionResult
func (a *DefaultAPIService) ExecuteProgramExecuteProgramPostExecute(r ApiExecuteProgramExecuteProgramPostRequest) (*PromptQlExecutionResult, *http.Response, error) {
	var (
		localVarReturnValue  *PromptQlExecutionResult
	)

	req, err := a.ExecuteProgramExecuteProgramPostRequest(r)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.ExecuteProgramExecuteProgramPostExecuteRequest(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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

type ApiQueryQueryPostRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	queryRequest *QueryRequest
}

func (r ApiQueryQueryPostRequest) QueryRequest(queryRequest QueryRequest) ApiQueryQueryPostRequest {
	r.queryRequest = &queryRequest
	return r
}

func (r ApiQueryQueryPostRequest) Execute() (*QueryResponse, *http.Response, error) {
	return r.ApiService.QueryQueryPostExecute(r)
}

/*
QueryQueryPost Query

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryQueryPostRequest
*/
func (a *DefaultAPIService) QueryQueryPost(ctx context.Context) ApiQueryQueryPostRequest {
	return ApiQueryQueryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}


// QueryQueryPostRequest creates an http request.
// @return QueryResponse
func (a *DefaultAPIService) QueryQueryPostRequest(r ApiQueryQueryPostRequest) (*http.Request, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.QueryQueryPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.queryRequest == nil {
		return nil, reportError("queryRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/event-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.queryRequest

	return a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
}

// QueryQueryPostExecuteRequest executes the http request. Returns the http response.
func (a *DefaultAPIService) QueryQueryPostExecuteRequest(req *http.Request) (*http.Response, error) {
	localVarHTTPResponse, err := a.client.cfg.HTTPClient.Do(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		
		if err != nil {
			return localVarHTTPResponse, err
		}
		
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// Execute executes the request
//  @return QueryResponse
func (a *DefaultAPIService) QueryQueryPostExecute(r ApiQueryQueryPostRequest) (*QueryResponse, *http.Response, error) {
	var (
		localVarReturnValue  *QueryResponse
	)

	req, err := a.QueryQueryPostRequest(r)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.QueryQueryPostExecuteRequest(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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
