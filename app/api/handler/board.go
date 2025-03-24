package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/api/apiutils"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"

)

type BoardRequest struct {
}

func BoardHandler(ctx context.Context, c *app.RequestContext) {
	var req BoardRequest
	// 解析请求体
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	resp, err := infra.BoardClient.GetBoard(ctx, &board.GetBoardRequest{})

	// 返回响应
	if apiutils.NotError(c, resp, err) {
		c.JSON(consts.StatusOK, resp.JudgeResult)
	}
}
