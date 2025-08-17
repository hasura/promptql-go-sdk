# ExecuteRequestV1ArtifactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Title** | **string** |  | 
**ArtifactType** | **string** |  | 
**Data** | [**AutomationArtifactData**](AutomationArtifactData.md) | Automation data containing Python code and optional sample input | 

## Methods

### NewExecuteRequestV1ArtifactsInner

`func NewExecuteRequestV1ArtifactsInner(identifier string, title string, artifactType string, data AutomationArtifactData, ) *ExecuteRequestV1ArtifactsInner`

NewExecuteRequestV1ArtifactsInner instantiates a new ExecuteRequestV1ArtifactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRequestV1ArtifactsInnerWithDefaults

`func NewExecuteRequestV1ArtifactsInnerWithDefaults() *ExecuteRequestV1ArtifactsInner`

NewExecuteRequestV1ArtifactsInnerWithDefaults instantiates a new ExecuteRequestV1ArtifactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ExecuteRequestV1ArtifactsInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ExecuteRequestV1ArtifactsInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ExecuteRequestV1ArtifactsInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTitle

`func (o *ExecuteRequestV1ArtifactsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExecuteRequestV1ArtifactsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExecuteRequestV1ArtifactsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetArtifactType

`func (o *ExecuteRequestV1ArtifactsInner) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ExecuteRequestV1ArtifactsInner) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ExecuteRequestV1ArtifactsInner) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetData

`func (o *ExecuteRequestV1ArtifactsInner) GetData() AutomationArtifactData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecuteRequestV1ArtifactsInner) GetDataOk() (*AutomationArtifactData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecuteRequestV1ArtifactsInner) SetData(v AutomationArtifactData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


