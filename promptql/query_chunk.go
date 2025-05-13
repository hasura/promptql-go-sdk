package promptql

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/hasura/promptql-go-sdk/api"
)

// QueryChunks are a convenient data type to help us join query response chunks together.
type QueryChunks struct {
	threadId          uuid.UUID
	assistantActions  []api.ApiThreadAssistantAction
	modifiedArtifacts map[string]api.ExecuteRequestArtifactsInner
	errorChunk        *api.ErrorChunk

	// cache the sorted list of artifacts for the consistent order
	modifiedArtifactKeys []string
}

// NewQueryChunks create an empty QueryChunks instance.
func NewQueryChunks() *QueryChunks {
	return &QueryChunks{
		modifiedArtifacts: make(map[string]api.ExecuteRequestArtifactsInner),
	}
}

// GetThreadID return the thread id.
func (qc QueryChunks) GetThreadId() uuid.UUID {
	return qc.threadId
}

// SetThreadID set the thread id.
func (qc *QueryChunks) SetThreadId(threadId uuid.UUID) {
	qc.threadId = threadId
}

// GetModifiedArtifacts return a list of artifacts.
func (qc QueryChunks) GetModifiedArtifacts() []api.ExecuteRequestArtifactsInner {
	results := make([]api.ExecuteRequestArtifactsInner, len(qc.modifiedArtifacts))

	for i, key := range qc.modifiedArtifactKeys {
		results[i] = qc.modifiedArtifacts[key]
	}

	return results
}

// SetModifiedArtifacts set modified artifacts.
func (qc *QueryChunks) SetModifiedArtifacts(inputs []api.ExecuteRequestArtifactsInner) {
	qc.modifiedArtifacts = make(map[string]api.ExecuteRequestArtifactsInner)
	qc.modifiedArtifactKeys = make([]string, len(inputs))

	for i, artifact := range inputs {
		key := buildArtifactKey(api.Artifact(artifact))
		qc.modifiedArtifacts[key] = artifact
		qc.modifiedArtifactKeys[i] = key
	}

	slices.Sort(qc.modifiedArtifactKeys)
}

// GetAssistantActions return a list of assistant actions.
func (qc QueryChunks) GetAssistantActions() []api.ApiThreadAssistantAction {
	return qc.assistantActions
}

// SetAssistantActions set assistant actions.
func (qc *QueryChunks) SetAssistantActions(inputs []api.ApiThreadAssistantAction) {
	qc.assistantActions = inputs
}

// GetErrorChunk returns a nullable error chunk.
func (qc QueryChunks) GetErrorChunk() *api.ErrorChunk {
	return qc.errorChunk
}

// SetErrorChunk set the error chunk.
func (qc *QueryChunks) SetErrorChunk(input *api.ErrorChunk) {
	qc.errorChunk = input
}

// IsError checks if the chunk is an error.
func (qc QueryChunks) IsError() bool {
	return qc.errorChunk != nil
}

// Add or merge the response chunk into the chunks list.
func (qc *QueryChunks) Add(chunk api.QueryResponseChunk) error {
	if chunk.IsNil() {
		return nil
	}

	switch chk := chunk.Interface().(type) {
	case *api.ThreadMetadataChunk:
		threadId, err := uuid.Parse(chk.ThreadId)
		if err != nil {
			return fmt.Errorf("invalid thread metadata chunk: %w", err)
		}

		qc.threadId = threadId
	case *api.ArtifactUpdateChunk:
		if chk == nil || chk.Artifact.IsNil() {
			return nil
		}

		key := buildArtifactKey(chk.Artifact)

		if qc.modifiedArtifacts == nil {
			qc.modifiedArtifacts = make(map[string]api.ExecuteRequestArtifactsInner)
		}

		artifact, ok := qc.modifiedArtifacts[key]
		if ok {
			mergedArtifact := qc.mergeArtifact(artifact, api.ExecuteRequestArtifactsInner(chk.Artifact))
			if mergedArtifact != nil {
				qc.modifiedArtifacts[key] = *mergedArtifact
			}
		} else {
			qc.modifiedArtifactKeys = append(qc.modifiedArtifactKeys, key)
			slices.Sort(qc.modifiedArtifactKeys)
			qc.modifiedArtifacts[key] = api.ExecuteRequestArtifactsInner(chk.Artifact)
		}
	case *api.AssistantActionChunk:
		qc.mergeAssistantActionChunk(chk)
	case *api.ErrorChunk:
		if qc.errorChunk == nil {
			qc.errorChunk = chk
		} else {
			qc.errorChunk.Error += chk.Error
		}
	}

	return nil
}

