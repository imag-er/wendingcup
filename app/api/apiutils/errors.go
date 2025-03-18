package apiutils


import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"fmt"
)

type rpcerror interface {
	GetCode() uint32
	GetMsg() string
}

func NotError(c *app.RequestContext, e rpcerror, err error) bool{
	failed := false
	if e.GetCode() != 0 {
		c.JSON(consts.StatusOK, utils.H{
			"message": fmt.Sprintf("%s 错误码:%d", e.GetMsg(), e.GetCode()),
		})
		failed = true
	}
	if err != nil {
		c.Error(err)
		failed = true
	}
	return !failed
}
