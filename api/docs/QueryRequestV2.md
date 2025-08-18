# QueryRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | **bool** |  | 
**Artifacts** | Pointer to [**[]ExecuteRequestV1ArtifactsInner**](ExecuteRequestV1ArtifactsInner.md) |  | [optional] 
**Timezone** | **string** | An IANA timezone used to interpret queries that implicitly require timezones | 
**Interactions** | [**[]ApiThreadInteraction**](ApiThreadInteraction.md) |  | 
**Version** | **string** |  | 
**Ddn** | [**DdnConfigV2**](DdnConfigV2.md) |  | 

## Methods

### NewQueryRequestV2

`func NewQueryRequestV2(stream bool, timezone string, interactions []ApiThreadInteraction, version string, ddn DdnConfigV2, ) *QueryRequestV2`

NewQueryRequestV2 instantiates a new QueryRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestV2WithDefaults

`func NewQueryRequestV2WithDefaults() *QueryRequestV2`

NewQueryRequestV2WithDefaults instantiates a new QueryRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *QueryRequestV2) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *QueryRequestV2) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *QueryRequestV2) SetStream(v bool)`

SetStream sets Stream field to given value.


### GetArtifacts

`func (o *QueryRequestV2) GetArtifacts() []ExecuteRequestV1ArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *QueryRequestV2) GetArtifactsOk() (*[]ExecuteRequestV1ArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *QueryRequestV2) SetArtifacts(v []ExecuteRequestV1ArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *QueryRequestV2) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetTimezone

`func (o *QueryRequestV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *QueryRequestV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *QueryRequestV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetInteractions

`func (o *QueryRequestV2) GetInteractions() []ApiThreadInteraction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *QueryRequestV2) GetInteractionsOk() (*[]ApiThreadInteraction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *QueryRequestV2) SetInteractions(v []ApiThreadInteraction)`

SetInteractions sets Interactions field to given value.


### GetVersion

`func (o *QueryRequestV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryRequestV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryRequestV2) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDdn

`func (o *QueryRequestV2) GetDdn() DdnConfigV2`

GetDdn returns the Ddn field if non-nil, zero value otherwise.

### GetDdnOk

`func (o *QueryRequestV2) GetDdnOk() (*DdnConfigV2, bool)`

GetDdnOk returns a tuple with the Ddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdn

`func (o *QueryRequestV2) SetDdn(v DdnConfigV2)`

SetDdn sets Ddn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


