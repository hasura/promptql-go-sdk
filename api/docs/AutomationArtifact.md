# AutomationArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Title** | **string** |  | 
**ArtifactType** | **string** |  | 
**Data** | [**AutomationArtifactData**](AutomationArtifactData.md) | Automation data containing Python code and optional sample input | 

## Methods

### NewAutomationArtifact

`func NewAutomationArtifact(identifier string, title string, artifactType string, data AutomationArtifactData, ) *AutomationArtifact`

NewAutomationArtifact instantiates a new AutomationArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationArtifactWithDefaults

`func NewAutomationArtifactWithDefaults() *AutomationArtifact`

NewAutomationArtifactWithDefaults instantiates a new AutomationArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *AutomationArtifact) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AutomationArtifact) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AutomationArtifact) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTitle

`func (o *AutomationArtifact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutomationArtifact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutomationArtifact) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetArtifactType

`func (o *AutomationArtifact) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *AutomationArtifact) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *AutomationArtifact) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetData

`func (o *AutomationArtifact) GetData() AutomationArtifactData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutomationArtifact) GetDataOk() (*AutomationArtifactData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutomationArtifact) SetData(v AutomationArtifactData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


