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
	GetBoard(ctx context.Context, Req *board.GetBoardRequest, callOptions ...callopt.Option) (r *board.GetBoardResponse, err error)
	AppendJudgeResult(ctx context.Context, Req *board.AppendJudgeResultRequest, callOptions ...callopt.Option) (r *board.AppendJudgeResultResponse, err error)
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

func (c *clientImpl) GetBoard(ctx context.Context, Req *board.GetBoardRequest, callOptions ...callopt.Option) (r *board.GetBoardResponse, err error) {
	return c.kitexClient.GetBoard(ctx, Req, callOptions...)
}

func (c *clientImpl) AppendJudgeResult(ctx context.Context, Req *board.AppendJudgeResultRequest, callOptions ...callopt.Option) (r *board.AppendJudgeResultResponse, err error) {
	return c.kitexClient.AppendJudgeResult(ctx, Req, callOptions...)
}
