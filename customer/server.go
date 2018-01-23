//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package customer

import (
	"github.com/hg2c/hellogrpc/customer/proto"
	"github.com/hwgo/pher/wgrpc"
	"golang.org/x/net/context"
)

type customerServer struct{}

func (s *customerServer) Get(context.Context, *proto.CustomerRequest) (*proto.CustomerReply, error) {
	return &proto.CustomerReply{
			Id:       "218",
			Name:     "Tom",
			Location: "ChongQing",
		},
		nil
}

func NewServer(name string, hostPort string) *wgrpc.Server {
	s := wgrpc.NewServerWithTracing(name, hostPort)
	proto.RegisterCustomerServer(s.Gs, &customerServer{})
	return s
}
