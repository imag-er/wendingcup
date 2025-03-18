package service

import (
	"context"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
	"github.com/imag-er/wendingcup/app/user/biz/dal/mysql"
	"errors"
)

type GetTeamInfoService struct {
	ctx context.Context
} // NewGetTeamInfoService new GetTeamInfoService
func NewGetTeamInfoService(ctx context.Context) *GetTeamInfoService {
	return &GetTeamInfoService{ctx: ctx}
}

// Run create note info
func (s *GetTeamInfoService) Run(req *user.GetTeamInfoRequest) (resp *user.GetTeamInfoResponse, err error) {
	// Finish your business logic.

	var val user.TeamInfo
	err = mysql.DB.Where("uuid =?", req.TeamId).Model(&val).Error
	if err != nil {
		return nil, errors.New("找不到队伍")
	}
	return &user.GetTeamInfoResponse{
		Teaminfo: &val,
	}, nil
	
}
