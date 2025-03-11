package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/imag-er/wendingcup/app/gateway/hertz_gen/gateway"
)

type SubmitService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSubmitService(Context context.Context, RequestContext *app.RequestContext) *SubmitService {
	return &SubmitService{RequestContext: RequestContext, Context: Context}
}

func (h *SubmitService) Run(req *gateway.SubmitRequest) (resp *gateway.SubmitResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
