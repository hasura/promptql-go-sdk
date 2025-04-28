# ApiAnthropicConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Model** | Pointer to **NullableString** |  | [optional] 
**BaseUrl** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | **string** |  | 

## Methods

### NewApiAnthropicConfig

`func NewApiAnthropicConfig(provider string, apiKey string, ) *ApiAnthropicConfig`

NewApiAnthropicConfig instantiates a new ApiAnthropicConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAnthropicConfigWithDefaults

`func NewApiAnthropicConfigWithDefaults() *ApiAnthropicConfig`

NewApiAnthropicConfigWithDefaults instantiates a new ApiAnthropicConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ApiAnthropicConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiAnthropicConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiAnthropicConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *ApiAnthropicConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApiAnthropicConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApiAnthropicConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApiAnthropicConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ApiAnthropicConfig) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ApiAnthropicConfig) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetBaseUrl

`func (o *ApiAnthropicConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ApiAnthropicConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ApiAnthropicConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ApiAnthropicConfig) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *ApiAnthropicConfig) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *ApiAnthropicConfig) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetApiKey

`func (o *ApiAnthropicConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiAnthropicConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiAnthropicConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


