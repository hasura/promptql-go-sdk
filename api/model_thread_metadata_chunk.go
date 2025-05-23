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

// checks if the ThreadMetadataChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadMetadataChunk{}

// ThreadMetadataChunk struct for ThreadMetadataChunk
type ThreadMetadataChunk struct {
	Type string `json:"type"`
	ThreadId string `json:"thread_id"`
}

type _ThreadMetadataChunk ThreadMetadataChunk

// NewThreadMetadataChunk instantiates a new ThreadMetadataChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadMetadataChunk(type_ string, threadId string) *ThreadMetadataChunk {
	this := ThreadMetadataChunk{}
	this.Type = type_
	this.ThreadId = threadId
	return &this
}

// NewThreadMetadataChunkWithDefaults instantiates a new ThreadMetadataChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadMetadataChunkWithDefaults() *ThreadMetadataChunk {
	this := ThreadMetadataChunk{}
	return &this
}

// GetType returns the Type field value
func (o *ThreadMetadataChunk) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ThreadMetadataChunk) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ThreadMetadataChunk) SetType(v string) {
	o.Type = v
}

// GetThreadId returns the ThreadId field value
func (o *ThreadMetadataChunk) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *ThreadMetadataChunk) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *ThreadMetadataChunk) SetThreadId(v string) {
	o.ThreadId = v
}

func (o ThreadMetadataChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadMetadataChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["thread_id"] = o.ThreadId
	return toSerialize, nil
}

func (o *ThreadMetadataChunk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"thread_id",
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

	varThreadMetadataChunk := _ThreadMetadataChunk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varThreadMetadataChunk)

	if err != nil {
		return err
	}

	*o = ThreadMetadataChunk(varThreadMetadataChunk)

	return err
}

type NullableThreadMetadataChunk struct {
	value *ThreadMetadataChunk
	isSet bool
}

func (v NullableThreadMetadataChunk) Get() *ThreadMetadataChunk {
	return v.value
}

func (v *NullableThreadMetadataChunk) Set(val *ThreadMetadataChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadMetadataChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadMetadataChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadMetadataChunk(val *ThreadMetadataChunk) *NullableThreadMetadataChunk {
	return &NullableThreadMetadataChunk{value: val, isSet: true}
}

func (v NullableThreadMetadataChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadMetadataChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

