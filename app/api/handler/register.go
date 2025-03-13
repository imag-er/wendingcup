package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
	"log"
)

type RegisterRequest struct {
	Teamname string         `form:"team_name"`
	Players  []*user.Player `form:"players"`
}
type Player struct {
	Name        string `form:"name"`
	Phonenumber string `form:"phonenumber"`
	StudentId   string `form:"student_id"`
}

type RegisterResponse struct {
	Teaminfo user.TeamInfo `form:"team_info"`
}

// 处理 Register 请求的 handler
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	log.Println("Get register request")
	var req RegisterRequest
	// 解析请求体
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}
	
	resp, err := infra.UserClient.Register(ctx, &user.RegisterRequest{
		Teamname: req.Teamname,
		Players:  req.Players,
	})
	if err != nil {
		c.Error(err)
		return
	}
	if resp.Code != 0 {
		c.JSON(consts.StatusOK, utils.H{
			"message": fmt.Sprintf("%s 错误码:%d", resp.Msg, resp.Code),
		})
		return
	}

	// 返回响应
	c.JSON(consts.StatusOK, RegisterResponse{
		// 解引用指针以匹配结构体字段类型
		// 复制 resp.Teaminfo 中的字段，避免直接复制包含 sync.Mutex 的结构体
		Teaminfo: user.TeamInfo{
			TeamId:   resp.Teaminfo.TeamId,
			Teamname: resp.Teaminfo.Teamname,
			Players:  resp.Teaminfo.Players,
		},
	})
}
