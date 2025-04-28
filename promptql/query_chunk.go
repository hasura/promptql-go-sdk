package promptql

import (
	"github.com/hasura/promptql-go-sdk/api"
)

// QueryChunks are a convenient data type to help us join query response chunks together.
type QueryChunks struct {
	ArtifactUpdateChunks  map[string]api.ArtifactUpdateChunk
	AssistantActionChunks []api.AssistantActionChunk
	ErrorChunk            *api.ErrorChunk
}

// NewQueryChunks create an empty QueryChunks instance.
func NewQueryChunks() *QueryChunks {
	return &QueryChunks{
		ArtifactUpdateChunks: make(map[string]api.ArtifactUpdateChunk),
	}
}

// GetArtifactUpdateChunks return the artifact update chunks as a slice.
func (qc *QueryChunks) GetArtifactUpdateChunks() []api.ArtifactUpdateChunk {
	results := make([]api.ArtifactUpdateChunk, 0, len(qc.ArtifactUpdateChunks))

	for _, item := range qc.ArtifactUpdateChunks {
		results = append(results, item)
	}

	return results
}

// IsError checks if the chunk is an error.
func (qc *QueryChunks) IsError() bool {
	return qc.ErrorChunk != nil
}

// Add or merge the response chunk into the chunks list.
func (qc *QueryChunks) Add(chunk api.QueryResponseChunk) error {
	if chunk.IsNil() {
		return nil
	}

	switch chk := chunk.Interface().(type) {
	case *api.ArtifactUpdateChunk:
		if chk == nil || chk.Artifact.IsNil() {
			return nil
		}

		identifier := chk.Artifact.GetIdentifier()
		key := chk.Type + ":" + identifier

		if qc.ArtifactUpdateChunks == nil {
			qc.ArtifactUpdateChunks = make(map[string]api.ArtifactUpdateChunk)
		}

		artifact, ok := qc.ArtifactUpdateChunks[key]
		if ok {
			mergedArtifact := qc.mergeArtifact(artifact.Artifact, chk.Artifact)
			if mergedArtifact != nil {
				artifact.Artifact = *mergedArtifact
				qc.ArtifactUpdateChunks[key] = artifact
			}
		} else {
			qc.ArtifactUpdateChunks[key] = *chk
		}
	case *api.AssistantActionChunk:
		qc.mergeAssistantActionChunk(chk)
	case *api.ErrorChunk:
		if qc.ErrorChunk == nil {
			qc.ErrorChunk = chk
		} else {
			qc.ErrorChunk.Error += chk.Error
		}
	}

	return nil
}

func (qc *QueryChunks) mergeArtifact(src, target api.Artifact) *api.Artifact {
	switch ta := target.Interface().(type) {
	case *api.TableArtifact:
		sa := src.AsTable()
		sa.Data = append(sa.Data, ta.Data...)
		result := api.TableArtifactAsArtifact(sa)

		return &result
	case *api.TextArtifact:
		sa := src.AsText()
		sa.Data += ta.Data
		result := api.TextArtifactAsArtifact(sa)

		return &result
	case *api.VisualizationArtifact:
		sa := src.AsVisualization()
		sa.Data.Html += ta.Data.Html
		result := api.VisualizationArtifactAsArtifact(sa)

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
	for i := len(qc.AssistantActionChunks); i <= chunkIndex; i++ {
		qc.AssistantActionChunks = append(
			qc.AssistantActionChunks,
			*NewAssistantActionChunk(int32(i)),
		)
	}

	src := qc.AssistantActionChunks[chunkIndex]
	src.Code = concatNullableString(src.Code, chk.Code)
	src.Message = concatNullableString(src.Message, chk.Message)
	src.Plan = concatNullableString(src.Plan, chk.Plan)
	src.CodeOutput = concatNullableString(src.CodeOutput, chk.CodeOutput)
	src.CodeError = concatNullableString(src.CodeError, chk.CodeError)

	qc.AssistantActionChunks[chunkIndex] = src
}

// will change when the set of required properties is changed.
func NewAssistantActionChunk(index int32) *api.AssistantActionChunk {
	return &api.AssistantActionChunk{
		Type:  string(api.QueryResponseTypeAssistantActionChunk),
		Index: index,
	}
}

// will change when the set of required properties is changed.
func NewErrorChunk(error_ string) *api.ErrorChunk {
	return &api.ErrorChunk{
		Type:  string(api.QueryResponseTypeErrorChunk),
		Error: error_,
	}
}

// will change when the set of required properties is changed.
func NewArtifactUpdateChunk[A api.ArtifactInterface](artifact A) *api.ArtifactUpdateChunk {
	return &api.ArtifactUpdateChunk{
		Type:     string(api.QueryResponseTypeArtifactUpdateChunk),
		Artifact: api.NewArtifact(artifact),
	}
}
