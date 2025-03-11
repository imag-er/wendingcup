package submit

import (
	"context"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SubmitFile(ctx context.Context, req *submit.SubmitFileRequest, callOptions ...callopt.Option) (resp *submit.SubmitFileResponse, err error) {
	resp, err = defaultClient.SubmitFile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SubmitFile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
