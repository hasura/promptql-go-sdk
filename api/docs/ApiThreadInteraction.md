# ApiThreadInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMessage** | [**ApiThreadUserMessage**](ApiThreadUserMessage.md) |  | 
**AssistantActions** | Pointer to [**[]ApiThreadAssistantAction**](ApiThreadAssistantAction.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


