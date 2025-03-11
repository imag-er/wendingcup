package board

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"

	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board/board"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() board.Client
	Service() string
	Board(ctx context.Context, Req *board.BoardRequest, callOptions ...callopt.Option) (r *board.BoardResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := board.NewClient(dstService, opts...)
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
	kitexClient board.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() board.Client {
	return c.kitexClient
}

func (c *clientImpl) Board(ctx context.Context, Req *board.BoardRequest, callOptions ...callopt.Option) (r *board.BoardResponse, err error) {
	return c.kitexClient.Board(ctx, Req, callOptions...)
}
