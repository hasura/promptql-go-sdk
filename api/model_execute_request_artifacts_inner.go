package api

import "encoding/json"

// ExecuteRequestArtifactsInner - struct for ExecuteRequestArtifactsInner.
type ExecuteRequestArtifactsInner Artifact

// TableArtifactAsExecuteRequestArtifactsInner is a convenience function that returns TableArtifact wrapped in ExecuteRequestArtifactsInner.
func TableArtifactAsExecuteRequestArtifactsInner(v *TableArtifact) ExecuteRequestArtifactsInner {
	return ExecuteRequestArtifactsInner{
		inner: v,
	}
}

// TextArtifactAsExecuteRequestArtifactsInner is a convenience function that returns TextArtifact wrapped in ExecuteRequestArtifactsInner.
func TextArtifactAsExecuteRequestArtifactsInner(v *TextArtifact) ExecuteRequestArtifactsInner {
	return ExecuteRequestArtifactsInner{
		inner: v,
	}
}

// VisualizationArtifactAsExecuteRequestArtifactsInner is a convenience function that returns VisualizationArtifact wrapped in ExecuteRequestArtifactsInner.
func VisualizationArtifactAsExecuteRequestArtifactsInner(
	v *VisualizationArtifact,
) ExecuteRequestArtifactsInner {
	return ExecuteRequestArtifactsInner{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *ExecuteRequestArtifactsInner) UnmarshalJSON(data []byte) error {
	var result Artifact

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	*dst = ExecuteRequestArtifactsInner(result)

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src ExecuteRequestArtifactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(Artifact(src))
}

// IsNil checks if the instance is empty.
func (src ExecuteRequestArtifactsInner) IsNil() bool {
	return src.inner == nil
}

// Interface returns the inner interface.
func (src ExecuteRequestArtifactsInner) Interface() ArtifactInterface {
	return src.inner
}

// AsText returns the instance as a nullable TextArtifact.
func (src ExecuteRequestArtifactsInner) AsText() *TextArtifact {
	return Artifact(src).AsText()
}

// AsText converts the instance to a nullable TableArtifact.
func (src ExecuteRequestArtifactsInner) AsTable() *TableArtifact {
	return Artifact(src).AsTable()
}

// AsText converts the instance to a nullable VisualizationArtifact.
func (src ExecuteRequestArtifactsInner) AsVisualization() *VisualizationArtifact {
	return Artifact(src).AsVisualization()
}

type NullableExecuteRequestArtifactsInner struct {
	value *ExecuteRequestArtifactsInner
	isSet bool
}

func (v NullableExecuteRequestArtifactsInner) Get() *ExecuteRequestArtifactsInner {
	return v.value
}

func (v *NullableExecuteRequestArtifactsInner) Set(val *ExecuteRequestArtifactsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteRequestArtifactsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteRequestArtifactsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteRequestArtifactsInner(
	val *ExecuteRequestArtifactsInner,
) *NullableExecuteRequestArtifactsInner {
	return &NullableExecuteRequestArtifactsInner{value: val, isSet: true}
}

func (v NullableExecuteRequestArtifactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteRequestArtifactsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true

	return json.Unmarshal(src, &v.value)
}
