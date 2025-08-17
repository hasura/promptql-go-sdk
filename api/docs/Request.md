# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Code** | **string** |  | 
**PromptqlApiKey** | Pointer to **string** |  | [optional] 
**AiPrimitivesLlm** | [**NullableExecuteRequestV1AiPrimitivesLlm**](ExecuteRequestV1AiPrimitivesLlm.md) |  | 
**Ddn** | [**DdnConfigV2**](DdnConfigV2.md) |  | 
**Artifacts** | [**[]ExecuteRequestV1ArtifactsInner**](ExecuteRequestV1ArtifactsInner.md) |  | 

## Methods

### NewRequest

`func NewRequest(version string, code string, aiPrimitivesLlm NullableExecuteRequestV1AiPrimitivesLlm, ddn DdnConfigV2, artifacts []ExecuteRequestV1ArtifactsInner, ) *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Request) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Request) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Request) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCode

`func (o *Request) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Request) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Request) SetCode(v string)`

SetCode sets Code field to given value.


### GetPromptqlApiKey

`func (o *Request) GetPromptqlApiKey() string`

GetPromptqlApiKey returns the PromptqlApiKey field if non-nil, zero value otherwise.

### GetPromptqlApiKeyOk

`func (o *Request) GetPromptqlApiKeyOk() (*string, bool)`

GetPromptqlApiKeyOk returns a tuple with the PromptqlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptqlApiKey

`func (o *Request) SetPromptqlApiKey(v string)`

SetPromptqlApiKey sets PromptqlApiKey field to given value.

### HasPromptqlApiKey

`func (o *Request) HasPromptqlApiKey() bool`

HasPromptqlApiKey returns a boolean if a field has been set.

### GetAiPrimitivesLlm

`func (o *Request) GetAiPrimitivesLlm() ExecuteRequestV1AiPrimitivesLlm`

GetAiPrimitivesLlm returns the AiPrimitivesLlm field if non-nil, zero value otherwise.

### GetAiPrimitivesLlmOk

`func (o *Request) GetAiPrimitivesLlmOk() (*ExecuteRequestV1AiPrimitivesLlm, bool)`

GetAiPrimitivesLlmOk returns a tuple with the AiPrimitivesLlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPrimitivesLlm

`func (o *Request) SetAiPrimitivesLlm(v ExecuteRequestV1AiPrimitivesLlm)`

SetAiPrimitivesLlm sets AiPrimitivesLlm field to given value.


### SetAiPrimitivesLlmNil

`func (o *Request) SetAiPrimitivesLlmNil(b bool)`

 SetAiPrimitivesLlmNil sets the value for AiPrimitivesLlm to be an explicit nil

### UnsetAiPrimitivesLlm
`func (o *Request) UnsetAiPrimitivesLlm()`

UnsetAiPrimitivesLlm ensures that no value is present for AiPrimitivesLlm, not even an explicit nil
### GetDdn

`func (o *Request) GetDdn() DdnConfigV2`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *Request) GetDdnOk() (*DdnConfigV2, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *Request) SetDdn(v DdnConfigV2)`

SetDdn sets Ddn field to given value.


### GetArtifacts

`func (o *Request) GetArtifacts() []ExecuteRequestV1ArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Request) GetArtifactsOk() (*[]ExecuteRequestV1ArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Request) SetArtifacts(v []ExecuteRequestV1ArtifactsInner)`

SetArtifacts sets Artifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


