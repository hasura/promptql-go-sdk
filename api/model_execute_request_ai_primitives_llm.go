package api

import (
	"encoding/json"
)

// ExecuteRequestAiPrimitivesLlm - struct for ExecuteRequestAiPrimitivesLlm.
type ExecuteRequestAiPrimitivesLlm Llm

// NewExecuteRequestAiPrimitivesLlm creates a generic ExecuteRequestAiPrimitivesLlm instance.
func NewExecuteRequestAiPrimitivesLlm(inner LlmInterface) ExecuteRequestAiPrimitivesLlm {
	return ExecuteRequestAiPrimitivesLlm{
		inner: inner,
	}
}

// ApiAnthropicConfigAsExecuteRequestAiPrimitivesLlm is a convenience function that returns ApiAnthropicConfig wrapped in ExecuteRequestAiPrimitivesLlm.
func ApiAnthropicConfigAsExecuteRequestAiPrimitivesLlm(
	v *ApiAnthropicConfig,
) ExecuteRequestAiPrimitivesLlm {
	return ExecuteRequestAiPrimitivesLlm{
		inner: v,
	}
}

// ApiOpenAIConfigAsExecuteRequestAiPrimitivesLlm is a convenience function that returns ApiOpenAIConfig wrapped in ExecuteRequestAiPrimitivesLlm.
func ApiOpenAIConfigAsExecuteRequestAiPrimitivesLlm(
	v *ApiOpenAIConfig,
) ExecuteRequestAiPrimitivesLlm {
	return ExecuteRequestAiPrimitivesLlm{
		inner: v,
	}
}

// HasuraLlmConfigAsExecuteRequestAiPrimitivesLlm is a convenience function that returns HasuraLlmConfig wrapped in ExecuteRequestAiPrimitivesLlm.
func HasuraLlmConfigAsExecuteRequestAiPrimitivesLlm(
	v *HasuraLlmConfig,
) ExecuteRequestAiPrimitivesLlm {
	return ExecuteRequestAiPrimitivesLlm{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *ExecuteRequestAiPrimitivesLlm) UnmarshalJSON(data []byte) error {
	var result Llm

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	*dst = ExecuteRequestAiPrimitivesLlm(result)

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src ExecuteRequestAiPrimitivesLlm) MarshalJSON() ([]byte, error) {
	return json.Marshal(Llm(src))
}

// IsNil checks if the instance is empty.
func (src ExecuteRequestAiPrimitivesLlm) IsNil() bool {
	return src.inner == nil
}

type NullableExecuteRequestAiPrimitivesLlm struct {
	value *ExecuteRequestAiPrimitivesLlm
	isSet bool
}

func (v NullableExecuteRequestAiPrimitivesLlm) Get() *ExecuteRequestAiPrimitivesLlm {
	return v.value
}

func (v *NullableExecuteRequestAiPrimitivesLlm) Set(val *ExecuteRequestAiPrimitivesLlm) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteRequestAiPrimitivesLlm) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteRequestAiPrimitivesLlm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteRequestAiPrimitivesLlm(
	val *ExecuteRequestAiPrimitivesLlm,
) *NullableExecuteRequestAiPrimitivesLlm {
	return &NullableExecuteRequestAiPrimitivesLlm{value: val, isSet: true}
}

func (v NullableExecuteRequestAiPrimitivesLlm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteRequestAiPrimitivesLlm) UnmarshalJSON(src []byte) error {
	v.isSet = true

	return json.Unmarshal(src, &v.value)
}
