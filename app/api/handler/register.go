package handler

import (
	"context"

	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/api/apiutils"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

type RegisterRequest struct {
	Teamname string         `json:"team_name,required"`
	Players  []*user.Player `json:"players,required"`
}

// 处理 Register 请求的 handler
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	var req RegisterRequest
	// 解析请求体
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}
	log.Printf("Get register request %v\n", req)

	resp, err := infra.UserClient.Register(ctx, &user.RegisterRequest{
		Teamname: req.Teamname,
		Players:  req.Players,
	})
	log.Printf("resp: %v\n", resp)
	
	if apiutils.NotError(c, resp, err) {
		c.JSON(consts.StatusOK, user.TeamInfo{
			TeamId:   resp.Teaminfo.TeamId,
			Teamname: resp.Teaminfo.Teamname,
			Players:  resp.Teaminfo.Players,
		})
	}
}
