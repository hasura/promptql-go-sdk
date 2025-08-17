package promptql

import (
	"context"
	"errors"

	"github.com/hasura/promptql-go-sdk/api"
	"go.opentelemetry.io/otel/codes"
)

// ExecuteProgram sends a request to the [Natural Language Query API] that allows you to interact with PromptQL directly,
// sending messages and receiving responses. The response is non-streaming.
//
// [Natural Language Query API]: https://hasura.io/docs/promptql/promptql-apis/natural-language-api/
func (c *Client) ExecuteProgram(
	ctx context.Context,
	body api.ExecuteRequestV1,
) (*api.PromptQlExecutionResult, error) {
	ctx, span := tracer.Start(ctx, "ExecuteProgram")
	defer span.End()

	if body.Code == "" {
		err := errors.New("'code' is required")

		span.SetStatus(codes.Error, "validation failure")
		span.RecordError(err)

		return nil, err
	}

	if body.Artifacts == nil {
		body.Artifacts = make([]api.ExecuteRequestV1ArtifactsInner, 0)
	}

	ddn := body.Ddn.Get()
	if ddn == nil {
		ddn = &api.DdnConfig{}
	}

	if err := c.validateDdnConfig(ddn); err != nil {
		span.SetStatus(codes.Error, "ddn config validation failure")
		span.RecordError(err)

		return nil, err
	}

	body.Ddn.Set(ddn)

	if body.AiPrimitivesLlm.Get() == nil {
		llm := api.NewExecuteRequestV1AiPrimitivesLlm(NewHasuraLlmConfig())
		body.AiPrimitivesLlm.Set(&llm)
	}

	req := c.APIClient.DefaultAPI.ExecuteProgramExecuteProgramPost(ctx).
		Request(api.ExecuteRequestV1AsQueryRequest(&body))

	result, _, err := req.Execute()
	if err != nil {
		span.SetStatus(codes.Error, "execution failure")
		span.RecordError(err)
	}

	return result, err
}

// NewExecuteProgramRequest is a convenient function to create a execute program request from code.
func NewExecuteProgramRequestV1(code string) api.ExecuteRequestV1 {
	return api.ExecuteRequestV1{
		Code: code,
	}
}