// AsQueryResponse converts to a QueryResponse instance.
func (qc QueryChunks) AsQueryResponse() api.QueryResponse {
	return api.QueryResponse{
		ThreadId:          qc.threadId.String(),
		AssistantActions:  qc.GetAssistantActions(),
		ModifiedArtifacts: qc.GetModifiedArtifacts(),
	}
}

func (qc *QueryChunks) mergeArtifact(
	src, target api.ExecuteRequestArtifactsInner,
) *api.ExecuteRequestArtifactsInner {
	switch ta := target.Interface().(type) {
	case *api.TableArtifact:
		sa := src.AsTable()
		sa.Data = append(sa.Data, ta.Data...)
		result := api.TableArtifactAsExecuteRequestArtifactsInner(sa)

		return &result
	case *api.TextArtifact:
		sa := src.AsText()
		sa.Data += ta.Data
		result := api.TextArtifactAsExecuteRequestArtifactsInner(sa)

		return &result
	case *api.VisualizationArtifact:
		sa := src.AsVisualization()
		sa.Data.Html += ta.Data.Html
		result := api.VisualizationArtifactAsExecuteRequestArtifactsInner(sa)

		return &result
	default:
		return nil
	}
}

func (qc *QueryChunks) mergeAssistantActionChunk(chk *api.AssistantActionChunk) {
	if chk == nil {
		return
	}

	chunkIndex := int(chk.Index)
	for i := len(qc.assistantActions); i <= chunkIndex; i++ {
		qc.assistantActions = append(
			qc.assistantActions,
			*api.NewApiThreadAssistantAction(),
		)
	}

	src := qc.assistantActions[chunkIndex]
	src.Code = concatNullableString(src.Code, chk.Code)
	src.Message = concatNullableString(src.Message, chk.Message)
	src.Plan = concatNullableString(src.Plan, chk.Plan)
	src.CodeOutput = concatNullableString(src.CodeOutput, chk.CodeOutput)
	src.CodeError = concatNullableString(src.CodeError, chk.CodeError)

	qc.assistantActions[chunkIndex] = src
}

// will change when the set of required properties is changed.
func NewAssistantActionChunk(index int32) *api.AssistantActionChunk {
	return &api.AssistantActionChunk{
		Type:  string(api.QueryResponseTypeAssistantActionChunk),
		Index: index,
	}
}

// NewErrorChunk create a new error chunk.
func NewErrorChunk(error_ string) *api.ErrorChunk {
	return &api.ErrorChunk{
		Type:  string(api.QueryResponseTypeErrorChunk),
		Error: error_,
	}
}

// will change when the set of required properties is changed.
func NewThreadMetadataChunk(threadId uuid.UUID) *api.ThreadMetadataChunk {
	return &api.ThreadMetadataChunk{
		Type:     string(api.QueryResponseTypeThreadMetadataChunk),
		ThreadId: threadId.String(),
	}
}

// NewArtifactUpdateChunk create a new artifact update chunk.
func NewArtifactUpdateChunk[A api.ArtifactInterface](artifact A) *api.ArtifactUpdateChunk {
	return &api.ArtifactUpdateChunk{
		Type:     string(api.QueryResponseTypeArtifactUpdateChunk),
		Artifact: api.NewArtifact(artifact),
	}
}

func buildArtifactKey(artifact api.Artifact) string {
	identifier := artifact.GetIdentifier()

	return string(artifact.GetArtifactType()) + ":" + identifier
}
