package main

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
	"github.com/imag-er/wendingcup/app/board/biz/service"
)

// BoardImpl implements the last service interface defined in the IDL.
type BoardImpl struct{}

// Board implements the BoardImpl interface.
func (s *BoardImpl) Board(ctx context.Context, req *board.BoardRequest) (resp *board.BoardResponse, err error) {
	resp, err = service.NewBoardService(ctx).Run(req)

	return resp, err
}
