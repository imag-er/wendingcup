package service

import (
	"context"
	"testing"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

func TestGetBoard_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetBoardService(ctx)
	// init req and assert value

	req := &board.GetBoardRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
