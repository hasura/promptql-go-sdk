# AssistantActionChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** |  | [optional] 
**Plan** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**CodeOutput** | Pointer to **NullableString** |  | [optional] 
**CodeError** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Index** | **int32** |  | 

## Methods

### NewAssistantActionChunk

`func NewAssistantActionChunk(type_ string, index int32, ) *AssistantActionChunk`

NewAssistantActionChunk instantiates a new AssistantActionChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantActionChunkWithDefaults

`func NewAssistantActionChunkWithDefaults() *AssistantActionChunk`

NewAssistantActionChunkWithDefaults instantiates a new AssistantActionChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AssistantActionChunk) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AssistantActionChunk) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AssistantActionChunk) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AssistantActionChunk) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AssistantActionChunk) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AssistantActionChunk) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPlan

`func (o *AssistantActionChunk) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AssistantActionChunk) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AssistantActionChunk) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AssistantActionChunk) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### SetPlanNil

`func (o *AssistantActionChunk) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *AssistantActionChunk) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetCode

`func (o *AssistantActionChunk) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AssistantActionChunk) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AssistantActionChunk) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AssistantActionChunk) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AssistantActionChunk) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AssistantActionChunk) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCodeOutput

`func (o *AssistantActionChunk) GetCodeOutput() string`

GetCodeOutput returns the CodeOutput field if non-nil, zero value otherwise.

### GetCodeOutputOk

`func (o *AssistantActionChunk) GetCodeOutputOk() (*string, bool)`

GetCodeOutputOk returns a tuple with the CodeOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOutput

`func (o *AssistantActionChunk) SetCodeOutput(v string)`

SetCodeOutput sets CodeOutput field to given value.

### HasCodeOutput

`func (o *AssistantActionChunk) HasCodeOutput() bool`

HasCodeOutput returns a boolean if a field has been set.

### SetCodeOutputNil

`func (o *AssistantActionChunk) SetCodeOutputNil(b bool)`

 SetCodeOutputNil sets the value for CodeOutput to be an explicit nil

### UnsetCodeOutput
`func (o *AssistantActionChunk) UnsetCodeOutput()`

UnsetCodeOutput ensures that no value is present for CodeOutput, not even an explicit nil
### GetCodeError

`func (o *AssistantActionChunk) GetCodeError() string`

GetCodeError returns the CodeError field if non-nil, zero value otherwise.

### GetCodeErrorOk

`func (o *AssistantActionChunk) GetCodeErrorOk() (*string, bool)`

GetCodeErrorOk returns a tuple with the CodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeError

`func (o *AssistantActionChunk) SetCodeError(v string)`

SetCodeError sets CodeError field to given value.

### HasCodeError

`func (o *AssistantActionChunk) HasCodeError() bool`

HasCodeError returns a boolean if a field has been set.

### SetCodeErrorNil

`func (o *AssistantActionChunk) SetCodeErrorNil(b bool)`

 SetCodeErrorNil sets the value for CodeError to be an explicit nil

### UnsetCodeError
`func (o *AssistantActionChunk) UnsetCodeError()`

UnsetCodeError ensures that no value is present for CodeError, not even an explicit nil
### GetType

`func (o *AssistantActionChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssistantActionChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssistantActionChunk) SetType(v string)`

SetType sets Type field to given value.


### GetIndex

`func (o *AssistantActionChunk) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AssistantActionChunk) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AssistantActionChunk) SetIndex(v int32)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


