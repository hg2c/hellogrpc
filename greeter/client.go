package greeter

import (
	pb "github.com/hg2c/hellogrpc/helloworld"
	config "github.com/hg2c/hellogrpc/lib"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
)

var (
	address     = config.GetConfig("greeter.server", "localhost") + ":50051"
	defaultName = config.GetConfig("greeter.message", "world")
)

type Client struct {
	tracer opentracing.Tracer
	logger log.Factory
	client pb.GreeterClient
}

func NewClient(tracer opentracing.Tracer, logger log.Factory) *Client {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	// conn, err := grpc.Dial(
	// 	address,
	// 	grpc.WithInsecure(),
	// 	// grpc.WithUnaryInterceptor(
	// 	// 	otgrpc.OpenTracingClientInterceptor(tracer)),
	// )
	if err != nil {
		logger.Bg().Fatal("did not connect: ", zap.Error(err))
	}

	// logger.Bg().Info("conn: ", zap.String(conn))

	// defer conn.Close()
	c := pb.NewGreeterClient(conn)

	return &Client{
		tracer: tracer,
		logger: logger,
		client: c,
	}
}

func (c *Client) Hello(name string) {
	c.logger.Bg().Info("hello", zap.String("name", name))
	r, err := c.client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		c.logger.Bg().Info("could not greet: ", zap.Error(err))
	} else {
		c.logger.Bg().Info("Greeting: ", zap.String("message", r.Message))
	}
}
