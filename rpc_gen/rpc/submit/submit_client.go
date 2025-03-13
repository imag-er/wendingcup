package submit

import (
	"context"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"

	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit/submit"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() submit.Client
	Service() string
	Submit(ctx context.Context, Req *submit.SubmitRequest, callOptions ...callopt.Option) (r *submit.SubmitResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := submit.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient submit.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() submit.Client {
	return c.kitexClient
}

func (c *clientImpl) Submit(ctx context.Context, Req *submit.SubmitRequest, callOptions ...callopt.Option) (r *submit.SubmitResponse, err error) {
	return c.kitexClient.Submit(ctx, Req, callOptions...)
}
