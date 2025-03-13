package user

import (
	"context"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"

	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user/user"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() user.Client
	Service() string
	Login(ctx context.Context, Req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	Register(ctx context.Context, Req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	GetTeamInfo(ctx context.Context, Req *user.GetTeamInfoRequst, callOptions ...callopt.Option) (r *user.TeamInfo, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := user.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient user.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() user.Client {
	return c.kitexClient
}

func (c *clientImpl) Login(ctx context.Context, Req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) Register(ctx context.Context, Req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) GetTeamInfo(ctx context.Context, Req *user.GetTeamInfoRequst, callOptions ...callopt.Option) (r *user.TeamInfo, err error) {
	return c.kitexClient.GetTeamInfo(ctx, Req, callOptions...)
}
