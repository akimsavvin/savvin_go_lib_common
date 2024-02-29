package sl

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type Log = *slog.Logger

func Pkg(pkg string) slog.Attr {
	return slog.String("package", pkg)
}

func Mdl(mdl string) slog.Attr {
	return slog.String("module", mdl)
}

func Op(op string) slog.Attr {
	return slog.String("operation", op)
}

func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func ReqId(ctx context.Context) slog.Attr {
	traceId := ctx.Value("reqId").(string)
	return slog.String("request_id", traceId)
}

func CorId(ctx context.Context) slog.Attr {
	spanId := ctx.Value("corId").(string)
	return slog.String("correlation_id", spanId)
}

func WithReqId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, "reqId", reqId)
}

func WithCorId(ctx context.Context, corId string) context.Context {
	return context.WithValue(ctx, "corId", corId)
}

func CtxFromReq(req *http.Request) context.Context {
	ctx := req.Context()

	reqId := req.Header.Get("X-Request-ID")
	corId := req.Header.Get("X-Correlation-ID")

	if reqId == "" {
		reqId = uuid.NewString()
	}

	if corId == "" {
		corId = uuid.NewString()
	}

	ctx = WithReqId(ctx, reqId)
	ctx = WithCorId(ctx, corId)

	return ctx
}
