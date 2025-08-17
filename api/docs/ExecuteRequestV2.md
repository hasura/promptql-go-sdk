# ExecuteRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Code** | **string** |  | 
**Ddn** | Pointer to [**NullableDdnConfigV2**](DdnConfigV2.md) |  | [optional] 
**Artifacts** | Pointer to [**[]ExecuteRequestV1ArtifactsInner**](ExecuteRequestV1ArtifactsInner.md) |  | [optional] 

## Methods

### NewExecuteRequestV2

`func NewExecuteRequestV2(version string, code string, ) *ExecuteRequestV2`

NewExecuteRequestV2 instantiates a new ExecuteRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRequestV2WithDefaults

`func NewExecuteRequestV2WithDefaults() *ExecuteRequestV2`

NewExecuteRequestV2WithDefaults instantiates a new ExecuteRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ExecuteRequestV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExecuteRequestV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExecuteRequestV2) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCode

`func (o *ExecuteRequestV2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExecuteRequestV2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExecuteRequestV2) SetCode(v string)`

SetCode sets Code field to given value.


### GetDdn

`func (o *ExecuteRequestV2) GetDdn() DdnConfigV2`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *ExecuteRequestV2) GetDdnOk() (*DdnConfigV2, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *ExecuteRequestV2) SetDdn(v DdnConfigV2)`

SetDdn sets Ddn field to given value.

### HasDdn

`func (o *ExecuteRequestV2) HasDdn() bool`

HasDdn returns a boolean if a field has been set.

### SetDdnNil

`func (o *ExecuteRequestV2) SetDdnNil(b bool)`

 SetDdnNil sets the value for Ddn to be an explicit nil

### UnsetDdn
`func (o *ExecuteRequestV2) UnsetDdn()`

UnsetDdn ensures that no value is present for Ddn, not even an explicit nil
### GetArtifacts

`func (o *ExecuteRequestV2) GetArtifacts() []ExecuteRequestV1ArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ExecuteRequestV2) GetArtifactsOk() (*[]ExecuteRequestV1ArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ExecuteRequestV2) SetArtifacts(v []ExecuteRequestV1ArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ExecuteRequestV2) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


