package submit

import (
	"context"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Submit(ctx context.Context, req *submit.SubmitRequest, callOptions ...callopt.Option) (resp *submit.SubmitResponse, err error) {
	resp, err = defaultClient.Submit(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Submit call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
