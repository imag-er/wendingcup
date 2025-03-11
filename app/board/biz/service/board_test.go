package service

import (
	"context"
	"testing"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

func TestBoard_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBoardService(ctx)
	// init req and assert value

	req := &board.BoardRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
