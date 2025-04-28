# ApiOpenAIConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Model** | Pointer to **NullableString** |  | [optional] 
**BaseUrl** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | **string** |  | 

## Methods

### NewApiOpenAIConfig

`func NewApiOpenAIConfig(provider string, apiKey string, ) *ApiOpenAIConfig`

NewApiOpenAIConfig instantiates a new ApiOpenAIConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOpenAIConfigWithDefaults

`func NewApiOpenAIConfigWithDefaults() *ApiOpenAIConfig`

NewApiOpenAIConfigWithDefaults instantiates a new ApiOpenAIConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ApiOpenAIConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiOpenAIConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiOpenAIConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *ApiOpenAIConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApiOpenAIConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApiOpenAIConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApiOpenAIConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ApiOpenAIConfig) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ApiOpenAIConfig) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetBaseUrl

`func (o *ApiOpenAIConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ApiOpenAIConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ApiOpenAIConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ApiOpenAIConfig) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *ApiOpenAIConfig) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *ApiOpenAIConfig) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetApiKey

`func (o *ApiOpenAIConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiOpenAIConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiOpenAIConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


