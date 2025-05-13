# QueryResponseChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ThreadId** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CodeOutput** | Pointer to **string** |  | [optional] 
**CodeError** | Pointer to **string** |  | [optional] 
**Index** | **int32** |  | 
**Artifact** | [**Artifact**](Artifact.md) |  | 
**Error** | **string** |  | 

## Methods

### NewQueryResponseChunk

`func NewQueryResponseChunk(type_ string, threadId string, index int32, artifact Artifact, error_ string, ) *QueryResponseChunk`

NewQueryResponseChunk instantiates a new QueryResponseChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseChunkWithDefaults

`func NewQueryResponseChunkWithDefaults() *QueryResponseChunk`

NewQueryResponseChunkWithDefaults instantiates a new QueryResponseChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryResponseChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryResponseChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryResponseChunk) SetType(v string)`

SetType sets Type field to given value.


### GetThreadId

`func (o *QueryResponseChunk) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *QueryResponseChunk) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *QueryResponseChunk) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMessage

`func (o *QueryResponseChunk) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QueryResponseChunk) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QueryResponseChunk) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QueryResponseChunk) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPlan

`func (o *QueryResponseChunk) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *QueryResponseChunk) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *QueryResponseChunk) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *QueryResponseChunk) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCode

`func (o *QueryResponseChunk) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *QueryResponseChunk) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *QueryResponseChunk) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *QueryResponseChunk) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodeOutput

`func (o *QueryResponseChunk) GetCodeOutput() string`

GetCodeOutput returns the CodeOutput field if non-nil, zero value otherwise.

### GetCodeOutputOk

`func (o *QueryResponseChunk) GetCodeOutputOk() (*string, bool)`

GetCodeOutputOk returns a tuple with the CodeOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOutput

`func (o *QueryResponseChunk) SetCodeOutput(v string)`

SetCodeOutput sets CodeOutput field to given value.

### HasCodeOutput

`func (o *QueryResponseChunk) HasCodeOutput() bool`

HasCodeOutput returns a boolean if a field has been set.

### GetCodeError

`func (o *QueryResponseChunk) GetCodeError() string`

GetCodeError returns the CodeError field if non-nil, zero value otherwise.

### GetCodeErrorOk

`func (o *QueryResponseChunk) GetCodeErrorOk() (*string, bool)`

GetCodeErrorOk returns a tuple with the CodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeError

`func (o *QueryResponseChunk) SetCodeError(v string)`

SetCodeError sets CodeError field to given value.

### HasCodeError

`func (o *QueryResponseChunk) HasCodeError() bool`

HasCodeError returns a boolean if a field has been set.

### GetIndex

`func (o *QueryResponseChunk) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *QueryResponseChunk) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *QueryResponseChunk) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetArtifact

`func (o *QueryResponseChunk) GetArtifact() Artifact`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *QueryResponseChunk) GetArtifactOk() (*Artifact, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *QueryResponseChunk) SetArtifact(v Artifact)`

SetArtifact sets Artifact field to given value.


### GetError

`func (o *QueryResponseChunk) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseChunk) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseChunk) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


