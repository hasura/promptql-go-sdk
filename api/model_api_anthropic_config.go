/*
PromptQL API

PromptQL API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiAnthropicConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAnthropicConfig{}

// ApiAnthropicConfig struct for ApiAnthropicConfig
type ApiAnthropicConfig struct {
	Provider string `json:"provider"`
	Model NullableString `json:"model,omitempty"`
	BaseUrl NullableString `json:"base_url,omitempty"`
	ApiKey string `json:"api_key"`
}

type _ApiAnthropicConfig ApiAnthropicConfig

// NewApiAnthropicConfig instantiates a new ApiAnthropicConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAnthropicConfig(provider string, apiKey string) *ApiAnthropicConfig {
	this := ApiAnthropicConfig{}
	this.Provider = provider
	this.ApiKey = apiKey
	return &this
}

// NewApiAnthropicConfigWithDefaults instantiates a new ApiAnthropicConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAnthropicConfigWithDefaults() *ApiAnthropicConfig {
	this := ApiAnthropicConfig{}
	return &this
}

// GetProvider returns the Provider field value
func (o *ApiAnthropicConfig) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ApiAnthropicConfig) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ApiAnthropicConfig) SetProvider(v string) {
	o.Provider = v
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAnthropicConfig) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAnthropicConfig) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *ApiAnthropicConfig) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *ApiAnthropicConfig) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *ApiAnthropicConfig) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *ApiAnthropicConfig) UnsetModel() {
	o.Model.Unset()
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAnthropicConfig) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BaseUrl.Get()
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAnthropicConfig) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseUrl.Get(), o.BaseUrl.IsSet()
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ApiAnthropicConfig) HasBaseUrl() bool {
	if o != nil && o.BaseUrl.IsSet() {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given NullableString and assigns it to the BaseUrl field.
func (o *ApiAnthropicConfig) SetBaseUrl(v string) {
	o.BaseUrl.Set(&v)
}
// SetBaseUrlNil sets the value for BaseUrl to be an explicit nil
func (o *ApiAnthropicConfig) SetBaseUrlNil() {
	o.BaseUrl.Set(nil)
}

// UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
func (o *ApiAnthropicConfig) UnsetBaseUrl() {
	o.BaseUrl.Unset()
}

// GetApiKey returns the ApiKey field value
func (o *ApiAnthropicConfig) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *ApiAnthropicConfig) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *ApiAnthropicConfig) SetApiKey(v string) {
	o.ApiKey = v
}

func (o ApiAnthropicConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAnthropicConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.BaseUrl.IsSet() {
		toSerialize["base_url"] = o.BaseUrl.Get()
	}
	toSerialize["api_key"] = o.ApiKey
	return toSerialize, nil
}

func (o *ApiAnthropicConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"api_key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiAnthropicConfig := _ApiAnthropicConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApiAnthropicConfig)

	if err != nil {
		return err
	}

	*o = ApiAnthropicConfig(varApiAnthropicConfig)

	return err
}

type NullableApiAnthropicConfig struct {
	value *ApiAnthropicConfig
	isSet bool
}

func (v NullableApiAnthropicConfig) Get() *ApiAnthropicConfig {
	return v.value
}

func (v *NullableApiAnthropicConfig) Set(val *ApiAnthropicConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAnthropicConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAnthropicConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAnthropicConfig(val *ApiAnthropicConfig) *NullableApiAnthropicConfig {
	return &NullableApiAnthropicConfig{value: val, isSet: true}
}

func (v NullableApiAnthropicConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAnthropicConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

