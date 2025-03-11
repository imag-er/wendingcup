package main

import (
	"context"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
	"github.com/imag-er/wendingcup/app/submit/biz/service"
)

// SubmitImpl implements the last service interface defined in the IDL.
type SubmitImpl struct{}

// SubmitFile implements the SubmitImpl interface.
func (s *SubmitImpl) SubmitFile(ctx context.Context, req *submit.SubmitFileRequest) (resp *submit.SubmitFileResponse, err error) {
	resp, err = service.NewSubmitFileService(ctx).Run(req)

	return resp, err
}
