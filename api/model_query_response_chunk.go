package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

// QueryResponseChunkType represents a query response chunk type enum.
type QueryResponseChunkType string

const (
	QueryResponseTypeAssistantActionChunk QueryResponseChunkType = "assistant_action_chunk"
	QueryResponseTypeArtifactUpdateChunk  QueryResponseChunkType = "artifact_update_chunk"
	QueryResponseTypeErrorChunk           QueryResponseChunkType = "error_chunk"
)

var enumValuesQueryResponseChunkType = []QueryResponseChunkType{
	QueryResponseTypeAssistantActionChunk,
	QueryResponseTypeArtifactUpdateChunk,
	QueryResponseTypeErrorChunk,
}

// ParseQueryResponseChunkType parses the QueryResponseChunk type enum from string.
func ParseQueryResponseChunkType(input string) (QueryResponseChunkType, error) {
	result := QueryResponseChunkType(input)

	if !slices.Contains(enumValuesQueryResponseChunkType, result) {
		return "", fmt.Errorf(
			"invalid QueryResponseChunk type, expected one of %v, got %s",
			enumValuesQueryResponseChunkType,
			input,
		)
	}

	return result, nil
}

// QueryResponseChunkInterface abstracts an underlying QueryResponseChunk struct.
type QueryResponseChunkInterface interface {
	GetType() string
}

// QueryResponseChunk the streaming response sends chunks of data in Server-Sent Events (SSE) format.
type QueryResponseChunk struct {
	inner QueryResponseChunkInterface
}

// ArtifactUpdateChunkAsQueryResponseChunk is a convenience function that returns ArtifactUpdateChunk wrapped in QueryResponseChunk.
func ArtifactUpdateChunkAsQueryResponseChunk(v *ArtifactUpdateChunk) QueryResponseChunk {
	return QueryResponseChunk{
		inner: v,
	}
}

// AssistantActionChunkAsQueryResponseChunk is a convenience function that returns AssistantActionChunk wrapped in QueryResponseChunk.
func AssistantActionChunkAsQueryResponseChunk(v *AssistantActionChunk) QueryResponseChunk {
	return QueryResponseChunk{
		inner: v,
	}
}

// ErrorChunkAsQueryResponseChunk is a convenience function that returns ErrorChunk wrapped in QueryResponseChunk.
func ErrorChunkAsQueryResponseChunk(v *ErrorChunk) QueryResponseChunk {
	return QueryResponseChunk{
		inner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *QueryResponseChunk) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) == 0 {
		return errors.New("unable to parse the QueryResponseChunk from empty object")
	}

	rawChunkType, ok := raw["type"]
	if !ok {
		return errors.New("unable to parse the QueryResponseChunk: field type is required")
	}

	var chunkType QueryResponseChunkType

	if err := json.Unmarshal(rawChunkType, &chunkType); err != nil {
		return fmt.Errorf("unable to parse the QueryResponseChunk: %w", err)
	}

	switch chunkType {
	case QueryResponseTypeAssistantActionChunk:
		result := AssistantActionChunk{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case QueryResponseTypeArtifactUpdateChunk:
		result := ArtifactUpdateChunk{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case QueryResponseTypeErrorChunk:
		result := ErrorChunk{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	default:
		return fmt.Errorf("unsupported QueryResponseChunk type: %s", rawChunkType)
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src QueryResponseChunk) MarshalJSON() ([]byte, error) {
	if src.inner != nil {
		return json.Marshal(src.inner)
	}

	return nil, nil // no data in oneOf schemas
}

// Interface returns the inner interface.
func (src QueryResponseChunk) Interface() QueryResponseChunkInterface {
	return src.inner
}

// IsNil checks if the instance is empty.
func (src QueryResponseChunk) IsNil() bool {
	return src.inner == nil
}

// String prints the instance to string.
func (src QueryResponseChunk) String() string {
	return fmt.Sprint(src.inner)
}

// AsAssistantActionChunk returns the instance as a nullable AssistantActionChunk.
func (src QueryResponseChunk) AsAssistantActionChunk() *AssistantActionChunk {
	result, ok := src.inner.(*AssistantActionChunk)
	if !ok {
		return nil
	}

	return result
}

// AsArtifactUpdateChunk returns the instance as a nullable ArtifactUpdateChunk.
func (src QueryResponseChunk) AsArtifactUpdateChunk() *ArtifactUpdateChunk {
	result, ok := src.inner.(*ArtifactUpdateChunk)
	if !ok {
		return nil
	}

	return result
}

// AsErrorChunk returns the instance as a nullable ErrorChunk.
func (src QueryResponseChunk) AsErrorChunk() *ErrorChunk {
	result, ok := src.inner.(*ErrorChunk)
	if !ok {
		return nil
	}

	return result
}
