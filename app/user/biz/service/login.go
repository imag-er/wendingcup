package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
	"github.com/imag-er/wendingcup/app/user/biz/dal/model"
	"github.com/imag-er/wendingcup/app/user/biz/dal/mysql"

)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// Finish your business logic.
	klog.Infof("LoginService: %+v", req)
	var val model.Team
	err = mysql.DB.Where("uuid = ?", req.TeamId).First(&val).Error
	if err != nil {
		resp = &user.LoginResponse{
			Code: 1004001,
			Msg:  "找不到队伍,请检查队伍ID是否正确",
		}
		return
	}
	klog.Infof("Loged in: %+v", req)
	resp = &user.LoginResponse{
		Code: 0,
		Msg:  "登录成功",
	}
	return
}
