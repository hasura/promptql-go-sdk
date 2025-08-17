package api

import (
	"encoding/json"
	"fmt"
	"slices"
)

// ExecuteRequestVersion represents an execute request version enum.
type ExecuteRequestVersion string

const (
	ExecuteRequestVersion1 ExecuteRequestVersion = "v1"
	ExecuteRequestVersion2 ExecuteRequestVersion = "v2"
)

var enumValuesExecuteRequestVersion = []ExecuteRequestVersion{
	ExecuteRequestVersion1,
	ExecuteRequestVersion2,
}

// ParseExecuteRequestVersion parses the execute request version enum from string.
func ParseExecuteRequestVersion(input string) (ExecuteRequestVersion, error) {
	result := ExecuteRequestVersion(input)

	if !slices.Contains(enumValuesExecuteRequestVersion, result) {
		return "", fmt.Errorf(
			"invalid execute request version, expected one of %v, got %s",
			enumValuesExecuteRequestVersion,
			input,
		)
	}

	return result, nil
}

// Unmarshal JSON data into one of the pointers in the enum.
func (dst *ExecuteRequestVersion) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	result, err := ParseExecuteRequestVersion(raw)
	if err != nil {
		return err
	}

	*dst = result

	return nil
}

// ExecuteRequestInterface abstracts the inner execute request object.
type ExecuteRequestInterface interface {
	GetVersion() string
	GetCode() string
}

// ExecuteRequest struct for an execute request object.
type ExecuteRequest struct {
	inner ExecuteRequestInterface
}

// ExecuteRequestV1AsQueryRequest is a convenience function that returns ExecuteRequestV1 wrapped in ExecuteRequest
func ExecuteRequestV1AsQueryRequest(inner *ExecuteRequestV1) ExecuteRequest {
	return ExecuteRequest{
		inner: inner,
	}
}

// ExecuteRequestV2AsQueryRequest is a convenience function that returns ExecuteRequestV2 wrapped in ExecuteRequest
func ExecuteRequestV2AsQueryRequest(inner *ExecuteRequestV2) ExecuteRequest {
	return ExecuteRequest{
		inner: inner,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExecuteRequest) UnmarshalJSON(data []byte) error {
	var rawVersion rawExecuteRequestVersion

	if err := json.Unmarshal(data, &rawVersion); err != nil {
		return err
	}

	switch rawVersion.Version {
	case ExecuteRequestVersion1:
		result := ExecuteRequestV1{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case ExecuteRequestVersion2:
		result := ExecuteRequestV2{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	default:
		return fmt.Errorf("unsupported execute request type: %s", rawVersion.Version)
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExecuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(src.inner)
}

// IsNil checks if the instance is empty.
func (src ExecuteRequest) IsNil() bool {
	return src.inner == nil
}

// Interface returns the inner interface.
func (src ExecuteRequest) Interface() ExecuteRequestInterface {
	return src.inner
}

// String prints the instance to string.
func (src ExecuteRequest) String() string {
	return fmt.Sprint(src.inner)
}

// AsExecuteRequestV1 returns the instance as a nullable ExecuteRequestV1.
func (src ExecuteRequest) AsExecuteRequestV1() *ExecuteRequestV1 {
	result, ok := src.inner.(*ExecuteRequestV1)
	if !ok {
		return nil
	}

	return result
}

// AsExecuteRequestV2 returns the instance as a nullable ExecuteRequestV2.
func (src ExecuteRequest) AsExecuteRequestV2() *ExecuteRequestV2 {
	result, ok := src.inner.(*ExecuteRequestV2)
	if !ok {
		return nil
	}

	return result
}

type rawExecuteRequestVersion struct {
	Version ExecuteRequestVersion `json:"version"`
}
