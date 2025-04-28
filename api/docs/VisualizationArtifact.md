# VisualizationArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Title** | **string** |  | 
**ArtifactType** | **string** |  | 
**Data** | [**VisualizationArtifactData**](VisualizationArtifactData.md) | Visualization data containing HTML and visualization data | 

## Methods

### NewVisualizationArtifact

`func NewVisualizationArtifact(identifier string, title string, artifactType string, data VisualizationArtifactData, ) *VisualizationArtifact`

NewVisualizationArtifact instantiates a new VisualizationArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationArtifactWithDefaults

`func NewVisualizationArtifactWithDefaults() *VisualizationArtifact`

NewVisualizationArtifactWithDefaults instantiates a new VisualizationArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VisualizationArtifact) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VisualizationArtifact) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VisualizationArtifact) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTitle

`func (o *VisualizationArtifact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VisualizationArtifact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VisualizationArtifact) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetArtifactType

`func (o *VisualizationArtifact) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *VisualizationArtifact) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *VisualizationArtifact) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetData

`func (o *VisualizationArtifact) GetData() VisualizationArtifactData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VisualizationArtifact) GetDataOk() (*VisualizationArtifactData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VisualizationArtifact) SetData(v VisualizationArtifactData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


