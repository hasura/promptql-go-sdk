# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteProgramExecuteProgramPost**](DefaultAPI.md#ExecuteProgramExecuteProgramPost) | **Post** /execute_program | Execute Program
[**QueryQueryPost**](DefaultAPI.md#QueryQueryPost) | **Post** /query | Query



## ExecuteProgramExecuteProgramPost

> PromptQlExecutionResult ExecuteProgramExecuteProgramPost(ctx).ExecuteRequest(executeRequest).Execute()

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
	executeRequest := *openapiclient.NewExecuteRequest("Code_example", "TODO", "TODO", []openapiclient.ExecuteRequestArtifactsInner{openapiclient.ExecuteRequest_artifacts_inner{TableArtifact: openapiclient.NewTableArtifact("Identifier_example", "Title_example", "ArtifactType_example", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}})}}) // ExecuteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExecuteProgramExecuteProgramPost(context.Background()).ExecuteRequest(executeRequest).Execute()
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
 **executeRequest** | [**ExecuteRequest**](ExecuteRequest.md) |  | 

### Return type

[**PromptQlExecutionResult**](PromptQlExecutionResult.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryQueryPost

> QueryResponse QueryQueryPost(ctx).QueryRequest(queryRequest).Execute()

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
	queryRequest := *openapiclient.NewQueryRequest("Version_example", *openapiclient.NewDdnConfig("Url_example"), "Timezone_example", []openapiclient.Interaction{*openapiclient.NewInteraction(*openapiclient.NewUserMessage("Text_example"))}, false) // QueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QueryQueryPost(context.Background()).QueryRequest(queryRequest).Execute()
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
 **queryRequest** | [**QueryRequest**](QueryRequest.md) |  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

