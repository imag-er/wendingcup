package service

import (
	"context"
	"testing"
	judge "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/judge"
)

func TestJudge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewJudgeService(ctx)
	// init req and assert value

	req := &judge.JudgeRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
