package customer

import (
	"golang.org/x/net/context"

	"go.uber.org/zap"

	"github.com/hg2c/hellogrpc/customer/proto"
	"github.com/hwgo/pher/wgrpc"
)

type Client struct {
	*wgrpc.Client
	client proto.CustomerClient
}

func NewClient(name string, host string, port int) *Client {
	ct := wgrpc.NewClientWithTracing(name, host, port)
	c := proto.NewCustomerClient(ct.Conn())

	return &Client{ct, c}
}

func (c *Client) Get() {
	defer c.Close()

	r, err := c.client.Get(context.Background(), &proto.CustomerRequest{Id: "760"})
	if err != nil {
		c.Logger().Info("could not greet: ", zap.Error(err))
	} else {
		c.Logger().Info("Customer: ", zap.String("customer name", r.Name))
	}
}
