# ExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**PromptqlApiKey** | Pointer to **NullableString** |  | [optional] 
**AiPrimitivesLlm** | [**NullableExecuteRequestAiPrimitivesLlm**](ExecuteRequestAiPrimitivesLlm.md) |  | 
**Ddn** | [**NullableDdnConfig**](DdnConfig.md) |  | 
**Artifacts** | [**[]ExecuteRequestArtifactsInner**](ExecuteRequestArtifactsInner.md) |  | 

## Methods

### NewExecuteRequest

`func NewExecuteRequest(code string, aiPrimitivesLlm NullableExecuteRequestAiPrimitivesLlm, ddn NullableDdnConfig, artifacts []ExecuteRequestArtifactsInner, ) *ExecuteRequest`

NewExecuteRequest instantiates a new ExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRequestWithDefaults

`func NewExecuteRequestWithDefaults() *ExecuteRequest`

NewExecuteRequestWithDefaults instantiates a new ExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExecuteRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExecuteRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExecuteRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetPromptqlApiKey

`func (o *ExecuteRequest) GetPromptqlApiKey() string`

GetPromptqlApiKey returns the PromptqlApiKey field if non-nil, zero value otherwise.

### GetPromptqlApiKeyOk

`func (o *ExecuteRequest) GetPromptqlApiKeyOk() (*string, bool)`

GetPromptqlApiKeyOk returns a tuple with the PromptqlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptqlApiKey

`func (o *ExecuteRequest) SetPromptqlApiKey(v string)`

SetPromptqlApiKey sets PromptqlApiKey field to given value.

### HasPromptqlApiKey

`func (o *ExecuteRequest) HasPromptqlApiKey() bool`

HasPromptqlApiKey returns a boolean if a field has been set.

### SetPromptqlApiKeyNil

`func (o *ExecuteRequest) SetPromptqlApiKeyNil(b bool)`

 SetPromptqlApiKeyNil sets the value for PromptqlApiKey to be an explicit nil

### UnsetPromptqlApiKey
`func (o *ExecuteRequest) UnsetPromptqlApiKey()`

UnsetPromptqlApiKey ensures that no value is present for PromptqlApiKey, not even an explicit nil
### GetAiPrimitivesLlm

`func (o *ExecuteRequest) GetAiPrimitivesLlm() ExecuteRequestAiPrimitivesLlm`

GetAiPrimitivesLlm returns the AiPrimitivesLlm field if non-nil, zero value otherwise.

### GetAiPrimitivesLlmOk

`func (o *ExecuteRequest) GetAiPrimitivesLlmOk() (*ExecuteRequestAiPrimitivesLlm, bool)`

GetAiPrimitivesLlmOk returns a tuple with the AiPrimitivesLlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPrimitivesLlm

`func (o *ExecuteRequest) SetAiPrimitivesLlm(v ExecuteRequestAiPrimitivesLlm)`

SetAiPrimitivesLlm sets AiPrimitivesLlm field to given value.


### SetAiPrimitivesLlmNil

`func (o *ExecuteRequest) SetAiPrimitivesLlmNil(b bool)`

 SetAiPrimitivesLlmNil sets the value for AiPrimitivesLlm to be an explicit nil

### UnsetAiPrimitivesLlm
`func (o *ExecuteRequest) UnsetAiPrimitivesLlm()`

UnsetAiPrimitivesLlm ensures that no value is present for AiPrimitivesLlm, not even an explicit nil
### GetDdn

`func (o *ExecuteRequest) GetDdn() DdnConfig`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *ExecuteRequest) GetDdnOk() (*DdnConfig, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *ExecuteRequest) SetDdn(v DdnConfig)`

SetDdn sets Ddn field to given value.


### SetDdnNil

`func (o *ExecuteRequest) SetDdnNil(b bool)`

 SetDdnNil sets the value for Ddn to be an explicit nil

### UnsetDdn
`func (o *ExecuteRequest) UnsetDdn()`

UnsetDdn ensures that no value is present for Ddn, not even an explicit nil
### GetArtifacts

`func (o *ExecuteRequest) GetArtifacts() []ExecuteRequestArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ExecuteRequest) GetArtifactsOk() (*[]ExecuteRequestArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ExecuteRequest) SetArtifacts(v []ExecuteRequestArtifactsInner)`

SetArtifacts sets Artifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


