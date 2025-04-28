package promptql

import (
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/hasura/promptql-go-sdk/api"
)

const (
	defaultPromptQLBaseURL = "https://api.promptql.pro.hasura.io"
	defaultTimezone        = "UTC"
)

// ClientConfig holds configurations of the PromptQL client.
type ClientConfig struct {
	// The optional base URL of PromptQL API. The default value is the endpoint to the public DDN.
	PromptQLBaseURL string

	// Custom headers to be injected into PromptQL requests.
	PromptQLHeaders map[string]string

	// The base URL of your Hasura DDN project.
	DdnBaseURL string

	// Default headers to be forwarded to your Hasura DDN service.
	DdnHeaders map[string]string

	// Lazy function to generate headers to be forwarded to your Hasura DDN service.
	DdnHeadersFn func() (map[string]string, error)

	// Optional default LLM configuration to be used for natural language API requests.
	Llm api.LlmInterface

	// Optional default configuration for the AI primitives LLM provider.
	AIPrimitivesLlmProvider api.LlmInterface

	// Optional default system instructions for the natural language API request.
	SystemInstructions string

	// An [IANA timezone] for interpreting time-based queries. Default is the UTC timezone.
	//
	// [IANA timezone]: https://data.iana.org/time-zones/tzdb-2021a/zone1970.tab
	Timezone string

	// Custom HTTP client or Doer implementation, default: http.DefaultClient
	HTTPClient api.Doer

	// Set a logger to enable logging.
	Logger *slog.Logger
}

// NewApiAnthropicConfig instantiates a new ApiAnthropicConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApiAnthropicConfig(apiKey string) *api.ApiAnthropicConfig {
	return &api.ApiAnthropicConfig{
		Provider: api.LlmProviderAnthropic,
		ApiKey:   apiKey,
	}
}

// will change when the set of required properties is changed.
func NewHasuraLlmConfig() *api.HasuraLlmConfig {
	return &api.HasuraLlmConfig{
		Provider: api.LlmProviderHasura,
	}
}

// will change when the set of required properties is changed.
func NewApiOpenAIConfig(apiKey string) *api.ApiOpenAIConfig {
	this := api.ApiOpenAIConfig{
		Provider: api.LlmProviderHasura,
		ApiKey:   apiKey,
	}

	return &this
}

func validateClientConfig(options *ClientConfig) (*ClientConfig, *url.URL, error) {
	if options == nil {
		options = &ClientConfig{}
	}

	var promptqlBaseURL *url.URL

	var err error

	if options.PromptQLBaseURL != "" {
		promptqlBaseURL, err = validateHttpURL(options.PromptQLBaseURL)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid PromptQL base url: %w", err)
		}
	} else {
		options.PromptQLBaseURL = defaultPromptQLBaseURL

		promptqlBaseURL, err = url.Parse(defaultPromptQLBaseURL)
		if err != nil {
			return nil, nil, err
		}
	}

	if options.Timezone == "" {
		options.Timezone = defaultTimezone
	}

	if options.DdnBaseURL != "" {
		baseURL, err := validateHttpURL(options.DdnBaseURL)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid DDN base url: %w", err)
		}

		baseURL.Path = strings.TrimRight(baseURL.Path, "/") + "/v1/sql"
		options.DdnBaseURL = baseURL.String()
	}

	return options, promptqlBaseURL, nil
}
