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
	SubmitFile(ctx context.Context, Req *submit.SubmitFileRequest, callOptions ...callopt.Option) (r *submit.SubmitFileResponse, err error)
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

func (c *clientImpl) SubmitFile(ctx context.Context, Req *submit.SubmitFileRequest, callOptions ...callopt.Option) (r *submit.SubmitFileResponse, err error) {
	return c.kitexClient.SubmitFile(ctx, Req, callOptions...)
}
