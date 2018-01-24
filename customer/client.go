package customer

import (
	"golang.org/x/net/context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/hwgo/pher/log"
	"github.com/hwgo/pher/wgrpc"

	"github.com/hg2c/hellogrpc/customer/proto"
)

type Client struct {
	*wgrpc.Client
	client proto.CustomerClient
}

func NewClient(name string, host string, port int, tracer opentracing.Tracer, logger log.Factory) *Client {
	// ct := wgrpc.NewClientWithTracing(name, host, port)
	ct := wgrpc.NewOTClient(host, port, tracer, logger)
	c := proto.NewCustomerClient(ct.Conn())

	return &Client{ct, c}
}

func NewClient2(name string, host string, port int) *Client {
	ct := wgrpc.NewClientWithTracing(name, host, port)
	// ct := wgrpc.NewOTClient(name, host, port, tracer, logger)
	c := proto.NewCustomerClient(ct.Conn())

	return &Client{ct, c}
}

func (c *Client) Get() *proto.CustomerReply {
	defer c.Close()

	r, err := c.client.Get(context.Background(), &proto.CustomerRequest{Id: "760"})
	if err != nil {
		c.Logger().Info("could not greet: ", zap.Error(err))
		return nil
	} else {
		c.Logger().Info("Customer: ", zap.String("customer name", r.Name))
		return r
	}
}
