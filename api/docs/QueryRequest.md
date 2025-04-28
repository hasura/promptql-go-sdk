# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**PromptqlApiKey** | Pointer to **NullableString** |  | [optional] 
**Llm** | Pointer to [**Llm**](Llm.md) |  | [optional] 
**AiPrimitivesLlm** | Pointer to [**NullableExecuteRequestAiPrimitivesLlm**](ExecuteRequestAiPrimitivesLlm.md) |  | [optional] 
**Ddn** | [**DdnConfig**](DdnConfig.md) |  | 
**Artifacts** | Pointer to [**[]ExecuteRequestArtifactsInner**](ExecuteRequestArtifactsInner.md) |  | [optional] 
**SystemInstructions** | Pointer to **NullableString** |  | [optional] 
**Timezone** | **string** | An IANA timezone used to interpret queries that implicitly require timezones | 
**Interactions** | [**[]Interaction**](Interaction.md) |  | 
**Stream** | **bool** |  | 

## Methods

### NewQueryRequest

`func NewQueryRequest(version string, ddn DdnConfig, timezone string, interactions []Interaction, stream bool, ) *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *QueryRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPromptqlApiKey

`func (o *QueryRequest) GetPromptqlApiKey() string`

GetPromptqlApiKey returns the PromptqlApiKey field if non-nil, zero value otherwise.

### GetPromptqlApiKeyOk

`func (o *QueryRequest) GetPromptqlApiKeyOk() (*string, bool)`

GetPromptqlApiKeyOk returns a tuple with the PromptqlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptqlApiKey

`func (o *QueryRequest) SetPromptqlApiKey(v string)`

SetPromptqlApiKey sets PromptqlApiKey field to given value.

### HasPromptqlApiKey

`func (o *QueryRequest) HasPromptqlApiKey() bool`

HasPromptqlApiKey returns a boolean if a field has been set.

### SetPromptqlApiKeyNil

`func (o *QueryRequest) SetPromptqlApiKeyNil(b bool)`

 SetPromptqlApiKeyNil sets the value for PromptqlApiKey to be an explicit nil

### UnsetPromptqlApiKey
`func (o *QueryRequest) UnsetPromptqlApiKey()`

UnsetPromptqlApiKey ensures that no value is present for PromptqlApiKey, not even an explicit nil
### GetLlm

`func (o *QueryRequest) GetLlm() Llm`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *QueryRequest) GetLlmOk() (*Llm, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *QueryRequest) SetLlm(v Llm)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *QueryRequest) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetAiPrimitivesLlm

`func (o *QueryRequest) GetAiPrimitivesLlm() ExecuteRequestAiPrimitivesLlm`

GetAiPrimitivesLlm returns the AiPrimitivesLlm field if non-nil, zero value otherwise.

### GetAiPrimitivesLlmOk

`func (o *QueryRequest) GetAiPrimitivesLlmOk() (*ExecuteRequestAiPrimitivesLlm, bool)`

GetAiPrimitivesLlmOk returns a tuple with the AiPrimitivesLlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPrimitivesLlm

`func (o *QueryRequest) SetAiPrimitivesLlm(v ExecuteRequestAiPrimitivesLlm)`

SetAiPrimitivesLlm sets AiPrimitivesLlm field to given value.

### HasAiPrimitivesLlm

`func (o *QueryRequest) HasAiPrimitivesLlm() bool`

HasAiPrimitivesLlm returns a boolean if a field has been set.

### SetAiPrimitivesLlmNil

`func (o *QueryRequest) SetAiPrimitivesLlmNil(b bool)`

 SetAiPrimitivesLlmNil sets the value for AiPrimitivesLlm to be an explicit nil

### UnsetAiPrimitivesLlm
`func (o *QueryRequest) UnsetAiPrimitivesLlm()`

UnsetAiPrimitivesLlm ensures that no value is present for AiPrimitivesLlm, not even an explicit nil
### GetDdn

`func (o *QueryRequest) GetDdn() DdnConfig`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *QueryRequest) GetDdnOk() (*DdnConfig, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *QueryRequest) SetDdn(v DdnConfig)`

SetDdn sets Ddn field to given value.


### GetArtifacts

`func (o *QueryRequest) GetArtifacts() []ExecuteRequestArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *QueryRequest) GetArtifactsOk() (*[]ExecuteRequestArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *QueryRequest) SetArtifacts(v []ExecuteRequestArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *QueryRequest) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetSystemInstructions

`func (o *QueryRequest) GetSystemInstructions() string`

GetSystemInstructions returns the SystemInstructions field if non-nil, zero value otherwise.

### GetSystemInstructionsOk

`func (o *QueryRequest) GetSystemInstructionsOk() (*string, bool)`

GetSystemInstructionsOk returns a tuple with the SystemInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInstructions

`func (o *QueryRequest) SetSystemInstructions(v string)`

SetSystemInstructions sets SystemInstructions field to given value.

### HasSystemInstructions

`func (o *QueryRequest) HasSystemInstructions() bool`

HasSystemInstructions returns a boolean if a field has been set.

### SetSystemInstructionsNil

`func (o *QueryRequest) SetSystemInstructionsNil(b bool)`

 SetSystemInstructionsNil sets the value for SystemInstructions to be an explicit nil

### UnsetSystemInstructions
`func (o *QueryRequest) UnsetSystemInstructions()`

UnsetSystemInstructions ensures that no value is present for SystemInstructions, not even an explicit nil
### GetTimezone

`func (o *QueryRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *QueryRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *QueryRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetInteractions

`func (o *QueryRequest) GetInteractions() []Interaction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *QueryRequest) GetInteractionsOk() (*[]Interaction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *QueryRequest) SetInteractions(v []Interaction)`

SetInteractions sets Interactions field to given value.


### GetStream

`func (o *QueryRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *QueryRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *QueryRequest) SetStream(v bool)`

SetStream sets Stream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


