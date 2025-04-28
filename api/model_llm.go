package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

// LlmProviderType represents a LLM provider type enum.
type LlmProviderType string

const (
	LlmProviderAnthropic = "anthropic"
	LlmProviderHasura    = "hasura"
	LlmProviderOpenAI    = "openai"
)

var enumValuesLlmProviderType = []LlmProviderType{
	LlmProviderAnthropic,
	LlmProviderHasura,
	LlmProviderOpenAI,
}

// ParseLlmProviderType parses the LLM provider type enum from string.
func ParseLlmProviderType(input string) (LlmProviderType, error) {
	result := LlmProviderType(input)

	if !slices.Contains(enumValuesLlmProviderType, result) {
		return "", fmt.Errorf(
			"invalid LLM provider type, expected one of %v, got %s",
			enumValuesLlmProviderType,
			input,
		)
	}

	return result, nil
}

// Unmarshal JSON data into one of the pointers in the enum.
func (dst *LlmProviderType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	result, err := ParseLlmProviderType(raw)
	if err != nil {
		return err
	}

	*dst = result

	return nil
}

// LlmInterface abstracts the LLM model as an interface.
type LlmInterface interface {
	GetProvider() string
}

// Llm - struct for Llm.
type Llm struct {
	inner LlmInterface
}

// NewLlm creates an Llm instance.
func NewLlm[L LlmInterface](inner L) Llm {
	return Llm{inner: inner}
}

// ApiAnthropicConfigAsLlm is a convenience function that returns ApiAnthropicConfig wrapped in Llm.
func ApiAnthropicConfigAsLlm(v *ApiAnthropicConfig) Llm {
	return Llm{
		inner: v,
	}
}

// ApiOpenAIConfigAsLlm is a convenience function that returns ApiOpenAIConfig wrapped in Llm.
func ApiOpenAIConfigAsLlm(v *ApiOpenAIConfig) Llm {
	return Llm{
		inner: v,
	}
}

// HasuraLlmConfigAsLlm is a convenience function that returns HasuraLlmConfig wrapped in Llm.
func HasuraLlmConfigAsLlm(v *HasuraLlmConfig) Llm {
	return Llm{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *Llm) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) == 0 {
		return errors.New("unable to parse the Llm from empty object")
	}

	rawArtifactType, ok := raw["provider"]
	if !ok {
		return errors.New("unable to parse the Llm: field 'provider' is required")
	}

	var providerType LlmProviderType

	if err := json.Unmarshal(rawArtifactType, &providerType); err != nil {
		return fmt.Errorf("unable to parse the Llm: %w", err)
	}

	switch providerType {
	case LlmProviderHasura:
		result := HasuraLlmConfig{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case LlmProviderOpenAI:
		result := ApiOpenAIConfig{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case LlmProviderAnthropic:
		result := ApiAnthropicConfig{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	default:
		return fmt.Errorf("unsupported LLM provider type: %s", providerType)
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src Llm) MarshalJSON() ([]byte, error) {
	if src.inner != nil {
		return json.Marshal(src.inner)
	}

	return nil, nil // no data in oneOf schemas
}

// IsNil checks if the instance is empty.
func (src Llm) IsNil() bool {
	return src.inner == nil
}

// Interface returns the inner interface.
func (src Llm) Interface() LlmInterface {
	return src.inner
}

// String prints the instance to string.
func (src Llm) String() string {
	return fmt.Sprint(src.inner)
}

// AsHasura returns the instance as a nullable HasuraLlmConfig.
func (src Llm) AsHasura() *HasuraLlmConfig {
	result, ok := src.inner.(*HasuraLlmConfig)
	if !ok {
		return nil
	}

	return result
}

// AsOpenAI returns the instance as a nullable ApiOpenAIConfig.
func (src Llm) AsOpenAI() *ApiOpenAIConfig {
	result, ok := src.inner.(*ApiOpenAIConfig)
	if !ok {
		return nil
	}

	return result
}

// AsAnthropic returns the instance as a nullable ApiAnthropicConfig.
func (src Llm) AsAnthropic() *ApiAnthropicConfig {
	result, ok := src.inner.(*ApiAnthropicConfig)
	if !ok {
		return nil
	}

	return result
}
