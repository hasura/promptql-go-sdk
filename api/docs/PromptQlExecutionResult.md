# PromptQlExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | **string** |  | 
**Error** | **NullableString** |  | 
**AccessedArtifactIds** | **[]string** |  | 
**ModifiedArtifacts** | [**[]ExecuteRequestArtifactsInner**](ExecuteRequestArtifactsInner.md) |  | 
**LlmUsages** | [**[]LlmUsage**](LlmUsage.md) |  | 

## Methods

### NewPromptQlExecutionResult

`func NewPromptQlExecutionResult(output string, error_ NullableString, accessedArtifactIds []string, modifiedArtifacts []ExecuteRequestArtifactsInner, llmUsages []LlmUsage, ) *PromptQlExecutionResult`

NewPromptQlExecutionResult instantiates a new PromptQlExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptQlExecutionResultWithDefaults

`func NewPromptQlExecutionResultWithDefaults() *PromptQlExecutionResult`

NewPromptQlExecutionResultWithDefaults instantiates a new PromptQlExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *PromptQlExecutionResult) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PromptQlExecutionResult) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PromptQlExecutionResult) SetOutput(v string)`

SetOutput sets Output field to given value.


### GetError

`func (o *PromptQlExecutionResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PromptQlExecutionResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PromptQlExecutionResult) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *PromptQlExecutionResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *PromptQlExecutionResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAccessedArtifactIds

`func (o *PromptQlExecutionResult) GetAccessedArtifactIds() []string`

GetAccessedArtifactIds returns the AccessedArtifactIds field if non-nil, zero value otherwise.

### GetAccessedArtifactIdsOk

`func (o *PromptQlExecutionResult) GetAccessedArtifactIdsOk() (*[]string, bool)`

GetAccessedArtifactIdsOk returns a tuple with the AccessedArtifactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedArtifactIds

`func (o *PromptQlExecutionResult) SetAccessedArtifactIds(v []string)`

SetAccessedArtifactIds sets AccessedArtifactIds field to given value.


### GetModifiedArtifacts

`func (o *PromptQlExecutionResult) GetModifiedArtifacts() []ExecuteRequestArtifactsInner`

GetModifiedArtifacts returns the ModifiedArtifacts field if non-nil, zero value otherwise.

### GetModifiedArtifactsOk

`func (o *PromptQlExecutionResult) GetModifiedArtifactsOk() (*[]ExecuteRequestArtifactsInner, bool)`

GetModifiedArtifactsOk returns a tuple with the ModifiedArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedArtifacts

`func (o *PromptQlExecutionResult) SetModifiedArtifacts(v []ExecuteRequestArtifactsInner)`

SetModifiedArtifacts sets ModifiedArtifacts field to given value.


### GetLlmUsages

`func (o *PromptQlExecutionResult) GetLlmUsages() []LlmUsage`

GetLlmUsages returns the LlmUsages field if non-nil, zero value otherwise.

### GetLlmUsagesOk

`func (o *PromptQlExecutionResult) GetLlmUsagesOk() (*[]LlmUsage, bool)`

GetLlmUsagesOk returns a tuple with the LlmUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmUsages

`func (o *PromptQlExecutionResult) SetLlmUsages(v []LlmUsage)`

SetLlmUsages sets LlmUsages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


