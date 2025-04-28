package promptql

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/hasura/promptql-go-sdk/api"
)

// will change when the set of required properties is changed.
func NewTableArtifact(identifier string, title string, data []map[string]any) *api.TableArtifact {
	return api.NewTableArtifact(identifier, title, string(api.ArtifactTypeTable), data)
}

// will change when the set of required properties is changed.
func NewTextArtifact(identifier string, title string, data string) *api.TextArtifact {
	return api.NewTextArtifact(identifier, title, string(api.ArtifactTypeText), data)
}

// will change when the set of required properties is changed.
func NewVisualizationArtifact(
	identifier string,
	title string,
	data api.VisualizationArtifactData,
) *api.VisualizationArtifact {
	return api.NewVisualizationArtifact(
		identifier,
		title,
		string(api.ArtifactTypeVisualization),
		data,
	)
}

func concatNullableString(a, b api.NullableString) api.NullableString {
	valueA := a.Get()
	valueB := b.Get()

	if valueA == nil {
		return b
	}

	if valueB == nil {
		return a
	}

	newValue := *valueA + *valueB
	result := api.NewNullableString(&newValue)

	return *result
}

func validateHttpURL(input string) (*url.URL, error) {
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		return nil, errors.New("invalid http url")
	}

	return url.Parse(input)
}

func parsePortFromURL(input *url.URL) (int, error) {
	port := input.Port()
	if port != "" {
		return strconv.Atoi(port)
	}

	if input.Scheme == "https" {
		return 443, nil
	}

	return 80, nil
}
