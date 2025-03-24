package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imag-er/wendingcup/app/user/biz/dal/model"
	"github.com/imag-er/wendingcup/app/user/biz/dal/mysql"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

type GetTeamInfoService struct {
	ctx context.Context
} // NewGetTeamInfoService new GetTeamInfoService
func NewGetTeamInfoService(ctx context.Context) *GetTeamInfoService {
	return &GetTeamInfoService{ctx: ctx}
}

// Run create note info
func (s *GetTeamInfoService) Run(req *user.GetTeamInfoRequest) (resp *user.GetTeamInfoResponse, err error) {
	var team_val model.Team
	err = mysql.DB.Where("uuid = ?", req.TeamId).First(&team_val).Error
	klog.Info(team_val)
	if err != nil {
		return &user.GetTeamInfoResponse{
			Code: 1004003,
			Msg:  "查询失队伍信息失败",
		}, nil
	}


	var player_val []*model.Player
	err = mysql.DB.Where("team_id = ?", req.TeamId).Find(&player_val).Error
	if err != nil {
		return &user.GetTeamInfoResponse{
			Code: 1004003,
			Msg:  "查询队伍成员信息失败",
		}, nil
	}

	

	resp = &user.GetTeamInfoResponse{
		Teaminfo: &user.TeamInfo{
			TeamId:   team_val.UUID,
			Teamname: team_val.Name,
			Players:  []*user.Player{},
		},
		Code: 0,
		Msg:  "success",
	}

	for _, v := range player_val {
		resp.Teaminfo.Players = append(resp.Teaminfo.Players, &user.Player{
			Phonenumber: v.Phone,
			StudentId:    v.StudentId,
			Name:    v.Name,
		})
	}

	return resp, nil
}
