package infra

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	
	user "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user"
)

var AuthMiddleware *jwt.HertzJWTMiddleware
var identityKey = "id"

type loginPayload struct {
	UUID string `form:"team_id,required" json:"team_id,required"`
}

func InitJWT() {
	// the jwt middleware
	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "wending-jwt",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals loginPayload
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			hlog.Infof("loginVals: %+v", loginVals)
			resp, err := UserClient.Login(ctx, &user.LoginRequest{
				TeamId: loginVals.UUID,
			})
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if resp.Code != 0 {
				return nil, errors.New(fmt.Sprintf("%s 错误码:%d", resp.Msg, resp.Code))
			}
			
			return &resp.Teaminfo, nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}
