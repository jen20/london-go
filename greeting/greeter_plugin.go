package greeting

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type GreeterPlugin struct {
	Impl Greeter
}

func (plugin GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{
		Impl: plugin.Impl,
	}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
