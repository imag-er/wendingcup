package service

import (
	"context"
	"testing"
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

func TestGetTeamInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetTeamInfoService(ctx)
	// init req and assert value

	req := &user.GetTeamInfoRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
