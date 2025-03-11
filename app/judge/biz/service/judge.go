package service

import (
	"context"
	judge "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/judge"
)

type JudgeService struct {
	ctx context.Context
} // NewJudgeService new JudgeService
func NewJudgeService(ctx context.Context) *JudgeService {
	return &JudgeService{ctx: ctx}
}

// Run create note info
func (s *JudgeService) Run(req *judge.JudgeRequest) (resp *judge.JudgeResponse, err error) {
	// Finish your business logic.
	
	return
}
