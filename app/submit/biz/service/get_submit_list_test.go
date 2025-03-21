package service

import (
	"context"
	"testing"
	submit "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit"
)

func TestGetSubmitList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetSubmitListService(ctx)
	// init req and assert value

	req := &submit.GetSubmitListRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
