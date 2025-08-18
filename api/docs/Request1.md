# Request1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | **bool** |  | 
**Artifacts** | Pointer to [**[]ExecuteRequestV1ArtifactsInner**](ExecuteRequestV1ArtifactsInner.md) |  | [optional] 
**Timezone** | **string** | An IANA timezone used to interpret queries that implicitly require timezones | 
**Interactions** | [**[]ApiThreadInteraction**](ApiThreadInteraction.md) |  | 
**Version** | **string** |  | 
**PromptqlApiKey** | Pointer to **string** |  | [optional] 
**Llm** | Pointer to [**Llm**](Llm.md) |  | [optional] 
**AiPrimitivesLlm** | Pointer to [**NullableExecuteRequestV1AiPrimitivesLlm**](ExecuteRequestV1AiPrimitivesLlm.md) |  | [optional] 
**Ddn** | [**DdnConfigV2**](DdnConfigV2.md) |  | 
**SystemInstructions** | Pointer to **string** |  | [optional] 

## Methods

### NewRequest1

`func NewRequest1(stream bool, timezone string, interactions []ApiThreadInteraction, version string, ddn DdnConfigV2, ) *Request1`

NewRequest1 instantiates a new Request1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequest1WithDefaults

`func NewRequest1WithDefaults() *Request1`

NewRequest1WithDefaults instantiates a new Request1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *Request1) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *Request1) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *Request1) SetStream(v bool)`

SetStream sets Stream field to given value.


### GetArtifacts

`func (o *Request1) GetArtifacts() []ExecuteRequestV1ArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Request1) GetArtifactsOk() (*[]ExecuteRequestV1ArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Request1) SetArtifacts(v []ExecuteRequestV1ArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *Request1) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetTimezone

`func (o *Request1) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Request1) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Request1) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetInteractions

`func (o *Request1) GetInteractions() []ApiThreadInteraction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *Request1) GetInteractionsOk() (*[]ApiThreadInteraction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *Request1) SetInteractions(v []ApiThreadInteraction)`

SetInteractions sets Interactions field to given value.


### GetVersion

`func (o *Request1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Request1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Request1) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPromptqlApiKey

`func (o *Request1) GetPromptqlApiKey() string`

GetPromptqlApiKey returns the PromptqlApiKey field if non-nil, zero value otherwise.

### GetPromptqlApiKeyOk

`func (o *Request1) GetPromptqlApiKeyOk() (*string, bool)`

GetPromptqlApiKeyOk returns a tuple with the PromptqlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptqlApiKey

`func (o *Request1) SetPromptqlApiKey(v string)`

SetPromptqlApiKey sets PromptqlApiKey field to given value.

### HasPromptqlApiKey

`func (o *Request1) HasPromptqlApiKey() bool`

HasPromptqlApiKey returns a boolean if a field has been set.

### GetLlm

`func (o *Request1) GetLlm() Llm`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *Request1) GetLlmOk() (*Llm, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *Request1) SetLlm(v Llm)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *Request1) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetAiPrimitivesLlm

`func (o *Request1) GetAiPrimitivesLlm() ExecuteRequestV1AiPrimitivesLlm`

GetAiPrimitivesLlm returns the AiPrimitivesLlm field if non-nil, zero value otherwise.

### GetAiPrimitivesLlmOk

`func (o *Request1) GetAiPrimitivesLlmOk() (*ExecuteRequestV1AiPrimitivesLlm, bool)`

GetAiPrimitivesLlmOk returns a tuple with the AiPrimitivesLlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPrimitivesLlm

`func (o *Request1) SetAiPrimitivesLlm(v ExecuteRequestV1AiPrimitivesLlm)`

SetAiPrimitivesLlm sets AiPrimitivesLlm field to given value.

### HasAiPrimitivesLlm

`func (o *Request1) HasAiPrimitivesLlm() bool`

HasAiPrimitivesLlm returns a boolean if a field has been set.

### SetAiPrimitivesLlmNil

`func (o *Request1) SetAiPrimitivesLlmNil(b bool)`

 SetAiPrimitivesLlmNil sets the value for AiPrimitivesLlm to be an explicit nil

### UnsetAiPrimitivesLlm
`func (o *Request1) UnsetAiPrimitivesLlm()`

UnsetAiPrimitivesLlm ensures that no value is present for AiPrimitivesLlm, not even an explicit nil
### GetDdn

`func (o *Request1) GetDdn() DdnConfigV2`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *Request1) GetDdnOk() (*DdnConfigV2, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *Request1) SetDdn(v DdnConfigV2)`

SetDdn sets Ddn field to given value.


### GetSystemInstructions

`func (o *Request1) GetSystemInstructions() string`

GetSystemInstructions returns the SystemInstructions field if non-nil, zero value otherwise.

### GetSystemInstructionsOk

`func (o *Request1) GetSystemInstructionsOk() (*string, bool)`

GetSystemInstructionsOk returns a tuple with the SystemInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInstructions

`func (o *Request1) SetSystemInstructions(v string)`

SetSystemInstructions sets SystemInstructions field to given value.

### HasSystemInstructions

`func (o *Request1) HasSystemInstructions() bool`

HasSystemInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


