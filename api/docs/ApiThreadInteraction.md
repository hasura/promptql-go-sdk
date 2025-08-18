# ApiThreadInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMessage** | [**ApiThreadUserMessage**](ApiThreadUserMessage.md) |  | 
**AssistantActions** | Pointer to [**[]ApiThreadAssistantAction**](ApiThreadAssistantAction.md) |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**CompletionTimestamp** | Pointer to **NullableTime** |  | [optional] 
**CanceledTimestamp** | Pointer to **NullableTime** |  | [optional] 
**ErrorTimestamp** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewApiThreadInteraction

`func NewApiThreadInteraction(userMessage ApiThreadUserMessage, ) *ApiThreadInteraction`

NewApiThreadInteraction instantiates a new ApiThreadInteraction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiThreadInteractionWithDefaults

`func NewApiThreadInteractionWithDefaults() *ApiThreadInteraction`

NewApiThreadInteractionWithDefaults instantiates a new ApiThreadInteraction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMessage

`func (o *ApiThreadInteraction) GetUserMessage() ApiThreadUserMessage`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *ApiThreadInteraction) GetUserMessageOk() (*ApiThreadUserMessage, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *ApiThreadInteraction) SetUserMessage(v ApiThreadUserMessage)`

SetUserMessage sets UserMessage field to given value.


### GetAssistantActions

`func (o *ApiThreadInteraction) GetAssistantActions() []ApiThreadAssistantAction`

GetAssistantActions returns the AssistantActions field if non-nil, zero value otherwise.

### GetAssistantActionsOk

`func (o *ApiThreadInteraction) GetAssistantActionsOk() (*[]ApiThreadAssistantAction, bool)`

GetAssistantActionsOk returns a tuple with the AssistantActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantActions

`func (o *ApiThreadInteraction) SetAssistantActions(v []ApiThreadAssistantAction)`

SetAssistantActions sets AssistantActions field to given value.

### HasAssistantActions

`func (o *ApiThreadInteraction) HasAssistantActions() bool`

HasAssistantActions returns a boolean if a field has been set.

### GetError

`func (o *ApiThreadInteraction) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiThreadInteraction) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiThreadInteraction) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiThreadInteraction) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ApiThreadInteraction) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ApiThreadInteraction) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetCompletionTimestamp

`func (o *ApiThreadInteraction) GetCompletionTimestamp() time.Time`

GetCompletionTimestamp returns the CompletionTimestamp field if non-nil, zero value otherwise.

### GetCompletionTimestampOk

`func (o *ApiThreadInteraction) GetCompletionTimestampOk() (*time.Time, bool)`

GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTimestamp

`func (o *ApiThreadInteraction) SetCompletionTimestamp(v time.Time)`

SetCompletionTimestamp sets CompletionTimestamp field to given value.

### HasCompletionTimestamp

`func (o *ApiThreadInteraction) HasCompletionTimestamp() bool`

HasCompletionTimestamp returns a boolean if a field has been set.

### SetCompletionTimestampNil

`func (o *ApiThreadInteraction) SetCompletionTimestampNil(b bool)`

 SetCompletionTimestampNil sets the value for CompletionTimestamp to be an explicit nil

### UnsetCompletionTimestamp
`func (o *ApiThreadInteraction) UnsetCompletionTimestamp()`

UnsetCompletionTimestamp ensures that no value is present for CompletionTimestamp, not even an explicit nil
### GetCanceledTimestamp

`func (o *ApiThreadInteraction) GetCanceledTimestamp() time.Time`

GetCanceledTimestamp returns the CanceledTimestamp field if non-nil, zero value otherwise.

### GetCanceledTimestampOk

`func (o *ApiThreadInteraction) GetCanceledTimestampOk() (*time.Time, bool)`

GetCanceledTimestampOk returns a tuple with the CanceledTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledTimestamp

`func (o *ApiThreadInteraction) SetCanceledTimestamp(v time.Time)`

SetCanceledTimestamp sets CanceledTimestamp field to given value.

### HasCanceledTimestamp

`func (o *ApiThreadInteraction) HasCanceledTimestamp() bool`

HasCanceledTimestamp returns a boolean if a field has been set.

### SetCanceledTimestampNil

`func (o *ApiThreadInteraction) SetCanceledTimestampNil(b bool)`

 SetCanceledTimestampNil sets the value for CanceledTimestamp to be an explicit nil

### UnsetCanceledTimestamp
`func (o *ApiThreadInteraction) UnsetCanceledTimestamp()`

UnsetCanceledTimestamp ensures that no value is present for CanceledTimestamp, not even an explicit nil
### GetErrorTimestamp

`func (o *ApiThreadInteraction) GetErrorTimestamp() time.Time`

GetErrorTimestamp returns the ErrorTimestamp field if non-nil, zero value otherwise.

### GetErrorTimestampOk

`func (o *ApiThreadInteraction) GetErrorTimestampOk() (*time.Time, bool)`

GetErrorTimestampOk returns a tuple with the ErrorTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTimestamp

`func (o *ApiThreadInteraction) SetErrorTimestamp(v time.Time)`

SetErrorTimestamp sets ErrorTimestamp field to given value.

### HasErrorTimestamp

`func (o *ApiThreadInteraction) HasErrorTimestamp() bool`

HasErrorTimestamp returns a boolean if a field has been set.

### SetErrorTimestampNil

`func (o *ApiThreadInteraction) SetErrorTimestampNil(b bool)`

 SetErrorTimestampNil sets the value for ErrorTimestamp to be an explicit nil

### UnsetErrorTimestamp
`func (o *ApiThreadInteraction) UnsetErrorTimestamp()`

UnsetErrorTimestamp ensures that no value is present for ErrorTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


