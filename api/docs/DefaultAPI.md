# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteProgramExecuteProgramPost**](DefaultAPI.md#ExecuteProgramExecuteProgramPost) | **Post** /execute_program | Execute Program
[**QueryQueryPost**](DefaultAPI.md#QueryQueryPost) | **Post** /query | Query



## ExecuteProgramExecuteProgramPost

> PromptQlExecutionResult ExecuteProgramExecuteProgramPost(ctx).Request(request).Execute()

Execute Program

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewRequest("Version_example", "Code_example", "TODO", *openapiclient.NewDdnConfigV2(), []openapiclient.ExecuteRequestV1ArtifactsInner{openapiclient.ExecuteRequestV1_artifacts_inner{AutomationArtifact: openapiclient.NewAutomationArtifact("Identifier_example", "Title_example", "ArtifactType_example", *openapiclient.NewAutomationArtifactData("Code_example"))}}) // Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExecuteProgramExecuteProgramPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExecuteProgramExecuteProgramPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteProgramExecuteProgramPost`: PromptQlExecutionResult
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExecuteProgramExecuteProgramPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteProgramExecuteProgramPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**Request**](Request.md) |  | 

### Return type

[**PromptQlExecutionResult**](PromptQlExecutionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryQueryPost

> QueryResponse QueryQueryPost(ctx).Request1(request1).Execute()

Query

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request1 := openapiclient.Request_1{QueryRequestV1: openapiclient.NewQueryRequestV1(false, "Timezone_example", []openapiclient.ApiThreadInteraction{*openapiclient.NewApiThreadInteraction(*openapiclient.NewApiThreadUserMessage("Text_example"))}, "Version_example", *openapiclient.NewDdnConfig("Url_example"))} // Request1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QueryQueryPost(context.Background()).Request1(request1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QueryQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryQueryPost`: QueryResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QueryQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request1** | [**Request1**](Request1.md) |  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

