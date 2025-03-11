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
