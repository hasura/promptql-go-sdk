package api

import (
	"encoding/json"
)

// ExecuteRequestV1AiPrimitivesLlm - struct for ExecuteRequestV1AiPrimitivesLlm.
type ExecuteRequestV1AiPrimitivesLlm Llm

// NewExecuteRequestV1AiPrimitivesLlm creates a generic ExecuteRequestAiPrimitivesLlm instance.
func NewExecuteRequestV1AiPrimitivesLlm(inner LlmInterface) ExecuteRequestV1AiPrimitivesLlm {
	return ExecuteRequestV1AiPrimitivesLlm{
		inner: inner,
	}
}

// ApiAnthropicConfigAsExecuteRequestV1AiPrimitivesLlm is a convenience function that returns ApiAnthropicConfig wrapped in ExecuteRequestV1AiPrimitivesLlm.
func ApiAnthropicConfigAsExecuteRequestV1AiPrimitivesLlm(
	v *ApiAnthropicConfig,
) ExecuteRequestV1AiPrimitivesLlm {
	return ExecuteRequestV1AiPrimitivesLlm{
		inner: v,
	}
}

// ApiOpenAIConfigAsExecuteRequestV1AiPrimitivesLlm is a convenience function that returns ApiOpenAIConfig wrapped in ExecuteRequestV1AiPrimitivesLlm.
func ApiOpenAIConfigAsExecuteRequestV1AiPrimitivesLlm(
	v *ApiOpenAIConfig,
) ExecuteRequestV1AiPrimitivesLlm {
	return ExecuteRequestV1AiPrimitivesLlm{
		inner: v,
	}
}

// HasuraLlmConfigV1AsExecuteRequestV1AiPrimitivesLlm is a convenience function that returns HasuraLlmConfig wrapped in ExecuteRequestV1AiPrimitivesLlm.
func HasuraLlmConfigV1AsExecuteRequestV1AiPrimitivesLlm(
	v *HasuraLlmConfig,
) ExecuteRequestV1AiPrimitivesLlm {
	return ExecuteRequestV1AiPrimitivesLlm{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *ExecuteRequestV1AiPrimitivesLlm) UnmarshalJSON(data []byte) error {
	var result Llm

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	*dst = ExecuteRequestV1AiPrimitivesLlm(result)

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src ExecuteRequestV1AiPrimitivesLlm) MarshalJSON() ([]byte, error) {
	return json.Marshal(Llm(src))
}

// IsNil checks if the instance is empty.
func (src ExecuteRequestV1AiPrimitivesLlm) IsNil() bool {
	return src.inner == nil
}

type NullableExecuteRequestV1AiPrimitivesLlm struct {
	value *ExecuteRequestV1AiPrimitivesLlm
	isSet bool
}

func (v NullableExecuteRequestV1AiPrimitivesLlm) Get() *ExecuteRequestV1AiPrimitivesLlm {
	return v.value
}

func (v *NullableExecuteRequestV1AiPrimitivesLlm) Set(val *ExecuteRequestV1AiPrimitivesLlm) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteRequestV1AiPrimitivesLlm) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteRequestV1AiPrimitivesLlm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteRequestV1AiPrimitivesLlm(
	val *ExecuteRequestV1AiPrimitivesLlm,
) *NullableExecuteRequestV1AiPrimitivesLlm {
	return &NullableExecuteRequestV1AiPrimitivesLlm{value: val, isSet: true}
}

func (v NullableExecuteRequestV1AiPrimitivesLlm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteRequestV1AiPrimitivesLlm) UnmarshalJSON(src []byte) error {
	v.isSet = true

	return json.Unmarshal(src, &v.value)
}
