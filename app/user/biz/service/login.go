package service

import (
	"context"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
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

	return
}
