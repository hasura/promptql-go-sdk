package promptql_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/hasura/promptql-go-sdk/api"
	"github.com/hasura/promptql-go-sdk/promptql"
)

func TestQueryAndExecuteProgram(t *testing.T) {
	client := createTestClient(t)

	result, err := client.Query(
		context.Background(),
		promptql.NewQueryRequestMessage("get list customers"),
	)
	if err != nil {
		t.Fatalf("expected non-error, got: %s", err)
	}

	if actionLen := len(result.AssistantActions); actionLen == 0 {
		t.Fatalf("expected len(assistant_actions) > 0 , got: %d", actionLen)
	}

	if msg := result.AssistantActions[0].Message.Get(); msg == nil || *msg == "" {
		t.Fatalf("expected non-empty action message, got: %+v", result.AssistantActions[0])
	}

	expectedArtifacts := []api.ExecuteRequestArtifactsInner{
		api.TableArtifactAsExecuteRequestArtifactsInner(&api.TableArtifact{
			Data: []map[string]any{
				{
					"id":   float64(1),
					"name": "John",
				},
				{
					"id":   float64(2),
					"name": "Tom",
				},
				{
					"id":   float64(3),
					"name": "Jerry",
				},
			},
		}),
	}

	assertTableArtifacts(t, expectedArtifacts, result.ModifiedArtifacts)

	var code string

	for _, action := range result.AssistantActions {
		rawCode := action.Code.Get()
		if rawCode != nil && *rawCode != "" {
			code = *rawCode
		}
	}

	if code == "" {
		t.Fatal("expected at least 1 action that contains code")
	}

	execResult, err := client.ExecuteProgram(
		context.Background(),
		promptql.NewExecuteProgramRequest(code),
	)
	if err != nil {
		t.Fatalf("expected non-error, got: %s", err)
	}

	assertTableArtifacts(t, result.ModifiedArtifacts, execResult.ModifiedArtifacts)
}

func TestQueryStream(t *testing.T) {
	client := createTestClient(t)

	err := client.QueryStream(
		context.Background(),
		promptql.NewQueryRequestMessage("get list customers"),
		func(chunk api.QueryResponseChunk) error {
			return nil
		},
	)
	if err != nil {
		t.Fatalf("expected non-error, got: %s", err)
	}
}

func createTestClient(t *testing.T) *promptql.Client {
	apiKey := assertGetenv(t, "PROMPTQL_API_KEY")
	ddnBaseURL := assertGetenv(t, "HASURA_DDN_BASE_URL")
	ddnAuthToken := assertGetenv(t, "DDN_AUTH_TOKEN")

	client, err := promptql.NewClient(apiKey, &promptql.ClientConfig{
		DdnBaseURL: ddnBaseURL,
		DdnHeaders: map[string]string{
			"Authorization": ddnAuthToken,
		},
		Timezone: "America/Los_Angeles",
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})),
	})
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}

	return client
}

func assertTableArtifacts(t *testing.T, expected, results []api.ExecuteRequestArtifactsInner) {
	if artifactLen := len(results); artifactLen == 0 {
		t.Fatalf("expected len(modified_artifacts) > 0 , got: %d", artifactLen)
	}

	switch artifact := results[0].Interface().(type) {
	case *api.TableArtifact:
		slices.SortFunc(artifact.Data, func(a, b map[string]any) int {
			return strings.Compare(fmt.Sprint(a["id"]), fmt.Sprint(b["id"]))
		})

		expectedData := expected[0].AsTable().Data
		assertDeepEqual(t, expectedData, artifact.Data, "artifact data not equal")
	default:
		t.Fatalf("expected table artifact, got: %s", artifact.GetArtifactType())
	}
}

func assertDeepEqual(t *testing.T, expected, result any, message string) {
	t.Helper()

	if api.IsNil(expected) && api.IsNil(result) {
		return
	}

	if !reflect.DeepEqual(expected, result) {
		eb, _ := json.Marshal(expected)
		ab, _ := json.Marshal(result)

		t.Errorf("%s,\nexpected: %s\ngot     : %s", message, string(eb), string(ab))
	}
}

func assertGetenv(t *testing.T, key string) string {
	t.Helper()

	value := os.Getenv(key)
	if value == "" {
		t.Fatalf("env %s is required", key)
	}

	return value
}
