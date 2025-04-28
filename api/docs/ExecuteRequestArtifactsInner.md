# ExecuteRequestArtifactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Title** | **string** |  | 
**ArtifactType** | **string** |  | 
**Data** | [**VisualizationArtifactData**](VisualizationArtifactData.md) | Visualization data containing HTML and visualization data | 

## Methods

### NewExecuteRequestArtifactsInner

`func NewExecuteRequestArtifactsInner(identifier string, title string, artifactType string, data VisualizationArtifactData, ) *ExecuteRequestArtifactsInner`

NewExecuteRequestArtifactsInner instantiates a new ExecuteRequestArtifactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRequestArtifactsInnerWithDefaults

`func NewExecuteRequestArtifactsInnerWithDefaults() *ExecuteRequestArtifactsInner`

NewExecuteRequestArtifactsInnerWithDefaults instantiates a new ExecuteRequestArtifactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ExecuteRequestArtifactsInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ExecuteRequestArtifactsInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ExecuteRequestArtifactsInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTitle

`func (o *ExecuteRequestArtifactsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExecuteRequestArtifactsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExecuteRequestArtifactsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetArtifactType

`func (o *ExecuteRequestArtifactsInner) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ExecuteRequestArtifactsInner) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ExecuteRequestArtifactsInner) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetData

`func (o *ExecuteRequestArtifactsInner) GetData() VisualizationArtifactData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecuteRequestArtifactsInner) GetDataOk() (*VisualizationArtifactData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecuteRequestArtifactsInner) SetData(v VisualizationArtifactData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


