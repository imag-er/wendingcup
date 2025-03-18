package handler

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/api/apiutils"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

// 定义 SubmitRequest 结构体
type SubmitRequest struct {
	TeamId string                `form:"team_id"`
	File   *multipart.FileHeader `form:"file"`
}

// 处理 Submit 请求的 handler
func SubmitHandler(ctx context.Context, c *app.RequestContext) {
	var req SubmitRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	// 读取文件内容
	file, err := req.File.Open()
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"error": "文件读取失败"})
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"error": "文件解析失败"})
		return
	}

	resp, err := infra.SubmitClient.Submit(ctx, &submit.SubmitRequest{
		TeamId: req.TeamId,
		File:   fileBytes, // 使用实际读取的文件字节
	})
	if apiutils.NotError(c, resp, err) {
		c.JSON(consts.StatusOK, utils.H{
			"id":      resp.SubmitId,
			"message": resp.Msg,
		})
	}

}
