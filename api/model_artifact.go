package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

// ArtifactType represents an artifact type enum.
type ArtifactType string

const (
	ArtifactTypeText          ArtifactType = "text"
	ArtifactTypeTable         ArtifactType = "table"
	ArtifactTypeVisualization ArtifactType = "visualization"
)

var enumValuesArtifactType = []ArtifactType{
	ArtifactTypeText,
	ArtifactTypeTable,
	ArtifactTypeVisualization,
}

// ParseArtifactType parses the artifact type enum from string.
func ParseArtifactType(input string) (ArtifactType, error) {
	result := ArtifactType(input)

	if !slices.Contains(enumValuesArtifactType, result) {
		return "", fmt.Errorf(
			"invalid artifact type, expected one of %v, got %s",
			enumValuesArtifactType,
			input,
		)
	}

	return result, nil
}

// Unmarshal JSON data into one of the pointers in the enum.
func (dst *ArtifactType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	result, err := ParseArtifactType(raw)
	if err != nil {
		return err
	}

	*dst = result

	return nil
}

// ArtifactInterface abstracts the inner artifact object.
type ArtifactInterface interface {
	GetArtifactType() string
	GetIdentifier() string
}

// Artifact represents a artifact response item that is one of table, table or visualization artifact.
type Artifact struct {
	inner ArtifactInterface
}

// NewArtifact creates a general Artifact.
func NewArtifact[A ArtifactInterface](v A) Artifact {
	return Artifact{
		inner: v,
	}
}

// TableArtifactAsArtifact is a convenience function that returns TableArtifact wrapped in Artifact.
func TableArtifactAsArtifact(v *TableArtifact) Artifact {
	return Artifact{
		inner: v,
	}
}

// TextArtifactAsArtifact is a convenience function that returns TextArtifact wrapped in Artifact.
func TextArtifactAsArtifact(v *TextArtifact) Artifact {
	return Artifact{
		inner: v,
	}
}

// VisualizationArtifactAsArtifact is a convenience function that returns VisualizationArtifact wrapped in Artifact.
func VisualizationArtifactAsArtifact(v *VisualizationArtifact) Artifact {
	return Artifact{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *Artifact) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) == 0 {
		return errors.New("unable to parse the Artifact from empty object")
	}

	rawArtifactType, ok := raw["artifact_type"]
	if !ok {
		return errors.New("unable to parse the Artifact: field artifact_type is required")
	}

	var artifactType ArtifactType

	if err := json.Unmarshal(rawArtifactType, &artifactType); err != nil {
		return fmt.Errorf("unable to parse the Artifact: %w", err)
	}

	switch artifactType {
	case ArtifactTypeText:
		result := TextArtifact{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case ArtifactTypeTable:
		result := TableArtifact{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case ArtifactTypeVisualization:
		result := VisualizationArtifact{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	default:
		return fmt.Errorf("unsupported artifact type: %s", artifactType)
	}

	return nil
}

// MarshalJSON Marshal data from the first non-nil pointers in the struct to JSON.
func (src Artifact) MarshalJSON() ([]byte, error) {
	if src.inner != nil {
		return json.Marshal(src.inner)
	}

	return nil, nil // no data in oneOf schemas
}

// GetArtifactType returns the artifact type of the underlying type.
func (src Artifact) GetArtifactType() ArtifactType {
	if src.IsNil() {
		return ""
	}

	return ArtifactType(src.inner.GetArtifactType())
}

// GetIdentifier returns the identifier of the underlying type.
func (src Artifact) GetIdentifier() string {
	if src.IsNil() {
		return ""
	}

	return src.inner.GetIdentifier()
}

// IsNil checks if the instance is empty.
func (src Artifact) IsNil() bool {
	return src.inner == nil
}

// Interface returns the inner interface.
func (src Artifact) Interface() ArtifactInterface {
	return src.inner
}

// String prints the instance to string.
func (src Artifact) String() string {
	return fmt.Sprint(src.inner)
}

// AsText returns the instance as a nullable TextArtifact.
func (src Artifact) AsText() *TextArtifact {
	result, ok := src.inner.(*TextArtifact)
	if !ok {
		return nil
	}

	return result
}

// AsText converts the instance to a nullable TableArtifact.
func (src Artifact) AsTable() *TableArtifact {
	result, ok := src.inner.(*TableArtifact)
	if !ok {
		return nil
	}

	return result
}

// AsText converts the instance to a nullable VisualizationArtifact.
func (src Artifact) AsVisualization() *VisualizationArtifact {
	result, ok := src.inner.(*VisualizationArtifact)
	if !ok {
		return nil
	}

	return result
}
