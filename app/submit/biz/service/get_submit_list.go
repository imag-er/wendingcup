package service

import (
	"context"

	"github.com/imag-er/wendingcup/app/submit/biz/dal/model"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

type GetSubmitListService struct {
	ctx context.Context
} // NewGetSubmitListService new GetSubmitListService
func NewGetSubmitListService(ctx context.Context) *GetSubmitListService {
	return &GetSubmitListService{ctx: ctx}
}

// Run create note info
func (s *GetSubmitListService) Run(req *submit.GetSubmitListRequest) (resp *submit.GetSubmitListResponse, err error) {
	// 把这里查询到的结果返回给客户端

	var submits []model.Submit
	err = mysql.DB.Where("team_id = ?", req.TeamId).Find(&submits).Error
	if err != nil {
		return nil, err
	}

	resp = &submit.GetSubmitListResponse{
		Code:    0,
		Msg:     "查询成功",
	}
	for _,v := range submits {
		resp.SubmitList = append(resp.SubmitList, &submit.SubmitInfo{
			TeamId:   v.TeamId,
			Status:   v.Message,
			Time:  v.CreatedAt.Format("01/02 15:04"),
		})
	}
	return
}
