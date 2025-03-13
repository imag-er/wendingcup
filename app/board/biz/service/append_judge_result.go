package service

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

type AppendJudgeResultService struct {
	ctx context.Context
} // NewAppendJudgeResultService new AppendJudgeResultService
func NewAppendJudgeResultService(ctx context.Context) *AppendJudgeResultService {
	return &AppendJudgeResultService{ctx: ctx}
}

// Run create note info
func (s *AppendJudgeResultService) Run(req *board.AppendJudgeResultRequest) (resp *board.AppendJudgeResultResponse, err error) {
	// Finish your business logic.

	return
}
