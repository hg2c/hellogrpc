package rpc

import (
	"google.golang.org/grpc"

	"github.com/charithe/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"
)

type Client struct {
	tracer opentracing.Tracer
	logger log.Factory
	cc     *grpc.ClientConn
}

func NewClientWithTracing(name string, hostPort string) *Client {
	logger := log.NewFactory(logger.With(zap.String("client", name)))
	tracer := tracing.Init(name, metricsFactory.Namespace(name, nil), logger)

	th := otgrpc.NewTraceHandler(tracer)
	conn, err := grpc.Dial(
		hostPort,
		grpc.WithStatsHandler(th),
		grpc.WithInsecure(),
	)

	if err != nil {
		logger.Bg().Fatal("did not connect: ", zap.Error(err))
	}

	return &Client{
		tracer: tracer,
		logger: logger,
		cc:     conn,
	}
}

func (c *Client) Conn() *grpc.ClientConn {
	return c.cc
}

func (c *Client) Logger() log.Logger {
	return c.logger.Bg()
}

func (c *Client) Close() {
	c.cc.Close()
}
