package service

import (
	"context"
	"testing"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

func TestAppendJudgeResult_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAppendJudgeResultService(ctx)
	// init req and assert value

	req := &board.AppendJudgeResultRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
