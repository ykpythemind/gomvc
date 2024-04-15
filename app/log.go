package gomvc

import (
	"context"
	"log/slog"
)

type key struct{}

var RequestIDKey = key{}

type LogHandler struct {
	slog.Handler
}

func NewLogHandler(h slog.Handler) *LogHandler {
	return &LogHandler{h}
}

func (h *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	if val, ok := ctx.Value(RequestIDKey).(string); ok {
		r.AddAttrs(slog.String("requestId", val))
	}
	return h.Handler.Handle(ctx, r)
}
