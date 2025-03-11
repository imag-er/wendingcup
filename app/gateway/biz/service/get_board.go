package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/imag-er/wendingcup/app/gateway/hertz_gen/gateway"
)

type GetBoardService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBoardService(Context context.Context, RequestContext *app.RequestContext) *GetBoardService {
	return &GetBoardService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBoardService) Run(req *gateway.GetBoardRequest) (resp *gateway.GetBoardResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
