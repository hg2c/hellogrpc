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

	"github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"
	"go.uber.org/zap"

	"net"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"net/http"
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

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Nihao " + in.Name}, nil
}

func main() {
	logger0, _ := zap.NewDevelopment()
	logger := log.NewFactory(logger0.With(zap.String("service", "greeter")))

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Bg().Fatal("failed to listen: %v", zap.Error(err))
	}

	metricsFactory := xkit.Wrap("", expvar.NewFactory(10)) // 10 buckets for histograms
	logger.Bg().Info("Using expvar as metrics backend")

	tracer := tracing.Init("greeter", metricsFactory.Namespace("greeter", nil), logger)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer)))

	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logger.Bg().Fatal("failed to serve: %v", zap.Error(err))
	}
}
