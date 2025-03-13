package service

import (
	"context"
	"errors"

	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

type GetTeamInfoService struct {
	ctx context.Context
} // NewGetTeamInfoService new GetTeamInfoService
func NewGetTeamInfoService(ctx context.Context) *GetTeamInfoService {
	return &GetTeamInfoService{ctx: ctx}
}

// Run create note info
func (s *GetTeamInfoService) Run(req *user.GetTeamInfoRequst) (resp *user.TeamInfo, err error) {
	// Finish your business logic.
	var val user.TeamInfo
	err = mysql.DB.Where("uuid =?", req.TeamId).Model(&val).Error
	if err!= nil {		
		return nil,errors.New("找不到队伍")
	}
	return &val, nil
}
