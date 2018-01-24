//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package greeter

import (
	"golang.org/x/net/context"

	"github.com/hwgo/pher/wgrpc"

	// "github.com/hg2c/hellogrpc/customer"
	pb "github.com/hg2c/hellogrpc/helloworld"
)

type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// client := customer.NewClient(
	// 	"customer_client",
	// 	"127.0.0.1",
	// 	50052,
	// )
	// defer client.Close()

	// client.LoggerFactory().For(ctx).Info("xxoo @ sayhello")
	// user := client.Get()

	// return &pb.HelloReply{Message: "Nihao " + user.Name}, nil
	return &pb.HelloReply{Message: "Nihao " + in.Name}, nil
}

func NewServer(name string, hostPort string) *wgrpc.Server {
	s := wgrpc.NewServerWithTracing(name, hostPort)
	pb.RegisterGreeterServer(s.Gs, &greeterServer{})
	return s
}
