package service

import (
	"context"

	"github.com/imag-er/wendingcup/app/board/biz/dal/model"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
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
	err = mysql.DB.Create(&model.Result{
		TeamId:          req.JudgeResult.TeamId,
		Score:           req.JudgeResult.Score,
		FileUploadTime:  req.JudgeResult.FileUploadTime,
		JudgeResultTime: req.JudgeResult.JudgeResultTime,
	}).Error
	if err != nil {
		return nil, err
	}
	resp = &board.AppendJudgeResultResponse{}

	return
}
