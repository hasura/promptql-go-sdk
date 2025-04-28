package promptql

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/hasura/promptql-go-sdk/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("PromptQLClient")

const (
	authorizationHeader        = "authorization"
	contentTypeHeader          = "Content-Type"
	contentTypeTextEventStream = "text/event-stream"
)

type transportWrapper struct {
	propagator    propagation.TextMapPropagator
	doer          api.Doer
	logger        *slog.Logger
	serverAddress string
	serverPort    int
}

func (tw *transportWrapper) Do(req *http.Request) (*http.Response, error) { //nolint:cyclop
	ctx, span := tracer.Start(
		req.Context(),
		req.Method+" "+req.URL.Path,
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	span.SetAttributes(
		attribute.String("http.request.method", req.Method),
		attribute.String("url.full", req.URL.String()),
		attribute.String("server.address", tw.serverAddress),
		attribute.Int("server.port", tw.serverPort),
		attribute.String("network.protocol.name", "http"),
	)

	requestLogAttrs := map[string]any{
		"url":    req.URL.String(),
		"method": req.Method,
	}

	requestHeaders := http.Header{}

	for _, key := range []string{contentTypeHeader, authorizationHeader, "accept"} {
		values := req.Header.Values(key)
		if len(values) == 0 {
			continue
		}

		if key == authorizationHeader {
			values = []string{"<censored>"}
		}

		span.SetAttributes(attribute.StringSlice("http.request.header."+key, values))
	}

	isDebug := tw.isDebug(ctx)
	if isDebug {
		requestHeaders = req.Header

		if req.Body != nil && req.ContentLength > 0 {
			rawBody, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}

			requestLogAttrs["body"] = string(rawBody)

			_ = req.Body.Close()
			req.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		}
	}

	requestLogAttrs["headers"] = requestHeaders

	logAttrs := []slog.Attr{
		slog.String("type", "promptql-client"),
		slog.Any("request", requestLogAttrs),
	}

	if req.ContentLength > 0 {
		span.SetAttributes(attribute.Int64("http.request.body.size", req.ContentLength))
	}

	tw.propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := tw.doer.Do(req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		tw.log(ctx, slog.LevelError, "failed to execute the request: "+err.Error(), logAttrs)

		return resp, err
	}

	span.SetAttributes(attribute.Int("http.response.status_code", resp.StatusCode))

	if resp.ContentLength >= 0 {
		span.SetAttributes(attribute.Int64("http.response.size", resp.ContentLength))
	}

	for key, values := range resp.Header {
		if len(values) > 0 {
			span.SetAttributes(attribute.StringSlice("http.request.header."+key, values))
		}
	}

	respLogAttrs := map[string]any{
		"http_status": resp.StatusCode,
		"headers":     resp.Header,
	}

	if isDebug && resp.Body != nil && resp.ContentLength > 0 &&
		resp.Header.Get(contentTypeHeader) != contentTypeTextEventStream {
		respBody, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()

		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)

			logAttrs = append(logAttrs, slog.Any("response", respLogAttrs))

			tw.log(ctx, slog.LevelError, "failed to read response body: "+err.Error(), logAttrs)

			return resp, err
		}

		respLogAttrs["body"] = string(respBody)
		resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

		span.SetAttributes(attribute.Int("http.response.size", len(respBody)))
	}

	logLevel := slog.LevelDebug
	logAttrs = append(logAttrs, slog.Any("response", respLogAttrs))

	if resp.StatusCode >= 400 {
		span.SetStatus(codes.Error, resp.Status)

		logLevel = slog.LevelError
	}

	tw.log(ctx, logLevel, resp.Status, logAttrs)

	return resp, err
}

func (tw *transportWrapper) isDebug(ctx context.Context) bool {
	return tw.logger != nil && tw.logger.Enabled(ctx, slog.LevelDebug)
}

func (tw *transportWrapper) log(
	ctx context.Context,
	level slog.Level,
	msg string,
	attrs []slog.Attr,
) {
	if tw.logger != nil {
		tw.logger.LogAttrs(ctx, level, msg, attrs...)
	}
}
