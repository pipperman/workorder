package server

import (
	v1 "github.com/pipperman/workorder/api/workorder/v1"
	"github.com/pipperman/workorder/internal/conf"
	"github.com/pipperman/workorder/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger, tp *tracesdk.TracerProvider) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterOrderHTTPServer(srv, greeter)
	return srv
}
