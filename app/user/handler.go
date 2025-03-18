package main

import (
	"context"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
	"github.com/imag-er/wendingcup/app/user/biz/service"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Register implements the UserImpl interface.
func (s *UserImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// GetTeamInfo implements the UserImpl interface.
func (s *UserImpl) GetTeamInfo(ctx context.Context, req *user.GetTeamInfoRequest) (resp *user.GetTeamInfoResponse, err error) {
	resp, err = service.NewGetTeamInfoService(ctx).Run(req)

	return resp, err
}
