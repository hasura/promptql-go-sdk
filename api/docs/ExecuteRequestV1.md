# ExecuteRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **NullableString** |  | [optional] 
**Code** | **string** |  | 
**PromptqlApiKey** | Pointer to **NullableString** |  | [optional] 
**AiPrimitivesLlm** | [**NullableExecuteRequestV1AiPrimitivesLlm**](ExecuteRequestV1AiPrimitivesLlm.md) |  | 
**Ddn** | [**NullableDdnConfig**](DdnConfig.md) |  | 
**Artifacts** | [**[]ExecuteRequestV1ArtifactsInner**](ExecuteRequestV1ArtifactsInner.md) |  | 

## Methods

### NewExecuteRequestV1

`func NewExecuteRequestV1(code string, aiPrimitivesLlm NullableExecuteRequestV1AiPrimitivesLlm, ddn NullableDdnConfig, artifacts []ExecuteRequestV1ArtifactsInner, ) *ExecuteRequestV1`

NewExecuteRequestV1 instantiates a new ExecuteRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRequestV1WithDefaults

`func NewExecuteRequestV1WithDefaults() *ExecuteRequestV1`

NewExecuteRequestV1WithDefaults instantiates a new ExecuteRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ExecuteRequestV1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExecuteRequestV1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExecuteRequestV1) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExecuteRequestV1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExecuteRequestV1) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExecuteRequestV1) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCode

`func (o *ExecuteRequestV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExecuteRequestV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExecuteRequestV1) SetCode(v string)`

SetCode sets Code field to given value.


### GetPromptqlApiKey

`func (o *ExecuteRequestV1) GetPromptqlApiKey() string`

GetPromptqlApiKey returns the PromptqlApiKey field if non-nil, zero value otherwise.

### GetPromptqlApiKeyOk

`func (o *ExecuteRequestV1) GetPromptqlApiKeyOk() (*string, bool)`

GetPromptqlApiKeyOk returns a tuple with the PromptqlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptqlApiKey

`func (o *ExecuteRequestV1) SetPromptqlApiKey(v string)`

SetPromptqlApiKey sets PromptqlApiKey field to given value.

### HasPromptqlApiKey

`func (o *ExecuteRequestV1) HasPromptqlApiKey() bool`

HasPromptqlApiKey returns a boolean if a field has been set.

### SetPromptqlApiKeyNil

`func (o *ExecuteRequestV1) SetPromptqlApiKeyNil(b bool)`

 SetPromptqlApiKeyNil sets the value for PromptqlApiKey to be an explicit nil

### UnsetPromptqlApiKey
`func (o *ExecuteRequestV1) UnsetPromptqlApiKey()`

UnsetPromptqlApiKey ensures that no value is present for PromptqlApiKey, not even an explicit nil
### GetAiPrimitivesLlm

`func (o *ExecuteRequestV1) GetAiPrimitivesLlm() ExecuteRequestV1AiPrimitivesLlm`

GetAiPrimitivesLlm returns the AiPrimitivesLlm field if non-nil, zero value otherwise.

### GetAiPrimitivesLlmOk

`func (o *ExecuteRequestV1) GetAiPrimitivesLlmOk() (*ExecuteRequestV1AiPrimitivesLlm, bool)`

GetAiPrimitivesLlmOk returns a tuple with the AiPrimitivesLlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPrimitivesLlm

`func (o *ExecuteRequestV1) SetAiPrimitivesLlm(v ExecuteRequestV1AiPrimitivesLlm)`

SetAiPrimitivesLlm sets AiPrimitivesLlm field to given value.


### SetAiPrimitivesLlmNil

`func (o *ExecuteRequestV1) SetAiPrimitivesLlmNil(b bool)`

 SetAiPrimitivesLlmNil sets the value for AiPrimitivesLlm to be an explicit nil

### UnsetAiPrimitivesLlm
`func (o *ExecuteRequestV1) UnsetAiPrimitivesLlm()`

UnsetAiPrimitivesLlm ensures that no value is present for AiPrimitivesLlm, not even an explicit nil
### GetDdn

`func (o *ExecuteRequestV1) GetDdn() DdnConfig`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *ExecuteRequestV1) GetDdnOk() (*DdnConfig, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *ExecuteRequestV1) SetDdn(v DdnConfig)`

SetDdn sets Ddn field to given value.


### SetDdnNil

`func (o *ExecuteRequestV1) SetDdnNil(b bool)`

 SetDdnNil sets the value for Ddn to be an explicit nil

### UnsetDdn
`func (o *ExecuteRequestV1) UnsetDdn()`

UnsetDdn ensures that no value is present for Ddn, not even an explicit nil
### GetArtifacts

`func (o *ExecuteRequestV1) GetArtifacts() []ExecuteRequestV1ArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ExecuteRequestV1) GetArtifactsOk() (*[]ExecuteRequestV1ArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ExecuteRequestV1) SetArtifacts(v []ExecuteRequestV1ArtifactsInner)`

SetArtifacts sets Artifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


