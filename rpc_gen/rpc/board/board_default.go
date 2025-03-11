package board

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Board(ctx context.Context, req *board.BoardRequest, callOptions ...callopt.Option) (resp *board.BoardResponse, err error) {
	resp, err = defaultClient.Board(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Board call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
