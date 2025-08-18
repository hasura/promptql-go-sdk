# DdnConfigV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildId** | Pointer to **NullableString** |  | [optional] 
**BuildVersion** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP headers that should be forwarded to DDN | [optional] 

## Methods

### NewDdnConfigV2

`func NewDdnConfigV2() *DdnConfigV2`

NewDdnConfigV2 instantiates a new DdnConfigV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdnConfigV2WithDefaults

`func NewDdnConfigV2WithDefaults() *DdnConfigV2`

NewDdnConfigV2WithDefaults instantiates a new DdnConfigV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildId

`func (o *DdnConfigV2) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *DdnConfigV2) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *DdnConfigV2) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *DdnConfigV2) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *DdnConfigV2) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *DdnConfigV2) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetBuildVersion

`func (o *DdnConfigV2) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *DdnConfigV2) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *DdnConfigV2) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *DdnConfigV2) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### SetBuildVersionNil

`func (o *DdnConfigV2) SetBuildVersionNil(b bool)`

 SetBuildVersionNil sets the value for BuildVersion to be an explicit nil

### UnsetBuildVersion
`func (o *DdnConfigV2) UnsetBuildVersion()`

UnsetBuildVersion ensures that no value is present for BuildVersion, not even an explicit nil
### GetHeaders

`func (o *DdnConfigV2) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DdnConfigV2) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DdnConfigV2) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DdnConfigV2) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


