package service

import (
	"context"
	board "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board"
)

type GetBoardService struct {
	ctx context.Context
} // NewGetBoardService new GetBoardService
func NewGetBoardService(ctx context.Context) *GetBoardService {
	return &GetBoardService{ctx: ctx}
}

// Run create note info
func (s *GetBoardService) Run(req *board.GetBoardRequest) (resp *board.GetBoardResponse, err error) {
	// Finish your business logic.

	return
}
