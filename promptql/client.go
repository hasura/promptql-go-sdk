// Package promptql provides a PromptQL client implementation.
package promptql

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hasura/promptql-go-sdk/api"
	"go.opentelemetry.io/otel"
)

// Client represents a PromptQL API client.
type Client struct {
	*api.APIClient

	config *ClientConfig
}

// NewClient create a new PromptQL client instance.
func NewClient(apiKey string, config *ClientConfig) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("promptql api key must not be empty")
	}

	validatedConfig, promptqlBaseURL, err := validateClientConfig(config)
	if err != nil {
		return nil, err
	}

	promptqlHeaders := map[string]string{
		"Authorization": "Bearer " + apiKey,
	}

	for key, value := range config.PromptQLHeaders {
		if key != "" {
			promptqlHeaders[key] = value
		}
	}

	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	promptqlServerPort, err := parsePortFromURL(promptqlBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid PromptQL URL port: %w", err)
	}

	clientWrapper := &transportWrapper{
		propagator:    otel.GetTextMapPropagator(),
		doer:          config.HTTPClient,
		serverAddress: promptqlBaseURL.Hostname(),
		serverPort:    promptqlServerPort,
	}

	client := api.NewAPIClient(&api.Configuration{
		DefaultHeader: promptqlHeaders,
		UserAgent:     "promptql-go/sdk/0.1.0",
		Servers: api.ServerConfigurations{
			{
				URL: validatedConfig.PromptQLBaseURL,
			},
		},
		HTTPClient: clientWrapper,
	})

	return &Client{
		APIClient: client,
		config:    validatedConfig,
	}, nil
}

func (c *Client) validateDdnConfig(ddn *api.DdnConfig) error {
	switch {
	case ddn.Url != "":
		_, err := validateHttpURL(ddn.Url)
		if err != nil {
			return fmt.Errorf("invalid ddn.url: %w", err)
		}
	case c.config.DdnBaseURL != "":
		ddn.Url = c.config.DdnBaseURL
	default:
		return errors.New("ddn.url is required")
	}

	if ddn.Headers == nil {
		ddn.Headers = make(map[string]string)
	}

	for key, value := range c.config.DdnHeaders {
		ddn.Headers[key] = value
	}

	if c.config.DdnHeadersFn != nil {
		headers, err := c.config.DdnHeadersFn()
		if err != nil {
			return err
		}

		for key, value := range headers {
			ddn.Headers[key] = value
		}
	}

	return nil
}
