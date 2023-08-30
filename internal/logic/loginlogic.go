package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"login/internal/svc"
	"login/internal/types"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Request) (resp *types.Response, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Username)
	if err != nil {
		return
	}
	resp = &types.Response{
		Message: user.Name,
	}
	claims := make(jwt.MapClaims)

	now := time.Now().Unix()
	claims["exp"] = now + l.svcCtx.Config.Auth.AccessExpire
	claims["alt"] = now
	claims["userId"] = req.Username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	fmt.Println("token")

	jwtToken, _ := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	fmt.Println(jwtToken)
	return
}
