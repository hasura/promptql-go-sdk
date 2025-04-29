# PromptQL Go SDK

A Go SDK for interacting with the [PromptQL Natural Language API](https://hasura.io/docs/promptql/promptql-apis/natural-language-api/).

## Installation

```bash
go get github.com/hasura/promptql-go-sdk
```

## Get started

### Prerequisite

- If you are new with PromptQL, follow [the quickstart guide of PromptQL](https://hasura.io/docs/promptql/quickstart/) to create a project.
- Create a PromptQL API Key in project settings tab on [https://console.hasura.io](https://console.hasura.io).
- Your Project API endpoint and security headers.

### Use PromptQL SDK

Create the PromptQL client with required configurations:

```go

import (
    "github.com/hasura/promptql-go-sdk/promptql"
)

func main() {
    client, err := promptql.NewClient("<promptql-api-key>", &promptql.ClientConfig{
        DdnBaseURL: "https://your-ddn-project",
        DdnHeaders: map[string]string{
            // authorization headers to your ddn project
            // "Authorization": "access-token",
        },
    })

    if err != nil {
        t.Fatalf("failed to create client: %s", err)
    }

    result, err := client.Query(context.Background(), promptql.NewQueryRequestMessage("what can you do?"))
    if err != nil {
        panic(err)
    }

    if len(result.AssistantActions) > 0 {
        if msg := result.AssistantActions[0].Message.Get(); msg != nil {
            log.Println(msg)
        }
    }

    // I can help you analyze data...
}
```

## Reference

### Natural Language 

The [Natural Language Query API](https://hasura.io/docs/promptql/promptql-apis/natural-language-api/) allows you to interact with PromptQL directly, sending messages and receiving responses.

#### Non-Streaming

```go
func (c *Client) Query(ctx context.Context, body api.QueryRequest) (*api.QueryResponse, error)
```

#### Streaming

The streaming response sends chunks of data in Server-Sent Events (SSE) format.
If the callback isn't set the client returns the raw response and you need to handle the response manually.

```go
func (c *Client) QueryStream(ctx context.Context, body api.QueryRequest, callback QueryStreamCallback) error
```

Example:

```go
err := client.QueryStream(
    context.Background(), 
    promptql.NewQueryRequestMessage("get list customers"), 
    func(chunk api.QueryResponseChunk) error {
        log.Println(chunk
        
        return nil
    },
)
```

### Execute Program

Execute a PromptQL program with your data.

```go
func (c *Client) ExecuteProgram(ctx context.Context, body api.ExecuteRequest) (*api.PromptQlExecutionResult, error)
```

## Development

### Generate codes

Use the following command to update Go types of PromptQL APIs from OpenAPI document.

```bash
make openapi-gen
```