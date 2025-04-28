# ArtifactUpdateChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Artifact** | [**Artifact**](Artifact.md) |  | 

## Methods

### NewArtifactUpdateChunk

`func NewArtifactUpdateChunk(type_ string, artifact Artifact, ) *ArtifactUpdateChunk`

NewArtifactUpdateChunk instantiates a new ArtifactUpdateChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactUpdateChunkWithDefaults

`func NewArtifactUpdateChunkWithDefaults() *ArtifactUpdateChunk`

NewArtifactUpdateChunkWithDefaults instantiates a new ArtifactUpdateChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArtifactUpdateChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactUpdateChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactUpdateChunk) SetType(v string)`

SetType sets Type field to given value.


### GetArtifact

`func (o *ArtifactUpdateChunk) GetArtifact() Artifact`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *ArtifactUpdateChunk) GetArtifactOk() (*Artifact, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *ArtifactUpdateChunk) SetArtifact(v Artifact)`

SetArtifact sets Artifact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


