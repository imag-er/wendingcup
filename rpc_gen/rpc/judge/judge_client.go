package judge

import (
	"context"
	judge "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/judge"

	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/judge/judge"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() judge.Client
	Service() string
	Judge(ctx context.Context, Req *judge.JudgeRequest, callOptions ...callopt.Option) (r *judge.JudgeResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := judge.NewClient(dstService, opts...)
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
	kitexClient judge.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() judge.Client {
	return c.kitexClient
}

func (c *clientImpl) Judge(ctx context.Context, Req *judge.JudgeRequest, callOptions ...callopt.Option) (r *judge.JudgeResponse, err error) {
	return c.kitexClient.Judge(ctx, Req, callOptions...)
}
