package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/imag-er/wendingcup/app/user/biz/dal/model"
	"github.com/imag-er/wendingcup/app/user/biz/dal/mysql"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	klog.Info("RegisterService: ", req.Teamname, " ", req.Players)
	// 判断队伍人数是否为2
	if len(req.Players) >= 3 {
		return &user.RegisterResponse{
			Teaminfo: nil,
			Code:     1004000,
			Msg:      "注册失败, 队伍人数最多2人",
		}, err
	}
	// 判断队伍名是否已存在
	var count int64
	err = mysql.DB.Model(&model.Team{}).Where("name = ?", req.Teamname).Count(&count).Error
	if err != nil || count > 0 {
		return &user.RegisterResponse{
			Teaminfo: nil,
			Code:     1004001,
			Msg:      "注册失败, 可能是以下原因: 1. 队伍名已存在 2. 服务器内部错误 \n 请更换队伍名重试 ",
		}, err
	}

	data := &model.Team{
		Name: req.Teamname,
		UUID: uuid.New().String(),
	}

	// 判断当前玩家是否已注册
	for _, player := range req.Players {
		var count int64
		err = mysql.DB.Model(&model.Player{}).Where("student_id = ?", player.StudentId).Count(&count).Error
		if err != nil || count > 0 {
			return &user.RegisterResponse{
				Teaminfo: nil,
				Code:     1004002,
				Msg:      "注册失败, 用户\"" + player.Name + "\"已存在 \n 请联系管理员 ",
			}, err
		}

		data.Players = append(data.Players, model.Player{
			TeamID:    data.UUID,
			Name:      player.Name,
			Phone:     player.Phonenumber,
			StudentId: player.StudentId,
		})
	}
	err = mysql.DB.Create(data).Error
	if err != nil {
		return &user.RegisterResponse{
			Teaminfo: nil,
			Code:     1004002,
			Msg:      "注册失败, 服务器内部错误 \n 请联系管理员 ",
		}, err
	}

	return &user.RegisterResponse{
		Code: 0,
		Msg:  "注册成功",
		Teaminfo: &user.TeamInfo{
			TeamId:   data.UUID,
			Teamname: req.Teamname,
			Players:  req.Players,
		},
	}, nil
}
