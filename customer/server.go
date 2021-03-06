//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package customer

import (
	"golang.org/x/net/context"
	"time"

	// "github.com/hwgo/pher/delay"
	"github.com/hwgo/pher/wgrpc"

	// "github.com/hg2c/hellogrpc/config"
	"github.com/hg2c/hellogrpc/customer/proto"
)

type customerServer struct{}

func (s *customerServer) Get(context.Context, *proto.CustomerRequest) (*proto.CustomerReply, error) {
	// simulate RPC delay
	time.Sleep(7 * time.Millisecond)
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
