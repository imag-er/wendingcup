package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imag-er/wendingcup/app/board/biz/dal/model"
	"github.com/imag-er/wendingcup/app/board/biz/dal/mysql"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
	"gorm.io/gorm"
)

type AppendJudgeResultService struct {
	ctx context.Context
} // NewAppendJudgeResultService new AppendJudgeResultService
func NewAppendJudgeResultService(ctx context.Context) *AppendJudgeResultService {
	return &AppendJudgeResultService{ctx: ctx}
}

// Run create note info
func (s *AppendJudgeResultService) Run(req *board.AppendJudgeResultRequest) (resp *board.AppendJudgeResultResponse, err error) {
	// 如果数据库中已经存在该记录 (TeamId为uniqueIndex,Primary)，则选择较高的Score更新
	// 否则插入新记录
	var val model.Result
	err = mysql.DB.Where("team_id = ?", req.JudgeResult.TeamId).First(&val).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		// 第一次提交
		klog.Info("第一次提交: ", req.JudgeResult.TeamId)
		val = model.Result{
			TeamId: req.JudgeResult.TeamId,
			Score:  req.JudgeResult.Score,
			FileUploadTime: req.JudgeResult.FileUploadTime,
			JudgeResultTime: req.JudgeResult.JudgeResultTime,
		}
		err = mysql.DB.Create(&val).Error
		if err != nil {
			klog.Error(err)
			return nil, err
		}
	} else {
		// 已经存在记录

		if req.JudgeResult.Score > val.Score {
			klog.Info("更新分数")
			err = mysql.DB.Model(&val).Update("score", req.JudgeResult.Score).Error
			if err != nil {
				klog.Error(err)
				return nil, err
			}
		}
	}


	resp = &board.AppendJudgeResultResponse{}

	return
}
