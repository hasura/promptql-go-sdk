package promptql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasura/promptql-go-sdk/api"
	"github.com/tmaxmax/go-sse"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// QueryStreamCallback is the callback function for /query streaming.
type QueryStreamCallback func(chunk api.QueryResponseChunk) error

// Query sends a request to the [Natural Language Query API] that allows you to interact with PromptQL directly,
// sending messages and receiving responses. The response is non-streaming.
//
// [Natural Language Query API]: https://hasura.io/docs/promptql/promptql-apis/natural-language-api/
func (c *Client) Query(ctx context.Context, body api.QueryRequest) (*api.QueryResponse, error) {
	ctx, span := tracer.Start(ctx, "Query")
	defer span.End()

	validatedBody, err := c.validateQueryRequest(body)
	if err != nil {
		span.SetStatus(codes.Error, "invalid request")
		span.RecordError(err)

		return nil, err
	}

	validatedBody.Stream = false

	req := c.APIClient.DefaultAPI.QueryQueryPost(ctx).QueryRequest(validatedBody)

	result, _, err := req.Execute()
	if err != nil {
		span.SetStatus(codes.Error, "failed to execute query")
		span.RecordError(err)
	}

	return result, err
}

// QueryStream sends a request to the [Natural Language Query API] that allows you to interact with PromptQL directly,
// sending messages and receiving responses. The response is streaming.
//
// [Natural Language Query API]: https://hasura.io/docs/promptql/promptql-apis/natural-language-api/
func (c *Client) QueryStream(
	ctx context.Context,
	body api.QueryRequest,
	callback QueryStreamCallback,
) error {
	ctx, span := tracer.Start(ctx, "QueryStream")
	defer span.End()

	if callback == nil {
		err := errors.New("'callback' function is required")
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	validatedBody, err := c.validateQueryRequest(body)
	if err != nil {
		span.SetStatus(codes.Error, "invalid request")
		span.RecordError(err)

		return err
	}

	validatedBody.Stream = true
	queryRequest := c.APIClient.DefaultAPI.QueryQueryPost(ctx).QueryRequest(validatedBody)

	req, err := c.DefaultAPI.QueryQueryPostRequest(queryRequest)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache", "no-cache")

	resp, err := c.DefaultAPI.QueryQueryPostExecuteRequest(req)
	if err != nil {
		span.SetStatus(codes.Error, "failed to execute query")
		span.RecordError(err)

		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var readErr error

	sse.Read(resp.Body, nil)(func(e sse.Event, err error) bool {
		if err != nil {
			readErr = err

			return false
		}

		if e.Data == "" {
			return true
		}

		if c.config.Logger != nil {
			c.config.Logger.Debug("received chunk", slog.String("data", e.Data))
		}

		var chunk api.QueryResponseChunk
		if err := json.NewDecoder(strings.NewReader(e.Data)).Decode(&chunk); err != nil {
			span.AddEvent("data_decode_failure", trace.WithAttributes(
				attribute.String("data", e.Data),
				attribute.String("error", err.Error()),
			))

			return true
		}

		if err := callback(chunk); err != nil {
			readErr = err

			return false
		}

		return true
	})

	if readErr != nil {
		span.SetStatus(codes.Error, "error happened when processing chunks")
		span.RecordError(readErr)
	}

	return readErr
}

func (c *Client) validateQueryRequest(body api.QueryRequest) (api.QueryRequest, error) {
	if len(body.Interactions) == 0 {
		return body, errors.New("require at least 1 interaction")
	}

	for i, interaction := range body.Interactions {
		if interaction.UserMessage.Text == "" {
			return body, fmt.Errorf("the user message in interactions[%d] must not be empty", i)
		}
	}

	if err := c.validateDdnConfig(&body.Ddn); err != nil {
		return body, err
	}

	body.Version = "v1"

	if body.Timezone == "" {
		body.Timezone = c.config.Timezone
	}

	if !body.SystemInstructions.IsSet() && c.config.SystemInstructions != "" {
		body.SystemInstructions.Set(&c.config.SystemInstructions)
	}

	if c.config.Llm != nil && (body.Llm == nil || body.Llm.IsNil()) {
		llm := api.NewLlm(c.config.Llm)
		body.Llm = &llm
	}

	if c.config.AIPrimitivesLlmProvider != nil &&
		(body.AiPrimitivesLlm.Get() == nil || body.AiPrimitivesLlm.Get().IsNil()) {
		llm := api.NewExecuteRequestAiPrimitivesLlm(c.config.AIPrimitivesLlmProvider)
		body.AiPrimitivesLlm.Set(&llm)
	}

	return body, nil
}

// NewQueryRequestMessage is a convenient function to create a simple query request with message only.
func NewQueryRequestMessage(message string) api.QueryRequest {
	return api.QueryRequest{
		Interactions: []api.Interaction{
			{
				UserMessage: api.UserMessage{
					Text: message,
				},
			},
		},
	}
}
