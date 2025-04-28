package promptql_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hasura/promptql-go-sdk/api"
	"github.com/hasura/promptql-go-sdk/promptql"
)

func TestQueryChunks(t *testing.T) {
	fixtures := []string{
		"{\"message\":\"I'll re\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":\"trieve the list of custome\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":\"rs for you.\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":\"1\",\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":\". Query the customers f\",\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":\"unction to get the list of all customers\",\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":\"\\n2. Store the results in\",\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":\" a table artifact showing customer IDs and names\",\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"sql = \\\"\\\"\\\"\\nSE\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"LECT id, name \\nFROM customers(\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\")\\nORDER BY name\\n\\\"\\\"\\\"\\n\\ncustom\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"ers = executor.r\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"un_sql(sql)\\n\\nif len(customers\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\") == 0:\\n    executor.print(\\\"No cust\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"omers found.\\\")\\nelse:\\n    executor.\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"store_artifact(\\n    \",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"    'customer_list',\\n        'List of Cus\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":\"tomers',\\n        'table',\\n        customers\\n    )\",\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":null,\"plan\":null,\"code\":null,\"code_output\":\"SQL statement returned 3 rows. Sample rows: [{'id': 3.0, 'name': 'Jerry'}, {'id': 1.0, 'name': 'John'}]\\n\",\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"type\":\"artifact_update_chunk\",\"artifact\":{\"identifier\":\"customer_list\",\"title\":\"List of Customers\",\"artifact_type\":\"table\",\"data\":[{\"id\":3,\"name\":\"Jerry\"},{\"id\":1,\"name\":\"John\"},{\"id\":2,\"name\":\"Tom\"}]}}",
		"{\"message\":null,\"plan\":null,\"code\":null,\"code_output\":\"Stored table artifact: identifier = 'customer_list', title = 'List of Customers', number of rows = 3, sample rows preview = [{'id': 3, 'name': 'Jerry'}, {'id': 1, 'name': 'John'}]\\n\",\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":0}",
		"{\"message\":\"H\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":1}",
		"{\"message\":\"ere are all the customers:\\n<artifact ide\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":1}",
		"{\"message\":\"ntifier='customer_list' warning='I cannot see \",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":1}",
		"{\"message\":\"the\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":1}",
		"{\"message\":\" full data so I must not make up observations' />\",\"plan\":null,\"code\":null,\"code_output\":null,\"code_error\":null,\"type\":\"assistant_action_chunk\",\"index\":1}",
	}

	expected := promptql.QueryChunks{
		ArtifactUpdateChunks: map[string]api.ArtifactUpdateChunk{
			"artifact_update_chunk:customer_list": {
				Type: string(api.QueryResponseTypeArtifactUpdateChunk),
				Artifact: api.TableArtifactAsArtifact(promptql.NewTableArtifact("customer_list", "List of Customers", []map[string]any{
					{"id": float64(3), "name": "Jerry"},
					{"id": float64(1), "name": "John"},
					{"id": float64(2), "name": "Tom"},
				})),
			},
		},
		AssistantActionChunks: []api.AssistantActionChunk{
			{
				Type:       string(api.QueryResponseTypeAssistantActionChunk),
				Code:       *api.NewNullableString(api.PtrString("sql = \"\"\"\nSELECT id, name \nFROM customers()\nORDER BY name\n\"\"\"\n\ncustomers = executor.run_sql(sql)\n\nif len(customers) == 0:\n    executor.print(\"No customers found.\")\nelse:\n    executor.store_artifact(\n        'customer_list',\n        'List of Customers',\n        'table',\n        customers\n    )")),
				CodeOutput: *api.NewNullableString(api.PtrString("SQL statement returned 3 rows. Sample rows: [{'id': 3.0, 'name': 'Jerry'}, {'id': 1.0, 'name': 'John'}]\nStored table artifact: identifier = 'customer_list', title = 'List of Customers', number of rows = 3, sample rows preview = [{'id': 3, 'name': 'Jerry'}, {'id': 1, 'name': 'John'}]\n")),
				CodeError:  *api.NewNullableString(nil),
				Index:      0,
				Message:    *api.NewNullableString(api.PtrString("I'll retrieve the list of customers for you.")),
				Plan:       *api.NewNullableString(api.PtrString("1. Query the customers function to get the list of all customers\n2. Store the results in a table artifact showing customer IDs and names")),
			},
			{
				Type:       string(api.QueryResponseTypeAssistantActionChunk),
				Index:      1,
				Message:    *api.NewNullableString(api.PtrString("Here are all the customers:\n\u003cartifact identifier='customer_list' warning='I cannot see the full data so I must not make up observations' /\u003e")),
				Plan:       *api.NewNullableString(nil),
				Code:       *api.NewNullableString(nil),
				CodeOutput: *api.NewNullableString(nil),
				CodeError:  *api.NewNullableString(nil),
			},
		},
		ErrorChunk: nil,
	}
	result := promptql.NewQueryChunks()

	for _, text := range fixtures {
		var chunk api.QueryResponseChunk
		if err := json.Unmarshal([]byte(text), &chunk); err != nil {
			t.Fatalf("failed to decode json chunk: %s", err)
		}

		if err := result.Add(chunk); err != nil {
			t.Errorf("failed to add chunk: %s, data: %s", err, chunk)
		}
	}

	if !reflect.DeepEqual(expected, *result) {
		assertDeepEqual(t, expected, *result, "not equal")
	}
}
