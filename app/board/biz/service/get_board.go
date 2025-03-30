package service

import (
	"context"
	"github.com/imag-er/wendingcup/app/board/biz/dal/model"
	"github.com/imag-er/wendingcup/app/board/biz/dal/mysql"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

type GetBoardService struct {
	ctx context.Context
} // NewGetBoardService new GetBoardService
func NewGetBoardService(ctx context.Context) *GetBoardService {
	return &GetBoardService{ctx: ctx}
}

// Run create note info
func (s *GetBoardService) Run(req *board.GetBoardRequest) (resp *board.GetBoardResponse, err error) {
	// Finish your business logic.
	// 获取mysql.DB中以model.Record为模型的所有记录

	var records []model.Result
	err = mysql.DB.Order("score desc").Find(&records).Error

	if err != nil {
		return nil, err
	}
	resp = &board.GetBoardResponse{
		Code: 0,
		Msg:  "查询成功",
	}
	for _, v := range records {
		resp.JudgeResult = append(resp.JudgeResult, &board.JudgeResult{
			Score:           v.Score,
			FileUploadTime:  v.FileUploadTime,
			JudgeResultTime: v.JudgeResultTime,
			TeamName:        v.TeamName,
		})
	}

	return
}
