package service

import (
	"context"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

type SubmitFileService struct {
	ctx context.Context
} // NewSubmitFileService new SubmitFileService
func NewSubmitFileService(ctx context.Context) *SubmitFileService {
	return &SubmitFileService{ctx: ctx}
}

// Run create note info
func (s *SubmitFileService) Run(req *submit.SubmitFileRequest) (resp *submit.SubmitFileResponse, err error) {
	// Finish your business logic.

	return
}
