# TableArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Title** | **string** |  | 
**ArtifactType** | **string** |  | 
**Data** | **[]map[string]interface{}** | Table data as a list of objects, with the object keys being column names | 

## Methods

### NewTableArtifact

`func NewTableArtifact(identifier string, title string, artifactType string, data []map[string]interface{}, ) *TableArtifact`

NewTableArtifact instantiates a new TableArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableArtifactWithDefaults

`func NewTableArtifactWithDefaults() *TableArtifact`

NewTableArtifactWithDefaults instantiates a new TableArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *TableArtifact) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TableArtifact) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TableArtifact) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTitle

`func (o *TableArtifact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TableArtifact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TableArtifact) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetArtifactType

`func (o *TableArtifact) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *TableArtifact) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *TableArtifact) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetData

`func (o *TableArtifact) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TableArtifact) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TableArtifact) SetData(v []map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


