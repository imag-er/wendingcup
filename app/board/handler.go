package main

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
	"github.com/imag-er/wendingcup/app/board/biz/service"

)

// BoardImpl implements the last service interface defined in the IDL.
type BoardImpl struct{}

// GetBoard implements the BoardImpl interface.
func (s *BoardImpl) GetBoard(ctx context.Context, req *board.GetBoardRequest) (resp *board.GetBoardResponse, err error) {
	resp, err = service.NewGetBoardService(ctx).Run(req)

	return resp, err
}

// AppendJudgeResult implements the BoardImpl interface.
func (s *BoardImpl) AppendJudgeResult(ctx context.Context, req *board.AppendJudgeResultRequest) (resp *board.AppendJudgeResultResponse, err error) {
	resp, err = service.NewAppendJudgeResultService(ctx).Run(req)

	return resp, err
}
