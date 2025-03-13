package board

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetBoard(ctx context.Context, req *board.GetBoardRequest, callOptions ...callopt.Option) (resp *board.GetBoardResponse, err error) {
	resp, err = defaultClient.GetBoard(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetBoard call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AppendJudgeResult(ctx context.Context, req *board.AppendJudgeResultRequest, callOptions ...callopt.Option) (resp *board.AppendJudgeResultResponse, err error) {
	resp, err = defaultClient.AppendJudgeResult(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AppendJudgeResult call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
