# QueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantActions** | [**[]AssistantAction**](AssistantAction.md) |  | 
**ModifiedArtifacts** | [**[]ExecuteRequestArtifactsInner**](ExecuteRequestArtifactsInner.md) | List of artifacts created or updated in this request. May contain duplicate artifact identifiers. | 

## Methods

### NewQueryResponse

`func NewQueryResponse(assistantActions []AssistantAction, modifiedArtifacts []ExecuteRequestArtifactsInner, ) *QueryResponse`

NewQueryResponse instantiates a new QueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseWithDefaults

`func NewQueryResponseWithDefaults() *QueryResponse`

NewQueryResponseWithDefaults instantiates a new QueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantActions

`func (o *QueryResponse) GetAssistantActions() []AssistantAction`

GetAssistantActions returns the AssistantActions field if non-nil, zero value otherwise.

### GetAssistantActionsOk

`func (o *QueryResponse) GetAssistantActionsOk() (*[]AssistantAction, bool)`

GetAssistantActionsOk returns a tuple with the AssistantActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantActions

`func (o *QueryResponse) SetAssistantActions(v []AssistantAction)`

SetAssistantActions sets AssistantActions field to given value.


### GetModifiedArtifacts

`func (o *QueryResponse) GetModifiedArtifacts() []ExecuteRequestArtifactsInner`

GetModifiedArtifacts returns the ModifiedArtifacts field if non-nil, zero value otherwise.

### GetModifiedArtifactsOk

`func (o *QueryResponse) GetModifiedArtifactsOk() (*[]ExecuteRequestArtifactsInner, bool)`

GetModifiedArtifactsOk returns a tuple with the ModifiedArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedArtifacts

`func (o *QueryResponse) SetModifiedArtifacts(v []ExecuteRequestArtifactsInner)`

SetModifiedArtifacts sets ModifiedArtifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


