# QueryRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | **bool** |  | 
**Artifacts** | Pointer to [**[]ExecuteRequestV1ArtifactsInner**](ExecuteRequestV1ArtifactsInner.md) |  | [optional] 
**Timezone** | **string** | An IANA timezone used to interpret queries that implicitly require timezones | 
**Interactions** | [**[]ApiThreadInteraction**](ApiThreadInteraction.md) |  | 
**Version** | **string** |  | 
**PromptqlApiKey** | Pointer to **NullableString** |  | [optional] 
**Llm** | Pointer to [**Llm**](Llm.md) |  | [optional] 
**AiPrimitivesLlm** | Pointer to [**NullableExecuteRequestV1AiPrimitivesLlm**](ExecuteRequestV1AiPrimitivesLlm.md) |  | [optional] 
**Ddn** | [**DdnConfig**](DdnConfig.md) |  | 
**SystemInstructions** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQueryRequestV1

`func NewQueryRequestV1(stream bool, timezone string, interactions []ApiThreadInteraction, version string, ddn DdnConfig, ) *QueryRequestV1`

NewQueryRequestV1 instantiates a new QueryRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestV1WithDefaults

`func NewQueryRequestV1WithDefaults() *QueryRequestV1`

NewQueryRequestV1WithDefaults instantiates a new QueryRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *QueryRequestV1) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *QueryRequestV1) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *QueryRequestV1) SetStream(v bool)`

SetStream sets Stream field to given value.


### GetArtifacts

`func (o *QueryRequestV1) GetArtifacts() []ExecuteRequestV1ArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *QueryRequestV1) GetArtifactsOk() (*[]ExecuteRequestV1ArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *QueryRequestV1) SetArtifacts(v []ExecuteRequestV1ArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *QueryRequestV1) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetTimezone

`func (o *QueryRequestV1) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *QueryRequestV1) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *QueryRequestV1) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetInteractions

`func (o *QueryRequestV1) GetInteractions() []ApiThreadInteraction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *QueryRequestV1) GetInteractionsOk() (*[]ApiThreadInteraction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *QueryRequestV1) SetInteractions(v []ApiThreadInteraction)`

SetInteractions sets Interactions field to given value.


### GetVersion

`func (o *QueryRequestV1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryRequestV1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryRequestV1) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPromptqlApiKey

`func (o *QueryRequestV1) GetPromptqlApiKey() string`

GetPromptqlApiKey returns the PromptqlApiKey field if non-nil, zero value otherwise.

### GetPromptqlApiKeyOk

`func (o *QueryRequestV1) GetPromptqlApiKeyOk() (*string, bool)`

GetPromptqlApiKeyOk returns a tuple with the PromptqlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptqlApiKey

`func (o *QueryRequestV1) SetPromptqlApiKey(v string)`

SetPromptqlApiKey sets PromptqlApiKey field to given value.

### HasPromptqlApiKey

`func (o *QueryRequestV1) HasPromptqlApiKey() bool`

HasPromptqlApiKey returns a boolean if a field has been set.

### SetPromptqlApiKeyNil

`func (o *QueryRequestV1) SetPromptqlApiKeyNil(b bool)`

 SetPromptqlApiKeyNil sets the value for PromptqlApiKey to be an explicit nil

### UnsetPromptqlApiKey
`func (o *QueryRequestV1) UnsetPromptqlApiKey()`

UnsetPromptqlApiKey ensures that no value is present for PromptqlApiKey, not even an explicit nil
### GetLlm

`func (o *QueryRequestV1) GetLlm() Llm`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *QueryRequestV1) GetLlmOk() (*Llm, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *QueryRequestV1) SetLlm(v Llm)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *QueryRequestV1) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetAiPrimitivesLlm

`func (o *QueryRequestV1) GetAiPrimitivesLlm() ExecuteRequestV1AiPrimitivesLlm`

GetAiPrimitivesLlm returns the AiPrimitivesLlm field if non-nil, zero value otherwise.

### GetAiPrimitivesLlmOk

`func (o *QueryRequestV1) GetAiPrimitivesLlmOk() (*ExecuteRequestV1AiPrimitivesLlm, bool)`

GetAiPrimitivesLlmOk returns a tuple with the AiPrimitivesLlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPrimitivesLlm

`func (o *QueryRequestV1) SetAiPrimitivesLlm(v ExecuteRequestV1AiPrimitivesLlm)`

SetAiPrimitivesLlm sets AiPrimitivesLlm field to given value.

### HasAiPrimitivesLlm

`func (o *QueryRequestV1) HasAiPrimitivesLlm() bool`

HasAiPrimitivesLlm returns a boolean if a field has been set.

### SetAiPrimitivesLlmNil

`func (o *QueryRequestV1) SetAiPrimitivesLlmNil(b bool)`

 SetAiPrimitivesLlmNil sets the value for AiPrimitivesLlm to be an explicit nil

### UnsetAiPrimitivesLlm
`func (o *QueryRequestV1) UnsetAiPrimitivesLlm()`

UnsetAiPrimitivesLlm ensures that no value is present for AiPrimitivesLlm, not even an explicit nil
### GetDdn

`func (o *QueryRequestV1) GetDdn() DdnConfig`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *QueryRequestV1) GetDdnOk() (*DdnConfig, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *QueryRequestV1) SetDdn(v DdnConfig)`

SetDdn sets Ddn field to given value.


### GetSystemInstructions

`func (o *QueryRequestV1) GetSystemInstructions() string`

GetSystemInstructions returns the SystemInstructions field if non-nil, zero value otherwise.

### GetSystemInstructionsOk

`func (o *QueryRequestV1) GetSystemInstructionsOk() (*string, bool)`

GetSystemInstructionsOk returns a tuple with the SystemInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInstructions

`func (o *QueryRequestV1) SetSystemInstructions(v string)`

SetSystemInstructions sets SystemInstructions field to given value.

### HasSystemInstructions

`func (o *QueryRequestV1) HasSystemInstructions() bool`

HasSystemInstructions returns a boolean if a field has been set.

### SetSystemInstructionsNil

`func (o *QueryRequestV1) SetSystemInstructionsNil(b bool)`

 SetSystemInstructionsNil sets the value for SystemInstructions to be an explicit nil

### UnsetSystemInstructions
`func (o *QueryRequestV1) UnsetSystemInstructions()`

UnsetSystemInstructions ensures that no value is present for SystemInstructions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


