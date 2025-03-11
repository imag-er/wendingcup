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
