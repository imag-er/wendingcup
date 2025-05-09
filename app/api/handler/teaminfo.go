package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imag-er/wendingcup/app/api/apiutils"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

// 处理 TeamInfo 请求的 handler
func TeamInfoHandler(ctx context.Context, c *app.RequestContext) {
	resp, err := infra.UserClient.GetTeamInfo(ctx, &user.GetTeamInfoRequest{
		TeamId: c.Param("team_id"),
	})
	klog.Infof("GetTeamInfo resp: %v, err: %v", resp, err)
	
	if apiutils.NotError(c, resp, err) {
		c.JSON(consts.StatusOK, resp.Teaminfo)
	}
}
