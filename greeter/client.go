package greeter

import (
	"golang.org/x/net/context"

	"go.uber.org/zap"

	pb "github.com/hg2c/hellogrpc/helloworld"
	"github.com/hg2c/hellogrpc/rpc"
)

type Client struct {
	*rpc.Client
	client pb.GreeterClient
}

func NewClient(name string, hostPort string) *Client {
	ct := rpc.NewClientWithTracing(name, hostPort)
	c := pb.NewGreeterClient(ct.Conn())

	return &Client{ct, c}
}

func (c *Client) Hello(name string) {
	defer c.Close()

	c.Logger().Info("hello", zap.String("name", name))
	r, err := c.client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		c.Logger().Info("could not greet: ", zap.Error(err))
	} else {
		c.Logger().Info("Greeting: ", zap.String("message", r.Message))
	}
}
