package greeter

import (
	"golang.org/x/net/context"

	"go.uber.org/zap"

	pb "github.com/hg2c/hellogrpc/helloworld"
	"github.com/hwgo/pher/wgrpc"
)

type Client struct {
	*wgrpc.Client
	client pb.GreeterClient
}

func NewClient(name string, host string, port int) *Client {
	ct := wgrpc.NewClientWithTracing(name, host, port)
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
