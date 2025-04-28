# Interaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMessage** | [**UserMessage**](UserMessage.md) |  | 
**AssistantActions** | Pointer to [**[]AssistantAction**](AssistantAction.md) |  | [optional] 

## Methods

### NewInteraction

`func NewInteraction(userMessage UserMessage, ) *Interaction`

NewInteraction instantiates a new Interaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionWithDefaults

`func NewInteractionWithDefaults() *Interaction`

NewInteractionWithDefaults instantiates a new Interaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMessage

`func (o *Interaction) GetUserMessage() UserMessage`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *Interaction) GetUserMessageOk() (*UserMessage, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *Interaction) SetUserMessage(v UserMessage)`

SetUserMessage sets UserMessage field to given value.


### GetAssistantActions

`func (o *Interaction) GetAssistantActions() []AssistantAction`

GetAssistantActions returns the AssistantActions field if non-nil, zero value otherwise.

### GetAssistantActionsOk

`func (o *Interaction) GetAssistantActionsOk() (*[]AssistantAction, bool)`

GetAssistantActionsOk returns a tuple with the AssistantActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantActions

`func (o *Interaction) SetAssistantActions(v []AssistantAction)`

SetAssistantActions sets AssistantActions field to given value.

### HasAssistantActions

`func (o *Interaction) HasAssistantActions() bool`

HasAssistantActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


