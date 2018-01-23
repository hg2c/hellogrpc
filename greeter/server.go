//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package greeter

import (
	pb "github.com/hg2c/hellogrpc/helloworld"
	"github.com/hg2c/hellogrpc/rpc"
	"golang.org/x/net/context"
)

type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Nihao " + in.Name}, nil
}

func NewServer(name string, hostPort string) *rpc.Server {
	s := rpc.NewServerWithTracing(name, hostPort)
	pb.RegisterGreeterServer(s.Gs, &greeterServer{})
	return s
}
