package gateway

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/gateway/biz/service"
	"github.com/imag-er/wendingcup/app/gateway/biz/utils"
	gateway "github.com/imag-er/wendingcup/app/gateway/hertz_gen/gateway"
)

// GetBoard .
// @router /board [GET]
func GetBoard(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.GetBoardRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.GetBoardResponse{}
	resp, err = service.NewGetBoardService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
