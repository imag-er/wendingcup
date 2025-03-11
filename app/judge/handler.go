package main

import (
	"context"
	judge "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/judge"
	"github.com/imag-er/wendingcup/app/judge/biz/service"
)

// JudgeImpl implements the last service interface defined in the IDL.
type JudgeImpl struct{}

// Judge implements the JudgeImpl interface.
func (s *JudgeImpl) Judge(ctx context.Context, req *judge.JudgeRequest) (resp *judge.JudgeResponse, err error) {
	resp, err = service.NewJudgeService(ctx).Run(req)

	return resp, err
}
