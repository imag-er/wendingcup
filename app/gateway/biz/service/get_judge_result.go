package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/imag-er/wendingcup/app/gateway/hertz_gen/gateway"
)

type GetJudgeResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetJudgeResultService(Context context.Context, RequestContext *app.RequestContext) *GetJudgeResultService {
	return &GetJudgeResultService{RequestContext: RequestContext, Context: Context}
}

func (h *GetJudgeResultService) Run(req *gateway.GetJudgeResultRequest) (resp *gateway.GetJudgeResultResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
