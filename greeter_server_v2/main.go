/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"

	"net"

	// "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/hwgo/pher/otgrpc"
	_ "net/http"
	_ "net/http/pprof"

	pb "github.com/hg2c/hellogrpc/helloworld"
	lib "github.com/hg2c/hellogrpc/lib"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = lib.GetConfig("greeter.port", ":50051")
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

func main() {
	logger0, _ := zap.NewDevelopment()
	logger := log.NewFactory(logger0.With(zap.String("service", "greeter")))

	metricsFactory := xkit.Wrap("", expvar.NewFactory(10)) // 10 buckets for histograms
	logger.Bg().Info("Using expvar as metrics backend")

	tracer := tracing.Init("greeter", metricsFactory.Namespace("greeter", nil), logger)

	server := NewServer(port, tracer, metricsFactory, logger)
	server.Run()
}
