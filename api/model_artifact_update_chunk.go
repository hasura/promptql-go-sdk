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

// checks if the ArtifactUpdateChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactUpdateChunk{}

// ArtifactUpdateChunk struct for ArtifactUpdateChunk
type ArtifactUpdateChunk struct {
	Type string `json:"type"`
	Artifact Artifact `json:"artifact"`
}

type _ArtifactUpdateChunk ArtifactUpdateChunk

// NewArtifactUpdateChunk instantiates a new ArtifactUpdateChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactUpdateChunk(type_ string, artifact Artifact) *ArtifactUpdateChunk {
	this := ArtifactUpdateChunk{}
	this.Type = type_
	this.Artifact = artifact
	return &this
}

// NewArtifactUpdateChunkWithDefaults instantiates a new ArtifactUpdateChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactUpdateChunkWithDefaults() *ArtifactUpdateChunk {
	this := ArtifactUpdateChunk{}
	return &this
}

// GetType returns the Type field value
func (o *ArtifactUpdateChunk) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArtifactUpdateChunk) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArtifactUpdateChunk) SetType(v string) {
	o.Type = v
}

// GetArtifact returns the Artifact field value
func (o *ArtifactUpdateChunk) GetArtifact() Artifact {
	if o == nil {
		var ret Artifact
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *ArtifactUpdateChunk) GetArtifactOk() (*Artifact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *ArtifactUpdateChunk) SetArtifact(v Artifact) {
	o.Artifact = v
}

func (o ArtifactUpdateChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactUpdateChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["artifact"] = o.Artifact
	return toSerialize, nil
}

func (o *ArtifactUpdateChunk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"artifact",
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

	varArtifactUpdateChunk := _ArtifactUpdateChunk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varArtifactUpdateChunk)

	if err != nil {
		return err
	}

	*o = ArtifactUpdateChunk(varArtifactUpdateChunk)

	return err
}

type NullableArtifactUpdateChunk struct {
	value *ArtifactUpdateChunk
	isSet bool
}

func (v NullableArtifactUpdateChunk) Get() *ArtifactUpdateChunk {
	return v.value
}

func (v *NullableArtifactUpdateChunk) Set(val *ArtifactUpdateChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactUpdateChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactUpdateChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactUpdateChunk(val *ArtifactUpdateChunk) *NullableArtifactUpdateChunk {
	return &NullableArtifactUpdateChunk{value: val, isSet: true}
}

func (v NullableArtifactUpdateChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactUpdateChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

