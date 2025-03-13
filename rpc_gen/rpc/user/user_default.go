package user

import (
	"context"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (resp *user.LoginResponse, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (resp *user.RegisterResponse, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetTeamInfo(ctx context.Context, req *user.GetTeamInfoRequst, callOptions ...callopt.Option) (resp *user.TeamInfo, err error) {
	resp, err = defaultClient.GetTeamInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetTeamInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
