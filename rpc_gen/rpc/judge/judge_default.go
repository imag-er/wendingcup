package judge

import (
	"context"
	judge "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/judge"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Judge(ctx context.Context, req *judge.JudgeRequest, callOptions ...callopt.Option) (resp *judge.JudgeResponse, err error) {
	resp, err = defaultClient.Judge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Judge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
