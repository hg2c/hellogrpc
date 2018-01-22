package greeter

import (
	pb "github.com/hg2c/hellogrpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/charithe/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
)

type Client struct {
	tracer opentracing.Tracer
	logger log.Factory
	client pb.GreeterClient
	cc     *grpc.ClientConn
}

func NewClient(hostPort string, tracer opentracing.Tracer, logger log.Factory) *Client {
	th := otgrpc.NewTraceHandler(tracer)
	conn, err := grpc.Dial(
		hostPort,
		grpc.WithStatsHandler(th),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Bg().Fatal("did not connect: ", zap.Error(err))
	}

	c := pb.NewGreeterClient(conn)

	return &Client{
		tracer: tracer,
		logger: logger,
		client: c,
		cc:     conn,
	}
}

func (c *Client) Hello(name string) {
	defer c.cc.Close()

	c.logger.Bg().Info("hello", zap.String("name", name))
	r, err := c.client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		c.logger.Bg().Info("could not greet: ", zap.Error(err))
	} else {
		c.logger.Bg().Info("Greeting: ", zap.String("message", r.Message))
	}
}
