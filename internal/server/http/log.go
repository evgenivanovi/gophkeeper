package http

import (
	"context"
	"log/slog"

	"github.com/evgenivanovi/gpl/std"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

// LogSuccessRequest ...
func LogSuccessRequest(request any) {
	if request != nil {
		slogx.Log().Debug(
			"Incoming request",
			slog.Any("request", &request),
		)
	} else {
		slogx.Log().Debug(
			"Incoming request",
			slog.Any("request", nil),
		)
	}
}

// LogSuccessRequestCtx ...
func LogSuccessRequestCtx(ctx context.Context, request any) {
	if request != nil {
		slogx.FromCtx(ctx).Debug(
			"Incoming request",
			slog.Any("request", &request),
		)
	} else {
		slogx.FromCtx(ctx).Debug(
			"Incoming request",
			slog.Any("request", nil),
		)
	}
}

// LogErrorRequest ...
func LogErrorRequest(err error) {
	if err != nil {
		slogx.Log().Error(
			"Incoming request",
			slog.String("error", err.Error()),
		)
	} else {
		slogx.Log().Error(
			"Incoming request",
			slog.String("error", std.Nil),
		)
	}
}

// LogErrorRequestCtx ...
func LogErrorRequestCtx(ctx context.Context, err error) {
	if err != nil {
		slogx.FromCtx(ctx).Error(
			"Incoming request",
			slog.String("error", err.Error()),
		)
	} else {
		slogx.FromCtx(ctx).Error(
			"Incoming request",
			slog.String("error", std.Nil),
		)
	}
}

// LogSuccessResponse ...
func LogSuccessResponse(response any) {
	if response != nil {
		slogx.Log().Debug(
			"Outcoming response",
			slog.Any("response", &response),
		)
	} else {
		slogx.Log().Debug(
			"Outcoming response",
			slog.Any("response", nil),
		)
	}
}

// LogSuccessResponseCtx ...
func LogSuccessResponseCtx(ctx context.Context, response any) {
	if response != nil {
		slogx.FromCtx(ctx).Debug(
			"Outcoming response",
			slog.Any("response", &response),
		)
	} else {
		slogx.FromCtx(ctx).Debug(
			"Outcoming response",
			slog.Any("response", nil),
		)
	}
}

// LogErrorResponse ...
func LogErrorResponse(err error) {
	if err != nil {
		slogx.Log().Error(
			"Outcoming response",
			slog.String("error", err.Error()),
		)
	} else {
		slogx.Log().Error(
			"Outcoming response",
			slog.String("error", std.Nil),
		)
	}
}

// LogErrorResponseCtx ...
func LogErrorResponseCtx(ctx context.Context, err error) {
	if err != nil {
		slogx.FromCtx(ctx).Error(
			"Outcoming response",
			slog.String("error", err.Error()),
		)
	} else {
		slogx.FromCtx(ctx).Error(
			"Outcoming response",
			slog.String("error", std.Nil),
		)
	}
}
