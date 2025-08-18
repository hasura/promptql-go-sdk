package api

import (
	"encoding/json"
	"fmt"
	"slices"
)

// QueryRequestVersion represents a query request version enum.
type QueryRequestVersion string

const (
	QueryRequestVersion1 QueryRequestVersion = "v1"
	QueryRequestVersion2 QueryRequestVersion = "v2"
)

var enumValuesQueryRequestVersion = []QueryRequestVersion{
	QueryRequestVersion1,
	QueryRequestVersion2,
}

// ParseQueryRequestVersion parses the query request version enum from string.
func ParseQueryRequestVersion(input string) (QueryRequestVersion, error) {
	result := QueryRequestVersion(input)

	if !slices.Contains(enumValuesQueryRequestVersion, result) {
		return "", fmt.Errorf(
			"invalid query request version, expected one of %v, got %s",
			enumValuesQueryRequestVersion,
			input,
		)
	}

	return result, nil
}

// Unmarshal JSON data into one of the pointers in the enum.
func (dst *QueryRequestVersion) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	result, err := ParseQueryRequestVersion(raw)
	if err != nil {
		return err
	}

	*dst = result

	return nil
}

// QueryRequestInterface abstracts the inner query request object.
type QueryRequestInterface interface {
	GetVersion() string
	GetInteractions() []ApiThreadInteraction
}

// QueryRequest struct for a query request object.
type QueryRequest struct {
	inner QueryRequestInterface
}

// QueryRequestV1AsQueryRequest is a convenience function that returns QueryRequestV1 wrapped in QueryRequest
func QueryRequestV1AsQueryRequest(inner *QueryRequestV1) QueryRequest {
	return QueryRequest{
		inner: inner,
	}
}

// QueryRequestV2AsQueryRequest is a convenience function that returns QueryRequestV2 wrapped in QueryRequest
func QueryRequestV2AsQueryRequest(inner *QueryRequestV2) QueryRequest {
	return QueryRequest{
		inner: inner,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *QueryRequest) UnmarshalJSON(data []byte) error {
	var rawVersion rawQueryRequestVersion

	if err := json.Unmarshal(data, &rawVersion); err != nil {
		return err
	}

	switch rawVersion.Version {
	case QueryRequestVersion1:
		result := QueryRequestV1{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	case QueryRequestVersion2:
		result := QueryRequestV2{}

		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		dst.inner = &result
	default:
		return fmt.Errorf("unsupported query request type: %s", rawVersion.Version)
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(src.inner)
}

// IsNil checks if the instance is empty.
func (src QueryRequest) IsNil() bool {
	return src.inner == nil
}

// Interface returns the inner interface.
func (src QueryRequest) Interface() QueryRequestInterface {
	return src.inner
}

// String prints the instance to string.
func (src QueryRequest) String() string {
	return fmt.Sprint(src.inner)
}

// AsQueryRequestV1 returns the instance as a nullable QueryRequestV1.
func (src QueryRequest) AsQueryRequestV1() *QueryRequestV1 {
	result, ok := src.inner.(*QueryRequestV1)
	if !ok {
		return nil
	}

	return result
}

// AsQueryRequestV2 returns the instance as a nullable QueryRequestV2.
func (src QueryRequest) AsQueryRequestV2() *QueryRequestV2 {
	result, ok := src.inner.(*QueryRequestV2)
	if !ok {
		return nil
	}

	return result
}

type rawQueryRequestVersion struct {
	Version QueryRequestVersion `json:"version"`
}
