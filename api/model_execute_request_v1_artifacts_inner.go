package api

import (
	"encoding/json"
)

// ExecuteRequestV1ArtifactsInner - struct for ExecuteRequestV1ArtifactsInner
type ExecuteRequestV1ArtifactsInner Artifact

// AutomationArtifactAsExecuteRequestV1ArtifactsInner is a convenience function that returns AutomationArtifact wrapped in ExecuteRequestV1ArtifactsInner
func AutomationArtifactAsExecuteRequestV1ArtifactsInner(v *AutomationArtifact) ExecuteRequestV1ArtifactsInner {
	return ExecuteRequestV1ArtifactsInner{
		inner: v,
	}
}

// TableArtifactAsExecuteRequestV1ArtifactsInner is a convenience function that returns TableArtifact wrapped in ExecuteRequestV1ArtifactsInner
func TableArtifactAsExecuteRequestV1ArtifactsInner(v *TableArtifact) ExecuteRequestV1ArtifactsInner {
	return ExecuteRequestV1ArtifactsInner{
		inner: v,
	}
}

// TextArtifactAsExecuteRequestV1ArtifactsInner is a convenience function that returns TextArtifact wrapped in ExecuteRequestV1ArtifactsInner
func TextArtifactAsExecuteRequestV1ArtifactsInner(v *TextArtifact) ExecuteRequestV1ArtifactsInner {
	return ExecuteRequestV1ArtifactsInner{
		inner: v,
	}
}

// VisualizationArtifactAsExecuteRequestV1ArtifactsInner is a convenience function that returns VisualizationArtifact wrapped in ExecuteRequestV1ArtifactsInner
func VisualizationArtifactAsExecuteRequestV1ArtifactsInner(v *VisualizationArtifact) ExecuteRequestV1ArtifactsInner {
	return ExecuteRequestV1ArtifactsInner{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExecuteRequestV1ArtifactsInner) UnmarshalJSON(data []byte) error {
	var result Artifact

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	*dst = ExecuteRequestV1ArtifactsInner(result)

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExecuteRequestV1ArtifactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(Artifact(src))
}

// GetArtifactType returns the artifact type of the underlying type.
func (src ExecuteRequestV1ArtifactsInner) GetArtifactType() ArtifactType {
	if src.IsNil() {
		return ""
	}

	return ArtifactType(src.inner.GetArtifactType())
}

// GetIdentifier returns the identifier of the underlying type.
func (src ExecuteRequestV1ArtifactsInner) GetIdentifier() string {
	if src.IsNil() {
		return ""
	}

	return src.inner.GetIdentifier()
}

// IsNil checks if the instance is empty.
func (src ExecuteRequestV1ArtifactsInner) IsNil() bool {
	return src.inner == nil
}

// Interface returns the inner interface.
func (src ExecuteRequestV1ArtifactsInner) Interface() ArtifactInterface {
	return src.inner
}

// AsText returns the instance as a nullable TextArtifact.
func (src ExecuteRequestV1ArtifactsInner) AsText() *TextArtifact {
	return Artifact(src).AsText()
}

// AsText converts the instance to a nullable TableArtifact.
func (src ExecuteRequestV1ArtifactsInner) AsTable() *TableArtifact {
	return Artifact(src).AsTable()
}

// AsText converts the instance to a nullable VisualizationArtifact.
func (src ExecuteRequestV1ArtifactsInner) AsVisualization() *VisualizationArtifact {
	return Artifact(src).AsVisualization()
}

// AsAutomation converts the instance to a nullable AutomationArtifact.
func (src ExecuteRequestV1ArtifactsInner) AsAutomation() *AutomationArtifact {
	return Artifact(src).AsAutomation()
}

type NullableExecuteRequestV1ArtifactsInner struct {
	value *ExecuteRequestV1ArtifactsInner
	isSet bool
}

func (v NullableExecuteRequestV1ArtifactsInner) Get() *ExecuteRequestV1ArtifactsInner {
	return v.value
}

func (v *NullableExecuteRequestV1ArtifactsInner) Set(val *ExecuteRequestV1ArtifactsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteRequestV1ArtifactsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteRequestV1ArtifactsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteRequestV1ArtifactsInner(val *ExecuteRequestV1ArtifactsInner) *NullableExecuteRequestV1ArtifactsInner {
	return &NullableExecuteRequestV1ArtifactsInner{value: val, isSet: true}
}

func (v NullableExecuteRequestV1ArtifactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteRequestV1ArtifactsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
