package main

import (
	"context"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
	"github.com/imag-er/wendingcup/app/submit/biz/service"
)

// SubmitImpl implements the last service interface defined in the IDL.
type SubmitImpl struct{}

// Submit implements the SubmitImpl interface.
func (s *SubmitImpl) Submit(ctx context.Context, req *submit.SubmitRequest) (resp *submit.SubmitResponse, err error) {
	resp, err = service.NewSubmitService(ctx).Run(req)

	return resp, err
}
