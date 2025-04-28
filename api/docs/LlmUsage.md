# LlmUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] [default to "unknown"]
**Model** | Pointer to **string** |  | [optional] [default to "unknown"]
**InputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**OutputTokens** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewLlmUsage

`func NewLlmUsage() *LlmUsage`

NewLlmUsage instantiates a new LlmUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmUsageWithDefaults

`func NewLlmUsageWithDefaults() *LlmUsage`

NewLlmUsageWithDefaults instantiates a new LlmUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *LlmUsage) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LlmUsage) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LlmUsage) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LlmUsage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetModel

`func (o *LlmUsage) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *LlmUsage) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *LlmUsage) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *LlmUsage) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInputTokens

`func (o *LlmUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *LlmUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *LlmUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *LlmUsage) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *LlmUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *LlmUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *LlmUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *LlmUsage) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


