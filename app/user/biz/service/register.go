package service

import (
	"context"

	"github.com/imag-er/wendingcup/app/user/biz/dal/model"
	"github.com/imag-er/wendingcup/app/user/biz/dal/mysql"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
	"github.com/google/uuid"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// Finish your business logic.
	var val model.Team
	err = mysql.DB.First(&val, "team_name = ?", req.Teamname).Error
	if err == nil {
		resp = &user.RegisterResponse{
			Code: 1004002,
			Msg:  "队伍名已存在,请更换队伍名",
			Teaminfo: nil,
		}
		return
	}
	data := &model.Team{
		TeamName: req.Teamname,
		UUID:     uuid.New().String(),
	}

	for _, v := range req.Players {
		var newv = model.Player{
			Name:        v.Name,
			PhoneNumber: v.Phonenumber,
			StudentId:   v.StudentId,
			TeamUUID:    data.UUID,
		}
		data.Players = append(data.Players, newv)
		err = mysql.DB.Create(newv).Error
		if err != nil {
			return nil, err
		}
	}
	err = mysql.DB.Create(data).Error
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		Teaminfo: &user.TeamInfo{
			TeamId: data.UUID,
			Teamname: req.Teamname,
			Players: req.Players,
		},	
	},nil
}
