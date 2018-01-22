//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package greeter

import (
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"net"

	"github.com/charithe/otgrpc"
	_ "net/http"
	_ "net/http/pprof"

	pb "github.com/hg2c/hellogrpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	hostPort string
	tracer   opentracing.Tracer
	logger   log.Factory
	server   *grpc.Server
}

type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Nihao " + in.Name}, nil
}

func NewServer(hostPort string, tracer opentracing.Tracer, metricsFactory metrics.Factory, logger log.Factory) *Server {
	th := otgrpc.NewTraceHandler(tracer)
	s := grpc.NewServer(grpc.StatsHandler(th))

	pb.RegisterGreeterServer(s, &greeterServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return &Server{
		hostPort: hostPort,
		tracer:   tracer,
		logger:   logger,
		server:   s,
	}
}

// Run starts the Customer server
func (s *Server) Run() error {
	bg := s.logger.Bg()
	lis, err := net.Listen("tcp", s.hostPort)

	if err != nil {
		bg.Fatal("Unable to start server", zap.Error(err))
		return err
	}

	bg.Info("Starting", zap.String("address", "tcp://"+s.hostPort))
	return s.server.Serve(lis)
}
