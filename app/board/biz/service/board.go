package service

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

type BoardService struct {
	ctx context.Context
} // NewBoardService new BoardService
func NewBoardService(ctx context.Context) *BoardService {
	return &BoardService{ctx: ctx}
}

// Run create note info
func (s *BoardService) Run(req *board.BoardRequest) (resp *board.BoardResponse, err error) {
	// Finish your business logic.
	
	return
}
