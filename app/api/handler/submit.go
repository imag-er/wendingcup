package handler

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

// 定义 SubmitRequest 结构体
type SubmitRequest struct {
	TeamId string                  `form:"team_id"`
	File   *multipart.FileHeader   `form:"file"`
}

// 定义 SubmitResponse 结构体
type SubmitResponse struct {
	SubmitID string `form:"submit_id"`
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
		c.JSON(consts.StatusInternalServerError, utils.H{"error": "文件读取失败"})
		return
	}
	defer file.Close()
	
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.H{"error": "文件解析失败"})
		return
	}

	hlog.Infof("submit: %v", fileBytes)

	resp, err := infra.SubmitClient.Submit(ctx, &submit.SubmitRequest{
		TeamId: req.TeamId,
		File:   fileBytes,  // 使用实际读取的文件字节
	})
	if err != nil {
		c.Error(err)
	}
	// 返回响应
	c.JSON(consts.StatusOK, resp.SubmitId)
}
